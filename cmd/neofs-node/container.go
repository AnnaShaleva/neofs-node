package main

import (
	"context"
	"crypto/ecdsa"
	"strconv"

	apiClient "github.com/nspcc-dev/neofs-api-go/pkg/client"
	containerSDK "github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	containerGRPC "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	containerCore "github.com/nspcc-dev/neofs-node/pkg/core/container"
	netmapCore "github.com/nspcc-dev/neofs-node/pkg/core/netmap"
	"github.com/nspcc-dev/neofs-node/pkg/local_object_storage/engine"
	"github.com/nspcc-dev/neofs-node/pkg/morph/client"
	"github.com/nspcc-dev/neofs-node/pkg/morph/client/container"
	"github.com/nspcc-dev/neofs-node/pkg/morph/client/container/wrapper"
	"github.com/nspcc-dev/neofs-node/pkg/network"
	"github.com/nspcc-dev/neofs-node/pkg/network/cache"
	containerTransportGRPC "github.com/nspcc-dev/neofs-node/pkg/network/transport/container/grpc"
	containerService "github.com/nspcc-dev/neofs-node/pkg/services/container"
	loadcontroller "github.com/nspcc-dev/neofs-node/pkg/services/container/announcement/load/controller"
	loadroute "github.com/nspcc-dev/neofs-node/pkg/services/container/announcement/load/route"
	containerMorph "github.com/nspcc-dev/neofs-node/pkg/services/container/morph"
	"github.com/nspcc-dev/neofs-node/pkg/util/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func initContainerService(c *cfg) {
	staticClient, err := client.NewStatic(
		c.cfgMorph.client,
		c.cfgContainer.scriptHash,
		c.cfgContainer.fee,
	)
	fatalOnErr(err)

	cnrClient, err := container.New(staticClient)
	fatalOnErr(err)

	wrap, err := wrapper.New(cnrClient)
	fatalOnErr(err)

	c.cfgObject.cnrStorage = wrap // use RPC node as source of containers
	c.cfgObject.cnrClient = wrap

	containerGRPC.RegisterContainerServiceServer(c.cfgGRPC.server,
		containerTransportGRPC.New(
			containerService.NewSignService(
				c.key,
				containerService.NewResponseService(
					containerService.NewExecutionService(
						containerMorph.NewExecutor(cnrClient),
					),
					c.respSvc,
				),
			),
		),
	)
}

type morphLoadWriter struct {
	log *logger.Logger

	cnrMorphClient *wrapper.Wrapper

	key []byte
}

func (w *morphLoadWriter) Put(a containerSDK.UsedSpaceAnnouncement) error {
	w.log.Debug("save used space announcement in contract",
		zap.Uint64("epoch", a.Epoch()),
		zap.Stringer("cid", a.ContainerID()),
		zap.Uint64("size", a.UsedSpace()),
	)

	return w.cnrMorphClient.AnnounceLoad(a, w.key)
}

func (*morphLoadWriter) Close() error {
	return nil
}

type nopLoadWriter struct{}

func (nopLoadWriter) Put(containerSDK.UsedSpaceAnnouncement) error {
	return nil
}

func (nopLoadWriter) Close() error {
	return nil
}

type remoteLoadAnnounceProvider struct {
	key *ecdsa.PrivateKey

	loadAddrSrc network.LocalAddressSource

	clientCache *cache.ClientCache

	deadEndProvider loadcontroller.WriterProvider
}

func (r *remoteLoadAnnounceProvider) InitRemote(srv loadroute.ServerInfo) (loadcontroller.WriterProvider, error) {
	if srv == nil {
		return r.deadEndProvider, nil
	}

	addr := srv.Address()

	if r.loadAddrSrc.LocalAddress().String() == srv.Address() {
		// if local => return no-op writer
		return loadcontroller.SimpleWriterProvider(new(nopLoadWriter)), nil
	}

	ipAddr, err := network.IPAddrFromMultiaddr(addr)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert address to IP format")
	}

	c, err := r.clientCache.Get(r.key, ipAddr)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize API client")
	}

	return &remoteLoadAnnounceWriterProvider{
		client: c,
	}, nil
}

type remoteLoadAnnounceWriterProvider struct {
	client *apiClient.Client
}

func (p *remoteLoadAnnounceWriterProvider) InitWriter(ctx context.Context) (loadcontroller.Writer, error) {
	return &remoteLoadAnnounceWriter{
		ctx:    ctx,
		client: p.client,
	}, nil
}

type remoteLoadAnnounceWriter struct {
	ctx context.Context

	client *apiClient.Client

	buf []containerSDK.UsedSpaceAnnouncement
}

func (r *remoteLoadAnnounceWriter) Put(a containerSDK.UsedSpaceAnnouncement) error {
	r.buf = append(r.buf, a)

	return nil
}

func (r *remoteLoadAnnounceWriter) Close() error {
	return r.client.AnnounceContainerUsedSpace(r.ctx, r.buf)
}

type loadPlacementBuilder struct {
	log *logger.Logger

	nmSrc netmapCore.Source

	cnrSrc containerCore.Source
}

func (l *loadPlacementBuilder) BuildPlacement(epoch uint64, cid *containerSDK.ID) ([]netmap.Nodes, error) {
	cnrNodes, nm, err := l.buildPlacement(epoch, cid)
	if err != nil {
		return nil, err
	}

	const pivotPrefix = "load_announcement_"

	pivot := []byte(
		pivotPrefix + strconv.FormatUint(epoch, 10),
	)

	placement, err := nm.GetPlacementVectors(cnrNodes, pivot)
	if err != nil {
		return nil, errors.Wrap(err, "could not build placement vectors")
	}

	return placement, nil
}

func (l *loadPlacementBuilder) buildPlacement(epoch uint64, cid *containerSDK.ID) (netmap.ContainerNodes, *netmap.Netmap, error) {
	cnr, err := l.cnrSrc.Get(cid)
	if err != nil {
		return nil, nil, err
	}

	nm, err := l.nmSrc.GetNetMapByEpoch(epoch)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not get network map")
	}

	cnrNodes, err := nm.GetContainerNodes(cnr.PlacementPolicy(), cid.ToV2().GetValue())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not build container nodes")
	}

	return cnrNodes, nm, nil
}

type localStorageLoad struct {
	log *logger.Logger

	engine *engine.StorageEngine
}

func (d *localStorageLoad) Iterate(f loadcontroller.UsedSpaceFilter, h loadcontroller.UsedSpaceHandler) error {
	idList := engine.ListContainers(d.engine)

	for i := range idList {
		sz := engine.ContainerSize(d.engine, idList[i])

		d.log.Debug("container size in storage engine calculated successfully",
			zap.Uint64("size", sz),
			zap.Stringer("cid", idList[i]),
		)

		a := containerSDK.NewAnnouncement()
		a.SetContainerID(idList[i])
		a.SetUsedSpace(sz)

		if f != nil && !f(*a) {
			continue
		}

		if err := h(*a); err != nil {
			return err
		}
	}

	return nil
}
