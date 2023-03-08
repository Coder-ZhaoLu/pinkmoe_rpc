package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/internal/svc"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/pinkmoe"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *pinkmoe.Empty) (*pinkmoe.BaseResp, error) {
	// todo: add your logic here and delete this line

	err := l.insertApiData()
	if err != nil {
		if strings.Contains(err.Error(), "common.createFailed") {
			return nil, statuserr.NewInvalidArgumentError(i18n.AlreadyInit)
		}
		return nil, statuserr.NewInternalError(err.Error())
	}

	err = l.insertMenuData()
	if err != nil {
		return nil, statuserr.NewInternalError(err.Error())
	}

	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeError(enum.Internal, err.Error())
	}

	err = l.insertSitemetaData()
	if err != nil {
		return nil, statuserr.NewInternalError(err.Error())
	}

	return &pinkmoe.BaseResp{
		Msg: i18n.Success,
	}, nil
}

func (l *InitDatabaseLogic) insertApiData() (err error) {

	// SITEMETA

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/sitemeta/update",
		Description: "apiDesc.updateSitemeta",
		ApiGroup:    "sitemeta",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/sitemeta",
		Description: "apiDesc.getSitemetaByKey",
		ApiGroup:    "sitemeta",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	// SERVICE

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/service/create",
		Description: "apiDesc.createService",
		ApiGroup:    "service",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/service/update",
		Description: "apiDesc.updateService",
		ApiGroup:    "service",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/service/delete",
		Description: "apiDesc.deleteService",
		ApiGroup:    "service",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/service/list",
		Description: "apiDesc.getServiceList",
		ApiGroup:    "service",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/service",
		Description: "apiDesc.getServiceByUuid",
		ApiGroup:    "service",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	// CATEGORY

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/category/create",
		Description: "apiDesc.createCategory",
		ApiGroup:    "category",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/category/update",
		Description: "apiDesc.updateCategory",
		ApiGroup:    "category",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/category/delete",
		Description: "apiDesc.deleteCategory",
		ApiGroup:    "category",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/category/list",
		Description: "apiDesc.getCategoryList",
		ApiGroup:    "category",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/category",
		Description: "apiDesc.getCategoryById",
		ApiGroup:    "category",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	// COMMENT

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/comment/create",
		Description: "apiDesc.createComment",
		ApiGroup:    "comment",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/comment/update",
		Description: "apiDesc.updateComment",
		ApiGroup:    "comment",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/comment/delete",
		Description: "apiDesc.deleteComment",
		ApiGroup:    "comment",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/comment/list",
		Description: "apiDesc.getCommentList",
		ApiGroup:    "comment",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/comment",
		Description: "apiDesc.getCommentByUuid",
		ApiGroup:    "comment",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *InitDatabaseLogic) insertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  enum.DefaultParentId,
		Path:      "",
		Name:      "PinkmoeManagementDirectory",
		Redirect:  "",
		Component: "LAYOUT",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.pinkmoeManagement",
			Icon:               "ic:twotone-settings-suggest",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 0,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/sitemeta",
		Name:      "SitemetaManagement",
		Redirect:  "",
		Component: "/pinkmoe/sitemeta/index",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.sitemetaManagement",
			Icon:               "ic:outline-settings-input-composite",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/category",
		Name:      "CategoryManagement",
		Redirect:  "",
		Component: "/pinkmoe/category/index",
		Sort:      2,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.categoryManagement",
			Icon:               "ic:baseline-menu",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/service",
		Name:      "ServiceManagement",
		Redirect:  "",
		Component: "/pinkmoe/service/index",
		Sort:      3,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.serviceManagement",
			Icon:               "ic:baseline-shopping-bag",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/comment",
		Name:      "CommentManagement",
		Redirect:  "",
		Component: "/pinkmoe/comment/index",
		Sort:      4,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.commentManagement",
			Icon:               "ic:baseline-insert-comment",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	return err
}

// insert init sitemeta data
func (l *InitDatabaseLogic) insertSitemetaData() error {
	var sitemeta []*ent.SitemetaCreate

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_header_meta").
		SetValue(""),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_home_meta").
		SetValue(""),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_footer_meta").
		SetValue(""),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_storage_meta").
		SetValue(""),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_sms_meta").
		SetValue(""),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_email_meta").
		SetValue(""),
	)

	err := l.svcCtx.DB.Sitemeta.CreateBulk(sitemeta...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert init category data
func (l *InitDatabaseLogic) insertCategoryData() error {
	var category []*ent.CategoryCreate

	category = append(category, l.svcCtx.DB.Category.Create().
		SetName("微服务").SetDesc("微服务扩展").SetSort(1).SetIcon("").SetSlug("service").SetStatus(1),
	)

	err := l.svcCtx.DB.Category.CreateBulk(category...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}
