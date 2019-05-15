// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo/echo.proto

package Echo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EchoMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMessage) Reset()         { *m = EchoMessage{} }
func (m *EchoMessage) String() string { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()    {}
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_071447ae7b5e538d, []int{0}
}

func (m *EchoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMessage.Unmarshal(m, b)
}
func (m *EchoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMessage.Marshal(b, m, deterministic)
}
func (m *EchoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMessage.Merge(m, src)
}
func (m *EchoMessage) XXX_Size() int {
	return xxx_messageInfo_EchoMessage.Size(m)
}
func (m *EchoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMessage proto.InternalMessageInfo

func (m *EchoMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "Echo.EchoMessage")
}

func init() { proto.RegisterFile("echo/echo.proto", fileDescriptor_071447ae7b5e538d) }

var fileDescriptor_071447ae7b5e538d = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4d, 0xce, 0xc8,
	0xd7, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0xae, 0xc9, 0x19, 0xf9, 0x52,
	0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9,
	0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x35, 0x4a, 0xca, 0x5c, 0xdc, 0x20, 0x55, 0xbe,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x91, 0x27, 0x17, 0xd8, 0x28, 0x21, 0x47, 0x28,
	0x2d, 0xa8, 0x07, 0xa2, 0xf4, 0x90, 0x34, 0x4a, 0x61, 0x0a, 0x29, 0x09, 0x37, 0x5d, 0x7e, 0x32,
	0x99, 0x89, 0x57, 0x89, 0x43, 0xbf, 0xcc, 0x10, 0xec, 0x2c, 0x2b, 0x46, 0xad, 0x24, 0x36, 0xb0,
	0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x73, 0xa2, 0x91, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/Echo.Echo/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Echo(ctx context.Context, req *EchoMessage) (*EchoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Echo.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo/echo.proto",
}
