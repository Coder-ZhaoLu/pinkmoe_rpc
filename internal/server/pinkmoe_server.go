// Code generated by goctl. DO NOT EDIT.
// Source: pinkmoe.proto

package server

import (
	"context"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/base"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/category"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/comment"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/service"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/sitemeta"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/logic/version"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
)

type PinkmoeServer struct {
	svcCtx *svc.ServiceContext
	pinkmoe.UnimplementedPinkmoeServer
}

func NewPinkmoeServer(svcCtx *svc.ServiceContext) *PinkmoeServer {
	return &PinkmoeServer{
		svcCtx: svcCtx,
	}
}

func (s *PinkmoeServer) InitDatabase(ctx context.Context, in *pinkmoe.Empty) (*pinkmoe.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Category management
func (s *PinkmoeServer) CreateCategory(ctx context.Context, in *pinkmoe.CategoryInfo) (*pinkmoe.BaseIDResp, error) {
	l := category.NewCreateCategoryLogic(ctx, s.svcCtx)
	return l.CreateCategory(in)
}

func (s *PinkmoeServer) UpdateCategory(ctx context.Context, in *pinkmoe.CategoryInfo) (*pinkmoe.BaseResp, error) {
	l := category.NewUpdateCategoryLogic(ctx, s.svcCtx)
	return l.UpdateCategory(in)
}

func (s *PinkmoeServer) GetCategoryList(ctx context.Context, in *pinkmoe.CategoryListReq) (*pinkmoe.CategoryListResp, error) {
	l := category.NewGetCategoryListLogic(ctx, s.svcCtx)
	return l.GetCategoryList(in)
}

func (s *PinkmoeServer) GetCategoryById(ctx context.Context, in *pinkmoe.IDReq) (*pinkmoe.CategoryInfo, error) {
	l := category.NewGetCategoryByIdLogic(ctx, s.svcCtx)
	return l.GetCategoryById(in)
}

func (s *PinkmoeServer) DeleteCategory(ctx context.Context, in *pinkmoe.IDsReq) (*pinkmoe.BaseResp, error) {
	l := category.NewDeleteCategoryLogic(ctx, s.svcCtx)
	return l.DeleteCategory(in)
}

// Comment management
func (s *PinkmoeServer) CreateComment(ctx context.Context, in *pinkmoe.CommentInfo) (*pinkmoe.BaseUUIDResp, error) {
	l := comment.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

func (s *PinkmoeServer) UpdateComment(ctx context.Context, in *pinkmoe.CommentInfo) (*pinkmoe.BaseResp, error) {
	l := comment.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *PinkmoeServer) GetCommentList(ctx context.Context, in *pinkmoe.CommentListReq) (*pinkmoe.CommentListResp, error) {
	l := comment.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}

func (s *PinkmoeServer) GetCommentById(ctx context.Context, in *pinkmoe.UUIDReq) (*pinkmoe.CommentInfo, error) {
	l := comment.NewGetCommentByIdLogic(ctx, s.svcCtx)
	return l.GetCommentById(in)
}

func (s *PinkmoeServer) DeleteComment(ctx context.Context, in *pinkmoe.UUIDsReq) (*pinkmoe.BaseResp, error) {
	l := comment.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// Service management
func (s *PinkmoeServer) CreateService(ctx context.Context, in *pinkmoe.ServiceInfo) (*pinkmoe.BaseUUIDResp, error) {
	l := service.NewCreateServiceLogic(ctx, s.svcCtx)
	return l.CreateService(in)
}

func (s *PinkmoeServer) UpdateService(ctx context.Context, in *pinkmoe.ServiceInfo) (*pinkmoe.BaseResp, error) {
	l := service.NewUpdateServiceLogic(ctx, s.svcCtx)
	return l.UpdateService(in)
}

func (s *PinkmoeServer) GetServiceList(ctx context.Context, in *pinkmoe.ServiceListReq) (*pinkmoe.ServiceListResp, error) {
	l := service.NewGetServiceListLogic(ctx, s.svcCtx)
	return l.GetServiceList(in)
}

func (s *PinkmoeServer) GetServiceById(ctx context.Context, in *pinkmoe.UUIDReq) (*pinkmoe.ServiceInfo, error) {
	l := service.NewGetServiceByIdLogic(ctx, s.svcCtx)
	return l.GetServiceById(in)
}

func (s *PinkmoeServer) GetServiceByIds(ctx context.Context, in *pinkmoe.UUIDsReq) (*pinkmoe.ServiceListResp, error) {
	l := service.NewGetServiceByIdsLogic(ctx, s.svcCtx)
	return l.GetServiceByIds(in)
}

func (s *PinkmoeServer) DeleteService(ctx context.Context, in *pinkmoe.UUIDsReq) (*pinkmoe.BaseResp, error) {
	l := service.NewDeleteServiceLogic(ctx, s.svcCtx)
	return l.DeleteService(in)
}

// Sitemeta management
func (s *PinkmoeServer) UpdateSitemeta(ctx context.Context, in *pinkmoe.SitemetaInfo) (*pinkmoe.BaseResp, error) {
	l := sitemeta.NewUpdateSitemetaLogic(ctx, s.svcCtx)
	return l.UpdateSitemeta(in)
}

func (s *PinkmoeServer) GetSitemetaById(ctx context.Context, in *pinkmoe.KeyReq) (*pinkmoe.SitemetaInfo, error) {
	l := sitemeta.NewGetSitemetaByIdLogic(ctx, s.svcCtx)
	return l.GetSitemetaById(in)
}

// Version management
func (s *PinkmoeServer) CreateVersion(ctx context.Context, in *pinkmoe.VersionInfo) (*pinkmoe.BaseUUIDResp, error) {
	l := version.NewCreateVersionLogic(ctx, s.svcCtx)
	return l.CreateVersion(in)
}

func (s *PinkmoeServer) UpdateVersion(ctx context.Context, in *pinkmoe.VersionInfo) (*pinkmoe.BaseResp, error) {
	l := version.NewUpdateVersionLogic(ctx, s.svcCtx)
	return l.UpdateVersion(in)
}

func (s *PinkmoeServer) GetVersionList(ctx context.Context, in *pinkmoe.VersionListReq) (*pinkmoe.VersionListResp, error) {
	l := version.NewGetVersionListLogic(ctx, s.svcCtx)
	return l.GetVersionList(in)
}

func (s *PinkmoeServer) GetVersionById(ctx context.Context, in *pinkmoe.UUIDReq) (*pinkmoe.VersionInfo, error) {
	l := version.NewGetVersionByIdLogic(ctx, s.svcCtx)
	return l.GetVersionById(in)
}

func (s *PinkmoeServer) DeleteVersion(ctx context.Context, in *pinkmoe.UUIDsReq) (*pinkmoe.BaseResp, error) {
	l := version.NewDeleteVersionLogic(ctx, s.svcCtx)
	return l.DeleteVersion(in)
}
