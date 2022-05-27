// Copyright © 2022 Durudex
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: durudex/v1/post.proto

package durudexv1

import (
	timestamp "github.com/durudex/dugopb/type/timestamp"
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

// Request for creating a new post.
type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post author uuid.
	AuthorId []byte `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// Post text.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetAuthorId() []byte {
	if x != nil {
		return x.AuthorId
	}
	return nil
}

func (x *CreatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Response for creating a new post.
type CreatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post uuid.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

// Request for getting a post by id.
type GetPostByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post uuid.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostByIdRequest) Reset() {
	*x = GetPostByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIdRequest) ProtoMessage() {}

func (x *GetPostByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIdRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostByIdRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

// Response for getting a post by id.
type GetPostByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post author uuid.
	AuthorId []byte `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// Post text.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Post creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Post update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *GetPostByIdResponse) Reset() {
	*x = GetPostByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIdResponse) ProtoMessage() {}

func (x *GetPostByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIdResponse.ProtoReflect.Descriptor instead.
func (*GetPostByIdResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostByIdResponse) GetAuthorId() []byte {
	if x != nil {
		return x.AuthorId
	}
	return nil
}

func (x *GetPostByIdResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetPostByIdResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetPostByIdResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request for deleting a post.
type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post uuid.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Post author uuid.
	AuthorId []byte `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePostRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DeletePostRequest) GetAuthorId() []byte {
	if x != nil {
		return x.AuthorId
	}
	return nil
}

// Response for deleting a post.
type DeletePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{5}
}

// Request for updating a post.
type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Post uuid.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Post author uuid.
	AuthorId []byte `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// Post text.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePostRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdatePostRequest) GetAuthorId() []byte {
	if x != nil {
		return x.AuthorId
	}
	return nil
}

func (x *UpdatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Response for updating a post.
type UpdatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePostResponse) Reset() {
	*x = UpdatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResponse) ProtoMessage() {}

func (x *UpdatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_post_proto_rawDescGZIP(), []int{7}
}

var File_durudex_v1_post_proto protoreflect.FileDescriptor

var file_durudex_v1_post_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x36, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64,
	0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0x40, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc4, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x1d, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1d, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xac, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f,
	0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x44, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x44,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_durudex_v1_post_proto_rawDescOnce sync.Once
	file_durudex_v1_post_proto_rawDescData = file_durudex_v1_post_proto_rawDesc
)

func file_durudex_v1_post_proto_rawDescGZIP() []byte {
	file_durudex_v1_post_proto_rawDescOnce.Do(func() {
		file_durudex_v1_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_durudex_v1_post_proto_rawDescData)
	})
	return file_durudex_v1_post_proto_rawDescData
}

var file_durudex_v1_post_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_durudex_v1_post_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil),   // 0: durudex.v1.CreatePostRequest
	(*CreatePostResponse)(nil),  // 1: durudex.v1.CreatePostResponse
	(*GetPostByIdRequest)(nil),  // 2: durudex.v1.GetPostByIdRequest
	(*GetPostByIdResponse)(nil), // 3: durudex.v1.GetPostByIdResponse
	(*DeletePostRequest)(nil),   // 4: durudex.v1.DeletePostRequest
	(*DeletePostResponse)(nil),  // 5: durudex.v1.DeletePostResponse
	(*UpdatePostRequest)(nil),   // 6: durudex.v1.UpdatePostRequest
	(*UpdatePostResponse)(nil),  // 7: durudex.v1.UpdatePostResponse
	(*timestamp.Timestamp)(nil), // 8: durudex.type.Timestamp
}
var file_durudex_v1_post_proto_depIdxs = []int32{
	8, // 0: durudex.v1.GetPostByIdResponse.created_at:type_name -> durudex.type.Timestamp
	8, // 1: durudex.v1.GetPostByIdResponse.updated_at:type_name -> durudex.type.Timestamp
	0, // 2: durudex.v1.PostService.CreatePost:input_type -> durudex.v1.CreatePostRequest
	2, // 3: durudex.v1.PostService.GetPostById:input_type -> durudex.v1.GetPostByIdRequest
	4, // 4: durudex.v1.PostService.DeletePost:input_type -> durudex.v1.DeletePostRequest
	6, // 5: durudex.v1.PostService.UpdatePost:input_type -> durudex.v1.UpdatePostRequest
	1, // 6: durudex.v1.PostService.CreatePost:output_type -> durudex.v1.CreatePostResponse
	3, // 7: durudex.v1.PostService.GetPostById:output_type -> durudex.v1.GetPostByIdResponse
	5, // 8: durudex.v1.PostService.DeletePost:output_type -> durudex.v1.DeletePostResponse
	7, // 9: durudex.v1.PostService.UpdatePost:output_type -> durudex.v1.UpdatePostResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_durudex_v1_post_proto_init() }
func file_durudex_v1_post_proto_init() {
	if File_durudex_v1_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_durudex_v1_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_durudex_v1_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostResponse); i {
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
		file_durudex_v1_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIdRequest); i {
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
		file_durudex_v1_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIdResponse); i {
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
		file_durudex_v1_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_durudex_v1_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResponse); i {
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
		file_durudex_v1_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_durudex_v1_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostResponse); i {
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
	file_durudex_v1_post_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_durudex_v1_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_durudex_v1_post_proto_goTypes,
		DependencyIndexes: file_durudex_v1_post_proto_depIdxs,
		MessageInfos:      file_durudex_v1_post_proto_msgTypes,
	}.Build()
	File_durudex_v1_post_proto = out.File
	file_durudex_v1_post_proto_rawDesc = nil
	file_durudex_v1_post_proto_goTypes = nil
	file_durudex_v1_post_proto_depIdxs = nil
}
