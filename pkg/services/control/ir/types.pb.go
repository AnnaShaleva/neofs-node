// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: pkg/services/control/ir/types.proto

package control

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Health status of the IR application.
type HealthStatus int32

const (
	// Undefined status, default value.
	HealthStatus_HEALTH_STATUS_UNDEFINED HealthStatus = 0
	// IR application is starting.
	HealthStatus_STARTING HealthStatus = 1
	// IR application is started and serves all services.
	HealthStatus_READY HealthStatus = 2
	// IR application is shutting down.
	HealthStatus_SHUTTING_DOWN HealthStatus = 3
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_UNDEFINED",
		1: "STARTING",
		2: "READY",
		3: "SHUTTING_DOWN",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_UNDEFINED": 0,
		"STARTING":                1,
		"READY":                   2,
		"SHUTTING_DOWN":           3,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_services_control_ir_types_proto_enumTypes[0].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_pkg_services_control_ir_types_proto_enumTypes[0]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_services_control_ir_types_proto_rawDescGZIP(), []int{0}
}

// Signature of some message.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key used for signing.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Binary signature.
	Sign []byte `protobuf:"bytes,2,opt,name=sign,json=signature,proto3" json:"sign,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_control_ir_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_control_ir_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_pkg_services_control_ir_types_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Signature) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

var File_pkg_services_control_ir_types_proto protoreflect.FileDescriptor

var file_pkg_services_control_ir_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x22, 0x36, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x17, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x57, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x48, 0x55, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x03, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_services_control_ir_types_proto_rawDescOnce sync.Once
	file_pkg_services_control_ir_types_proto_rawDescData = file_pkg_services_control_ir_types_proto_rawDesc
)

func file_pkg_services_control_ir_types_proto_rawDescGZIP() []byte {
	file_pkg_services_control_ir_types_proto_rawDescOnce.Do(func() {
		file_pkg_services_control_ir_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_services_control_ir_types_proto_rawDescData)
	})
	return file_pkg_services_control_ir_types_proto_rawDescData
}

var file_pkg_services_control_ir_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_services_control_ir_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_services_control_ir_types_proto_goTypes = []interface{}{
	(HealthStatus)(0), // 0: ircontrol.HealthStatus
	(*Signature)(nil), // 1: ircontrol.Signature
}
var file_pkg_services_control_ir_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_services_control_ir_types_proto_init() }
func file_pkg_services_control_ir_types_proto_init() {
	if File_pkg_services_control_ir_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_services_control_ir_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_services_control_ir_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_services_control_ir_types_proto_goTypes,
		DependencyIndexes: file_pkg_services_control_ir_types_proto_depIdxs,
		EnumInfos:         file_pkg_services_control_ir_types_proto_enumTypes,
		MessageInfos:      file_pkg_services_control_ir_types_proto_msgTypes,
	}.Build()
	File_pkg_services_control_ir_types_proto = out.File
	file_pkg_services_control_ir_types_proto_rawDesc = nil
	file_pkg_services_control_ir_types_proto_goTypes = nil
	file_pkg_services_control_ir_types_proto_depIdxs = nil
}
