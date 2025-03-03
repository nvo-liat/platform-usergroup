// Code generated by protoc-gen-nvo. DO NOT EDIT.
// source: protos/usergroup.proto

package protos

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Service Name
const ServiceName string = "liat.platform.usergroup"

// Api Endpoints for UsergroupService service

func NewUsergroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UsergroupService service

type UsergroupService interface {
	Show(ctx context.Context, in *ShowRequest, opts ...client.CallOption) (*UsergroupResponse, error)
}

type usergroupService struct {
	c    client.Client
	name string
}

func NewUsergroupService(c client.Client) UsergroupService {
	return &usergroupService{
		c:    c,
		name: ServiceName,
	}
}

func (c *usergroupService) Show(ctx context.Context, in *ShowRequest, opts ...client.CallOption) (*UsergroupResponse, error) {
	req := c.c.NewRequest(c.name, "UsergroupService.Show", in)
	out := new(UsergroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UsergroupService service

type UsergroupServiceHandler interface {
	Show(context.Context, *ShowRequest, *UsergroupResponse) error
}

func RegisterUsergroupServiceHandler(s server.Server, hdlr UsergroupServiceHandler, opts ...server.HandlerOption) error {
	type usergroupService interface {
		Show(ctx context.Context, in *ShowRequest, out *UsergroupResponse) error
	}
	type UsergroupService struct {
		usergroupService
	}
	h := &usergroupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UsergroupService{h}, opts...))
}

type usergroupServiceHandler struct {
	UsergroupServiceHandler
}

func (h *usergroupServiceHandler) Show(ctx context.Context, in *ShowRequest, out *UsergroupResponse) error {
	return h.UsergroupServiceHandler.Show(ctx, in, out)
}
