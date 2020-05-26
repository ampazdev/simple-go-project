// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package productproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Product struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_0d0bac2fcd735923, []int{0}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetProductsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductsRequest) Reset()         { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()    {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_0d0bac2fcd735923, []int{1}
}
func (m *GetProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsRequest.Unmarshal(m, b)
}
func (m *GetProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsRequest.Merge(dst, src)
}
func (m *GetProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductsRequest.Size(m)
}
func (m *GetProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsRequest proto.InternalMessageInfo

type GetProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProductsResponse) Reset()         { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()    {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_0d0bac2fcd735923, []int{2}
}
func (m *GetProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse.Unmarshal(m, b)
}
func (m *GetProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse.Merge(dst, src)
}
func (m *GetProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse.Size(m)
}
func (m *GetProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse proto.InternalMessageInfo

func (m *GetProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "productproto.Product")
	proto.RegisterType((*GetProductsRequest)(nil), "productproto.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "productproto.GetProductsResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_product_0d0bac2fcd735923) }

var fileDescriptor_product_0d0bac2fcd735923 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x3d, 0x0e, 0xc2, 0x30,
	0x0c, 0x85, 0x95, 0x16, 0xf1, 0xe3, 0x02, 0x83, 0x01, 0x29, 0x63, 0xd4, 0xa9, 0x53, 0x25, 0xe0,
	0x10, 0xb0, 0x81, 0x72, 0x03, 0x68, 0x3c, 0x64, 0x20, 0x09, 0x71, 0x7a, 0x7f, 0xa4, 0x10, 0xa1,
	0xb2, 0xf9, 0x7b, 0xcf, 0xb6, 0x3e, 0xd8, 0x84, 0xe8, 0xcd, 0x38, 0xa4, 0x3e, 0x44, 0x9f, 0x3c,
	0xae, 0x0b, 0x66, 0x6a, 0x6f, 0xb0, 0xb8, 0x7f, 0x19, 0xb7, 0x50, 0x59, 0x23, 0x85, 0x12, 0x5d,
	0xad, 0x2b, 0x6b, 0x10, 0x61, 0xe6, 0x1e, 0x2f, 0x92, 0x95, 0x12, 0xdd, 0x4a, 0xe7, 0x19, 0x15,
	0x34, 0x86, 0x78, 0x88, 0x36, 0x24, 0xeb, 0x9d, 0xac, 0x73, 0x35, 0x8d, 0xda, 0x3d, 0xe0, 0x85,
	0x52, 0xf9, 0xc9, 0x9a, 0xde, 0x23, 0x71, 0x6a, 0xaf, 0xb0, 0xfb, 0x4b, 0x39, 0x78, 0xc7, 0x84,
	0x47, 0x58, 0x16, 0x1b, 0x96, 0x42, 0xd5, 0x5d, 0x73, 0x3a, 0xf4, 0x53, 0xbd, 0xbe, 0x5c, 0xe8,
	0xdf, 0xda, 0x73, 0x9e, 0x8b, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xc6, 0x23, 0xd6, 0xd6,
	0x00, 0x00, 0x00,
}