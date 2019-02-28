// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scrape/grpc_scrape.proto

package plagiari_sm_grpc_scrape

/*
Package name (~DNS)
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RequestArticle for SimpleScrape
type RequestArticle struct {
	Feed                 string   `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Lang                 string   `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestArticle) Reset()         { *m = RequestArticle{} }
func (m *RequestArticle) String() string { return proto.CompactTextString(m) }
func (*RequestArticle) ProtoMessage()    {}
func (*RequestArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{0}
}
func (m *RequestArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestArticle.Unmarshal(m, b)
}
func (m *RequestArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestArticle.Marshal(b, m, deterministic)
}
func (dst *RequestArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestArticle.Merge(dst, src)
}
func (m *RequestArticle) XXX_Size() int {
	return xxx_messageInfo_RequestArticle.Size(m)
}
func (m *RequestArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestArticle.DiscardUnknown(m)
}

var xxx_messageInfo_RequestArticle proto.InternalMessageInfo

func (m *RequestArticle) GetFeed() string {
	if m != nil {
		return m.Feed
	}
	return ""
}

func (m *RequestArticle) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RequestArticle) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

// RequestStream for Scrape in Streaming mode (ex. svc-listen)
type RequestStream struct {
	Feed                 string   `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	TweetId              string   `protobuf:"bytes,3,opt,name=tweetId,proto3" json:"tweetId,omitempty"`
	Lang                 string   `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	ScreenName           string   `protobuf:"bytes,5,opt,name=screenName,proto3" json:"screenName,omitempty"`
	CrawledAt            string   `protobuf:"bytes,6,opt,name=crawledAt,proto3" json:"crawledAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestStream) Reset()         { *m = RequestStream{} }
func (m *RequestStream) String() string { return proto.CompactTextString(m) }
func (*RequestStream) ProtoMessage()    {}
func (*RequestStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{1}
}
func (m *RequestStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestStream.Unmarshal(m, b)
}
func (m *RequestStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestStream.Marshal(b, m, deterministic)
}
func (dst *RequestStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestStream.Merge(dst, src)
}
func (m *RequestStream) XXX_Size() int {
	return xxx_messageInfo_RequestStream.Size(m)
}
func (m *RequestStream) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestStream.DiscardUnknown(m)
}

var xxx_messageInfo_RequestStream proto.InternalMessageInfo

func (m *RequestStream) GetFeed() string {
	if m != nil {
		return m.Feed
	}
	return ""
}

func (m *RequestStream) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RequestStream) GetTweetId() string {
	if m != nil {
		return m.TweetId
	}
	return ""
}

func (m *RequestStream) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *RequestStream) GetScreenName() string {
	if m != nil {
		return m.ScreenName
	}
	return ""
}

func (m *RequestStream) GetCrawledAt() string {
	if m != nil {
		return m.CrawledAt
	}
	return ""
}

// Content of the Scraped Article or Stream
type Content struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Authors              []string `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	PublishedAt          string   `protobuf:"bytes,5,opt,name=publishedAt,proto3" json:"publishedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{2}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (dst *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(dst, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Content) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Content) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Content) GetPublishedAt() string {
	if m != nil {
		return m.PublishedAt
	}
	return ""
}

// NLP data of the Scraped Article or Stream
type NLP struct {
	Keywords             []string `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	StopWords            []string `protobuf:"bytes,2,rep,name=stopWords,proto3" json:"stopWords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NLP) Reset()         { *m = NLP{} }
func (m *NLP) String() string { return proto.CompactTextString(m) }
func (*NLP) ProtoMessage()    {}
func (*NLP) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{3}
}
func (m *NLP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NLP.Unmarshal(m, b)
}
func (m *NLP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NLP.Marshal(b, m, deterministic)
}
func (dst *NLP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NLP.Merge(dst, src)
}
func (m *NLP) XXX_Size() int {
	return xxx_messageInfo_NLP.Size(m)
}
func (m *NLP) XXX_DiscardUnknown() {
	xxx_messageInfo_NLP.DiscardUnknown(m)
}

var xxx_messageInfo_NLP proto.InternalMessageInfo

func (m *NLP) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *NLP) GetStopWords() []string {
	if m != nil {
		return m.StopWords
	}
	return nil
}

// Data for the Response Object
type Data struct {
	Content              *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Nlp                  *NLP     `protobuf:"bytes,2,opt,name=nlp,proto3" json:"nlp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{4}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Data) GetNlp() *NLP {
	if m != nil {
		return m.Nlp
	}
	return nil
}

// Response Object
type Response struct {
	// success, error
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 200, 500
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// message
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// data Object {}
	Data                 *Data    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_scrape_503903a20e92b3b0, []int{5}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestArticle)(nil), "plagiari.sm.grpc.scrape.RequestArticle")
	proto.RegisterType((*RequestStream)(nil), "plagiari.sm.grpc.scrape.RequestStream")
	proto.RegisterType((*Content)(nil), "plagiari.sm.grpc.scrape.Content")
	proto.RegisterType((*NLP)(nil), "plagiari.sm.grpc.scrape.NLP")
	proto.RegisterType((*Data)(nil), "plagiari.sm.grpc.scrape.Data")
	proto.RegisterType((*Response)(nil), "plagiari.sm.grpc.scrape.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPCScrapeClient is the client API for GRPCScrape service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCScrapeClient interface {
	// Endpoint Scrape
	Scrape(ctx context.Context, in *RequestStream, opts ...grpc.CallOption) (*Response, error)
	// Endpoint SimpleScrape
	SimpleScrape(ctx context.Context, in *RequestArticle, opts ...grpc.CallOption) (*Response, error)
}

type gRPCScrapeClient struct {
	cc *grpc.ClientConn
}

func NewGRPCScrapeClient(cc *grpc.ClientConn) GRPCScrapeClient {
	return &gRPCScrapeClient{cc}
}

func (c *gRPCScrapeClient) Scrape(ctx context.Context, in *RequestStream, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/plagiari.sm.grpc.scrape.GRPCScrape/Scrape", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCScrapeClient) SimpleScrape(ctx context.Context, in *RequestArticle, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/plagiari.sm.grpc.scrape.GRPCScrape/SimpleScrape", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCScrapeServer is the server API for GRPCScrape service.
type GRPCScrapeServer interface {
	// Endpoint Scrape
	Scrape(context.Context, *RequestStream) (*Response, error)
	// Endpoint SimpleScrape
	SimpleScrape(context.Context, *RequestArticle) (*Response, error)
}

func RegisterGRPCScrapeServer(s *grpc.Server, srv GRPCScrapeServer) {
	s.RegisterService(&_GRPCScrape_serviceDesc, srv)
}

func _GRPCScrape_Scrape_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCScrapeServer).Scrape(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plagiari.sm.grpc.scrape.GRPCScrape/Scrape",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCScrapeServer).Scrape(ctx, req.(*RequestStream))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCScrape_SimpleScrape_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCScrapeServer).SimpleScrape(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plagiari.sm.grpc.scrape.GRPCScrape/SimpleScrape",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCScrapeServer).SimpleScrape(ctx, req.(*RequestArticle))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCScrape_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plagiari.sm.grpc.scrape.GRPCScrape",
	HandlerType: (*GRPCScrapeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scrape",
			Handler:    _GRPCScrape_Scrape_Handler,
		},
		{
			MethodName: "SimpleScrape",
			Handler:    _GRPCScrape_SimpleScrape_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrape/grpc_scrape.proto",
}

func init() {
	proto.RegisterFile("scrape/grpc_scrape.proto", fileDescriptor_grpc_scrape_503903a20e92b3b0)
}

var fileDescriptor_grpc_scrape_503903a20e92b3b0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xb5, 0x93, 0xb4, 0x13, 0x40, 0x68, 0x85, 0x60, 0x55, 0x15, 0x14, 0x7c, 0x00, 0x4e,
	0x46, 0x84, 0x1b, 0x17, 0x54, 0x15, 0x09, 0x81, 0xaa, 0x28, 0x72, 0x84, 0xb8, 0x20, 0xa1, 0x8d,
	0x3d, 0xb8, 0x16, 0xb6, 0x77, 0xd9, 0x1d, 0x2b, 0xea, 0x99, 0x2f, 0xe1, 0x53, 0xf8, 0x33, 0xb4,
	0xe3, 0x4d, 0x53, 0x0e, 0xa1, 0xbd, 0xbd, 0x99, 0x37, 0xfb, 0xf6, 0xbd, 0x59, 0x1b, 0xa4, 0x2b,
	0xac, 0x32, 0xf8, 0xaa, 0xb2, 0xa6, 0xf8, 0x36, 0xe0, 0xcc, 0x58, 0x4d, 0x5a, 0x3c, 0x36, 0x8d,
	0xaa, 0x6a, 0x65, 0xeb, 0xcc, 0xb5, 0x99, 0xa7, 0xb3, 0x81, 0x4e, 0x3f, 0xc1, 0xfd, 0x1c, 0x7f,
	0xf6, 0xe8, 0xe8, 0xd4, 0x52, 0x5d, 0x34, 0x28, 0x04, 0x24, 0xdf, 0x11, 0x4b, 0x19, 0xcd, 0xa2,
	0x97, 0x47, 0x39, 0x63, 0xf1, 0x00, 0xe2, 0xde, 0x36, 0xf2, 0x80, 0x5b, 0x1e, 0xfa, 0xa9, 0x46,
	0x75, 0x95, 0x8c, 0x87, 0x29, 0x8f, 0xd3, 0xdf, 0x11, 0xdc, 0x0b, 0x62, 0x2b, 0xb2, 0xa8, 0xda,
	0x5b, 0x6a, 0x49, 0x98, 0xd0, 0x06, 0x91, 0x3e, 0x96, 0x41, 0x6e, 0x5b, 0x5e, 0xdd, 0x92, 0xec,
	0x6e, 0x11, 0x4f, 0x01, 0x5c, 0x61, 0x11, 0xbb, 0x85, 0x6a, 0x51, 0x8e, 0x98, 0xb9, 0xd6, 0x11,
	0x27, 0x70, 0x54, 0x58, 0xb5, 0x69, 0xb0, 0x3c, 0x25, 0x39, 0x66, 0x7a, 0xd7, 0x48, 0x7f, 0x45,
	0x30, 0x39, 0xd3, 0x1d, 0x61, 0x47, 0xe2, 0x21, 0x8c, 0xa8, 0xa6, 0x06, 0x83, 0xbd, 0xa1, 0xf0,
	0x77, 0xae, 0x75, 0x79, 0x19, 0x0c, 0x32, 0xf6, 0x0e, 0x55, 0x4f, 0x17, 0xda, 0x3a, 0x19, 0xcf,
	0x62, 0xef, 0x30, 0x94, 0x7e, 0x9a, 0x54, 0xe5, 0x64, 0xc2, 0x6d, 0xc6, 0x62, 0x06, 0x53, 0xd3,
	0xaf, 0x9b, 0xda, 0x5d, 0xb0, 0x87, 0xc1, 0xe2, 0xf5, 0x56, 0xfa, 0x0e, 0xe2, 0xc5, 0xf9, 0x52,
	0x1c, 0xc3, 0xe1, 0x0f, 0xbc, 0xdc, 0x68, 0x5b, 0x3a, 0x19, 0xb1, 0xc0, 0x55, 0xed, 0x63, 0x38,
	0xd2, 0xe6, 0x0b, 0x93, 0x07, 0x4c, 0xee, 0x1a, 0xa9, 0x85, 0xe4, 0xbd, 0x22, 0x25, 0xde, 0xc2,
	0xa4, 0x18, 0xd2, 0x70, 0x88, 0xe9, 0x7c, 0x96, 0xed, 0x79, 0xe9, 0x2c, 0xa4, 0xce, 0xb7, 0x07,
	0x44, 0x06, 0x71, 0xd7, 0x18, 0xce, 0x39, 0x9d, 0x9f, 0xec, 0x3d, 0xb7, 0x38, 0x5f, 0xe6, 0x7e,
	0xd0, 0xaf, 0xee, 0x30, 0x47, 0x67, 0x74, 0xe7, 0x50, 0x3c, 0x82, 0xb1, 0x23, 0x45, 0xbd, 0x0b,
	0xcb, 0x0b, 0x95, 0xdf, 0x47, 0xa1, 0x4b, 0x64, 0xd5, 0x51, 0xce, 0xd8, 0x6f, 0xaf, 0x45, 0xe7,
	0x54, 0x85, 0xdb, 0xf7, 0x0d, 0xa5, 0x78, 0x0d, 0x49, 0xa9, 0x48, 0xf1, 0xfb, 0x4e, 0xe7, 0x4f,
	0xf6, 0x7a, 0xf0, 0x59, 0x73, 0x1e, 0x9d, 0xff, 0x89, 0x00, 0x3e, 0xe4, 0xcb, 0xb3, 0x15, 0x33,
	0xe2, 0x33, 0x8c, 0x03, 0x7a, 0xbe, 0xf7, 0xf4, 0x3f, 0xdf, 0xe4, 0xf1, 0xb3, 0xff, 0xcc, 0x0d,
	0xe1, 0xd2, 0x3b, 0xe2, 0x2b, 0xdc, 0x5d, 0xd5, 0xad, 0x69, 0x30, 0x88, 0xbf, 0xb8, 0x49, 0x3c,
	0xfc, 0x3d, 0xb7, 0x52, 0x5f, 0x8f, 0xf9, 0xa7, 0x7c, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0x1c, 0x35, 0xe8, 0xb0, 0x03, 0x00, 0x00,
}
