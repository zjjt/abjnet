// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/payment/payment.proto

package payment

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

type Payment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Transacno            string   `protobuf:"bytes,2,opt,name=transacno,proto3" json:"transacno,omitempty"`
	Nomclient            string   `protobuf:"bytes,3,opt,name=nomclient,proto3" json:"nomclient,omitempty"`
	Prenomclient         string   `protobuf:"bytes,4,opt,name=prenomclient,proto3" json:"prenomclient,omitempty"`
	Telephone            string   `protobuf:"bytes,5,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Datepaymentuser      string   `protobuf:"bytes,6,opt,name=datepaymentuser,proto3" json:"datepaymentuser,omitempty"`
	Conventionno         string   `protobuf:"bytes,7,opt,name=conventionno,proto3" json:"conventionno,omitempty"`
	Policeno             string   `protobuf:"bytes,8,opt,name=policeno,proto3" json:"policeno,omitempty"`
	Montant              string   `protobuf:"bytes,9,opt,name=montant,proto3" json:"montant,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ce6d11f0608d17, []int{0}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Payment) GetTransacno() string {
	if m != nil {
		return m.Transacno
	}
	return ""
}

func (m *Payment) GetNomclient() string {
	if m != nil {
		return m.Nomclient
	}
	return ""
}

func (m *Payment) GetPrenomclient() string {
	if m != nil {
		return m.Prenomclient
	}
	return ""
}

func (m *Payment) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *Payment) GetDatepaymentuser() string {
	if m != nil {
		return m.Datepaymentuser
	}
	return ""
}

func (m *Payment) GetConventionno() string {
	if m != nil {
		return m.Conventionno
	}
	return ""
}

func (m *Payment) GetPoliceno() string {
	if m != nil {
		return m.Policeno
	}
	return ""
}

func (m *Payment) GetMontant() string {
	if m != nil {
		return m.Montant
	}
	return ""
}

func (m *Payment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Response struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Payment              *Payment `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
	Errors               []*Error `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ce6d11f0608d17, []int{1}
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

func (m *Response) GetPayment() *Payment {
	if m != nil {
		return m.Payment
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
	return fileDescriptor_78ce6d11f0608d17, []int{2}
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
	proto.RegisterType((*Payment)(nil), "payment.Payment")
	proto.RegisterType((*Response)(nil), "payment.Response")
	proto.RegisterType((*Error)(nil), "payment.Error")
}

func init() {
	proto.RegisterFile("proto/payment/payment.proto", fileDescriptor_78ce6d11f0608d17)
}

var fileDescriptor_78ce6d11f0608d17 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xe3, 0x40,
	0x0c, 0xc6, 0xb7, 0xe9, 0x9f, 0xa4, 0xee, 0xaa, 0xbb, 0xeb, 0xd3, 0xa8, 0xcb, 0xa1, 0xca, 0x01,
	0x55, 0x08, 0x15, 0xa9, 0x9c, 0x41, 0xe2, 0xc0, 0xbd, 0x0a, 0x4f, 0x10, 0x26, 0x96, 0x88, 0xd4,
	0xda, 0xd1, 0x64, 0xa8, 0xd4, 0xf7, 0xe0, 0x65, 0x78, 0x3b, 0x34, 0x93, 0x99, 0xb4, 0xc0, 0x81,
	0x53, 0xed, 0xdf, 0xf7, 0x69, 0xfc, 0xd5, 0x0e, 0xfc, 0x6f, 0x8c, 0x58, 0xb9, 0x69, 0xca, 0xe3,
	0x9e, 0xd8, 0xc6, 0xdf, 0xb5, 0xa7, 0x98, 0x86, 0x36, 0x7f, 0x4f, 0x20, 0xdd, 0x76, 0x35, 0xce,
	0x21, 0xa9, 0x2b, 0x35, 0x58, 0x0e, 0x56, 0xd3, 0x22, 0xa9, 0x2b, 0xbc, 0x80, 0xa9, 0x35, 0x25,
	0xb7, 0xa5, 0x66, 0x51, 0x89, 0xc7, 0x27, 0xe0, 0x54, 0x96, 0xbd, 0xde, 0xd5, 0xc4, 0x56, 0x0d,
	0x3b, 0xb5, 0x07, 0x98, 0xc3, 0xef, 0xc6, 0xd0, 0xc9, 0x30, 0xf2, 0x86, 0x4f, 0xcc, 0xbf, 0x4f,
	0x3b, 0x6a, 0x5e, 0x84, 0x49, 0x8d, 0xc3, 0xfb, 0x11, 0xe0, 0x0a, 0xfe, 0x54, 0xa5, 0xa5, 0x10,
	0xf4, 0xb5, 0x25, 0xa3, 0x26, 0xde, 0xf3, 0x15, 0xbb, 0x59, 0x5a, 0xf8, 0x40, 0x6c, 0x6b, 0x61,
	0x16, 0x95, 0x76, 0xb3, 0xce, 0x19, 0x2e, 0x20, 0x6b, 0x64, 0x57, 0x6b, 0x62, 0x51, 0x99, 0xd7,
	0xfb, 0x1e, 0x15, 0xa4, 0x7b, 0x61, 0x5b, 0xb2, 0x55, 0x53, 0x2f, 0xc5, 0xd6, 0x25, 0xd4, 0x86,
	0x4a, 0x4b, 0xd5, 0x83, 0x55, 0xd0, 0x25, 0xec, 0x41, 0xfe, 0x36, 0x80, 0xac, 0xa0, 0xb6, 0x11,
	0x6e, 0x09, 0x11, 0x46, 0x95, 0xfb, 0x1f, 0x6e, 0x7d, 0x59, 0xe1, 0x6b, 0x5c, 0xc2, 0xac, 0xa2,
	0x56, 0x9b, 0xba, 0x71, 0x29, 0xc2, 0x0a, 0xcf, 0x11, 0x5e, 0x41, 0xbc, 0x84, 0x5f, 0xe1, 0x6c,
	0xf3, 0x77, 0x1d, 0x0f, 0x15, 0xae, 0x52, 0x44, 0x03, 0x5e, 0xc2, 0x84, 0x8c, 0x11, 0xd3, 0xaa,
	0xd1, 0x72, 0xb8, 0x9a, 0x6d, 0xe6, 0xbd, 0xf5, 0xd1, 0xe1, 0x22, 0xa8, 0xf9, 0x1d, 0x8c, 0x3d,
	0x70, 0x91, 0xb4, 0x54, 0x5d, 0xa4, 0x71, 0xe1, 0xeb, 0x9f, 0x23, 0x6d, 0xee, 0x61, 0x1e, 0x46,
	0x3f, 0x91, 0x39, 0xd4, 0x9a, 0xf0, 0x1a, 0x86, 0xdb, 0xf2, 0x88, 0xdf, 0xa2, 0x2d, 0xfe, 0xf5,
	0x24, 0xae, 0x21, 0xff, 0xf5, 0x3c, 0xf1, 0x5f, 0xd8, 0xed, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x05, 0xee, 0x69, 0x80, 0x02, 0x00, 0x00,
}
