package category

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateCategory management
func (l *CreateCategoryLogic) CreateCategory(in *pinkmoe.CategoryInfo) (*pinkmoe.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Category.Create().
		SetIcon(in.Icon).
		SetSlug(in.Slug).
		SetStatus(uint8(in.Status)).
		SetSort(in.Sort).
		SetDesc(in.Desc).
		SetName(in.Name).
		SetID(in.Id).
		SetType(in.Type).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}
	return &pinkmoe.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
