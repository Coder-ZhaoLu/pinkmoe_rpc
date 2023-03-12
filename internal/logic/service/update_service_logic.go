package service

import (
	"context"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateServiceLogic) UpdateService(in *pinkmoe.ServiceInfo) (*pinkmoe.BaseResp, error) {
	err := l.svcCtx.DB.Service.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
		SetNotEmptyTitle(in.Title).
		SetNotEmptyDesc(in.Desc).
		SetNotEmptyContent(in.Content).
		SetNotEmptyCategoryID(uint64(in.CategoryId)).
		SetNotEmptyAuthorUUID(in.AuthorUuid).
		SetNotEmptyType(in.Type).
		SetNotEmptyCover(in.Cover).
		SetNotEmptyPrice(in.Price).
		SetNotEmptyView(in.View).
		SetNotEmptyStatus(uint8(in.Status)).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
