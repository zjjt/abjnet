// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/prestation/prestation.proto

package prestation

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

type Prestation struct {
	Transacno            int64    `protobuf:"varint,1,opt,name=transacno,proto3" json:"transacno,omitempty"`
	Nomclient            string   `protobuf:"bytes,2,opt,name=nomclient,proto3" json:"nomclient,omitempty"`
	Prenomclient         string   `protobuf:"bytes,3,opt,name=prenomclient,proto3" json:"prenomclient,omitempty"`
	Telephone            string   `protobuf:"bytes,4,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Datedamandeuser      string   `protobuf:"bytes,5,opt,name=datedamandeuser,proto3" json:"datedamandeuser,omitempty"`
	Conventionno         int32    `protobuf:"varint,6,opt,name=conventionno,proto3" json:"conventionno,omitempty"`
	Policeno             int32    `protobuf:"varint,7,opt,name=policeno,proto3" json:"policeno,omitempty"`
	Montantdemande       int64    `protobuf:"varint,8,opt,name=montantdemande,proto3" json:"montantdemande,omitempty"`
	Montantrestant       int64    `protobuf:"varint,9,opt,name=montantrestant,proto3" json:"montantrestant,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prestation) Reset()         { *m = Prestation{} }
func (m *Prestation) String() string { return proto.CompactTextString(m) }
func (*Prestation) ProtoMessage()    {}
func (*Prestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{0}
}

func (m *Prestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prestation.Unmarshal(m, b)
}
func (m *Prestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prestation.Marshal(b, m, deterministic)
}
func (m *Prestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prestation.Merge(m, src)
}
func (m *Prestation) XXX_Size() int {
	return xxx_messageInfo_Prestation.Size(m)
}
func (m *Prestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Prestation.DiscardUnknown(m)
}

var xxx_messageInfo_Prestation proto.InternalMessageInfo

func (m *Prestation) GetTransacno() int64 {
	if m != nil {
		return m.Transacno
	}
	return 0
}

func (m *Prestation) GetNomclient() string {
	if m != nil {
		return m.Nomclient
	}
	return ""
}

func (m *Prestation) GetPrenomclient() string {
	if m != nil {
		return m.Prenomclient
	}
	return ""
}

func (m *Prestation) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *Prestation) GetDatedamandeuser() string {
	if m != nil {
		return m.Datedamandeuser
	}
	return ""
}

func (m *Prestation) GetConventionno() int32 {
	if m != nil {
		return m.Conventionno
	}
	return 0
}

func (m *Prestation) GetPoliceno() int32 {
	if m != nil {
		return m.Policeno
	}
	return 0
}

func (m *Prestation) GetMontantdemande() int64 {
	if m != nil {
		return m.Montantdemande
	}
	return 0
}

func (m *Prestation) GetMontantrestant() int64 {
	if m != nil {
		return m.Montantrestant
	}
	return 0
}

func (m *Prestation) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Request struct {
	Police               int32    `protobuf:"varint,1,opt,name=police,proto3" json:"police,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPolice() int32 {
	if m != nil {
		return m.Police
	}
	return 0
}

type Response struct {
	Done                 bool        `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	Description          string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Prestation           *Prestation `protobuf:"bytes,3,opt,name=prestation,proto3" json:"prestation,omitempty"`
	Errors               []*Error    `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Response) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Response) GetPrestation() *Prestation {
	if m != nil {
		return m.Prestation
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4002c97688aaf2, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Prestation)(nil), "prestation.Prestation")
	proto.RegisterType((*Request)(nil), "prestation.Request")
	proto.RegisterType((*Response)(nil), "prestation.Response")
	proto.RegisterType((*Error)(nil), "prestation.Error")
}

func init() {
	proto.RegisterFile("proto/prestation/prestation.proto", fileDescriptor_9d4002c97688aaf2)
}

var fileDescriptor_9d4002c97688aaf2 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0xeb, 0x24, 0x76, 0x9c, 0x49, 0x68, 0x89, 0x5a, 0x82, 0x08, 0x3d, 0x38, 0x3e, 0x14,
	0xf7, 0x92, 0x42, 0x0a, 0xa5, 0x50, 0x7a, 0xe8, 0xa1, 0xf7, 0x45, 0x0b, 0x7b, 0xd7, 0xca, 0x03,
	0x31, 0x38, 0x92, 0x57, 0x52, 0xf2, 0x10, 0xfb, 0x1a, 0xfb, 0x66, 0xfb, 0x24, 0x8b, 0x14, 0xc7,
	0x96, 0x03, 0x61, 0x6f, 0xd2, 0x37, 0xbf, 0xc6, 0xff, 0xfc, 0x83, 0x61, 0xd3, 0x68, 0x65, 0xd5,
	0x8f, 0x46, 0xa3, 0xb1, 0xdc, 0x56, 0x4a, 0x06, 0xc7, 0xad, 0xaf, 0x11, 0xe8, 0x49, 0xfe, 0x3a,
	0x02, 0xb8, 0xeb, 0xae, 0xe4, 0x2b, 0xcc, 0xac, 0xe6, 0xd2, 0x70, 0x21, 0x15, 0x8d, 0xb2, 0xa8,
	0x18, 0xb3, 0x1e, 0xb8, 0xaa, 0x54, 0x07, 0x51, 0x57, 0x28, 0x2d, 0x1d, 0x65, 0x51, 0x31, 0x63,
	0x3d, 0x20, 0x39, 0x2c, 0x1a, 0x8d, 0xbd, 0x60, 0xec, 0x05, 0x03, 0xe6, 0xfb, 0x63, 0x8d, 0xcd,
	0x5e, 0x49, 0xa4, 0x93, 0x73, 0x87, 0x0e, 0x90, 0x02, 0x3e, 0x95, 0xdc, 0x62, 0xc9, 0x0f, 0x5c,
	0x96, 0x78, 0x34, 0xa8, 0x69, 0xec, 0x35, 0xd7, 0xd8, 0x7d, 0x4b, 0x28, 0x79, 0x42, 0xe9, 0x5c,
	0x4b, 0x45, 0x93, 0x2c, 0x2a, 0x62, 0x36, 0x60, 0x64, 0x0d, 0x69, 0xa3, 0xea, 0x4a, 0xa0, 0x54,
	0x74, 0xea, 0xeb, 0xdd, 0x9d, 0x7c, 0x83, 0x8f, 0x07, 0x25, 0x2d, 0x97, 0xb6, 0x44, 0xdf, 0x95,
	0xa6, 0x7e, 0xd8, 0x2b, 0x1a, 0xe8, 0x7c, 0x46, 0xd2, 0xd2, 0xd9, 0x40, 0xd7, 0x52, 0x37, 0x97,
	0xd0, 0xe8, 0x4c, 0xfe, 0xb3, 0x14, 0xce, 0x73, 0x75, 0x20, 0xdf, 0xc0, 0x94, 0xe1, 0xd3, 0x11,
	0x8d, 0x25, 0x2b, 0x48, 0xce, 0x26, 0x7c, 0xba, 0x31, 0x6b, 0x6f, 0xf9, 0x4b, 0x04, 0x29, 0x43,
	0xd3, 0x28, 0x69, 0x90, 0x10, 0x98, 0x94, 0x2e, 0x20, 0x27, 0x49, 0x99, 0x3f, 0x93, 0x0c, 0xe6,
	0x25, 0x1a, 0xa1, 0xab, 0xc6, 0x8d, 0xd7, 0xa6, 0x1f, 0x22, 0xf2, 0x0b, 0x82, 0xc5, 0xfa, 0xf4,
	0xe7, 0xbb, 0xd5, 0x36, 0xd8, 0x7e, 0xbf, 0x67, 0x16, 0x28, 0xc9, 0x77, 0x48, 0x50, 0x6b, 0xa5,
	0x0d, 0x9d, 0x64, 0xe3, 0x62, 0xbe, 0x5b, 0x86, 0x6f, 0xfe, 0xbb, 0x0a, 0x6b, 0x05, 0xf9, 0x5f,
	0x88, 0x3d, 0x70, 0x0e, 0x85, 0x2a, 0x2f, 0x43, 0xf8, 0xf3, 0xfb, 0x0e, 0x77, 0xcf, 0x11, 0x2c,
	0x7b, 0x13, 0xf7, 0xa8, 0x4f, 0x95, 0x40, 0xf2, 0x1b, 0x12, 0xc6, 0xc5, 0x9e, 0x5b, 0x72, 0xc3,
	0xed, 0xfa, 0x4b, 0xc8, 0x2f, 0x29, 0xe5, 0x1f, 0xc8, 0x1f, 0x58, 0x3c, 0xf0, 0x1a, 0x8f, 0xba,
	0x7d, 0xff, 0x79, 0xa8, 0xf3, 0x89, 0xdf, 0x7a, 0xfc, 0x98, 0xf8, 0x9f, 0xe1, 0xe7, 0x5b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x23, 0x4d, 0xcd, 0xa1, 0x31, 0x03, 0x00, 0x00,
}
