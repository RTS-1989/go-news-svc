// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gonews/gonews.proto

package gonews

import (
	comment "github.com/RTS-1989/go-news-svc/pkg/pb/comment"
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
		mi := &file_gonews_gonews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsRequest) ProtoMessage() {}

func (x *PostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[0]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{0}
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
		mi := &file_gonews_gonews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneNewsRequest) ProtoMessage() {}

func (x *OneNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[1]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{1}
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

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilterValue string `protobuf:"bytes,2,opt,name=filterValue,proto3" json:"filterValue,omitempty"`
	PageSize    int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page        int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FilterNewsRequest) Reset() {
	*x = FilterNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gonews_gonews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterNewsRequest) ProtoMessage() {}

func (x *FilterNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[2]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{2}
}

func (x *FilterNewsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FilterNewsRequest) GetFilterValue() string {
	if x != nil {
		return x.FilterValue
	}
	return ""
}

func (x *FilterNewsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FilterNewsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
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
		mi := &file_gonews_gonews_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[3]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{3}
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
		mi := &file_gonews_gonews_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnePostResponse) ProtoMessage() {}

func (x *OnePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[4]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{4}
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
		mi := &file_gonews_gonews_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse) ProtoMessage() {}

func (x *PostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[5]
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
	return file_gonews_gonews_proto_rawDescGZIP(), []int{5}
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

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pages       int32 `protobuf:"varint,1,opt,name=pages,proto3" json:"pages,omitempty"`
	CurrentPage int32 `protobuf:"varint,2,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	PostsOnPage int32 `protobuf:"varint,3,opt,name=postsOnPage,proto3" json:"postsOnPage,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gonews_gonews_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_gonews_gonews_proto_rawDescGZIP(), []int{6}
}

func (x *Pagination) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetPostsOnPage() int32 {
	if x != nil {
		return x.PostsOnPage
	}
	return 0
}

type ListPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsCountGet int64 `protobuf:"varint,1,opt,name=newsCountGet,proto3" json:"newsCountGet,omitempty"`
	UserId       int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageSize     int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page         int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gonews_gonews_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_gonews_gonews_proto_rawDescGZIP(), []int{7}
}

func (x *ListPostsRequest) GetNewsCountGet() int64 {
	if x != nil {
		return x.NewsCountGet
	}
	return 0
}

func (x *ListPostsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPostsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         int64       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error          string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PaginationInfo *Pagination `protobuf:"bytes,3,opt,name=paginationInfo,proto3" json:"paginationInfo,omitempty"`
	Posts          []*Post     `protobuf:"bytes,4,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gonews_gonews_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_gonews_gonews_proto_rawDescGZIP(), []int{8}
}

func (x *ListPostsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListPostsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ListPostsResponse) GetPaginationInfo() *Pagination {
	if x != nil {
		return x.PaginationInfo
	}
	return nil
}

func (x *ListPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type DetailedNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64              `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Post     *Post              `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Comments []*comment.Comment `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *DetailedNewsResponse) Reset() {
	*x = DetailedNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gonews_gonews_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailedNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedNewsResponse) ProtoMessage() {}

func (x *DetailedNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gonews_gonews_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedNewsResponse.ProtoReflect.Descriptor instead.
func (*DetailedNewsResponse) Descriptor() ([]byte, []int) {
	return file_gonews_gonews_proto_rawDescGZIP(), []int{9}
}

func (x *DetailedNewsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DetailedNewsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DetailedNewsResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *DetailedNewsResponse) GetComments() []*comment.Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

var File_gonews_gonews_proto protoreflect.FileDescriptor

var file_gonews_gonews_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x1a, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65,
	0x77, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73,
	0x49, 0x64, 0x22, 0x7f, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x75, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70,
	0x75, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x58, 0x6d, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x58, 0x6d, 0x6c, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x63, 0x0a, 0x0f, 0x4f, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x4f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x4f, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x47, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x65, 0x77,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xe4,
	0x02, 0x0a, 0x0d, 0x47, 0x6f, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x73,
	0x46, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x19, 0x2e, 0x67, 0x6f, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x18,
	0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x54, 0x53, 0x2d, 0x31, 0x39, 0x38, 0x39, 0x2f, 0x67, 0x6f, 0x2d,
	0x6e, 0x65, 0x77, 0x73, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x6b, 0x62, 0x2f, 0x70, 0x62, 0x2f,
	0x67, 0x6f, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gonews_gonews_proto_rawDescOnce sync.Once
	file_gonews_gonews_proto_rawDescData = file_gonews_gonews_proto_rawDesc
)

func file_gonews_gonews_proto_rawDescGZIP() []byte {
	file_gonews_gonews_proto_rawDescOnce.Do(func() {
		file_gonews_gonews_proto_rawDescData = protoimpl.X.CompressGZIP(file_gonews_gonews_proto_rawDescData)
	})
	return file_gonews_gonews_proto_rawDescData
}

var file_gonews_gonews_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gonews_gonews_proto_goTypes = []interface{}{
	(*PostsRequest)(nil),         // 0: gonews.PostsRequest
	(*OneNewsRequest)(nil),       // 1: gonews.OneNewsRequest
	(*FilterNewsRequest)(nil),    // 2: gonews.FilterNewsRequest
	(*Post)(nil),                 // 3: gonews.Post
	(*OnePostResponse)(nil),      // 4: gonews.OnePostResponse
	(*PostsResponse)(nil),        // 5: gonews.PostsResponse
	(*Pagination)(nil),           // 6: gonews.Pagination
	(*ListPostsRequest)(nil),     // 7: gonews.ListPostsRequest
	(*ListPostsResponse)(nil),    // 8: gonews.ListPostsResponse
	(*DetailedNewsResponse)(nil), // 9: gonews.DetailedNewsResponse
	(*comment.Comment)(nil),      // 10: comment.Comment
}
var file_gonews_gonews_proto_depIdxs = []int32{
	3,  // 0: gonews.OnePostResponse.posts:type_name -> gonews.Post
	3,  // 1: gonews.PostsResponse.posts:type_name -> gonews.Post
	6,  // 2: gonews.ListPostsResponse.paginationInfo:type_name -> gonews.Pagination
	3,  // 3: gonews.ListPostsResponse.posts:type_name -> gonews.Post
	3,  // 4: gonews.DetailedNewsResponse.post:type_name -> gonews.Post
	10, // 5: gonews.DetailedNewsResponse.comments:type_name -> comment.Comment
	0,  // 6: gonews.GoNewsService.Posts:input_type -> gonews.PostsRequest
	1,  // 7: gonews.GoNewsService.NewsFullDetailed:input_type -> gonews.OneNewsRequest
	1,  // 8: gonews.GoNewsService.NewsShortDetailed:input_type -> gonews.OneNewsRequest
	2,  // 9: gonews.GoNewsService.FilterNews:input_type -> gonews.FilterNewsRequest
	7,  // 10: gonews.GoNewsService.ListNews:input_type -> gonews.ListPostsRequest
	5,  // 11: gonews.GoNewsService.Posts:output_type -> gonews.PostsResponse
	9,  // 12: gonews.GoNewsService.NewsFullDetailed:output_type -> gonews.DetailedNewsResponse
	4,  // 13: gonews.GoNewsService.NewsShortDetailed:output_type -> gonews.OnePostResponse
	8,  // 14: gonews.GoNewsService.FilterNews:output_type -> gonews.ListPostsResponse
	8,  // 15: gonews.GoNewsService.ListNews:output_type -> gonews.ListPostsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gonews_gonews_proto_init() }
func file_gonews_gonews_proto_init() {
	if File_gonews_gonews_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gonews_gonews_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_gonews_gonews_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_gonews_gonews_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsRequest); i {
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
		file_gonews_gonews_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsResponse); i {
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
		file_gonews_gonews_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailedNewsResponse); i {
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
			RawDescriptor: file_gonews_gonews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gonews_gonews_proto_goTypes,
		DependencyIndexes: file_gonews_gonews_proto_depIdxs,
		MessageInfos:      file_gonews_gonews_proto_msgTypes,
	}.Build()
	File_gonews_gonews_proto = out.File
	file_gonews_gonews_proto_rawDesc = nil
	file_gonews_gonews_proto_goTypes = nil
	file_gonews_gonews_proto_depIdxs = nil
}
