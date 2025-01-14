// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pod/pod.proto

package pod

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

// Api Endpoints for Pod service

func NewPodEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Pod service

type PodService interface {
	AddPod(ctx context.Context, in *PodInfo, opts ...client.CallOption) (*Response, error)
	DeletePod(ctx context.Context, in *PodID, opts ...client.CallOption) (*Response, error)
	FindPodByID(ctx context.Context, in *PodID, opts ...client.CallOption) (*PodInfo, error)
	UpdatePod(ctx context.Context, in *PodInfo, opts ...client.CallOption) (*Response, error)
	FindAllPod(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllPod, error)
}

type podService struct {
	c    client.Client
	name string
}

func NewPodService(name string, c client.Client) PodService {
	return &podService{
		c:    c,
		name: name,
	}
}

func (c *podService) AddPod(ctx context.Context, in *PodInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pod.AddPod", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podService) DeletePod(ctx context.Context, in *PodID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pod.DeletePod", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podService) FindPodByID(ctx context.Context, in *PodID, opts ...client.CallOption) (*PodInfo, error) {
	req := c.c.NewRequest(c.name, "Pod.FindPodByID", in)
	out := new(PodInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podService) UpdatePod(ctx context.Context, in *PodInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Pod.UpdatePod", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podService) FindAllPod(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllPod, error) {
	req := c.c.NewRequest(c.name, "Pod.FindAllPod", in)
	out := new(AllPod)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pod service

type PodHandler interface {
	AddPod(context.Context, *PodInfo, *Response) error
	DeletePod(context.Context, *PodID, *Response) error
	FindPodByID(context.Context, *PodID, *PodInfo) error
	UpdatePod(context.Context, *PodInfo, *Response) error
	FindAllPod(context.Context, *FindAll, *AllPod) error
}

func RegisterPodHandler(s server.Server, hdlr PodHandler, opts ...server.HandlerOption) error {
	type pod interface {
		AddPod(ctx context.Context, in *PodInfo, out *Response) error
		DeletePod(ctx context.Context, in *PodID, out *Response) error
		FindPodByID(ctx context.Context, in *PodID, out *PodInfo) error
		UpdatePod(ctx context.Context, in *PodInfo, out *Response) error
		FindAllPod(ctx context.Context, in *FindAll, out *AllPod) error
	}
	type Pod struct {
		pod
	}
	h := &podHandler{hdlr}
	return s.Handle(s.NewHandler(&Pod{h}, opts...))
}

type podHandler struct {
	PodHandler
}

func (h *podHandler) AddPod(ctx context.Context, in *PodInfo, out *Response) error {
	return h.PodHandler.AddPod(ctx, in, out)
}

func (h *podHandler) DeletePod(ctx context.Context, in *PodID, out *Response) error {
	return h.PodHandler.DeletePod(ctx, in, out)
}

func (h *podHandler) FindPodByID(ctx context.Context, in *PodID, out *PodInfo) error {
	return h.PodHandler.FindPodByID(ctx, in, out)
}

func (h *podHandler) UpdatePod(ctx context.Context, in *PodInfo, out *Response) error {
	return h.PodHandler.UpdatePod(ctx, in, out)
}

func (h *podHandler) FindAllPod(ctx context.Context, in *FindAll, out *AllPod) error {
	return h.PodHandler.FindAllPod(ctx, in, out)
}
