// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ApiCall struct {
	ConvId               string   `protobuf:"bytes,1,opt,name=conv_id,json=convId,proto3" json:"conv_id,omitempty"`
	Uid                  bool     `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiCall) Reset()         { *m = ApiCall{} }
func (m *ApiCall) String() string { return proto.CompactTextString(m) }
func (*ApiCall) ProtoMessage()    {}
func (*ApiCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ApiCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiCall.Unmarshal(m, b)
}
func (m *ApiCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiCall.Marshal(b, m, deterministic)
}
func (m *ApiCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiCall.Merge(m, src)
}
func (m *ApiCall) XXX_Size() int {
	return xxx_messageInfo_ApiCall.Size(m)
}
func (m *ApiCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiCall.DiscardUnknown(m)
}

var xxx_messageInfo_ApiCall proto.InternalMessageInfo

func (m *ApiCall) GetConvId() string {
	if m != nil {
		return m.ConvId
	}
	return ""
}

func (m *ApiCall) GetUid() bool {
	if m != nil {
		return m.Uid
	}
	return false
}

func (m *ApiCall) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ApiResponse struct {
	ConvId               string   `protobuf:"bytes,1,opt,name=conv_id,json=convId,proto3" json:"conv_id,omitempty"`
	Uid                  bool     `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Score                float32  `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	RollingScore         float32  `protobuf:"fixed32,4,opt,name=rolling_score,json=rollingScore,proto3" json:"rolling_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiResponse) Reset()         { *m = ApiResponse{} }
func (m *ApiResponse) String() string { return proto.CompactTextString(m) }
func (*ApiResponse) ProtoMessage()    {}
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ApiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiResponse.Unmarshal(m, b)
}
func (m *ApiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiResponse.Marshal(b, m, deterministic)
}
func (m *ApiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiResponse.Merge(m, src)
}
func (m *ApiResponse) XXX_Size() int {
	return xxx_messageInfo_ApiResponse.Size(m)
}
func (m *ApiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiResponse proto.InternalMessageInfo

func (m *ApiResponse) GetConvId() string {
	if m != nil {
		return m.ConvId
	}
	return ""
}

func (m *ApiResponse) GetUid() bool {
	if m != nil {
		return m.Uid
	}
	return false
}

func (m *ApiResponse) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ApiResponse) GetRollingScore() float32 {
	if m != nil {
		return m.RollingScore
	}
	return 0
}

func init() {
	proto.RegisterType((*ApiCall)(nil), "apiCall")
	proto.RegisterType((*ApiResponse)(nil), "apiResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe3, 0x62, 0x4f, 0x2c, 0xc8, 0x74, 0x4e, 0xcc, 0xc9,
	0x11, 0x12, 0xe7, 0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x98,
	0x14, 0x18, 0x35, 0x38, 0x82, 0x40, 0x4c, 0x90, 0x48, 0x6e, 0x71, 0xba, 0x04, 0x33, 0x58, 0x19,
	0x88, 0xa9, 0x54, 0xcc, 0xc5, 0x9d, 0x58, 0x90, 0x19, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x4a, 0x8a, 0x59, 0x22, 0x5c, 0xac, 0xc5, 0xc9, 0xf9, 0x45, 0xa9, 0x60, 0xd3, 0x98, 0x82, 0x20,
	0x1c, 0x21, 0x65, 0x2e, 0xde, 0xa2, 0xfc, 0x9c, 0x9c, 0xcc, 0xbc, 0xf4, 0x78, 0x88, 0x2c, 0x0b,
	0x58, 0x96, 0x07, 0x2a, 0x18, 0x0c, 0x12, 0x33, 0x8a, 0xe6, 0x62, 0x0f, 0x28, 0xca, 0x4f, 0x4e,
	0x2d, 0x2e, 0x16, 0xd2, 0xe4, 0xe2, 0x0c, 0x4a, 0xcd, 0x49, 0x2d, 0x4b, 0xcc, 0x4b, 0x4e, 0x15,
	0xe2, 0xd0, 0x83, 0xfa, 0x49, 0x8a, 0x47, 0x0f, 0xc9, 0x55, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c,
	0x42, 0xaa, 0x5c, 0xac, 0x21, 0x20, 0x63, 0xf0, 0x2b, 0x4b, 0x62, 0x03, 0x07, 0x90, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x96, 0xec, 0x31, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessClient is the client API for Process service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessClient interface {
	Relevance(ctx context.Context, opts ...grpc.CallOption) (Process_RelevanceClient, error)
	Troll(ctx context.Context, opts ...grpc.CallOption) (Process_TrollClient, error)
}

type processClient struct {
	cc *grpc.ClientConn
}

func NewProcessClient(cc *grpc.ClientConn) ProcessClient {
	return &processClient{cc}
}

func (c *processClient) Relevance(ctx context.Context, opts ...grpc.CallOption) (Process_RelevanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Process_serviceDesc.Streams[0], "/Process/Relevance", opts...)
	if err != nil {
		return nil, err
	}
	x := &processRelevanceClient{stream}
	return x, nil
}

type Process_RelevanceClient interface {
	Send(*ApiCall) error
	Recv() (*ApiResponse, error)
	grpc.ClientStream
}

type processRelevanceClient struct {
	grpc.ClientStream
}

func (x *processRelevanceClient) Send(m *ApiCall) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processRelevanceClient) Recv() (*ApiResponse, error) {
	m := new(ApiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processClient) Troll(ctx context.Context, opts ...grpc.CallOption) (Process_TrollClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Process_serviceDesc.Streams[1], "/Process/Troll", opts...)
	if err != nil {
		return nil, err
	}
	x := &processTrollClient{stream}
	return x, nil
}

type Process_TrollClient interface {
	Send(*ApiCall) error
	Recv() (*ApiResponse, error)
	grpc.ClientStream
}

type processTrollClient struct {
	grpc.ClientStream
}

func (x *processTrollClient) Send(m *ApiCall) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processTrollClient) Recv() (*ApiResponse, error) {
	m := new(ApiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessServer is the server API for Process service.
type ProcessServer interface {
	Relevance(Process_RelevanceServer) error
	Troll(Process_TrollServer) error
}

func RegisterProcessServer(s *grpc.Server, srv ProcessServer) {
	s.RegisterService(&_Process_serviceDesc, srv)
}

func _Process_Relevance_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessServer).Relevance(&processRelevanceServer{stream})
}

type Process_RelevanceServer interface {
	Send(*ApiResponse) error
	Recv() (*ApiCall, error)
	grpc.ServerStream
}

type processRelevanceServer struct {
	grpc.ServerStream
}

func (x *processRelevanceServer) Send(m *ApiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processRelevanceServer) Recv() (*ApiCall, error) {
	m := new(ApiCall)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Process_Troll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessServer).Troll(&processTrollServer{stream})
}

type Process_TrollServer interface {
	Send(*ApiResponse) error
	Recv() (*ApiCall, error)
	grpc.ServerStream
}

type processTrollServer struct {
	grpc.ServerStream
}

func (x *processTrollServer) Send(m *ApiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processTrollServer) Recv() (*ApiCall, error) {
	m := new(ApiCall)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Process_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Process",
	HandlerType: (*ProcessServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Relevance",
			Handler:       _Process_Relevance_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Troll",
			Handler:       _Process_Troll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
