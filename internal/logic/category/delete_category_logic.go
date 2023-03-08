package category

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/category"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *pinkmoe.IDsReq) (*pinkmoe.BaseResp, error) {
	_, err := l.svcCtx.DB.Category.Delete().Where(category.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
