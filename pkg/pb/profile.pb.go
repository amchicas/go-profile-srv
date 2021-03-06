// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: profile.proto

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

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lastname    string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Votes       uint64 `protobuf:"varint,6,opt,name=votes,proto3" json:"votes,omitempty"`
	Students    uint64 `protobuf:"varint,7,opt,name=students,proto3" json:"students,omitempty"`
	Website     string `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Youtube     string `protobuf:"bytes,9,opt,name=youtube,proto3" json:"youtube,omitempty"`
	Linkedin    string `protobuf:"bytes,10,opt,name=linkedin,proto3" json:"linkedin,omitempty"`
	Twitter     string `protobuf:"bytes,11,opt,name=twitter,proto3" json:"twitter,omitempty"`
	Facebook    string `protobuf:"bytes,12,opt,name=facebook,proto3" json:"facebook,omitempty"`
	Created     int64  `protobuf:"varint,13,opt,name=created,proto3" json:"created,omitempty"`
	Modified    int64  `protobuf:"varint,14,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Profile) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Profile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Profile) GetVotes() uint64 {
	if x != nil {
		return x.Votes
	}
	return 0
}

func (x *Profile) GetStudents() uint64 {
	if x != nil {
		return x.Students
	}
	return 0
}

func (x *Profile) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Profile) GetYoutube() string {
	if x != nil {
		return x.Youtube
	}
	return ""
}

func (x *Profile) GetLinkedin() string {
	if x != nil {
		return x.Linkedin
	}
	return ""
}

func (x *Profile) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *Profile) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *Profile) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Profile) GetModified() int64 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

func (x *AddReq) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Profile *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpdateReq) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Msg    string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{5}
}

func (x *Resp) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Resp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type FindProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *FindProfileReq) Reset() {
	*x = FindProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProfileReq) ProtoMessage() {}

func (x *FindProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProfileReq.ProtoReflect.Descriptor instead.
func (*FindProfileReq) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{6}
}

func (x *FindProfileReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type FindProfileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Profile *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *FindProfileResp) Reset() {
	*x = FindProfileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProfileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProfileResp) ProtoMessage() {}

func (x *FindProfileResp) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProfileResp.ProtoReflect.Descriptor instead.
func (*FindProfileResp) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{7}
}

func (x *FindProfileResp) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindProfileResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindProfileResp) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x22, 0xef, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22,
	0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32,
	0xe2, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_profile_proto_goTypes = []interface{}{
	(*EmptyReq)(nil),        // 0: profile.EmptyReq
	(*Profile)(nil),         // 1: profile.Profile
	(*AddReq)(nil),          // 2: profile.AddReq
	(*UpdateReq)(nil),       // 3: profile.UpdateReq
	(*DeleteReq)(nil),       // 4: profile.DeleteReq
	(*Resp)(nil),            // 5: profile.Resp
	(*FindProfileReq)(nil),  // 6: profile.FindProfileReq
	(*FindProfileResp)(nil), // 7: profile.FindProfileResp
}
var file_profile_proto_depIdxs = []int32{
	1, // 0: profile.AddReq.profile:type_name -> profile.Profile
	1, // 1: profile.UpdateReq.profile:type_name -> profile.Profile
	1, // 2: profile.FindProfileResp.profile:type_name -> profile.Profile
	2, // 3: profile.ProfileService.AddProfile:input_type -> profile.AddReq
	3, // 4: profile.ProfileService.Update:input_type -> profile.UpdateReq
	4, // 5: profile.ProfileService.Remove:input_type -> profile.DeleteReq
	6, // 6: profile.ProfileService.FindProfile:input_type -> profile.FindProfileReq
	5, // 7: profile.ProfileService.AddProfile:output_type -> profile.Resp
	5, // 8: profile.ProfileService.Update:output_type -> profile.Resp
	5, // 9: profile.ProfileService.Remove:output_type -> profile.Resp
	7, // 10: profile.ProfileService.FindProfile:output_type -> profile.FindProfileResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
		file_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProfileReq); i {
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
		file_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProfileResp); i {
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
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
