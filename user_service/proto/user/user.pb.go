// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nom                  string   `protobuf:"bytes,2,opt,name=nom,proto3" json:"nom,omitempty"`
	Prenom               string   `protobuf:"bytes,3,opt,name=prenom,proto3" json:"prenom,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetNom() string {
	if m != nil {
		return m.Nom
	}
	return ""
}

func (m *User) GetPrenom() string {
	if m != nil {
		return m.Prenom
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
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

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
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

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
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
	return fileDescriptor_9b283a848145d6b7, []int{4}
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
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() {
	proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7)
}

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xf3, 0x65, 0x3b, 0xa1, 0x45, 0x06, 0x95, 0xd0, 0x83, 0xc4, 0x14, 0x44, 0x11, 0x2a,
	0xd4, 0xb3, 0x87, 0x22, 0xd2, 0x7b, 0xfc, 0xc0, 0x6b, 0x6c, 0x06, 0x5c, 0x4c, 0xb3, 0xe9, 0xee,
	0xb6, 0xfe, 0x52, 0xff, 0x8f, 0xec, 0x6c, 0x2a, 0xb5, 0x08, 0x7a, 0x49, 0xe6, 0xbd, 0x79, 0x0c,
	0xef, 0xbd, 0x04, 0x8e, 0x1b, 0x25, 0x8d, 0xbc, 0x5e, 0x6b, 0x52, 0xfc, 0x98, 0x30, 0xc6, 0xc0,
	0xce, 0x99, 0x82, 0xe0, 0x49, 0x93, 0xc2, 0x21, 0x78, 0xa2, 0x4c, 0xba, 0x69, 0xf7, 0xa2, 0x9f,
	0x7b, 0xa2, 0xc4, 0x43, 0xf0, 0x6b, 0xb9, 0x4c, 0x3c, 0x26, 0xec, 0x88, 0x27, 0x10, 0x35, 0x8a,
	0x2c, 0xe9, 0x33, 0xd9, 0x22, 0x3c, 0x82, 0x90, 0x96, 0x85, 0xa8, 0x92, 0x80, 0x69, 0x07, 0x70,
	0x04, 0xbd, 0xa6, 0xd0, 0xfa, 0x43, 0xaa, 0x32, 0x09, 0x79, 0xf1, 0x8d, 0xb3, 0x3e, 0x1c, 0xe4,
	0xb4, 0x5a, 0x93, 0x36, 0xd9, 0x0a, 0x7a, 0x39, 0xe9, 0x46, 0xd6, 0x9a, 0xf0, 0x14, 0xd8, 0x12,
	0x9b, 0x88, 0xa7, 0x30, 0x61, 0xaf, 0xd6, 0x5c, 0xce, 0x3c, 0xa6, 0x10, 0xda, 0xb7, 0x4e, 0xbc,
	0xd4, 0xdf, 0x13, 0xb8, 0x05, 0x8e, 0x21, 0x22, 0xa5, 0xa4, 0xd2, 0x89, 0xcf, 0x92, 0xd8, 0x49,
	0xee, 0x2d, 0x97, 0xb7, 0xab, 0xec, 0x05, 0xc2, 0x47, 0xf9, 0x4e, 0xb5, 0x35, 0x6e, 0xec, 0xd0,
	0xa6, 0x76, 0xc0, 0xb2, 0x9b, 0xa2, 0x12, 0x25, 0x47, 0xef, 0xe5, 0x0e, 0xfc, 0xef, 0xf2, 0x2d,
	0x84, 0x4c, 0x20, 0x42, 0xb0, 0x90, 0x25, 0xf1, 0xe1, 0x30, 0xe7, 0x19, 0x53, 0x88, 0x4b, 0xd2,
	0x0b, 0x25, 0x1a, 0x23, 0x64, 0xdd, 0x16, 0xbb, 0x4b, 0x4d, 0x3f, 0xbb, 0x10, 0xdb, 0x34, 0x0f,
	0xa4, 0x36, 0x62, 0x41, 0x78, 0x0e, 0xd1, 0x9d, 0xa2, 0xc2, 0x10, 0xee, 0x44, 0x1d, 0x0d, 0xdd,
	0xbc, 0x6d, 0x2d, 0xeb, 0xe0, 0x18, 0xfc, 0x39, 0x99, 0x3f, 0x44, 0x97, 0x10, 0xcd, 0xc9, 0xcc,
	0xaa, 0x0a, 0x07, 0xdb, 0x1d, 0x7f, 0x81, 0x5f, 0xa4, 0x67, 0x10, 0xcc, 0xd6, 0xe6, 0xed, 0xc7,
	0xc1, 0x36, 0x2f, 0x17, 0x97, 0x75, 0xf0, 0x0a, 0x06, 0xcf, 0xb6, 0x97, 0xc2, 0x90, 0xeb, 0x72,
	0x77, 0xbf, 0x27, 0x7e, 0x8d, 0xf8, 0x7f, 0xbb, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xf6,
	0x1a, 0xff, 0x88, 0x02, 0x00, 0x00,
}
