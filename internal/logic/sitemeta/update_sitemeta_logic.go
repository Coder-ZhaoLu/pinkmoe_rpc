package sitemeta

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/sitemeta"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSitemetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSitemetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSitemetaLogic {
	return &UpdateSitemetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSitemeta management
func (l *UpdateSitemetaLogic) UpdateSitemeta(in *pinkmoe.SitemetaInfo) (*pinkmoe.BaseResp, error) {
	err := l.svcCtx.DB.Sitemeta.Update().
		Where(sitemeta.KeyEQ(in.Key)).
		SetNotEmptyValue(in.Value).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
