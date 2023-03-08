// Code generated by goctl. DO NOT EDIT.
// Source: pinkmoe.proto

package pinkmoeclient

import (
	"context"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp       = pinkmoe.BaseIDResp
	BaseResp         = pinkmoe.BaseResp
	BaseUUIDResp     = pinkmoe.BaseUUIDResp
	CategoryInfo     = pinkmoe.CategoryInfo
	CategoryListReq  = pinkmoe.CategoryListReq
	CategoryListResp = pinkmoe.CategoryListResp
	CommentInfo      = pinkmoe.CommentInfo
	CommentListReq   = pinkmoe.CommentListReq
	CommentListResp  = pinkmoe.CommentListResp
	Empty            = pinkmoe.Empty
	IDReq            = pinkmoe.IDReq
	IDsReq           = pinkmoe.IDsReq
	KeyReq           = pinkmoe.KeyReq
	PageInfoReq      = pinkmoe.PageInfoReq
	ServiceInfo      = pinkmoe.ServiceInfo
	ServiceListReq   = pinkmoe.ServiceListReq
	ServiceListResp  = pinkmoe.ServiceListResp
	SitemetaInfo     = pinkmoe.SitemetaInfo
	UUIDReq          = pinkmoe.UUIDReq
	UUIDsReq         = pinkmoe.UUIDsReq

	Pinkmoe interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Category management
		CreateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetCategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		GetCategoryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryInfo, error)
		DeleteCategory(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Comment management
		CreateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
		GetCommentById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*CommentInfo, error)
		DeleteComment(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Service management
		CreateService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetServiceList(ctx context.Context, in *ServiceListReq, opts ...grpc.CallOption) (*ServiceListResp, error)
		GetServiceById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*ServiceInfo, error)
		DeleteService(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Sitemeta management
		UpdateSitemeta(ctx context.Context, in *SitemetaInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetSitemetaById(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*SitemetaInfo, error)
	}

	defaultPinkmoe struct {
		cli zrpc.Client
	}
)

func NewPinkmoe(cli zrpc.Client) Pinkmoe {
	return &defaultPinkmoe{
		cli: cli,
	}
}

func (m *defaultPinkmoe) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Category management
func (m *defaultPinkmoe) CreateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

func (m *defaultPinkmoe) UpdateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.UpdateCategory(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetCategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetCategoryList(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetCategoryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryInfo, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetCategoryById(ctx, in, opts...)
}

func (m *defaultPinkmoe) DeleteCategory(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.DeleteCategory(ctx, in, opts...)
}

// Comment management
func (m *defaultPinkmoe) CreateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

func (m *defaultPinkmoe) UpdateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetCommentById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*CommentInfo, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetCommentById(ctx, in, opts...)
}

func (m *defaultPinkmoe) DeleteComment(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// Service management
func (m *defaultPinkmoe) CreateService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.CreateService(ctx, in, opts...)
}

func (m *defaultPinkmoe) UpdateService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.UpdateService(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetServiceList(ctx context.Context, in *ServiceListReq, opts ...grpc.CallOption) (*ServiceListResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetServiceList(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetServiceById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*ServiceInfo, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetServiceById(ctx, in, opts...)
}

func (m *defaultPinkmoe) DeleteService(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.DeleteService(ctx, in, opts...)
}

// Sitemeta management
func (m *defaultPinkmoe) UpdateSitemeta(ctx context.Context, in *SitemetaInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.UpdateSitemeta(ctx, in, opts...)
}

func (m *defaultPinkmoe) GetSitemetaById(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*SitemetaInfo, error) {
	client := pinkmoe.NewPinkmoeClient(m.cli.Conn())
	return client.GetSitemetaById(ctx, in, opts...)
}