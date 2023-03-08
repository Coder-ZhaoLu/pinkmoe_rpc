package service

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/service"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceByIdLogic {
	return &GetServiceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceByIdLogic) GetServiceById(in *pinkmoe.UUIDReq) (*pinkmoe.ServiceInfo, error) {
	result, err := l.svcCtx.DB.Service.Query().Where(service.IDEQ(uuidx.ParseUUIDString(in.Id))).First(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.ServiceInfo{
		Id:         result.ID.String(),
		CreatedAt:  result.CreatedAt.UnixMilli(),
		UpdatedAt:  result.CreatedAt.UnixMilli(),
		Status:     uint32(result.Status),
		Title:      result.Title,
		Content:    result.Content,
		CategoryId: int64(result.CategoryID),
		AuthorUuid: result.AuthorUUID,
		Cover:      result.Cover,
		Type:       result.Type,
		View:       result.View,
		Price:      result.Price,
	}, nil
}
