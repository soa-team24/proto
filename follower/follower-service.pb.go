// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: follower-service.proto

package follower

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

type AddFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follow *Profile `protobuf:"bytes,1,opt,name=follow,proto3" json:"follow,omitempty"`
}

func (x *AddFollowRequest) Reset() {
	*x = AddFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFollowRequest) ProtoMessage() {}

func (x *AddFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFollowRequest.ProtoReflect.Descriptor instead.
func (*AddFollowRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddFollowRequest) GetFollow() *Profile {
	if x != nil {
		return x.Follow
	}
	return nil
}

type AddFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddFollowResponse) Reset() {
	*x = AddFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFollowResponse) ProtoMessage() {}

func (x *AddFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFollowResponse.ProtoReflect.Descriptor instead.
func (*AddFollowResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{1}
}

type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowersRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*Profile `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowersResponse) GetFollowers() []*Profile {
	if x != nil {
		return x.Followers
	}
	return nil
}

type GetFollowersOfMyFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFollowersOfMyFollowersRequest) Reset() {
	*x = GetFollowersOfMyFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersOfMyFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersOfMyFollowersRequest) ProtoMessage() {}

func (x *GetFollowersOfMyFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersOfMyFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersOfMyFollowersRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetFollowersOfMyFollowersRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllProfilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllProfilesRequest) Reset() {
	*x = GetAllProfilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProfilesRequest) ProtoMessage() {}

func (x *GetAllProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProfilesRequest.ProtoReflect.Descriptor instead.
func (*GetAllProfilesRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{5}
}

type GetAllProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *GetAllProfilesResponse) Reset() {
	*x = GetAllProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProfilesResponse) ProtoMessage() {}

func (x *GetAllProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProfilesResponse.ProtoReflect.Descriptor instead.
func (*GetAllProfilesResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllProfilesResponse) GetProfiles() []*Profile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type CheckIfUserFollowsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId uint32 `protobuf:"varint,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckIfUserFollowsRequest) Reset() {
	*x = CheckIfUserFollowsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfUserFollowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfUserFollowsRequest) ProtoMessage() {}

func (x *CheckIfUserFollowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfUserFollowsRequest.ProtoReflect.Descriptor instead.
func (*CheckIfUserFollowsRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{7}
}

func (x *CheckIfUserFollowsRequest) GetFollowerId() uint32 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

func (x *CheckIfUserFollowsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckIfUserFollowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFollowing bool `protobuf:"varint,1,opt,name=is_following,json=isFollowing,proto3" json:"is_following,omitempty"`
}

func (x *CheckIfUserFollowsResponse) Reset() {
	*x = CheckIfUserFollowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfUserFollowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfUserFollowsResponse) ProtoMessage() {}

func (x *CheckIfUserFollowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfUserFollowsResponse.ProtoReflect.Descriptor instead.
func (*CheckIfUserFollowsResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{8}
}

func (x *CheckIfUserFollowsResponse) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName      string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	ProfilePicture string     `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	UserId         int64      `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Followers      []*Profile `protobuf:"bytes,6,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[9]
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
	return file_follower_service_proto_rawDescGZIP(), []int{9}
}

func (x *Profile) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Profile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Profile) GetProfilePicture() string {
	if x != nil {
		return x.ProfilePicture
	}
	return ""
}

func (x *Profile) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Profile) GetFollowers() []*Profile {
	if x != nil {
		return x.Followers
	}
	return nil
}

type Follow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId  uint32 `protobuf:"varint,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	FollowerId uint32 `protobuf:"varint,2,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
}

func (x *Follow) Reset() {
	*x = Follow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follow) ProtoMessage() {}

func (x *Follow) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follow.ProtoReflect.Descriptor instead.
func (*Follow) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{10}
}

func (x *Follow) GetProfileId() uint32 {
	if x != nil {
		return x.ProfileId
	}
	return 0
}

func (x *Follow) GetFollowerId() uint32 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

var File_follower_service_proto protoreflect.FileDescriptor

var file_follower_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x13, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x22, 0x3b, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x4f, 0x66, 0x4d, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22,
	0xbf, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x22, 0x48, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x32, 0x98, 0x04, 0x0a, 0x0d,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x66, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x12, 0x82, 0x01,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x4f, 0x66, 0x4d, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x21,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x4f, 0x66, 0x4d,
	0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x12, 0x20, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x7d, 0x12, 0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x7c, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1a,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x66, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12,
	0x25, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73,
	0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follower_service_proto_rawDescOnce sync.Once
	file_follower_service_proto_rawDescData = file_follower_service_proto_rawDesc
)

func file_follower_service_proto_rawDescGZIP() []byte {
	file_follower_service_proto_rawDescOnce.Do(func() {
		file_follower_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_service_proto_rawDescData)
	})
	return file_follower_service_proto_rawDescData
}

var file_follower_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_follower_service_proto_goTypes = []interface{}{
	(*AddFollowRequest)(nil),                 // 0: AddFollowRequest
	(*AddFollowResponse)(nil),                // 1: AddFollowResponse
	(*GetFollowersRequest)(nil),              // 2: GetFollowersRequest
	(*GetFollowersResponse)(nil),             // 3: GetFollowersResponse
	(*GetFollowersOfMyFollowersRequest)(nil), // 4: GetFollowersOfMyFollowersRequest
	(*GetAllProfilesRequest)(nil),            // 5: GetAllProfilesRequest
	(*GetAllProfilesResponse)(nil),           // 6: GetAllProfilesResponse
	(*CheckIfUserFollowsRequest)(nil),        // 7: CheckIfUserFollowsRequest
	(*CheckIfUserFollowsResponse)(nil),       // 8: CheckIfUserFollowsResponse
	(*Profile)(nil),                          // 9: Profile
	(*Follow)(nil),                           // 10: Follow
}
var file_follower_service_proto_depIdxs = []int32{
	9, // 0: AddFollowRequest.follow:type_name -> Profile
	9, // 1: GetFollowersResponse.followers:type_name -> Profile
	9, // 2: GetAllProfilesResponse.profiles:type_name -> Profile
	9, // 3: Profile.followers:type_name -> Profile
	0, // 4: FollowService.AddFollow:input_type -> AddFollowRequest
	2, // 5: FollowService.GetAllFollowersForUser:input_type -> GetFollowersRequest
	4, // 6: FollowService.GetAllFollowersOfMyFollowers:input_type -> GetFollowersOfMyFollowersRequest
	5, // 7: FollowService.GetAllProfiles:input_type -> GetAllProfilesRequest
	7, // 8: FollowService.CheckIfUserFollows:input_type -> CheckIfUserFollowsRequest
	1, // 9: FollowService.AddFollow:output_type -> AddFollowResponse
	3, // 10: FollowService.GetAllFollowersForUser:output_type -> GetFollowersResponse
	3, // 11: FollowService.GetAllFollowersOfMyFollowers:output_type -> GetFollowersResponse
	6, // 12: FollowService.GetAllProfiles:output_type -> GetAllProfilesResponse
	8, // 13: FollowService.CheckIfUserFollows:output_type -> CheckIfUserFollowsResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_follower_service_proto_init() }
func file_follower_service_proto_init() {
	if File_follower_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follower_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFollowRequest); i {
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
		file_follower_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFollowResponse); i {
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
		file_follower_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersRequest); i {
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
		file_follower_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersResponse); i {
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
		file_follower_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersOfMyFollowersRequest); i {
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
		file_follower_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProfilesRequest); i {
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
		file_follower_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProfilesResponse); i {
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
		file_follower_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfUserFollowsRequest); i {
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
		file_follower_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfUserFollowsResponse); i {
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
		file_follower_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_follower_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follow); i {
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
			RawDescriptor: file_follower_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_service_proto_goTypes,
		DependencyIndexes: file_follower_service_proto_depIdxs,
		MessageInfos:      file_follower_service_proto_msgTypes,
	}.Build()
	File_follower_service_proto = out.File
	file_follower_service_proto_rawDesc = nil
	file_follower_service_proto_goTypes = nil
	file_follower_service_proto_depIdxs = nil
}