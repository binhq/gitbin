// Code generated by protoc-gen-go.
// source: binstack/v1alpha1/binstack.proto
// DO NOT EDIT!

/*
Package binstack is a generated protocol buffer package.

It is generated from these files:
	binstack/v1alpha1/binstack.proto

It has these top-level messages:
	BinarySearch
	Binary
	DownloadInfo
*/
package binstack

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

// Format of the binary.
//
// For example archived or just plain binary format.
type DownloadInfo_Format int32

const (
	DownloadInfo_UNKNOWN DownloadInfo_Format = 0
	DownloadInfo_BINARY  DownloadInfo_Format = 1
	DownloadInfo_TARGZ   DownloadInfo_Format = 2
	DownloadInfo_ZIP     DownloadInfo_Format = 3
)

var DownloadInfo_Format_name = map[int32]string{
	0: "UNKNOWN",
	1: "BINARY",
	2: "TARGZ",
	3: "ZIP",
}
var DownloadInfo_Format_value = map[string]int32{
	"UNKNOWN": 0,
	"BINARY":  1,
	"TARGZ":   2,
	"ZIP":     3,
}

func (x DownloadInfo_Format) String() string {
	return proto.EnumName(DownloadInfo_Format_name, int32(x))
}
func (DownloadInfo_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Search criteria for a binary.
type BinarySearch struct {
	// Binary or repository name which uniquely identifies the binary.
	//
	// Examples:
	//      - Github repo name in the form of owner/repository
	//      - Vanity name in the provider
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Exact version or version constraint.
	//
	// If empty, try to find the latest version.
	//
	// Types that are valid to be assigned to Version:
	//	*BinarySearch_ExactVersion
	//	*BinarySearch_VersionConstraint
	Version isBinarySearch_Version `protobuf_oneof:"version"`
	// The name of the OS this binary is built for.
	//
	// eg. darwin, linux, windows, etc.
	Os string `protobuf:"bytes,4,opt,name=os" json:"os,omitempty"`
	// The arch this binary is built for.
	//
	// eg. i386, amd64, etc.
	Arch string `protobuf:"bytes,5,opt,name=arch" json:"arch,omitempty"`
}

func (m *BinarySearch) Reset()                    { *m = BinarySearch{} }
func (m *BinarySearch) String() string            { return proto.CompactTextString(m) }
func (*BinarySearch) ProtoMessage()               {}
func (*BinarySearch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBinarySearch_Version interface {
	isBinarySearch_Version()
}

type BinarySearch_ExactVersion struct {
	ExactVersion string `protobuf:"bytes,2,opt,name=exact_version,json=exactVersion,oneof"`
}
type BinarySearch_VersionConstraint struct {
	VersionConstraint string `protobuf:"bytes,3,opt,name=version_constraint,json=versionConstraint,oneof"`
}

func (*BinarySearch_ExactVersion) isBinarySearch_Version()      {}
func (*BinarySearch_VersionConstraint) isBinarySearch_Version() {}

func (m *BinarySearch) GetVersion() isBinarySearch_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *BinarySearch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BinarySearch) GetExactVersion() string {
	if x, ok := m.GetVersion().(*BinarySearch_ExactVersion); ok {
		return x.ExactVersion
	}
	return ""
}

func (m *BinarySearch) GetVersionConstraint() string {
	if x, ok := m.GetVersion().(*BinarySearch_VersionConstraint); ok {
		return x.VersionConstraint
	}
	return ""
}

func (m *BinarySearch) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *BinarySearch) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BinarySearch) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BinarySearch_OneofMarshaler, _BinarySearch_OneofUnmarshaler, _BinarySearch_OneofSizer, []interface{}{
		(*BinarySearch_ExactVersion)(nil),
		(*BinarySearch_VersionConstraint)(nil),
	}
}

func _BinarySearch_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BinarySearch)
	// version
	switch x := m.Version.(type) {
	case *BinarySearch_ExactVersion:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ExactVersion)
	case *BinarySearch_VersionConstraint:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.VersionConstraint)
	case nil:
	default:
		return fmt.Errorf("BinarySearch.Version has unexpected type %T", x)
	}
	return nil
}

func _BinarySearch_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BinarySearch)
	switch tag {
	case 2: // version.exact_version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Version = &BinarySearch_ExactVersion{x}
		return true, err
	case 3: // version.version_constraint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Version = &BinarySearch_VersionConstraint{x}
		return true, err
	default:
		return false, nil
	}
}

func _BinarySearch_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BinarySearch)
	// version
	switch x := m.Version.(type) {
	case *BinarySearch_ExactVersion:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ExactVersion)))
		n += len(x.ExactVersion)
	case *BinarySearch_VersionConstraint:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.VersionConstraint)))
		n += len(x.VersionConstraint)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Binary information.
type Binary struct {
	// Binary or repository name which uniquely identifies the binary.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Project description.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Project homepage.
	Homepage string `protobuf:"bytes,3,opt,name=homepage" json:"homepage,omitempty"`
	// Binary version.
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// Download information.
	DownloadInfo *DownloadInfo `protobuf:"bytes,5,opt,name=download_info,json=downloadInfo" json:"download_info,omitempty"`
}

func (m *Binary) Reset()                    { *m = Binary{} }
func (m *Binary) String() string            { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()               {}
func (*Binary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Binary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Binary) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Binary) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *Binary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Binary) GetDownloadInfo() *DownloadInfo {
	if m != nil {
		return m.DownloadInfo
	}
	return nil
}

// Binary download information.
type DownloadInfo struct {
	// URL of the binary.
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// Archive format.
	Format DownloadInfo_Format `protobuf:"varint,2,opt,name=format,enum=binhq.binstack.v1alpha1.DownloadInfo_Format" json:"format,omitempty"`
	// Path to the actual binary when extracting from an archive.
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *DownloadInfo) Reset()                    { *m = DownloadInfo{} }
func (m *DownloadInfo) String() string            { return proto.CompactTextString(m) }
func (*DownloadInfo) ProtoMessage()               {}
func (*DownloadInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DownloadInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DownloadInfo) GetFormat() DownloadInfo_Format {
	if m != nil {
		return m.Format
	}
	return DownloadInfo_UNKNOWN
}

func (m *DownloadInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*BinarySearch)(nil), "binhq.binstack.v1alpha1.BinarySearch")
	proto.RegisterType((*Binary)(nil), "binhq.binstack.v1alpha1.Binary")
	proto.RegisterType((*DownloadInfo)(nil), "binhq.binstack.v1alpha1.DownloadInfo")
	proto.RegisterEnum("binhq.binstack.v1alpha1.DownloadInfo_Format", DownloadInfo_Format_name, DownloadInfo_Format_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Binstack service

type BinstackClient interface {
	// Look for binary information
	FindBinary(ctx context.Context, in *BinarySearch, opts ...grpc.CallOption) (*Binary, error)
}

type binstackClient struct {
	cc *grpc.ClientConn
}

func NewBinstackClient(cc *grpc.ClientConn) BinstackClient {
	return &binstackClient{cc}
}

func (c *binstackClient) FindBinary(ctx context.Context, in *BinarySearch, opts ...grpc.CallOption) (*Binary, error) {
	out := new(Binary)
	err := grpc.Invoke(ctx, "/binhq.binstack.v1alpha1.Binstack/FindBinary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Binstack service

type BinstackServer interface {
	// Look for binary information
	FindBinary(context.Context, *BinarySearch) (*Binary, error)
}

func RegisterBinstackServer(s *grpc.Server, srv BinstackServer) {
	s.RegisterService(&_Binstack_serviceDesc, srv)
}

func _Binstack_FindBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinarySearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinstackServer).FindBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binhq.binstack.v1alpha1.Binstack/FindBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinstackServer).FindBinary(ctx, req.(*BinarySearch))
	}
	return interceptor(ctx, in, info, handler)
}

var _Binstack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binhq.binstack.v1alpha1.Binstack",
	HandlerType: (*BinstackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBinary",
			Handler:    _Binstack_FindBinary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "binstack/v1alpha1/binstack.proto",
}

func init() { proto.RegisterFile("binstack/v1alpha1/binstack.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xdd, 0x24, 0xbb, 0x69, 0x3b, 0xcd, 0xae, 0xc2, 0x5c, 0x88, 0xf6, 0x42, 0x15, 0x69, 0x25,
	0x0e, 0x28, 0xd5, 0x16, 0xf1, 0x01, 0x1b, 0xaa, 0x42, 0x41, 0x0a, 0xc8, 0x14, 0x10, 0xbd, 0x14,
	0x37, 0x49, 0x89, 0x45, 0x6b, 0x07, 0xc7, 0x14, 0xf8, 0x24, 0x8e, 0xfc, 0x01, 0x9f, 0x86, 0xe2,
	0x38, 0x51, 0x0e, 0x54, 0x70, 0x9b, 0x79, 0xef, 0x8d, 0xde, 0x9b, 0xb1, 0x61, 0xb2, 0x65, 0xbc,
	0x52, 0x34, 0xfd, 0x3c, 0x3d, 0xde, 0xd2, 0x7d, 0x59, 0xd0, 0xdb, 0x69, 0x8b, 0x44, 0xa5, 0x14,
	0x4a, 0xe0, 0xfd, 0x2d, 0xe3, 0xc5, 0x97, 0xa8, 0x43, 0x5b, 0x5d, 0xf8, 0xd3, 0x02, 0x2f, 0x66,
	0x9c, 0xca, 0x1f, 0x6f, 0x72, 0x2a, 0xd3, 0x02, 0x11, 0xce, 0x39, 0x3d, 0xe4, 0x81, 0x35, 0xb1,
	0x1e, 0x8e, 0x88, 0xae, 0xf1, 0x06, 0x2e, 0xf3, 0xef, 0x34, 0x55, 0x9b, 0x63, 0x2e, 0x2b, 0x26,
	0x78, 0x60, 0xd7, 0xe4, 0xf3, 0x33, 0xe2, 0x69, 0xf8, 0x5d, 0x83, 0xe2, 0x14, 0xd0, 0x08, 0x36,
	0xa9, 0xe0, 0x95, 0x92, 0x94, 0x71, 0x15, 0x38, 0x46, 0x7b, 0xcf, 0x70, 0x4f, 0x3b, 0x0a, 0xaf,
	0xc0, 0x16, 0x55, 0x70, 0xae, 0x9d, 0x6c, 0x51, 0xd5, 0xde, 0x75, 0x86, 0xe0, 0xa2, 0xf1, 0xae,
	0xeb, 0x78, 0x04, 0x03, 0x33, 0x18, 0xfe, 0xb6, 0xc0, 0x6d, 0xb2, 0xfe, 0x35, 0xe5, 0x04, 0xc6,
	0x59, 0x5e, 0xa5, 0x92, 0x95, 0xaa, 0xcb, 0x48, 0xfa, 0x10, 0x5e, 0xc3, 0xb0, 0x10, 0x87, 0xbc,
	0xa4, 0x9f, 0xf2, 0x26, 0x16, 0xe9, 0x7a, 0x0c, 0x3a, 0x1f, 0x13, 0xa8, 0x6d, 0xf1, 0x05, 0x5c,
	0x66, 0xe2, 0x1b, 0xdf, 0x0b, 0x9a, 0x6d, 0x18, 0xdf, 0x09, 0x1d, 0x6f, 0x3c, 0xbb, 0x89, 0x4e,
	0xdc, 0x34, 0x9a, 0x1b, 0xf5, 0x92, 0xef, 0x04, 0xf1, 0xb2, 0x5e, 0x17, 0xfe, 0xb2, 0xc0, 0xeb,
	0xd3, 0xe8, 0x83, 0xf3, 0x55, 0xee, 0xcd, 0x1e, 0x75, 0x89, 0x73, 0x70, 0x77, 0x42, 0x1e, 0xa8,
	0xd2, 0x1b, 0x5c, 0xcd, 0x1e, 0xfd, 0x97, 0x4f, 0xb4, 0xd0, 0x33, 0xc4, 0xcc, 0xd6, 0x07, 0x2a,
	0xa9, 0x2a, 0xcc, 0x9a, 0xba, 0x0e, 0x9f, 0x80, 0xdb, 0xa8, 0x70, 0x0c, 0x83, 0xb7, 0xc9, 0xcb,
	0xe4, 0xd5, 0xfb, 0xc4, 0x3f, 0x43, 0x00, 0x37, 0x5e, 0x26, 0x77, 0xe4, 0x83, 0x6f, 0xe1, 0x08,
	0x2e, 0x56, 0x77, 0xe4, 0xd9, 0xda, 0xb7, 0x71, 0x00, 0xce, 0x7a, 0xf9, 0xda, 0x77, 0x66, 0x1f,
	0x61, 0x18, 0x1b, 0x6f, 0x5c, 0x01, 0x2c, 0x18, 0xcf, 0xcc, 0x2b, 0x9c, 0x3e, 0x41, 0xff, 0x4b,
	0x5d, 0x3f, 0xf8, 0x87, 0x2c, 0x86, 0xf5, 0xb0, 0xe5, 0xb6, 0xae, 0xfe, 0xb0, 0x8f, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0x39, 0xd0, 0xc1, 0xd4, 0x02, 0x00, 0x00,
}
