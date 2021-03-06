// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buku.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	buku.proto

It has these top-level messages:
	AddBukuReq
	ReadBukuByKodeReq
	ReadBukuByKodeResp
	ReadBukuResp
	UpdateBukuReq
	ReadBukuByKeteranganReq
	ReadBukuByKeteranganResp
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddBukuReq struct {
	KodeBuku   string `protobuf:"bytes,1,opt,name=KodeBuku" json:"KodeBuku,omitempty"`
	NamaBuku   string `protobuf:"bytes,2,opt,name=NamaBuku" json:"NamaBuku,omitempty"`
	Status     int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	Keterangan string `protobuf:"bytes,4,opt,name=Keterangan" json:"Keterangan,omitempty"`
	CreateBy   string `protobuf:"bytes,5,opt,name=CreateBy" json:"CreateBy,omitempty"`
	CretateOn  string `protobuf:"bytes,6,opt,name=CretateOn" json:"CretateOn,omitempty"`
}

func (m *AddBukuReq) Reset()                    { *m = AddBukuReq{} }
func (m *AddBukuReq) String() string            { return proto.CompactTextString(m) }
func (*AddBukuReq) ProtoMessage()               {}
func (*AddBukuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddBukuReq) GetKodeBuku() string {
	if m != nil {
		return m.KodeBuku
	}
	return ""
}

func (m *AddBukuReq) GetNamaBuku() string {
	if m != nil {
		return m.NamaBuku
	}
	return ""
}

func (m *AddBukuReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddBukuReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func (m *AddBukuReq) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

func (m *AddBukuReq) GetCretateOn() string {
	if m != nil {
		return m.CretateOn
	}
	return ""
}

type ReadBukuByKodeReq struct {
	KodeBuku string `protobuf:"bytes,1,opt,name=KodeBuku" json:"KodeBuku,omitempty"`
}

func (m *ReadBukuByKodeReq) Reset()                    { *m = ReadBukuByKodeReq{} }
func (m *ReadBukuByKodeReq) String() string            { return proto.CompactTextString(m) }
func (*ReadBukuByKodeReq) ProtoMessage()               {}
func (*ReadBukuByKodeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadBukuByKodeReq) GetKodeBuku() string {
	if m != nil {
		return m.KodeBuku
	}
	return ""
}

type ReadBukuByKodeResp struct {
	KodeBuku   string `protobuf:"bytes,1,opt,name=KodeBuku" json:"KodeBuku,omitempty"`
	NamaBuku   string `protobuf:"bytes,2,opt,name=NamaBuku" json:"NamaBuku,omitempty"`
	Status     int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	Keterangan string `protobuf:"bytes,4,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *ReadBukuByKodeResp) Reset()                    { *m = ReadBukuByKodeResp{} }
func (m *ReadBukuByKodeResp) String() string            { return proto.CompactTextString(m) }
func (*ReadBukuByKodeResp) ProtoMessage()               {}
func (*ReadBukuByKodeResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadBukuByKodeResp) GetKodeBuku() string {
	if m != nil {
		return m.KodeBuku
	}
	return ""
}

func (m *ReadBukuByKodeResp) GetNamaBuku() string {
	if m != nil {
		return m.NamaBuku
	}
	return ""
}

func (m *ReadBukuByKodeResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadBukuByKodeResp) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadBukuResp struct {
	AllBuku []*ReadBukuByKodeResp `protobuf:"bytes,1,rep,name=AllBuku" json:"AllBuku,omitempty"`
}

func (m *ReadBukuResp) Reset()                    { *m = ReadBukuResp{} }
func (m *ReadBukuResp) String() string            { return proto.CompactTextString(m) }
func (*ReadBukuResp) ProtoMessage()               {}
func (*ReadBukuResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadBukuResp) GetAllBuku() []*ReadBukuByKodeResp {
	if m != nil {
		return m.AllBuku
	}
	return nil
}

type UpdateBukuReq struct {
	KodeBuku   string `protobuf:"bytes,1,opt,name=KodeBuku" json:"KodeBuku,omitempty"`
	NamaBuku   string `protobuf:"bytes,2,opt,name=NamaBuku" json:"NamaBuku,omitempty"`
	Status     int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	Keterangan string `protobuf:"bytes,4,opt,name=Keterangan" json:"Keterangan,omitempty"`
	UpdateBy   string `protobuf:"bytes,5,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
}

func (m *UpdateBukuReq) Reset()                    { *m = UpdateBukuReq{} }
func (m *UpdateBukuReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateBukuReq) ProtoMessage()               {}
func (*UpdateBukuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateBukuReq) GetKodeBuku() string {
	if m != nil {
		return m.KodeBuku
	}
	return ""
}

func (m *UpdateBukuReq) GetNamaBuku() string {
	if m != nil {
		return m.NamaBuku
	}
	return ""
}

func (m *UpdateBukuReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateBukuReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func (m *UpdateBukuReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type ReadBukuByKeteranganReq struct {
	Keterangan string `protobuf:"bytes,1,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *ReadBukuByKeteranganReq) Reset()                    { *m = ReadBukuByKeteranganReq{} }
func (m *ReadBukuByKeteranganReq) String() string            { return proto.CompactTextString(m) }
func (*ReadBukuByKeteranganReq) ProtoMessage()               {}
func (*ReadBukuByKeteranganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadBukuByKeteranganReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadBukuByKeteranganResp struct {
	AllBukuKet []*ReadBukuByKodeResp `protobuf:"bytes,1,rep,name=AllBukuKet" json:"AllBukuKet,omitempty"`
}

func (m *ReadBukuByKeteranganResp) Reset()                    { *m = ReadBukuByKeteranganResp{} }
func (m *ReadBukuByKeteranganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadBukuByKeteranganResp) ProtoMessage()               {}
func (*ReadBukuByKeteranganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadBukuByKeteranganResp) GetAllBukuKet() []*ReadBukuByKodeResp {
	if m != nil {
		return m.AllBukuKet
	}
	return nil
}

func init() {
	proto.RegisterType((*AddBukuReq)(nil), "grpc.AddBukuReq")
	proto.RegisterType((*ReadBukuByKodeReq)(nil), "grpc.ReadBukuByKodeReq")
	proto.RegisterType((*ReadBukuByKodeResp)(nil), "grpc.ReadBukuByKodeResp")
	proto.RegisterType((*ReadBukuResp)(nil), "grpc.ReadBukuResp")
	proto.RegisterType((*UpdateBukuReq)(nil), "grpc.UpdateBukuReq")
	proto.RegisterType((*ReadBukuByKeteranganReq)(nil), "grpc.ReadBukuByKeteranganReq")
	proto.RegisterType((*ReadBukuByKeteranganResp)(nil), "grpc.ReadBukuByKeteranganResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for BukuService service

type BukuServiceClient interface {
	AddBuku(ctx context.Context, in *AddBukuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadBukuByKode(ctx context.Context, in *ReadBukuByKodeReq, opts ...grpc1.CallOption) (*ReadBukuByKodeResp, error)
	ReadBuku(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadBukuResp, error)
	UpdateBuku(ctx context.Context, in *UpdateBukuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadBukuByKeterangan(ctx context.Context, in *ReadBukuByKeteranganReq, opts ...grpc1.CallOption) (*ReadBukuByKeteranganResp, error)
}

type bukuServiceClient struct {
	cc *grpc1.ClientConn
}

func NewBukuServiceClient(cc *grpc1.ClientConn) BukuServiceClient {
	return &bukuServiceClient{cc}
}

func (c *bukuServiceClient) AddBuku(ctx context.Context, in *AddBukuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.BukuService/AddBuku", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bukuServiceClient) ReadBukuByKode(ctx context.Context, in *ReadBukuByKodeReq, opts ...grpc1.CallOption) (*ReadBukuByKodeResp, error) {
	out := new(ReadBukuByKodeResp)
	err := grpc1.Invoke(ctx, "/grpc.BukuService/ReadBukuByKode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bukuServiceClient) ReadBuku(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadBukuResp, error) {
	out := new(ReadBukuResp)
	err := grpc1.Invoke(ctx, "/grpc.BukuService/ReadBuku", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bukuServiceClient) UpdateBuku(ctx context.Context, in *UpdateBukuReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.BukuService/UpdateBuku", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bukuServiceClient) ReadBukuByKeterangan(ctx context.Context, in *ReadBukuByKeteranganReq, opts ...grpc1.CallOption) (*ReadBukuByKeteranganResp, error) {
	out := new(ReadBukuByKeteranganResp)
	err := grpc1.Invoke(ctx, "/grpc.BukuService/ReadBukuByKeterangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BukuService service

type BukuServiceServer interface {
	AddBuku(context.Context, *AddBukuReq) (*google_protobuf.Empty, error)
	ReadBukuByKode(context.Context, *ReadBukuByKodeReq) (*ReadBukuByKodeResp, error)
	ReadBuku(context.Context, *google_protobuf.Empty) (*ReadBukuResp, error)
	UpdateBuku(context.Context, *UpdateBukuReq) (*google_protobuf.Empty, error)
	ReadBukuByKeterangan(context.Context, *ReadBukuByKeteranganReq) (*ReadBukuByKeteranganResp, error)
}

func RegisterBukuServiceServer(s *grpc1.Server, srv BukuServiceServer) {
	s.RegisterService(&_BukuService_serviceDesc, srv)
}

func _BukuService_AddBuku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBukuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BukuServiceServer).AddBuku(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BukuService/AddBuku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BukuServiceServer).AddBuku(ctx, req.(*AddBukuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BukuService_ReadBukuByKode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBukuByKodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BukuServiceServer).ReadBukuByKode(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BukuService/ReadBukuByKode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BukuServiceServer).ReadBukuByKode(ctx, req.(*ReadBukuByKodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BukuService_ReadBuku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BukuServiceServer).ReadBuku(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BukuService/ReadBuku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BukuServiceServer).ReadBuku(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BukuService_UpdateBuku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBukuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BukuServiceServer).UpdateBuku(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BukuService/UpdateBuku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BukuServiceServer).UpdateBuku(ctx, req.(*UpdateBukuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BukuService_ReadBukuByKeterangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBukuByKeteranganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BukuServiceServer).ReadBukuByKeterangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BukuService/ReadBukuByKeterangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BukuServiceServer).ReadBukuByKeterangan(ctx, req.(*ReadBukuByKeteranganReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BukuService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.BukuService",
	HandlerType: (*BukuServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddBuku",
			Handler:    _BukuService_AddBuku_Handler,
		},
		{
			MethodName: "ReadBukuByKode",
			Handler:    _BukuService_ReadBukuByKode_Handler,
		},
		{
			MethodName: "ReadBuku",
			Handler:    _BukuService_ReadBuku_Handler,
		},
		{
			MethodName: "UpdateBuku",
			Handler:    _BukuService_UpdateBuku_Handler,
		},
		{
			MethodName: "ReadBukuByKeterangan",
			Handler:    _BukuService_ReadBukuByKeterangan_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "buku.proto",
}

func init() { proto.RegisterFile("buku.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0x29, 0xdf, 0x0c, 0x6a, 0x74, 0x34, 0xd0, 0x54, 0x25, 0x64, 0x4f, 0x9c, 0x4a, 0x82,
	0x31, 0xc1, 0x78, 0xa2, 0x84, 0x13, 0x89, 0x26, 0x45, 0xe3, 0x79, 0xa1, 0x6b, 0x63, 0xf8, 0x68,
	0x29, 0x5b, 0x13, 0xee, 0xbe, 0x85, 0xef, 0xe1, 0xcb, 0xf8, 0x32, 0x66, 0xdb, 0x2e, 0xb4, 0xf2,
	0x75, 0xd3, 0x5b, 0xe7, 0xf3, 0xff, 0xeb, 0xcc, 0x0e, 0xc0, 0xd0, 0x1f, 0xfb, 0xba, 0xeb, 0x39,
	0xdc, 0xc1, 0xac, 0xed, 0xb9, 0x23, 0xed, 0xd2, 0x76, 0x1c, 0x7b, 0xc2, 0x9a, 0x81, 0x6f, 0xe8,
	0xbf, 0x36, 0xd9, 0xd4, 0xe5, 0xcb, 0x30, 0x85, 0x7c, 0x29, 0x00, 0x1d, 0xcb, 0x32, 0xfc, 0xb1,
	0x6f, 0xb2, 0x39, 0x6a, 0x50, 0xec, 0x3b, 0x16, 0x13, 0xa6, 0xaa, 0xd4, 0x95, 0x46, 0xc9, 0x5c,
	0xd9, 0x22, 0xf6, 0x40, 0xa7, 0x34, 0x88, 0xa5, 0xc3, 0x98, 0xb4, 0xb1, 0x02, 0xf9, 0x01, 0xa7,
	0xdc, 0x5f, 0xa8, 0x99, 0xba, 0xd2, 0xc8, 0x99, 0x91, 0x85, 0x35, 0x80, 0x3e, 0xe3, 0xcc, 0xa3,
	0x33, 0x9b, 0xce, 0xd4, 0x6c, 0x50, 0x15, 0xf3, 0x88, 0x9e, 0x5d, 0x8f, 0x51, 0xce, 0x8c, 0xa5,
	0x9a, 0x0b, 0x7b, 0x4a, 0x1b, 0xaf, 0xa0, 0xd4, 0xf5, 0x18, 0xa7, 0x9c, 0x3d, 0xce, 0xd4, 0x7c,
	0x10, 0x5c, 0x3b, 0x48, 0x13, 0xce, 0x4c, 0x46, 0x03, 0x70, 0x63, 0x29, 0x18, 0x0f, 0xe0, 0x93,
	0x0f, 0x05, 0xf0, 0x77, 0xc5, 0xc2, 0xfd, 0xeb, 0x3f, 0x26, 0x06, 0x1c, 0x49, 0x8a, 0x40, 0xbf,
	0x05, 0x85, 0xce, 0x64, 0x12, 0xc9, 0x67, 0x1a, 0xe5, 0x96, 0xaa, 0x8b, 0xad, 0xe9, 0x9b, 0xa8,
	0xa6, 0x4c, 0x24, 0x9f, 0x0a, 0x1c, 0x3f, 0xbb, 0x96, 0x18, 0xd3, 0xff, 0xed, 0x2d, 0x02, 0x58,
	0xed, 0x4d, 0xda, 0xe4, 0x0e, 0xaa, 0x31, 0xf8, 0x55, 0x8d, 0xc0, 0x4c, 0xb6, 0x55, 0x36, 0x86,
	0xf3, 0x04, 0xea, 0xf6, 0xd2, 0x85, 0x8b, 0x6d, 0x80, 0xe8, 0xff, 0xfb, 0x8c, 0x1f, 0x9c, 0x55,
	0x2c, 0xb7, 0xf5, 0x9d, 0x86, 0xb2, 0xf8, 0x1e, 0x30, 0xef, 0xfd, 0x6d, 0xc4, 0xf0, 0x16, 0x0a,
	0xd1, 0x93, 0xc7, 0xd3, 0xb0, 0xc1, 0xfa, 0x02, 0xb4, 0x8a, 0x1e, 0x9e, 0x8b, 0x2e, 0xcf, 0x45,
	0xef, 0x89, 0x73, 0x21, 0x29, 0xec, 0xc1, 0x49, 0x52, 0x08, 0xab, 0xdb, 0xe5, 0xe7, 0xda, 0x4e,
	0x2e, 0x92, 0xc2, 0x36, 0x14, 0xa5, 0x1f, 0x77, 0x88, 0x69, 0x98, 0xac, 0x8f, 0x2a, 0xef, 0x01,
	0xd6, 0x5b, 0xc7, 0xf3, 0x30, 0x27, 0xf1, 0x0e, 0xf6, 0xd0, 0xbf, 0xc0, 0xc5, 0xb6, 0xd1, 0xe2,
	0xf5, 0x06, 0x6a, 0x7c, 0x63, 0x5a, 0x6d, 0x5f, 0x58, 0x50, 0x0d, 0xf3, 0x81, 0xd4, 0xcd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xd4, 0x57, 0x92, 0x79, 0x04, 0x00, 0x00,
}
