// Code generated by protoc-gen-go. DO NOT EDIT.
// source: diffservice.proto

package diffstore

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GetDiffsRequest struct {
	Priority             int64    `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	MainDigest           string   `protobuf:"bytes,2,opt,name=mainDigest,proto3" json:"mainDigest,omitempty"`
	RightDigests         []string `protobuf:"bytes,3,rep,name=rightDigests,proto3" json:"rightDigests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffsRequest) Reset()         { *m = GetDiffsRequest{} }
func (m *GetDiffsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiffsRequest) ProtoMessage()    {}
func (*GetDiffsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{1}
}
func (m *GetDiffsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffsRequest.Unmarshal(m, b)
}
func (m *GetDiffsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffsRequest.Marshal(b, m, deterministic)
}
func (dst *GetDiffsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffsRequest.Merge(dst, src)
}
func (m *GetDiffsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiffsRequest.Size(m)
}
func (m *GetDiffsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffsRequest proto.InternalMessageInfo

func (m *GetDiffsRequest) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *GetDiffsRequest) GetMainDigest() string {
	if m != nil {
		return m.MainDigest
	}
	return ""
}

func (m *GetDiffsRequest) GetRightDigests() []string {
	if m != nil {
		return m.RightDigests
	}
	return nil
}

type GetDiffsResponse struct {
	Diffs                []byte   `protobuf:"bytes,1,opt,name=diffs,proto3" json:"diffs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffsResponse) Reset()         { *m = GetDiffsResponse{} }
func (m *GetDiffsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDiffsResponse) ProtoMessage()    {}
func (*GetDiffsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{2}
}
func (m *GetDiffsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffsResponse.Unmarshal(m, b)
}
func (m *GetDiffsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffsResponse.Marshal(b, m, deterministic)
}
func (dst *GetDiffsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffsResponse.Merge(dst, src)
}
func (m *GetDiffsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDiffsResponse.Size(m)
}
func (m *GetDiffsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffsResponse proto.InternalMessageInfo

func (m *GetDiffsResponse) GetDiffs() []byte {
	if m != nil {
		return m.Diffs
	}
	return nil
}

type PurgeDigestsRequest struct {
	Digests              []string `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
	PurgeGCS             bool     `protobuf:"varint,2,opt,name=purgeGCS,proto3" json:"purgeGCS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeDigestsRequest) Reset()         { *m = PurgeDigestsRequest{} }
func (m *PurgeDigestsRequest) String() string { return proto.CompactTextString(m) }
func (*PurgeDigestsRequest) ProtoMessage()    {}
func (*PurgeDigestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{3}
}
func (m *PurgeDigestsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeDigestsRequest.Unmarshal(m, b)
}
func (m *PurgeDigestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeDigestsRequest.Marshal(b, m, deterministic)
}
func (dst *PurgeDigestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeDigestsRequest.Merge(dst, src)
}
func (m *PurgeDigestsRequest) XXX_Size() int {
	return xxx_messageInfo_PurgeDigestsRequest.Size(m)
}
func (m *PurgeDigestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeDigestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeDigestsRequest proto.InternalMessageInfo

func (m *PurgeDigestsRequest) GetDigests() []string {
	if m != nil {
		return m.Digests
	}
	return nil
}

func (m *PurgeDigestsRequest) GetPurgeGCS() bool {
	if m != nil {
		return m.PurgeGCS
	}
	return false
}

type UnavailableDigestsResponse struct {
	DigestFailures       map[string]*DigestFailureResponse `protobuf:"bytes,1,rep,name=digestFailures,proto3" json:"digestFailures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *UnavailableDigestsResponse) Reset()         { *m = UnavailableDigestsResponse{} }
func (m *UnavailableDigestsResponse) String() string { return proto.CompactTextString(m) }
func (*UnavailableDigestsResponse) ProtoMessage()    {}
func (*UnavailableDigestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{4}
}
func (m *UnavailableDigestsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnavailableDigestsResponse.Unmarshal(m, b)
}
func (m *UnavailableDigestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnavailableDigestsResponse.Marshal(b, m, deterministic)
}
func (dst *UnavailableDigestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnavailableDigestsResponse.Merge(dst, src)
}
func (m *UnavailableDigestsResponse) XXX_Size() int {
	return xxx_messageInfo_UnavailableDigestsResponse.Size(m)
}
func (m *UnavailableDigestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnavailableDigestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnavailableDigestsResponse proto.InternalMessageInfo

func (m *UnavailableDigestsResponse) GetDigestFailures() map[string]*DigestFailureResponse {
	if m != nil {
		return m.DigestFailures
	}
	return nil
}

type WarmDigestsRequest struct {
	Priority             int64    `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	Digests              []string `protobuf:"bytes,2,rep,name=digests,proto3" json:"digests,omitempty"`
	Sync                 bool     `protobuf:"varint,3,opt,name=sync,proto3" json:"sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WarmDigestsRequest) Reset()         { *m = WarmDigestsRequest{} }
func (m *WarmDigestsRequest) String() string { return proto.CompactTextString(m) }
func (*WarmDigestsRequest) ProtoMessage()    {}
func (*WarmDigestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{5}
}
func (m *WarmDigestsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarmDigestsRequest.Unmarshal(m, b)
}
func (m *WarmDigestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarmDigestsRequest.Marshal(b, m, deterministic)
}
func (dst *WarmDigestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarmDigestsRequest.Merge(dst, src)
}
func (m *WarmDigestsRequest) XXX_Size() int {
	return xxx_messageInfo_WarmDigestsRequest.Size(m)
}
func (m *WarmDigestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WarmDigestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WarmDigestsRequest proto.InternalMessageInfo

func (m *WarmDigestsRequest) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *WarmDigestsRequest) GetDigests() []string {
	if m != nil {
		return m.Digests
	}
	return nil
}

func (m *WarmDigestsRequest) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type WarmDiffsRequest struct {
	Priority             int64    `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	LeftDigests          []string `protobuf:"bytes,2,rep,name=leftDigests,proto3" json:"leftDigests,omitempty"`
	RightDigests         []string `protobuf:"bytes,3,rep,name=rightDigests,proto3" json:"rightDigests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WarmDiffsRequest) Reset()         { *m = WarmDiffsRequest{} }
func (m *WarmDiffsRequest) String() string { return proto.CompactTextString(m) }
func (*WarmDiffsRequest) ProtoMessage()    {}
func (*WarmDiffsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{6}
}
func (m *WarmDiffsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarmDiffsRequest.Unmarshal(m, b)
}
func (m *WarmDiffsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarmDiffsRequest.Marshal(b, m, deterministic)
}
func (dst *WarmDiffsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarmDiffsRequest.Merge(dst, src)
}
func (m *WarmDiffsRequest) XXX_Size() int {
	return xxx_messageInfo_WarmDiffsRequest.Size(m)
}
func (m *WarmDiffsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WarmDiffsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WarmDiffsRequest proto.InternalMessageInfo

func (m *WarmDiffsRequest) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *WarmDiffsRequest) GetLeftDigests() []string {
	if m != nil {
		return m.LeftDigests
	}
	return nil
}

func (m *WarmDiffsRequest) GetRightDigests() []string {
	if m != nil {
		return m.RightDigests
	}
	return nil
}

type DigestFailureResponse struct {
	Digest               string   `protobuf:"bytes,1,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
	TS                   int64    `protobuf:"varint,3,opt,name=TS,proto3" json:"TS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DigestFailureResponse) Reset()         { *m = DigestFailureResponse{} }
func (m *DigestFailureResponse) String() string { return proto.CompactTextString(m) }
func (*DigestFailureResponse) ProtoMessage()    {}
func (*DigestFailureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_diffservice_d7943ba5cd994fc1, []int{7}
}
func (m *DigestFailureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DigestFailureResponse.Unmarshal(m, b)
}
func (m *DigestFailureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DigestFailureResponse.Marshal(b, m, deterministic)
}
func (dst *DigestFailureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DigestFailureResponse.Merge(dst, src)
}
func (m *DigestFailureResponse) XXX_Size() int {
	return xxx_messageInfo_DigestFailureResponse.Size(m)
}
func (m *DigestFailureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DigestFailureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DigestFailureResponse proto.InternalMessageInfo

func (m *DigestFailureResponse) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *DigestFailureResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *DigestFailureResponse) GetTS() int64 {
	if m != nil {
		return m.TS
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "diffstore.Empty")
	proto.RegisterType((*GetDiffsRequest)(nil), "diffstore.GetDiffsRequest")
	proto.RegisterType((*GetDiffsResponse)(nil), "diffstore.GetDiffsResponse")
	proto.RegisterType((*PurgeDigestsRequest)(nil), "diffstore.PurgeDigestsRequest")
	proto.RegisterType((*UnavailableDigestsResponse)(nil), "diffstore.UnavailableDigestsResponse")
	proto.RegisterMapType((map[string]*DigestFailureResponse)(nil), "diffstore.UnavailableDigestsResponse.DigestFailuresEntry")
	proto.RegisterType((*WarmDigestsRequest)(nil), "diffstore.WarmDigestsRequest")
	proto.RegisterType((*WarmDiffsRequest)(nil), "diffstore.WarmDiffsRequest")
	proto.RegisterType((*DigestFailureResponse)(nil), "diffstore.DigestFailureResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiffServiceClient is the client API for DiffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiffServiceClient interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error)
	// Same functionality as WarmDigests in the diff.DiffStore interface.
	WarmDigests(ctx context.Context, in *WarmDigestsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Same functionality as WarmDiffs in the diff.DiffStore interface.
	WarmDiffs(ctx context.Context, in *WarmDiffsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error)
	// Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Ping is used to test connection.
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type diffServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiffServiceClient(cc *grpc.ClientConn) DiffServiceClient {
	return &diffServiceClient{cc}
}

func (c *diffServiceClient) GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error) {
	out := new(GetDiffsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/GetDiffs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) WarmDigests(ctx context.Context, in *WarmDigestsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/WarmDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) WarmDiffs(ctx context.Context, in *WarmDiffsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/WarmDiffs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error) {
	out := new(UnavailableDigestsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/UnavailableDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/PurgeDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServiceServer is the server API for DiffService service.
type DiffServiceServer interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(context.Context, *GetDiffsRequest) (*GetDiffsResponse, error)
	// Same functionality as WarmDigests in the diff.DiffStore interface.
	WarmDigests(context.Context, *WarmDigestsRequest) (*Empty, error)
	// Same functionality as WarmDiffs in the diff.DiffStore interface.
	WarmDiffs(context.Context, *WarmDiffsRequest) (*Empty, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(context.Context, *Empty) (*UnavailableDigestsResponse, error)
	// Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(context.Context, *PurgeDigestsRequest) (*Empty, error)
	// Ping is used to test connection.
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterDiffServiceServer(s *grpc.Server, srv DiffServiceServer) {
	s.RegisterService(&_DiffService_serviceDesc, srv)
}

func _DiffService_GetDiffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).GetDiffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/GetDiffs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).GetDiffs(ctx, req.(*GetDiffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_WarmDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarmDigestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).WarmDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/WarmDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).WarmDigests(ctx, req.(*WarmDigestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_WarmDiffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarmDiffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).WarmDiffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/WarmDiffs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).WarmDiffs(ctx, req.(*WarmDiffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_UnavailableDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/UnavailableDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_PurgeDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeDigestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).PurgeDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/PurgeDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).PurgeDigests(ctx, req.(*PurgeDigestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "diffstore.DiffService",
	HandlerType: (*DiffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiffs",
			Handler:    _DiffService_GetDiffs_Handler,
		},
		{
			MethodName: "WarmDigests",
			Handler:    _DiffService_WarmDigests_Handler,
		},
		{
			MethodName: "WarmDiffs",
			Handler:    _DiffService_WarmDiffs_Handler,
		},
		{
			MethodName: "UnavailableDigests",
			Handler:    _DiffService_UnavailableDigests_Handler,
		},
		{
			MethodName: "PurgeDigests",
			Handler:    _DiffService_PurgeDigests_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DiffService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diffservice.proto",
}

func init() { proto.RegisterFile("diffservice.proto", fileDescriptor_diffservice_d7943ba5cd994fc1) }

var fileDescriptor_diffservice_d7943ba5cd994fc1 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xed, 0xa6, 0x8d, 0xc7, 0x51, 0x09, 0x53, 0x40, 0x96, 0x2b, 0x2a, 0x6b, 0x25, 0xa4,
	0x1c, 0x50, 0x0e, 0x41, 0x42, 0x80, 0x38, 0x20, 0x48, 0xe8, 0x81, 0x03, 0xd5, 0xa6, 0xa8, 0x37,
	0xa4, 0x6d, 0xba, 0x09, 0x2b, 0x1c, 0x3b, 0x5d, 0x6f, 0x22, 0xe5, 0x4b, 0xf8, 0x41, 0x3e, 0x04,
	0x79, 0xd7, 0x71, 0x37, 0xb1, 0xa9, 0x72, 0xf3, 0xcc, 0xce, 0xbc, 0x7d, 0xf3, 0xf6, 0x8d, 0xe1,
	0xe9, 0x9d, 0x98, 0xcd, 0x72, 0x2e, 0xd7, 0x62, 0xca, 0x07, 0x4b, 0x99, 0xa9, 0x0c, 0x7d, 0x9d,
	0x52, 0x99, 0xe4, 0xe4, 0x04, 0xda, 0xe3, 0xc5, 0x52, 0x6d, 0xc8, 0x3d, 0x3c, 0xb9, 0xe4, 0x6a,
	0x54, 0x1c, 0x50, 0x7e, 0xbf, 0xe2, 0xb9, 0xc2, 0x08, 0x3a, 0x4b, 0x29, 0x32, 0x29, 0xd4, 0x26,
	0x74, 0x62, 0xa7, 0xef, 0xd1, 0x2a, 0xc6, 0x0b, 0x80, 0x05, 0x13, 0xe9, 0x48, 0xcc, 0x79, 0xae,
	0x42, 0x37, 0x76, 0xfa, 0x3e, 0xb5, 0x32, 0x48, 0xa0, 0x2b, 0xc5, 0xfc, 0x97, 0x32, 0x61, 0x1e,
	0x7a, 0xb1, 0xd7, 0xf7, 0xe9, 0x4e, 0x8e, 0xf4, 0xa1, 0xf7, 0x70, 0x65, 0xbe, 0xcc, 0xd2, 0x9c,
	0xe3, 0x33, 0x68, 0x6b, 0x72, 0xfa, 0xc2, 0x2e, 0x35, 0x01, 0xf9, 0x06, 0x67, 0x57, 0x2b, 0x39,
	0xe7, 0x65, 0xe7, 0x96, 0x60, 0x08, 0x27, 0x77, 0x25, 0xbe, 0xa3, 0xf1, 0xb7, 0xa1, 0xa6, 0x5e,
	0x34, 0x5c, 0x7e, 0x99, 0x68, 0x72, 0x1d, 0x5a, 0xc5, 0xe4, 0xaf, 0x03, 0xd1, 0x8f, 0x94, 0xad,
	0x99, 0x48, 0xd8, 0x6d, 0xf2, 0x80, 0x59, 0x32, 0x60, 0x70, 0x6a, 0x50, 0xbe, 0x32, 0x91, 0xac,
	0x24, 0x37, 0xd8, 0xc1, 0xf0, 0xfd, 0xa0, 0x52, 0x6d, 0xf0, 0xff, 0xf6, 0xc1, 0x68, 0xa7, 0x77,
	0x9c, 0x2a, 0xb9, 0xa1, 0x7b, 0x80, 0xd1, 0x14, 0xce, 0x1a, 0xca, 0xb0, 0x07, 0xde, 0x6f, 0x6e,
	0xa4, 0xf6, 0x69, 0xf1, 0x89, 0x6f, 0xa1, 0xbd, 0x66, 0xc9, 0x8a, 0xeb, 0x19, 0x82, 0x61, 0x6c,
	0x51, 0xd8, 0x01, 0xd8, 0xde, 0x4e, 0x4d, 0xf9, 0x07, 0xf7, 0x9d, 0x43, 0x7e, 0x02, 0xde, 0x30,
	0xb9, 0xd8, 0x93, 0xec, 0xb1, 0x37, 0xb5, 0xe4, 0x74, 0x77, 0xe5, 0x44, 0x38, 0xca, 0x37, 0xe9,
	0x34, 0xf4, 0xb4, 0x94, 0xfa, 0x9b, 0x28, 0xe8, 0x19, 0xfc, 0x03, 0x1d, 0x13, 0x43, 0x90, 0xf0,
	0x59, 0x65, 0x08, 0x73, 0x83, 0x9d, 0x3a, 0xc8, 0x33, 0x37, 0xf0, 0xbc, 0x71, 0x72, 0x7c, 0x01,
	0xc7, 0xa5, 0x19, 0x8d, 0x7e, 0x65, 0x54, 0xe4, 0x29, 0x67, 0x79, 0x96, 0x96, 0x26, 0x2d, 0x23,
	0x3c, 0x05, 0xf7, 0x7a, 0xa2, 0x07, 0xf2, 0xa8, 0x7b, 0x3d, 0x19, 0xfe, 0xf1, 0x20, 0x28, 0x66,
	0x99, 0x98, 0x4d, 0xc1, 0x31, 0x74, 0xb6, 0xe6, 0xc4, 0xc8, 0xd2, 0x7d, 0x6f, 0x49, 0xa2, 0xf3,
	0xc6, 0x33, 0x43, 0x8a, 0xb4, 0xf0, 0x13, 0x04, 0xd6, 0x2b, 0xe0, 0x4b, 0xab, 0xba, 0xfe, 0x3a,
	0x51, 0xcf, 0x3a, 0x36, 0x6b, 0xd9, 0xc2, 0x8f, 0xe0, 0x57, 0x3a, 0xe3, 0x79, 0xad, 0xdf, 0xa2,
	0xd2, 0xd4, 0xfd, 0x1d, 0xb0, 0x6e, 0x56, 0xac, 0x55, 0x46, 0xaf, 0x0e, 0x72, 0x37, 0x69, 0xe1,
	0x67, 0xe8, 0xda, 0xab, 0x88, 0x17, 0x56, 0x63, 0xc3, 0x8e, 0x36, 0x92, 0x7a, 0x0d, 0x47, 0x57,
	0x22, 0x9d, 0x37, 0xd0, 0x68, 0xa8, 0xbe, 0x3d, 0xd6, 0x3f, 0xad, 0x37, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0x68, 0x23, 0x23, 0xc9, 0x04, 0x00, 0x00,
}
