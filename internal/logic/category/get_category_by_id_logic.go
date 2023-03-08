package category

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *pinkmoe.IDReq) (*pinkmoe.CategoryInfo, error) {
	result, err := l.svcCtx.DB.Category.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	return &pinkmoe.CategoryInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.CreatedAt.UnixMilli(),
		Sort:      result.Sort,
		Status:    uint32(result.Status),
		Name:      result.Name,
		Slug:      result.Slug,
		Icon:      result.Icon,
		Desc:      result.Desc,
	}, nil
}
