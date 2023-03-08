package category

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *pinkmoe.CategoryInfo) (*pinkmoe.BaseResp, error) {
	err := l.svcCtx.DB.Category.UpdateOneID(in.Id).
		SetNotEmptyName(in.Name).
		SetNotEmptyDesc(in.Desc).
		SetNotEmptyIcon(in.Icon).
		SetNotEmptySlug(in.Slug).
		SetNotEmptySort(in.Sort).
		SetNotEmptyStatus(uint8(in.Status)).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
