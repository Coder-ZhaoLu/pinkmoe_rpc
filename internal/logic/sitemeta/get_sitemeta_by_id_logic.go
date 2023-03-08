package sitemeta

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/sitemeta"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSitemetaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSitemetaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSitemetaByIdLogic {
	return &GetSitemetaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSitemetaByIdLogic) GetSitemetaById(in *pinkmoe.KeyReq) (*pinkmoe.SitemetaInfo, error) {
	result, err := l.svcCtx.DB.Sitemeta.Query().Where(sitemeta.KeyEQ(in.Key)).First(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.SitemetaInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.CreatedAt.UnixMilli(),
		Key:       result.Key,
		Value:     result.Value,
	}, nil
}
