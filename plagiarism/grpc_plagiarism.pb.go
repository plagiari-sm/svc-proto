// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plagiarism/grpc_plagiarism.proto

package plagiari_sm_grpc_plagiarism

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

// RequestDocument from Elasticsearch
type RequestDocument struct {
	// article id
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDocument) Reset()         { *m = RequestDocument{} }
func (m *RequestDocument) String() string { return proto.CompactTextString(m) }
func (*RequestDocument) ProtoMessage()    {}
func (*RequestDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_plagiarism_17ea1748f5675512, []int{0}
}
func (m *RequestDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestDocument.Unmarshal(m, b)
}
func (m *RequestDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestDocument.Marshal(b, m, deterministic)
}
func (dst *RequestDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDocument.Merge(dst, src)
}
func (m *RequestDocument) XXX_Size() int {
	return xxx_messageInfo_RequestDocument.Size(m)
}
func (m *RequestDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDocument.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDocument proto.InternalMessageInfo

func (m *RequestDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Notification
type Notification struct {
	// success, error
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// 200, 500
	Code int32 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	// message
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	// label
	Label string `protobuf:"bytes,4,opt,name=label" json:"label,omitempty"`
	// article id
	Id                   string   `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_plagiarism_17ea1748f5675512, []int{1}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Notification) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Notification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Notification) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Notification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestDocument)(nil), "plagiari.sm.grpc.plagiarism.RequestDocument")
	proto.RegisterType((*Notification)(nil), "plagiari.sm.grpc.plagiarism.Notification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlagiarismClient is the client API for Plagiarism service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlagiarismClient interface {
	// Endpoint Detect
	Detect(ctx context.Context, in *RequestDocument, opts ...grpc.CallOption) (*Notification, error)
}

type plagiarismClient struct {
	cc *grpc.ClientConn
}

func NewPlagiarismClient(cc *grpc.ClientConn) PlagiarismClient {
	return &plagiarismClient{cc}
}

func (c *plagiarismClient) Detect(ctx context.Context, in *RequestDocument, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/plagiari.sm.grpc.plagiarism.Plagiarism/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlagiarismServer is the server API for Plagiarism service.
type PlagiarismServer interface {
	// Endpoint Detect
	Detect(context.Context, *RequestDocument) (*Notification, error)
}

func RegisterPlagiarismServer(s *grpc.Server, srv PlagiarismServer) {
	s.RegisterService(&_Plagiarism_serviceDesc, srv)
}

func _Plagiarism_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDocument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlagiarismServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plagiari.sm.grpc.plagiarism.Plagiarism/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlagiarismServer).Detect(ctx, req.(*RequestDocument))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plagiarism_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plagiari.sm.grpc.plagiarism.Plagiarism",
	HandlerType: (*PlagiarismServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _Plagiarism_Detect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plagiarism/grpc_plagiarism.proto",
}

func init() {
	proto.RegisterFile("plagiarism/grpc_plagiarism.proto", fileDescriptor_grpc_plagiarism_17ea1748f5675512)
}

var fileDescriptor_grpc_plagiarism_17ea1748f5675512 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0x80, 0x30,
	0x14, 0x86, 0xd3, 0xd4, 0xe8, 0x10, 0x05, 0x87, 0x88, 0x51, 0x37, 0xe6, 0x95, 0x41, 0x2c, 0xa8,
	0x57, 0xf0, 0x3a, 0xc2, 0x17, 0x88, 0x39, 0x4f, 0x32, 0x70, 0x4e, 0xdd, 0xf1, 0xa6, 0xa7, 0x8f,
	0x96, 0x66, 0x74, 0xe1, 0xdd, 0xbe, 0x9f, 0x0f, 0xfe, 0xfd, 0x07, 0xf2, 0xb1, 0x57, 0x9d, 0x51,
	0xb3, 0xf1, 0xf6, 0xa9, 0x9b, 0x47, 0xfd, 0xbe, 0xb3, 0x1c, 0x67, 0xc7, 0x0e, 0xef, 0xb6, 0x44,
	0x7a, 0x2b, 0xbf, 0x15, 0xb9, 0x2b, 0xc5, 0x3d, 0x5c, 0xd5, 0x34, 0x2d, 0xe4, 0xb9, 0x72, 0x7a,
	0xb1, 0x34, 0x30, 0x5e, 0x42, 0x6c, 0x5a, 0x11, 0xe5, 0x51, 0x79, 0x5e, 0xc7, 0xa6, 0x2d, 0x3e,
	0xe1, 0xe2, 0xd5, 0xb1, 0xf9, 0x30, 0x5a, 0xb1, 0x71, 0x03, 0xde, 0x40, 0xe6, 0x59, 0xf1, 0xe2,
	0x57, 0x67, 0x25, 0x44, 0x48, 0xb4, 0x6b, 0x49, 0xc4, 0x79, 0x54, 0xa6, 0x75, 0x78, 0xa3, 0x80,
	0x33, 0x4b, 0xde, 0xab, 0x8e, 0xc4, 0x69, 0x90, 0x37, 0xc4, 0x6b, 0x48, 0x7b, 0xd5, 0x50, 0x2f,
	0x92, 0x90, 0xff, 0xc0, 0xda, 0x9d, 0x6e, 0xdd, 0xcf, 0x13, 0xc0, 0xdb, 0xef, 0x67, 0x51, 0x43,
	0x56, 0x11, 0x93, 0x66, 0x7c, 0x94, 0x07, 0xa3, 0xe4, 0xbf, 0x45, 0xb7, 0x0f, 0x87, 0xf6, 0xdf,
	0x71, 0xc5, 0x49, 0x93, 0x85, 0xab, 0xbd, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x54, 0x78, 0xd1,
	0xe6, 0x59, 0x01, 0x00, 0x00,
}
