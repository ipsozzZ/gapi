// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: fuxi/v1/fuxi.proto

package v1

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

// 获取首页
type IndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexReq) Reset() {
	*x = IndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fuxi_v1_fuxi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReq) ProtoMessage() {}

func (x *IndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_fuxi_v1_fuxi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReq.ProtoReflect.Descriptor instead.
func (*IndexReq) Descriptor() ([]byte, []int) {
	return file_fuxi_v1_fuxi_proto_rawDescGZIP(), []int{0}
}

// 首页内容
type IndexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IndexResp) Reset() {
	*x = IndexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fuxi_v1_fuxi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexResp) ProtoMessage() {}

func (x *IndexResp) ProtoReflect() protoreflect.Message {
	mi := &file_fuxi_v1_fuxi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexResp.ProtoReflect.Descriptor instead.
func (*IndexResp) Descriptor() ([]byte, []int) {
	return file_fuxi_v1_fuxi_proto_rawDescGZIP(), []int{1}
}

func (x *IndexResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_fuxi_v1_fuxi_proto protoreflect.FileDescriptor

var file_fuxi_v1_fuxi_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x75, 0x78, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x78, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x78, 0x69, 0x2e, 0x76,
	0x31, 0x22, 0x0a, 0x0a, 0x08, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x22, 0x25, 0x0a,
	0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x43, 0x0a, 0x04, 0x46, 0x75, 0x78, 0x69, 0x12, 0x3b, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x75, 0x78, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x78, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x75, 0x78, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_fuxi_v1_fuxi_proto_rawDescOnce sync.Once
	file_fuxi_v1_fuxi_proto_rawDescData = file_fuxi_v1_fuxi_proto_rawDesc
)

func file_fuxi_v1_fuxi_proto_rawDescGZIP() []byte {
	file_fuxi_v1_fuxi_proto_rawDescOnce.Do(func() {
		file_fuxi_v1_fuxi_proto_rawDescData = protoimpl.X.CompressGZIP(file_fuxi_v1_fuxi_proto_rawDescData)
	})
	return file_fuxi_v1_fuxi_proto_rawDescData
}

var file_fuxi_v1_fuxi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fuxi_v1_fuxi_proto_goTypes = []interface{}{
	(*IndexReq)(nil),  // 0: api.fuxi.v1.IndexReq
	(*IndexResp)(nil), // 1: api.fuxi.v1.IndexResp
}
var file_fuxi_v1_fuxi_proto_depIdxs = []int32{
	0, // 0: api.fuxi.v1.Fuxi.GetIndex:input_type -> api.fuxi.v1.IndexReq
	1, // 1: api.fuxi.v1.Fuxi.GetIndex:output_type -> api.fuxi.v1.IndexResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fuxi_v1_fuxi_proto_init() }
func file_fuxi_v1_fuxi_proto_init() {
	if File_fuxi_v1_fuxi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fuxi_v1_fuxi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexReq); i {
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
		file_fuxi_v1_fuxi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexResp); i {
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
			RawDescriptor: file_fuxi_v1_fuxi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fuxi_v1_fuxi_proto_goTypes,
		DependencyIndexes: file_fuxi_v1_fuxi_proto_depIdxs,
		MessageInfos:      file_fuxi_v1_fuxi_proto_msgTypes,
	}.Build()
	File_fuxi_v1_fuxi_proto = out.File
	file_fuxi_v1_fuxi_proto_rawDesc = nil
	file_fuxi_v1_fuxi_proto_goTypes = nil
	file_fuxi_v1_fuxi_proto_depIdxs = nil
}
