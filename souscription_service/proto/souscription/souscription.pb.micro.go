// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/souscription/souscription.proto

package souscription

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SouscriptionService service

type SouscriptionService interface {
	Subscribe(ctx context.Context, in *Souscription, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type souscriptionService struct {
	c    client.Client
	name string
}

func NewSouscriptionService(name string, c client.Client) SouscriptionService {
	return &souscriptionService{
		c:    c,
		name: name,
	}
}

func (c *souscriptionService) Subscribe(ctx context.Context, in *Souscription, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SouscriptionService.Subscribe", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souscriptionService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SouscriptionService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souscriptionService) DeleteAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SouscriptionService.DeleteAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SouscriptionService service

type SouscriptionServiceHandler interface {
	Subscribe(context.Context, *Souscription, *Response) error
	GetAll(context.Context, *Request, *Response) error
	DeleteAll(context.Context, *Request, *Response) error
}

func RegisterSouscriptionServiceHandler(s server.Server, hdlr SouscriptionServiceHandler, opts ...server.HandlerOption) error {
	type souscriptionService interface {
		Subscribe(ctx context.Context, in *Souscription, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
		DeleteAll(ctx context.Context, in *Request, out *Response) error
	}
	type SouscriptionService struct {
		souscriptionService
	}
	h := &souscriptionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SouscriptionService{h}, opts...))
}

type souscriptionServiceHandler struct {
	SouscriptionServiceHandler
}

func (h *souscriptionServiceHandler) Subscribe(ctx context.Context, in *Souscription, out *Response) error {
	return h.SouscriptionServiceHandler.Subscribe(ctx, in, out)
}

func (h *souscriptionServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.SouscriptionServiceHandler.GetAll(ctx, in, out)
}

func (h *souscriptionServiceHandler) DeleteAll(ctx context.Context, in *Request, out *Response) error {
	return h.SouscriptionServiceHandler.DeleteAll(ctx, in, out)
}