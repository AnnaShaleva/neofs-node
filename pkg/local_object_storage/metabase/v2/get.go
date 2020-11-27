package meta

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	objectSDK "github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/nspcc-dev/neofs-node/pkg/core/object"
	"go.etcd.io/bbolt"
)

// Get returns object header for specified address.
func (db *DB) Get(addr *objectSDK.Address) (obj *object.Object, err error) {
	err = db.boltDB.View(func(tx *bbolt.Tx) error {
		obj, err = db.get(tx, addr, true)

		return err
	})

	return obj, err
}

func (db *DB) get(tx *bbolt.Tx, addr *objectSDK.Address, checkGraveyard bool) (*object.Object, error) {
	obj := object.New()
	key := objectKey(addr.ObjectID())
	cid := addr.ContainerID()

	if checkGraveyard && inGraveyard(tx, addr) {
		return nil, ErrAlreadyRemoved
	}

	// check in primary index
	data := getFromBucket(tx, primaryBucketName(cid), key)
	if len(data) != 0 {
		return obj, obj.Unmarshal(data)
	}

	// if not found then check in tombstone index
	data = getFromBucket(tx, tombstoneBucketName(cid), key)
	if len(data) != 0 {
		return obj, obj.Unmarshal(data)
	}

	// if not found then check in storage group index
	data = getFromBucket(tx, storageGroupBucketName(cid), key)
	if len(data) != 0 {
		return obj, obj.Unmarshal(data)
	}

	// if not found then check if object is a virtual
	return getVirtualObject(tx, cid, key)
}

func getFromBucket(tx *bbolt.Tx, name, key []byte) []byte {
	bkt := tx.Bucket(name)
	if bkt == nil {
		return nil
	}

	return bkt.Get(key)
}

func getVirtualObject(tx *bbolt.Tx, cid *container.ID, key []byte) (*object.Object, error) {
	parentBucket := tx.Bucket(parentBucketName(cid))
	if parentBucket == nil {
		return nil, ErrNotFound
	}

	relativeLst, err := decodeList(parentBucket.Get(key))
	if err != nil {
		return nil, err
	}

	if len(relativeLst) == 0 { // this should never happen though
		return nil, ErrNotFound
	}

	// pick last item, for now there is not difference which address to pick
	// but later list might be sorted so first or last value can be more
	// prioritized to choose
	virtualOID := relativeLst[len(relativeLst)-1]
	data := getFromBucket(tx, primaryBucketName(cid), virtualOID)

	child := object.New()

	err = child.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal child with parent: %w", err)
	}

	if child.GetParent() == nil { // this should never happen though
		return nil, ErrNotFound
	}

	return child.GetParent(), nil
}
