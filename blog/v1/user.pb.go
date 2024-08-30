// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.1
// source: blog/v1/user.proto

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

// 保存玩家
type ReqSaveUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User *CUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ReqSaveUser) Reset() {
	*x = ReqSaveUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSaveUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSaveUser) ProtoMessage() {}

func (x *ReqSaveUser) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSaveUser.ProtoReflect.Descriptor instead.
func (*ReqSaveUser) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *ReqSaveUser) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqSaveUser) GetUser() *CUser {
	if x != nil {
		return x.User
	}
	return nil
}

type RespSaveUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32 `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
}

func (x *RespSaveUser) Reset() {
	*x = RespSaveUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespSaveUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespSaveUser) ProtoMessage() {}

func (x *RespSaveUser) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespSaveUser.ProtoReflect.Descriptor instead.
func (*RespSaveUser) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *RespSaveUser) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

// 修改玩家状态
type ReqChangeUserState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TargetUid int32 `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	State     int32 `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ReqChangeUserState) Reset() {
	*x = ReqChangeUserState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqChangeUserState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqChangeUserState) ProtoMessage() {}

func (x *ReqChangeUserState) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqChangeUserState.ProtoReflect.Descriptor instead.
func (*ReqChangeUserState) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *ReqChangeUserState) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqChangeUserState) GetTargetUid() int32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *ReqChangeUserState) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type RespChangeUserState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32 `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
}

func (x *RespChangeUserState) Reset() {
	*x = RespChangeUserState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespChangeUserState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespChangeUserState) ProtoMessage() {}

func (x *RespChangeUserState) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespChangeUserState.ProtoReflect.Descriptor instead.
func (*RespChangeUserState) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *RespChangeUserState) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

// 请求玩家详细数据
type ReqUserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TargetUid int32 `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *ReqUserDetail) Reset() {
	*x = ReqUserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserDetail) ProtoMessage() {}

func (x *ReqUserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserDetail.ProtoReflect.Descriptor instead.
func (*ReqUserDetail) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *ReqUserDetail) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqUserDetail) GetTargetUid() int32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

type RespUserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Detail  *CUser `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *RespUserDetail) Reset() {
	*x = RespUserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespUserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespUserDetail) ProtoMessage() {}

func (x *RespUserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespUserDetail.ProtoReflect.Descriptor instead.
func (*RespUserDetail) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *RespUserDetail) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *RespUserDetail) GetDetail() *CUser {
	if x != nil {
		return x.Detail
	}
	return nil
}

// 请求玩家
type ReqListUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ReqListUser) Reset() {
	*x = ReqListUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqListUser) ProtoMessage() {}

func (x *ReqListUser) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqListUser.ProtoReflect.Descriptor instead.
func (*ReqListUser) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *ReqListUser) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReqListUser) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type RespListUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32    `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Users   []*CUser `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *RespListUser) Reset() {
	*x = RespListUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespListUser) ProtoMessage() {}

func (x *RespListUser) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespListUser.ProtoReflect.Descriptor instead.
func (*RespListUser) Descriptor() ([]byte, []int) {
	return file_blog_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *RespListUser) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *RespListUser) GetUsers() []*CUser {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_blog_v1_user_proto protoreflect.FileDescriptor

var file_blog_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x62,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x53, 0x61, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x5b, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x33, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x4f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_v1_user_proto_rawDescOnce sync.Once
	file_blog_v1_user_proto_rawDescData = file_blog_v1_user_proto_rawDesc
)

func file_blog_v1_user_proto_rawDescGZIP() []byte {
	file_blog_v1_user_proto_rawDescOnce.Do(func() {
		file_blog_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_user_proto_rawDescData)
	})
	return file_blog_v1_user_proto_rawDescData
}

var file_blog_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blog_v1_user_proto_goTypes = []interface{}{
	(*ReqSaveUser)(nil),         // 0: blog.v1.ReqSaveUser
	(*RespSaveUser)(nil),        // 1: blog.v1.RespSaveUser
	(*ReqChangeUserState)(nil),  // 2: blog.v1.ReqChangeUserState
	(*RespChangeUserState)(nil), // 3: blog.v1.RespChangeUserState
	(*ReqUserDetail)(nil),       // 4: blog.v1.ReqUserDetail
	(*RespUserDetail)(nil),      // 5: blog.v1.RespUserDetail
	(*ReqListUser)(nil),         // 6: blog.v1.ReqListUser
	(*RespListUser)(nil),        // 7: blog.v1.RespListUser
	(*CUser)(nil),               // 8: blog.v1.CUser
}
var file_blog_v1_user_proto_depIdxs = []int32{
	8, // 0: blog.v1.ReqSaveUser.user:type_name -> blog.v1.CUser
	8, // 1: blog.v1.RespUserDetail.detail:type_name -> blog.v1.CUser
	8, // 2: blog.v1.RespListUser.users:type_name -> blog.v1.CUser
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blog_v1_user_proto_init() }
func file_blog_v1_user_proto_init() {
	if File_blog_v1_user_proto != nil {
		return
	}
	file_blog_v1_blog_def_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blog_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSaveUser); i {
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
		file_blog_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespSaveUser); i {
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
		file_blog_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqChangeUserState); i {
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
		file_blog_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespChangeUserState); i {
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
		file_blog_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserDetail); i {
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
		file_blog_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespUserDetail); i {
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
		file_blog_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqListUser); i {
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
		file_blog_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespListUser); i {
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
			RawDescriptor: file_blog_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blog_v1_user_proto_goTypes,
		DependencyIndexes: file_blog_v1_user_proto_depIdxs,
		MessageInfos:      file_blog_v1_user_proto_msgTypes,
	}.Build()
	File_blog_v1_user_proto = out.File
	file_blog_v1_user_proto_rawDesc = nil
	file_blog_v1_user_proto_goTypes = nil
	file_blog_v1_user_proto_depIdxs = nil
}
