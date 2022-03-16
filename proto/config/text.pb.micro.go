// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/config/text.proto

package config

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

// Api Endpoints for Text service

func NewTextEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Text service

type TextService interface {
	// 写入
	Write(ctx context.Context, in *TextWriteRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 读取
	Read(ctx context.Context, in *TextReadRequest, opts ...client.CallOption) (*TextReadResponse, error)
	// 删除
	Delete(ctx context.Context, in *TextDeleteRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取
	Get(ctx context.Context, in *TextGetRequest, opts ...client.CallOption) (*TextGetResponse, error)
	// 列举
	List(ctx context.Context, in *TextListRequest, opts ...client.CallOption) (*TextListResponse, error)
	// 搜索
	Search(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error)
}

type textService struct {
	c    client.Client
	name string
}

func NewTextService(name string, c client.Client) TextService {
	return &textService{
		c:    c,
		name: name,
	}
}

func (c *textService) Write(ctx context.Context, in *TextWriteRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Text.Write", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textService) Read(ctx context.Context, in *TextReadRequest, opts ...client.CallOption) (*TextReadResponse, error) {
	req := c.c.NewRequest(c.name, "Text.Read", in)
	out := new(TextReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textService) Delete(ctx context.Context, in *TextDeleteRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Text.Delete", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textService) Get(ctx context.Context, in *TextGetRequest, opts ...client.CallOption) (*TextGetResponse, error) {
	req := c.c.NewRequest(c.name, "Text.Get", in)
	out := new(TextGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textService) List(ctx context.Context, in *TextListRequest, opts ...client.CallOption) (*TextListResponse, error) {
	req := c.c.NewRequest(c.name, "Text.List", in)
	out := new(TextListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textService) Search(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Text.Search", in)
	out := new(TextSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Text service

type TextHandler interface {
	// 写入
	Write(context.Context, *TextWriteRequest, *UuidResponse) error
	// 读取
	Read(context.Context, *TextReadRequest, *TextReadResponse) error
	// 删除
	Delete(context.Context, *TextDeleteRequest, *UuidResponse) error
	// 获取
	Get(context.Context, *TextGetRequest, *TextGetResponse) error
	// 列举
	List(context.Context, *TextListRequest, *TextListResponse) error
	// 搜索
	Search(context.Context, *TextSearchRequest, *TextSearchResponse) error
}

func RegisterTextHandler(s server.Server, hdlr TextHandler, opts ...server.HandlerOption) error {
	type text interface {
		Write(ctx context.Context, in *TextWriteRequest, out *UuidResponse) error
		Read(ctx context.Context, in *TextReadRequest, out *TextReadResponse) error
		Delete(ctx context.Context, in *TextDeleteRequest, out *UuidResponse) error
		Get(ctx context.Context, in *TextGetRequest, out *TextGetResponse) error
		List(ctx context.Context, in *TextListRequest, out *TextListResponse) error
		Search(ctx context.Context, in *TextSearchRequest, out *TextSearchResponse) error
	}
	type Text struct {
		text
	}
	h := &textHandler{hdlr}
	return s.Handle(s.NewHandler(&Text{h}, opts...))
}

type textHandler struct {
	TextHandler
}

func (h *textHandler) Write(ctx context.Context, in *TextWriteRequest, out *UuidResponse) error {
	return h.TextHandler.Write(ctx, in, out)
}

func (h *textHandler) Read(ctx context.Context, in *TextReadRequest, out *TextReadResponse) error {
	return h.TextHandler.Read(ctx, in, out)
}

func (h *textHandler) Delete(ctx context.Context, in *TextDeleteRequest, out *UuidResponse) error {
	return h.TextHandler.Delete(ctx, in, out)
}

func (h *textHandler) Get(ctx context.Context, in *TextGetRequest, out *TextGetResponse) error {
	return h.TextHandler.Get(ctx, in, out)
}

func (h *textHandler) List(ctx context.Context, in *TextListRequest, out *TextListResponse) error {
	return h.TextHandler.List(ctx, in, out)
}

func (h *textHandler) Search(ctx context.Context, in *TextSearchRequest, out *TextSearchResponse) error {
	return h.TextHandler.Search(ctx, in, out)
}
