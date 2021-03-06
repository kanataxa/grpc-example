// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetInputRequest struct {
	Value                int32    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInputRequest) Reset()         { *m = GetInputRequest{} }
func (m *GetInputRequest) String() string { return proto.CompactTextString(m) }
func (*GetInputRequest) ProtoMessage()    {}
func (*GetInputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *GetInputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInputRequest.Unmarshal(m, b)
}
func (m *GetInputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInputRequest.Marshal(b, m, deterministic)
}
func (m *GetInputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInputRequest.Merge(m, src)
}
func (m *GetInputRequest) XXX_Size() int {
	return xxx_messageInfo_GetInputRequest.Size(m)
}
func (m *GetInputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInputRequest proto.InternalMessageInfo

func (m *GetInputRequest) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*GetInputRequest)(nil), "GetInputRequest")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x92, 0x4a, 0xea,
	0x5c, 0xfc, 0xee, 0xa9, 0x25, 0x9e, 0x79, 0x05, 0xa5, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac,
	0x41, 0x10, 0x8e, 0x91, 0x3b, 0x17, 0x6f, 0x50, 0x6a, 0x4e, 0x62, 0xa5, 0x6f, 0x6a, 0x4a, 0x66,
	0x62, 0x49, 0x7e, 0x91, 0x90, 0x19, 0x17, 0x07, 0x4c, 0xa7, 0x90, 0x80, 0x1e, 0x9a, 0x21, 0x52,
	0x62, 0x7a, 0x10, 0x5b, 0xf5, 0x60, 0xb6, 0xea, 0xb9, 0x82, 0x6c, 0x55, 0x62, 0x48, 0x62, 0x03,
	0x8b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf7, 0x9b, 0x3a, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RelayMediatorClient is the client API for RelayMediator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelayMediatorClient interface {
	GetInput(ctx context.Context, in *GetInputRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type relayMediatorClient struct {
	cc *grpc.ClientConn
}

func NewRelayMediatorClient(cc *grpc.ClientConn) RelayMediatorClient {
	return &relayMediatorClient{cc}
}

func (c *relayMediatorClient) GetInput(ctx context.Context, in *GetInputRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/RelayMediator/GetInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayMediatorServer is the server API for RelayMediator service.
type RelayMediatorServer interface {
	GetInput(context.Context, *GetInputRequest) (*empty.Empty, error)
}

// UnimplementedRelayMediatorServer can be embedded to have forward compatible implementations.
type UnimplementedRelayMediatorServer struct {
}

func (*UnimplementedRelayMediatorServer) GetInput(ctx context.Context, req *GetInputRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInput not implemented")
}

func RegisterRelayMediatorServer(s *grpc.Server, srv RelayMediatorServer) {
	s.RegisterService(&_RelayMediator_serviceDesc, srv)
}

func _RelayMediator_GetInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayMediatorServer).GetInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelayMediator/GetInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayMediatorServer).GetInput(ctx, req.(*GetInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RelayMediator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RelayMediator",
	HandlerType: (*RelayMediatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInput",
			Handler:    _RelayMediator_GetInput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
