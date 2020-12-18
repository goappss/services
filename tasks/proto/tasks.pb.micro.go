// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tasks.proto

package tasks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Tasks service

func NewTasksEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Tasks service

type TasksService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...client.CallOption) (*CancelResponse, error)
	Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*CompleteResponse, error)
	Defer(ctx context.Context, in *DeferRequest, opts ...client.CallOption) (*DeferResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...client.CallOption) (*NextResponse, error)
	Unassign(ctx context.Context, in *UnassignRequest, opts ...client.CallOption) (*UnassignResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error)
}

type tasksService struct {
	c    client.Client
	name string
}

func NewTasksService(name string, c client.Client) TasksService {
	return &tasksService{
		c:    c,
		name: name,
	}
}

func (c *tasksService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Cancel(ctx context.Context, in *CancelRequest, opts ...client.CallOption) (*CancelResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Cancel", in)
	out := new(CancelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*CompleteResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Complete", in)
	out := new(CompleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Defer(ctx context.Context, in *DeferRequest, opts ...client.CallOption) (*DeferResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Defer", in)
	out := new(DeferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Next(ctx context.Context, in *NextRequest, opts ...client.CallOption) (*NextResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Next", in)
	out := new(NextResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Unassign(ctx context.Context, in *UnassignRequest, opts ...client.CallOption) (*UnassignResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Unassign", in)
	out := new(UnassignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error) {
	req := c.c.NewRequest(c.name, "Tasks.Remove", in)
	out := new(RemoveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Cancel(context.Context, *CancelRequest, *CancelResponse) error
	Complete(context.Context, *CompleteRequest, *CompleteResponse) error
	Defer(context.Context, *DeferRequest, *DeferResponse) error
	Next(context.Context, *NextRequest, *NextResponse) error
	Unassign(context.Context, *UnassignRequest, *UnassignResponse) error
	Remove(context.Context, *RemoveRequest, *RemoveResponse) error
}

func RegisterTasksHandler(s server.Server, hdlr TasksHandler, opts ...server.HandlerOption) error {
	type tasks interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Cancel(ctx context.Context, in *CancelRequest, out *CancelResponse) error
		Complete(ctx context.Context, in *CompleteRequest, out *CompleteResponse) error
		Defer(ctx context.Context, in *DeferRequest, out *DeferResponse) error
		Next(ctx context.Context, in *NextRequest, out *NextResponse) error
		Unassign(ctx context.Context, in *UnassignRequest, out *UnassignResponse) error
		Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error
	}
	type Tasks struct {
		tasks
	}
	h := &tasksHandler{hdlr}
	return s.Handle(s.NewHandler(&Tasks{h}, opts...))
}

type tasksHandler struct {
	TasksHandler
}

func (h *tasksHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.TasksHandler.Create(ctx, in, out)
}

func (h *tasksHandler) Cancel(ctx context.Context, in *CancelRequest, out *CancelResponse) error {
	return h.TasksHandler.Cancel(ctx, in, out)
}

func (h *tasksHandler) Complete(ctx context.Context, in *CompleteRequest, out *CompleteResponse) error {
	return h.TasksHandler.Complete(ctx, in, out)
}

func (h *tasksHandler) Defer(ctx context.Context, in *DeferRequest, out *DeferResponse) error {
	return h.TasksHandler.Defer(ctx, in, out)
}

func (h *tasksHandler) Next(ctx context.Context, in *NextRequest, out *NextResponse) error {
	return h.TasksHandler.Next(ctx, in, out)
}

func (h *tasksHandler) Unassign(ctx context.Context, in *UnassignRequest, out *UnassignResponse) error {
	return h.TasksHandler.Unassign(ctx, in, out)
}

func (h *tasksHandler) Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error {
	return h.TasksHandler.Remove(ctx, in, out)
}