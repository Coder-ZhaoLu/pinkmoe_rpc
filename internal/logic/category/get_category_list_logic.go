package category

import (
	"context"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/category"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/predicate"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryListLogic) GetCategoryList(in *pinkmoe.CategoryListReq) (*pinkmoe.CategoryListResp, error) {
	var predicates []predicate.Category
	if in.Name != "" {
		predicates = append(predicates, category.NameContains(in.Name))
	}
	if in.Slug != "" {
		predicates = append(predicates, category.SlugContains(in.Slug))
	}
	if in.Icon != "" {
		predicates = append(predicates, category.IconContains(in.Icon))
	}
	result, err := l.svcCtx.DB.Category.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(pager *ent.CategoryPager) {
		pager.Order = ent.Asc(category.FieldSort)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(err, in)
	}

	resp := &pinkmoe.CategoryListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &pinkmoe.CategoryInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			Sort:      v.Sort,
			Status:    uint32(v.Status),
			Name:      v.Name,
			Slug:      v.Slug,
			Icon:      v.Icon,
			Desc:      v.Desc,
			Type:      v.Type,
		})
	}

	return resp, nil
}
