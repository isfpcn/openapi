// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/douyin/client_token.proto

package douyin

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetClientTokenRequest struct {
	// 是应用唯一标识 aw05az2qjv******
	ClientKey            string   `protobuf:"bytes,1,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty" form:"client_key" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClientTokenRequest) Reset()         { *m = GetClientTokenRequest{} }
func (m *GetClientTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetClientTokenRequest) ProtoMessage()    {}
func (*GetClientTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a5001b5f0223dd, []int{0}
}
func (m *GetClientTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientTokenRequest.Unmarshal(m, b)
}
func (m *GetClientTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetClientTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientTokenRequest.Merge(m, src)
}
func (m *GetClientTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetClientTokenRequest.Size(m)
}
func (m *GetClientTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientTokenRequest proto.InternalMessageInfo

func (m *GetClientTokenRequest) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

type GetClientTokenResponse struct {
	Message              string                       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *GetClientTokenResponse_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetClientTokenResponse) Reset()         { *m = GetClientTokenResponse{} }
func (m *GetClientTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetClientTokenResponse) ProtoMessage()    {}
func (*GetClientTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a5001b5f0223dd, []int{1}
}
func (m *GetClientTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientTokenResponse.Unmarshal(m, b)
}
func (m *GetClientTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetClientTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientTokenResponse.Merge(m, src)
}
func (m *GetClientTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetClientTokenResponse.Size(m)
}
func (m *GetClientTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientTokenResponse proto.InternalMessageInfo

func (m *GetClientTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetClientTokenResponse) GetData() *GetClientTokenResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetClientTokenResponse_Data struct {
	// 获取的 access_token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// access_token 有效时间，单位：秒
	ExpiresIn int64 `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	// error_code
	ErrorCode int64 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code"`
	// 错误表述
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClientTokenResponse_Data) Reset()         { *m = GetClientTokenResponse_Data{} }
func (m *GetClientTokenResponse_Data) String() string { return proto.CompactTextString(m) }
func (*GetClientTokenResponse_Data) ProtoMessage()    {}
func (*GetClientTokenResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a5001b5f0223dd, []int{1, 0}
}
func (m *GetClientTokenResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientTokenResponse_Data.Unmarshal(m, b)
}
func (m *GetClientTokenResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientTokenResponse_Data.Marshal(b, m, deterministic)
}
func (m *GetClientTokenResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientTokenResponse_Data.Merge(m, src)
}
func (m *GetClientTokenResponse_Data) XXX_Size() int {
	return xxx_messageInfo_GetClientTokenResponse_Data.Size(m)
}
func (m *GetClientTokenResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientTokenResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientTokenResponse_Data proto.InternalMessageInfo

func (m *GetClientTokenResponse_Data) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetClientTokenResponse_Data) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *GetClientTokenResponse_Data) GetErrorCode() int64 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetClientTokenResponse_Data) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClientTokenRequest)(nil), "douyin.GetClientTokenRequest")
	proto.RegisterType((*GetClientTokenResponse)(nil), "douyin.GetClientTokenResponse")
	proto.RegisterType((*GetClientTokenResponse_Data)(nil), "douyin.GetClientTokenResponse.Data")
}

func init() { proto.RegisterFile("api/douyin/client_token.proto", fileDescriptor_b6a5001b5f0223dd) }

var fileDescriptor_b6a5001b5f0223dd = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xb1, 0x8e, 0xd3, 0x40,
	0x14, 0x94, 0x73, 0xd1, 0x21, 0xbf, 0xa0, 0x43, 0x2c, 0x3a, 0x64, 0x59, 0x17, 0x12, 0x16, 0x21,
	0x1d, 0xc5, 0xd9, 0xe2, 0x28, 0x90, 0x8e, 0x2e, 0x77, 0x12, 0x20, 0x3a, 0x8b, 0x8a, 0x26, 0xda,
	0xd8, 0xef, 0x9c, 0x55, 0x92, 0x7d, 0xce, 0xee, 0x1a, 0x91, 0x96, 0x8a, 0x9e, 0x86, 0x6f, 0xe0,
	0x77, 0xe8, 0x53, 0x51, 0xa5, 0xe4, 0x0b, 0x90, 0x77, 0x13, 0xc5, 0x41, 0x40, 0xe5, 0x9d, 0x37,
	0x7e, 0x3b, 0x9e, 0xf1, 0x40, 0x5f, 0x54, 0x32, 0x2d, 0xa8, 0x5e, 0x49, 0x95, 0xe6, 0x73, 0x89,
	0xca, 0x8e, 0x2d, 0xcd, 0x50, 0x25, 0x95, 0x26, 0x4b, 0xec, 0xd8, 0x53, 0xf1, 0x45, 0x29, 0xed,
	0xb4, 0x9e, 0x24, 0x39, 0x2d, 0xd2, 0x92, 0x4a, 0x4a, 0x1d, 0x3d, 0xa9, 0x6f, 0x1d, 0x72, 0xc0,
	0x9d, 0xfc, 0x5a, 0x7c, 0x56, 0x12, 0x95, 0x73, 0x4c, 0x9b, 0xcb, 0x85, 0x52, 0x64, 0x85, 0x95,
	0xa4, 0x8c, 0x67, 0xb9, 0x80, 0xd3, 0xd7, 0x68, 0xaf, 0x9d, 0xda, 0xfb, 0x46, 0x2c, 0xc3, 0x65,
	0x8d, 0xc6, 0xb2, 0x37, 0x00, 0xdb, 0x6f, 0x98, 0xe1, 0x2a, 0x0a, 0x86, 0xc1, 0x79, 0x38, 0x7a,
	0xf6, 0x6b, 0x3d, 0x78, 0x7a, 0x4b, 0x7a, 0x71, 0xc5, 0xf7, 0x1c, 0x1f, 0x7e, 0x14, 0x73, 0x59,
	0x08, 0x8b, 0x57, 0x5c, 0xe3, 0xb2, 0x96, 0x1a, 0x0b, 0x9e, 0x85, 0xfe, 0x85, 0x77, 0xb8, 0xe2,
	0xdf, 0x3a, 0xf0, 0xf0, 0x4f, 0x0d, 0x53, 0x91, 0x32, 0xc8, 0x22, 0xb8, 0xb3, 0x40, 0x63, 0x44,
	0x89, 0x5e, 0x21, 0xdb, 0x41, 0xf6, 0x12, 0xba, 0x85, 0xb0, 0x22, 0xea, 0x0c, 0x83, 0xf3, 0xde,
	0xe5, 0x93, 0xc4, 0x7b, 0x4f, 0xfe, 0x7e, 0x4f, 0x72, 0x23, 0xac, 0xc8, 0xdc, 0x42, 0xfc, 0x3d,
	0x80, 0x6e, 0x03, 0xd9, 0x63, 0xb8, 0x2b, 0xf2, 0x1c, 0x8d, 0xf1, 0x21, 0x6e, 0x05, 0x7a, 0x7e,
	0xe6, 0xd6, 0x59, 0x1f, 0x00, 0x3f, 0x55, 0x52, 0xa3, 0x19, 0x4b, 0xe5, 0xa4, 0x8e, 0xb2, 0x70,
	0x3b, 0x79, 0xab, 0xd8, 0x05, 0x00, 0x6a, 0x4d, 0x7a, 0x9c, 0x53, 0x81, 0xd1, 0x51, 0x43, 0x8f,
	0x4e, 0x36, 0xeb, 0x41, 0x6b, 0x9a, 0x85, 0xee, 0x7c, 0x4d, 0x05, 0xb2, 0xe7, 0xd0, 0x2b, 0xd0,
	0xe4, 0x5a, 0x56, 0x4d, 0xc0, 0x51, 0xd7, 0x45, 0x76, 0x6f, 0xb3, 0x1e, 0xb4, 0xc7, 0x59, 0x1b,
	0x5c, 0x7e, 0x09, 0xe0, 0xfe, 0x8d, 0x73, 0xd6, 0x72, 0xc5, 0x0c, 0x9c, 0x1c, 0xfa, 0x64, 0xfd,
	0x7f, 0xf9, 0x77, 0xff, 0x2a, 0x7e, 0xf4, 0xff, 0x78, 0x38, 0xff, 0xfc, 0xe3, 0xe7, 0xd7, 0xce,
	0x19, 0x8f, 0x77, 0xed, 0x22, 0x51, 0xdb, 0xe9, 0x41, 0xc7, 0x46, 0xa7, 0x1f, 0x1e, 0x24, 0xe9,
	0xbe, 0x80, 0xaf, 0xfc, 0x63, 0x72, 0xec, 0x6a, 0xf2, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4d, 0xfc, 0xf5, 0x09, 0x9c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DouyinClientTokenClient is the client API for DouyinClientToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DouyinClientTokenClient interface {
	// client_token 用于不需要用户授权就可以调用的接口。client_token 的有效时间为 2 个小时，重复获取 client_token 后会使上次的 client_token 失效（但有 5 分钟的缓冲时间，连续多次获取 client_token 只会保留最新的两个 client_token）
	GetClientToken(ctx context.Context, in *GetClientTokenRequest, opts ...grpc.CallOption) (*GetClientTokenResponse, error)
}

type douyinClientTokenClient struct {
	cc *grpc.ClientConn
}

func NewDouyinClientTokenClient(cc *grpc.ClientConn) DouyinClientTokenClient {
	return &douyinClientTokenClient{cc}
}

func (c *douyinClientTokenClient) GetClientToken(ctx context.Context, in *GetClientTokenRequest, opts ...grpc.CallOption) (*GetClientTokenResponse, error) {
	out := new(GetClientTokenResponse)
	err := c.cc.Invoke(ctx, "/douyin.DouyinClientToken/GetClientToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DouyinClientTokenServer is the server API for DouyinClientToken service.
type DouyinClientTokenServer interface {
	// client_token 用于不需要用户授权就可以调用的接口。client_token 的有效时间为 2 个小时，重复获取 client_token 后会使上次的 client_token 失效（但有 5 分钟的缓冲时间，连续多次获取 client_token 只会保留最新的两个 client_token）
	GetClientToken(context.Context, *GetClientTokenRequest) (*GetClientTokenResponse, error)
}

// UnimplementedDouyinClientTokenServer can be embedded to have forward compatible implementations.
type UnimplementedDouyinClientTokenServer struct {
}

func (*UnimplementedDouyinClientTokenServer) GetClientToken(ctx context.Context, req *GetClientTokenRequest) (*GetClientTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientToken not implemented")
}

func RegisterDouyinClientTokenServer(s *grpc.Server, srv DouyinClientTokenServer) {
	s.RegisterService(&_DouyinClientToken_serviceDesc, srv)
}

func _DouyinClientToken_GetClientToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DouyinClientTokenServer).GetClientToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.DouyinClientToken/GetClientToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DouyinClientTokenServer).GetClientToken(ctx, req.(*GetClientTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DouyinClientToken_serviceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.DouyinClientToken",
	HandlerType: (*DouyinClientTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientToken",
			Handler:    _DouyinClientToken_GetClientToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/douyin/client_token.proto",
}
