package service

import (
	"context"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Service management
func (l *CreateServiceLogic) CreateService(in *pinkmoe.ServiceInfo) (*pinkmoe.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Service.Create().
		SetTitle(in.Title).
		SetContent(in.Content).
		SetStatus(uint8(in.Status)).
		SetCategoryID(uint64(in.CategoryId)).
		SetAuthorUUID(in.AuthorUuid).
		SetCover(in.Cover).
		SetType(in.Type).
		SetPrice(in.Price).
		SetView(in.View).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
