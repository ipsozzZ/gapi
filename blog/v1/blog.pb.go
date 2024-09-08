// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.1
// source: blog/v1/blog.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 保存文章
type ReqSaveSystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int32      `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Cid   int32      `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	CType int32      `protobuf:"varint,3,opt,name=c_type,json=cType,proto3" json:"c_type,omitempty"`
	Intro string     `protobuf:"bytes,4,opt,name=intro,proto3" json:"intro,omitempty"`
	Conf  *SysConfig `protobuf:"bytes,5,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *ReqSaveSystemConfig) Reset() {
	*x = ReqSaveSystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSaveSystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSaveSystemConfig) ProtoMessage() {}

func (x *ReqSaveSystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSaveSystemConfig.ProtoReflect.Descriptor instead.
func (*ReqSaveSystemConfig) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{0}
}

func (x *ReqSaveSystemConfig) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqSaveSystemConfig) GetCid() int32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *ReqSaveSystemConfig) GetCType() int32 {
	if x != nil {
		return x.CType
	}
	return 0
}

func (x *ReqSaveSystemConfig) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *ReqSaveSystemConfig) GetConf() *SysConfig {
	if x != nil {
		return x.Conf
	}
	return nil
}

type RespSaveSystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32 `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
}

func (x *RespSaveSystemConfig) Reset() {
	*x = RespSaveSystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespSaveSystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespSaveSystemConfig) ProtoMessage() {}

func (x *RespSaveSystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespSaveSystemConfig.ProtoReflect.Descriptor instead.
func (*RespSaveSystemConfig) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{1}
}

func (x *RespSaveSystemConfig) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

// 请求配置
type ReqGetSystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CType int32 `protobuf:"varint,3,opt,name=c_type,json=cType,proto3" json:"c_type,omitempty"`
}

func (x *ReqGetSystemConfig) Reset() {
	*x = ReqGetSystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetSystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetSystemConfig) ProtoMessage() {}

func (x *ReqGetSystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetSystemConfig.ProtoReflect.Descriptor instead.
func (*ReqGetSystemConfig) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGetSystemConfig) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqGetSystemConfig) GetCType() int32 {
	if x != nil {
		return x.CType
	}
	return 0
}

type RespGetSystemConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32      `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Cid     int32      `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	CType   int32      `protobuf:"varint,3,opt,name=c_type,json=cType,proto3" json:"c_type,omitempty"`
	Intro   string     `protobuf:"bytes,4,opt,name=intro,proto3" json:"intro,omitempty"`
	Conf    *SysConfig `protobuf:"bytes,5,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *RespGetSystemConfig) Reset() {
	*x = RespGetSystemConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespGetSystemConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespGetSystemConfig) ProtoMessage() {}

func (x *RespGetSystemConfig) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespGetSystemConfig.ProtoReflect.Descriptor instead.
func (*RespGetSystemConfig) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{3}
}

func (x *RespGetSystemConfig) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *RespGetSystemConfig) GetCid() int32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *RespGetSystemConfig) GetCType() int32 {
	if x != nil {
		return x.CType
	}
	return 0
}

func (x *RespGetSystemConfig) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *RespGetSystemConfig) GetConf() *SysConfig {
	if x != nil {
		return x.Conf
	}
	return nil
}

var File_blog_v1_blog_proto protoreflect.FileDescriptor

var file_blog_v1_blog_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x53, 0x61, 0x76, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22,
	0x31, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x70, 0x53, 0x61, 0x76, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x3d, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x32, 0xc6, 0x09, 0x0a, 0x0b,
	0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x12, 0x4e, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x6a, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a,
	0x22, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x68, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x53, 0x61, 0x76, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x53, 0x61, 0x76, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x60, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65,
	0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1c, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x57, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x53, 0x61, 0x76, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x53,
	0x61, 0x76, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x73, 0x61, 0x76, 0x65, 0x12, 0x5f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2f, 0x64, 0x65, 0x6c, 0x12, 0x65, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x7b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5a, 0x0a, 0x0b,
	0x4c, 0x69, 0x6b, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x48, 0x69, 0x74, 0x73,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x69, 0x74, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x1a, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x48,
	0x69, 0x74, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x68, 0x69, 0x74, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_v1_blog_proto_rawDescOnce sync.Once
	file_blog_v1_blog_proto_rawDescData = file_blog_v1_blog_proto_rawDesc
)

func file_blog_v1_blog_proto_rawDescGZIP() []byte {
	file_blog_v1_blog_proto_rawDescOnce.Do(func() {
		file_blog_v1_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_blog_proto_rawDescData)
	})
	return file_blog_v1_blog_proto_rawDescData
}

var file_blog_v1_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blog_v1_blog_proto_goTypes = []interface{}{
	(*ReqSaveSystemConfig)(nil),  // 0: blog.v1.ReqSaveSystemConfig
	(*RespSaveSystemConfig)(nil), // 1: blog.v1.RespSaveSystemConfig
	(*ReqGetSystemConfig)(nil),   // 2: blog.v1.ReqGetSystemConfig
	(*RespGetSystemConfig)(nil),  // 3: blog.v1.RespGetSystemConfig
	(*SysConfig)(nil),            // 4: blog.v1.SysConfig
	(*ReqUserDetail)(nil),        // 5: blog.v1.ReqUserDetail
	(*ReqUserByAccount)(nil),     // 6: blog.v1.ReqUserByAccount
	(*ReqSaveUser)(nil),          // 7: blog.v1.ReqSaveUser
	(*ReqChangeUserState)(nil),   // 8: blog.v1.ReqChangeUserState
	(*ReqListUser)(nil),          // 9: blog.v1.ReqListUser
	(*ReqListArticle)(nil),       // 10: blog.v1.ReqListArticle
	(*ReqSaveArticle)(nil),       // 11: blog.v1.ReqSaveArticle
	(*ReqDeleteArticle)(nil),     // 12: blog.v1.ReqDeleteArticle
	(*ReqArticleDetail)(nil),     // 13: blog.v1.ReqArticleDetail
	(*ReqLikeArticle)(nil),       // 14: blog.v1.ReqLikeArticle
	(*ReqHitsArticle)(nil),       // 15: blog.v1.ReqHitsArticle
	(*RespUserDetail)(nil),       // 16: blog.v1.RespUserDetail
	(*RespSaveUser)(nil),         // 17: blog.v1.RespSaveUser
	(*RespChangeUserState)(nil),  // 18: blog.v1.RespChangeUserState
	(*RespListUser)(nil),         // 19: blog.v1.RespListUser
	(*RespListArticle)(nil),      // 20: blog.v1.RespListArticle
	(*RespSaveArticle)(nil),      // 21: blog.v1.RespSaveArticle
	(*RespDeleteArticle)(nil),    // 22: blog.v1.RespDeleteArticle
	(*RespArticleDetail)(nil),    // 23: blog.v1.RespArticleDetail
	(*RespLikeArticle)(nil),      // 24: blog.v1.RespLikeArticle
	(*RespHitsArticle)(nil),      // 25: blog.v1.RespHitsArticle
}
var file_blog_v1_blog_proto_depIdxs = []int32{
	4,  // 0: blog.v1.ReqSaveSystemConfig.conf:type_name -> blog.v1.SysConfig
	4,  // 1: blog.v1.RespGetSystemConfig.conf:type_name -> blog.v1.SysConfig
	5,  // 2: blog.v1.BlogService.UserDetail:input_type -> blog.v1.ReqUserDetail
	6,  // 3: blog.v1.BlogService.UserByAccount:input_type -> blog.v1.ReqUserByAccount
	7,  // 4: blog.v1.BlogService.SaveUser:input_type -> blog.v1.ReqSaveUser
	8,  // 5: blog.v1.BlogService.ChangeUserState:input_type -> blog.v1.ReqChangeUserState
	9,  // 6: blog.v1.BlogService.ListUser:input_type -> blog.v1.ReqListUser
	0,  // 7: blog.v1.BlogService.SaveSystemConfig:input_type -> blog.v1.ReqSaveSystemConfig
	2,  // 8: blog.v1.BlogService.GetSystemConfig:input_type -> blog.v1.ReqGetSystemConfig
	10, // 9: blog.v1.BlogService.ListArticle:input_type -> blog.v1.ReqListArticle
	11, // 10: blog.v1.BlogService.SaveArticle:input_type -> blog.v1.ReqSaveArticle
	12, // 11: blog.v1.BlogService.DeleteArticle:input_type -> blog.v1.ReqDeleteArticle
	13, // 12: blog.v1.BlogService.ArticleDetail:input_type -> blog.v1.ReqArticleDetail
	14, // 13: blog.v1.BlogService.LikeArticle:input_type -> blog.v1.ReqLikeArticle
	15, // 14: blog.v1.BlogService.HitsArticle:input_type -> blog.v1.ReqHitsArticle
	16, // 15: blog.v1.BlogService.UserDetail:output_type -> blog.v1.RespUserDetail
	16, // 16: blog.v1.BlogService.UserByAccount:output_type -> blog.v1.RespUserDetail
	17, // 17: blog.v1.BlogService.SaveUser:output_type -> blog.v1.RespSaveUser
	18, // 18: blog.v1.BlogService.ChangeUserState:output_type -> blog.v1.RespChangeUserState
	19, // 19: blog.v1.BlogService.ListUser:output_type -> blog.v1.RespListUser
	1,  // 20: blog.v1.BlogService.SaveSystemConfig:output_type -> blog.v1.RespSaveSystemConfig
	3,  // 21: blog.v1.BlogService.GetSystemConfig:output_type -> blog.v1.RespGetSystemConfig
	20, // 22: blog.v1.BlogService.ListArticle:output_type -> blog.v1.RespListArticle
	21, // 23: blog.v1.BlogService.SaveArticle:output_type -> blog.v1.RespSaveArticle
	22, // 24: blog.v1.BlogService.DeleteArticle:output_type -> blog.v1.RespDeleteArticle
	23, // 25: blog.v1.BlogService.ArticleDetail:output_type -> blog.v1.RespArticleDetail
	24, // 26: blog.v1.BlogService.LikeArticle:output_type -> blog.v1.RespLikeArticle
	25, // 27: blog.v1.BlogService.HitsArticle:output_type -> blog.v1.RespHitsArticle
	15, // [15:28] is the sub-list for method output_type
	2,  // [2:15] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_blog_v1_blog_proto_init() }
func file_blog_v1_blog_proto_init() {
	if File_blog_v1_blog_proto != nil {
		return
	}
	file_blog_v1_blog_def_proto_init()
	file_blog_v1_user_proto_init()
	file_blog_v1_article_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blog_v1_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSaveSystemConfig); i {
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
		file_blog_v1_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespSaveSystemConfig); i {
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
		file_blog_v1_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetSystemConfig); i {
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
		file_blog_v1_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespGetSystemConfig); i {
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
			RawDescriptor: file_blog_v1_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_v1_blog_proto_goTypes,
		DependencyIndexes: file_blog_v1_blog_proto_depIdxs,
		MessageInfos:      file_blog_v1_blog_proto_msgTypes,
	}.Build()
	File_blog_v1_blog_proto = out.File
	file_blog_v1_blog_proto_rawDesc = nil
	file_blog_v1_blog_proto_goTypes = nil
	file_blog_v1_blog_proto_depIdxs = nil
}
