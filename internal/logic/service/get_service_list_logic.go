package service

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/predicate"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/service"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceListLogic) GetServiceList(in *pinkmoe.ServiceListReq) (*pinkmoe.ServiceListResp, error) {
	var predicates []predicate.Service
	if in.Title != "" {
		predicates = append(predicates, service.TitleContains(in.Title))
	}
	if in.Content != "" {
		predicates = append(predicates, service.ContentContains(in.Content))
	}
	result, err := l.svcCtx.DB.Service.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	resp := &pinkmoe.ServiceListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &pinkmoe.ServiceInfo{
			Id:         v.ID.String(),
			CreatedAt:  v.CreatedAt.UnixMilli(),
			Status:     uint32(v.Status),
			Title:      v.Title,
			Content:    v.Content,
			CategoryId: int64(v.CategoryID),
			AuthorUuid: v.AuthorUUID,
			Type:       v.Type,
			Price:      v.Price,
			View:       v.View,
		})
	}

	return resp, nil
}
