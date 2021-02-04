// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: compare/grpc_compare.proto

// Package name (~DNS)

package compare

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang    string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Meta) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Language Model Selection
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Required Values
	DocId       string `protobuf:"bytes,2,opt,name=docId,proto3" json:"docId,omitempty"`
	PublishedAt string `protobuf:"bytes,3,opt,name=publishedAt,proto3" json:"publishedAt,omitempty"`
	Body        string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// Optional Values
	Title  string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Url    string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Source string   `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Tags   []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{1}
}

func (x *Article) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Article) GetDocId() string {
	if x != nil {
		return x.DocId
	}
	return ""
}

func (x *Article) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Article) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Article) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Case    *Article   `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	Compare []*Article `protobuf:"bytes,2,rep,name=compare,proto3" json:"compare,omitempty"`
	Deep    bool       `protobuf:"varint,3,opt,name=deep,proto3" json:"deep,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{2}
}

func (x *List) GetCase() *Article {
	if x != nil {
		return x.Case
	}
	return nil
}

func (x *List) GetCompare() []*Article {
	if x != nil {
		return x.Compare
	}
	return nil
}

func (x *List) GetDeep() bool {
	if x != nil {
		return x.Deep
	}
	return false
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DocId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocId string `protobuf:"bytes,1,opt,name=docId,proto3" json:"docId,omitempty"`
}

func (x *DocId) Reset() {
	*x = DocId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocId) ProtoMessage() {}

func (x *DocId) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocId.ProtoReflect.Descriptor instead.
func (*DocId) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{4}
}

func (x *DocId) GetDocId() string {
	if x != nil {
		return x.DocId
	}
	return ""
}

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Data   *Article `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{5}
}

func (x *Case) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Case) GetData() *Article {
	if x != nil {
		return x.Data
	}
	return nil
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{6}
}

func (x *Score) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Score `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{7}
}

func (x *Relationship) GetData() *Score {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Relationship) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Relationship) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Case          *Case           `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	Related       []*Case         `protobuf:"bytes,2,rep,name=related,proto3" json:"related,omitempty"`
	Relationships []*Relationship `protobuf:"bytes,3,rep,name=relationships,proto3" json:"relationships,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{8}
}

func (x *Data) GetCase() *Case {
	if x != nil {
		return x.Case
	}
	return nil
}

func (x *Data) GetRelated() []*Case {
	if x != nil {
		return x.Related
	}
	return nil
}

func (x *Data) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

// Response Object
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success, error
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 200, 500
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// message
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// data Object {}
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compare_grpc_compare_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_compare_grpc_compare_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_compare_grpc_compare_proto_rawDescGZIP(), []int{9}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_compare_grpc_compare_proto protoreflect.FileDescriptor

var file_compare_grpc_compare_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x22, 0x34, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xcc, 0x01, 0x0a, 0x07,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f,
	0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x6c, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x04, 0x63, 0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x65, 0x70, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d,
	0x0a, 0x05, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64, 0x22, 0x44, 0x0a,
	0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x8f,
	0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x63, 0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x22, 0x73, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb0, 0x01, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x42, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x42, 0x79, 0x4e, 0x65, 0x77, 0x44, 0x6f,
	0x63, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_compare_grpc_compare_proto_rawDescOnce sync.Once
	file_compare_grpc_compare_proto_rawDescData = file_compare_grpc_compare_proto_rawDesc
)

func file_compare_grpc_compare_proto_rawDescGZIP() []byte {
	file_compare_grpc_compare_proto_rawDescOnce.Do(func() {
		file_compare_grpc_compare_proto_rawDescData = protoimpl.X.CompressGZIP(file_compare_grpc_compare_proto_rawDescData)
	})
	return file_compare_grpc_compare_proto_rawDescData
}

var file_compare_grpc_compare_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_compare_grpc_compare_proto_goTypes = []interface{}{
	(*Meta)(nil),         // 0: compare.Meta
	(*Article)(nil),      // 1: compare.Article
	(*List)(nil),         // 2: compare.List
	(*Id)(nil),           // 3: compare.Id
	(*DocId)(nil),        // 4: compare.DocId
	(*Case)(nil),         // 5: compare.Case
	(*Score)(nil),        // 6: compare.Score
	(*Relationship)(nil), // 7: compare.Relationship
	(*Data)(nil),         // 8: compare.Data
	(*Response)(nil),     // 9: compare.Response
}
var file_compare_grpc_compare_proto_depIdxs = []int32{
	0,  // 0: compare.Article.meta:type_name -> compare.Meta
	1,  // 1: compare.List.case:type_name -> compare.Article
	1,  // 2: compare.List.compare:type_name -> compare.Article
	1,  // 3: compare.Case.data:type_name -> compare.Article
	6,  // 4: compare.Relationship.data:type_name -> compare.Score
	5,  // 5: compare.Data.case:type_name -> compare.Case
	5,  // 6: compare.Data.related:type_name -> compare.Case
	7,  // 7: compare.Data.relationships:type_name -> compare.Relationship
	8,  // 8: compare.Response.data:type_name -> compare.Data
	2,  // 9: compare.GRPCCompare.CompareByList:input_type -> compare.List
	4,  // 10: compare.GRPCCompare.CompareById:input_type -> compare.DocId
	1,  // 11: compare.GRPCCompare.CompareByNewDoc:input_type -> compare.Article
	9,  // 12: compare.GRPCCompare.CompareByList:output_type -> compare.Response
	9,  // 13: compare.GRPCCompare.CompareById:output_type -> compare.Response
	9,  // 14: compare.GRPCCompare.CompareByNewDoc:output_type -> compare.Response
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_compare_grpc_compare_proto_init() }
func file_compare_grpc_compare_proto_init() {
	if File_compare_grpc_compare_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compare_grpc_compare_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_compare_grpc_compare_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_compare_grpc_compare_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
		file_compare_grpc_compare_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_compare_grpc_compare_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocId); i {
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
		file_compare_grpc_compare_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
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
		file_compare_grpc_compare_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_compare_grpc_compare_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relationship); i {
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
		file_compare_grpc_compare_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_compare_grpc_compare_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_compare_grpc_compare_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compare_grpc_compare_proto_goTypes,
		DependencyIndexes: file_compare_grpc_compare_proto_depIdxs,
		MessageInfos:      file_compare_grpc_compare_proto_msgTypes,
	}.Build()
	File_compare_grpc_compare_proto = out.File
	file_compare_grpc_compare_proto_rawDesc = nil
	file_compare_grpc_compare_proto_goTypes = nil
	file_compare_grpc_compare_proto_depIdxs = nil
}
