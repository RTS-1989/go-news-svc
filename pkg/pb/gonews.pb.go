// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/gonews.proto

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

type PostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsCountGet int64 `protobuf:"varint,1,opt,name=newsCountGet,proto3" json:"newsCountGet,omitempty"`
}

func (x *PostsRequest) Reset() {
	*x = PostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsRequest) ProtoMessage() {}

func (x *PostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsRequest.ProtoReflect.Descriptor instead.
func (*PostsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{0}
}

func (x *PostsRequest) GetNewsCountGet() int64 {
	if x != nil {
		return x.NewsCountGet
	}
	return 0
}

type OneNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId int64 `protobuf:"varint,1,opt,name=newsId,proto3" json:"newsId,omitempty"`
}

func (x *OneNewsRequest) Reset() {
	*x = OneNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneNewsRequest) ProtoMessage() {}

func (x *OneNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneNewsRequest.ProtoReflect.Descriptor instead.
func (*OneNewsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{1}
}

func (x *OneNewsRequest) GetNewsId() int64 {
	if x != nil {
		return x.NewsId
	}
	return 0
}

type FilterNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterValue string `protobuf:"bytes,2,opt,name=filterValue,proto3" json:"filterValue,omitempty"`
}

func (x *FilterNewsRequest) Reset() {
	*x = FilterNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterNewsRequest) ProtoMessage() {}

func (x *FilterNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterNewsRequest.ProtoReflect.Descriptor instead.
func (*FilterNewsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{2}
}

func (x *FilterNewsRequest) GetFilterValue() string {
	if x != nil {
		return x.FilterValue
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PubTime       int64  `protobuf:"varint,4,opt,name=pubTime,proto3" json:"pubTime,omitempty"`
	Link          string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	SourceXmlLink string `protobuf:"bytes,6,opt,name=sourceXmlLink,proto3" json:"sourceXmlLink,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{3}
}

func (x *Post) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetPubTime() int64 {
	if x != nil {
		return x.PubTime
	}
	return 0
}

func (x *Post) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Post) GetSourceXmlLink() string {
	if x != nil {
		return x.SourceXmlLink
	}
	return ""
}

type OnePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Posts  *Post  `protobuf:"bytes,3,opt,name=posts,proto3" json:"posts,omitempty"`
}

func (x *OnePostResponse) Reset() {
	*x = OnePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnePostResponse) ProtoMessage() {}

func (x *OnePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnePostResponse.ProtoReflect.Descriptor instead.
func (*OnePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{4}
}

func (x *OnePostResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OnePostResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *OnePostResponse) GetPosts() *Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type PostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Posts  []*Post `protobuf:"bytes,3,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostsResponse) Reset() {
	*x = PostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_gonews_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse) ProtoMessage() {}

func (x *PostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_gonews_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsResponse.ProtoReflect.Descriptor instead.
func (*PostsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_gonews_proto_rawDescGZIP(), []int{5}
}

func (x *PostsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PostsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_pkg_pb_gonews_proto protoreflect.FileDescriptor

var file_pkg_pb_gonews_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x22, 0x32, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x22, 0x28, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x11, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x75,
	0x62, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x58, 0x6d, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x58, 0x6d, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x22,
	0x63, 0x0a, 0x0f, 0x4f, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0x98, 0x02, 0x0a, 0x0d, 0x47, 0x6f, 0x4e, 0x65,
	0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f,
	0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x73,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f,
	0x6e, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x19,
	0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_gonews_proto_rawDescOnce sync.Once
	file_pkg_pb_gonews_proto_rawDescData = file_pkg_pb_gonews_proto_rawDesc
)

func file_pkg_pb_gonews_proto_rawDescGZIP() []byte {
	file_pkg_pb_gonews_proto_rawDescOnce.Do(func() {
		file_pkg_pb_gonews_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_gonews_proto_rawDescData)
	})
	return file_pkg_pb_gonews_proto_rawDescData
}

var file_pkg_pb_gonews_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_pb_gonews_proto_goTypes = []interface{}{
	(*PostsRequest)(nil),      // 0: gonews.PostsRequest
	(*OneNewsRequest)(nil),    // 1: gonews.OneNewsRequest
	(*FilterNewsRequest)(nil), // 2: gonews.FilterNewsRequest
	(*Post)(nil),              // 3: gonews.Post
	(*OnePostResponse)(nil),   // 4: gonews.OnePostResponse
	(*PostsResponse)(nil),     // 5: gonews.PostsResponse
}
var file_pkg_pb_gonews_proto_depIdxs = []int32{
	3, // 0: gonews.OnePostResponse.posts:type_name -> gonews.Post
	3, // 1: gonews.PostsResponse.posts:type_name -> gonews.Post
	0, // 2: gonews.GoNewsService.Posts:input_type -> gonews.PostsRequest
	1, // 3: gonews.GoNewsService.NewsFullDetailed:input_type -> gonews.OneNewsRequest
	1, // 4: gonews.GoNewsService.NewsShortDetailed:input_type -> gonews.OneNewsRequest
	2, // 5: gonews.GoNewsService.FilterNews:input_type -> gonews.FilterNewsRequest
	5, // 6: gonews.GoNewsService.Posts:output_type -> gonews.PostsResponse
	4, // 7: gonews.GoNewsService.NewsFullDetailed:output_type -> gonews.OnePostResponse
	4, // 8: gonews.GoNewsService.NewsShortDetailed:output_type -> gonews.OnePostResponse
	5, // 9: gonews.GoNewsService.FilterNews:output_type -> gonews.PostsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_pb_gonews_proto_init() }
func file_pkg_pb_gonews_proto_init() {
	if File_pkg_pb_gonews_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_gonews_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsRequest); i {
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
		file_pkg_pb_gonews_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneNewsRequest); i {
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
		file_pkg_pb_gonews_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterNewsRequest); i {
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
		file_pkg_pb_gonews_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_pkg_pb_gonews_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnePostResponse); i {
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
		file_pkg_pb_gonews_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsResponse); i {
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
			RawDescriptor: file_pkg_pb_gonews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_gonews_proto_goTypes,
		DependencyIndexes: file_pkg_pb_gonews_proto_depIdxs,
		MessageInfos:      file_pkg_pb_gonews_proto_msgTypes,
	}.Build()
	File_pkg_pb_gonews_proto = out.File
	file_pkg_pb_gonews_proto_rawDesc = nil
	file_pkg_pb_gonews_proto_goTypes = nil
	file_pkg_pb_gonews_proto_depIdxs = nil
}