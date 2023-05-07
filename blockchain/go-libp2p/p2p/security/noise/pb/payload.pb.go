// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/payload.proto

package pb

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

type NoiseExtensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebtransportCerthashes [][]byte `protobuf:"bytes,1,rep,name=webtransport_certhashes,json=webtransportCerthashes" json:"webtransport_certhashes,omitempty"`
	StreamMuxers           []string `protobuf:"bytes,2,rep,name=stream_muxers,json=streamMuxers" json:"stream_muxers,omitempty"`
}

func (x *NoiseExtensions) Reset() {
	*x = NoiseExtensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoiseExtensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseExtensions) ProtoMessage() {}

func (x *NoiseExtensions) ProtoReflect() protoreflect.Message {
	mi := &file_pb_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseExtensions.ProtoReflect.Descriptor instead.
func (*NoiseExtensions) Descriptor() ([]byte, []int) {
	return file_pb_payload_proto_rawDescGZIP(), []int{0}
}

func (x *NoiseExtensions) GetWebtransportCerthashes() [][]byte {
	if x != nil {
		return x.WebtransportCerthashes
	}
	return nil
}

func (x *NoiseExtensions) GetStreamMuxers() []string {
	if x != nil {
		return x.StreamMuxers
	}
	return nil
}

type NoiseHandshakePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityKey []byte           `protobuf:"bytes,1,opt,name=identity_key,json=identityKey" json:"identity_key,omitempty"`
	IdentitySig []byte           `protobuf:"bytes,2,opt,name=identity_sig,json=identitySig" json:"identity_sig,omitempty"`
	Extensions  *NoiseExtensions `protobuf:"bytes,4,opt,name=extensions" json:"extensions,omitempty"`
}

func (x *NoiseHandshakePayload) Reset() {
	*x = NoiseHandshakePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoiseHandshakePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseHandshakePayload) ProtoMessage() {}

func (x *NoiseHandshakePayload) ProtoReflect() protoreflect.Message {
	mi := &file_pb_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseHandshakePayload.ProtoReflect.Descriptor instead.
func (*NoiseHandshakePayload) Descriptor() ([]byte, []int) {
	return file_pb_payload_proto_rawDescGZIP(), []int{1}
}

func (x *NoiseHandshakePayload) GetIdentityKey() []byte {
	if x != nil {
		return x.IdentityKey
	}
	return nil
}

func (x *NoiseHandshakePayload) GetIdentitySig() []byte {
	if x != nil {
		return x.IdentitySig
	}
	return nil
}

func (x *NoiseHandshakePayload) GetExtensions() *NoiseExtensions {
	if x != nil {
		return x.Extensions
	}
	return nil
}

var File_pb_payload_proto protoreflect.FileDescriptor

var file_pb_payload_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x6f, 0x0a, 0x0f, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x17, 0x77, 0x65, 0x62,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x16, 0x77, 0x65, 0x62, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x65, 0x72, 0x74, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x75, 0x78,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x69, 0x73,
	0x65, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
}

var (
	file_pb_payload_proto_rawDescOnce sync.Once
	file_pb_payload_proto_rawDescData = file_pb_payload_proto_rawDesc
)

func file_pb_payload_proto_rawDescGZIP() []byte {
	file_pb_payload_proto_rawDescOnce.Do(func() {
		file_pb_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_payload_proto_rawDescData)
	})
	return file_pb_payload_proto_rawDescData
}

var file_pb_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_payload_proto_goTypes = []interface{}{
	(*NoiseExtensions)(nil),       // 0: pb.NoiseExtensions
	(*NoiseHandshakePayload)(nil), // 1: pb.NoiseHandshakePayload
}
var file_pb_payload_proto_depIdxs = []int32{
	0, // 0: pb.NoiseHandshakePayload.extensions:type_name -> pb.NoiseExtensions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_payload_proto_init() }
func file_pb_payload_proto_init() {
	if File_pb_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoiseExtensions); i {
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
		file_pb_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoiseHandshakePayload); i {
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
			RawDescriptor: file_pb_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_payload_proto_goTypes,
		DependencyIndexes: file_pb_payload_proto_depIdxs,
		MessageInfos:      file_pb_payload_proto_msgTypes,
	}.Build()
	File_pb_payload_proto = out.File
	file_pb_payload_proto_rawDesc = nil
	file_pb_payload_proto_goTypes = nil
	file_pb_payload_proto_depIdxs = nil
}
