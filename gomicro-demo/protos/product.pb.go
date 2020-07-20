// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ProdModel struct {
	// @inject_tag:json:"id"
	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag:json:"name"
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ProdModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProdsRequest struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdsRequest) Reset()         { *m = ProdsRequest{} }
func (m *ProdsRequest) String() string { return proto.CompactTextString(m) }
func (*ProdsRequest) ProtoMessage()    {}
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdsRequest.Unmarshal(m, b)
}
func (m *ProdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdsRequest.Marshal(b, m, deterministic)
}
func (m *ProdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdsRequest.Merge(m, src)
}
func (m *ProdsRequest) XXX_Size() int {
	return xxx_messageInfo_ProdsRequest.Size(m)
}
func (m *ProdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdsRequest proto.InternalMessageInfo

func (m *ProdsRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ProdListResponse struct {
	Data                 []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProdListResponse) Reset()         { *m = ProdListResponse{} }
func (m *ProdListResponse) String() string { return proto.CompactTextString(m) }
func (*ProdListResponse) ProtoMessage()    {}
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *ProdListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdListResponse.Unmarshal(m, b)
}
func (m *ProdListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdListResponse.Marshal(b, m, deterministic)
}
func (m *ProdListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdListResponse.Merge(m, src)
}
func (m *ProdListResponse) XXX_Size() int {
	return xxx_messageInfo_ProdListResponse.Size(m)
}
func (m *ProdListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdListResponse proto.InternalMessageInfo

func (m *ProdListResponse) GetData() []*ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "protos.ProdModel")
	proto.RegisterType((*ProdsRequest)(nil), "protos.ProdsRequest")
	proto.RegisterType((*ProdListResponse)(nil), "protos.ProdListResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xfa,
	0x5c, 0x9c, 0x01, 0x45, 0xf9, 0x29, 0xbe, 0xf9, 0x29, 0xa9, 0x39, 0x42, 0x7c, 0x5c, 0x4c, 0x9e,
	0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x7e,
	0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x12, 0x17, 0x0f,
	0x48, 0x43, 0x71, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0x48, 0x4d, 0x71, 0x66, 0x55, 0x2a,
	0x54, 0x17, 0x98, 0xad, 0x64, 0xc9, 0x25, 0x00, 0x52, 0xe3, 0x93, 0x59, 0x5c, 0x12, 0x94, 0x5a,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xca, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0x24, 0x08, 0x71, 0x46, 0xb1, 0x1e, 0xdc, 0xf2, 0x20, 0xb0, 0xb4, 0x51,
	0x20, 0x17, 0x5f, 0x00, 0xc4, 0xa1, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xf6, 0x5c,
	0xdc, 0xee, 0xa9, 0x25, 0x30, 0xf3, 0x84, 0x44, 0x90, 0x75, 0xc2, 0x5c, 0x21, 0x25, 0x81, 0x2c,
	0x8a, 0x6c, 0xaf, 0x12, 0x43, 0x12, 0xc4, 0xab, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x58, 0x9e, 0x71, 0x02, 0x01, 0x00, 0x00,
}
