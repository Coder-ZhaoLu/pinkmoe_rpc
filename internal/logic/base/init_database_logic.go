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

	err = l.insertCategoryData()
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

	// VERSION

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/version/create",
		Description: "apiDesc.createVersion",
		ApiGroup:    "version",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/version/update",
		Description: "apiDesc.updateVersion",
		ApiGroup:    "version",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/version/delete",
		Description: "apiDesc.deleteVersion",
		ApiGroup:    "version",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/version/list",
		Description: "apiDesc.getVersionList",
		ApiGroup:    "version",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/version",
		Description: "apiDesc.getVersionByUuid",
		ApiGroup:    "version",
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

	siteData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
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
		ParentId:  siteData.Id,
		Path:      "/sitemeta/headermeta",
		Name:      "HeadermetaManagement",
		Redirect:  "",
		Component: "/pinkmoe/sitemeta/header_meta/index",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.headermetaManagement",
			Icon:               "ic:round-view-headline",
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
		ParentId:  siteData.Id,
		Path:      "/sitemeta/footermeta",
		Name:      "FootermetaManagement",
		Redirect:  "",
		Component: "/pinkmoe/sitemeta/footer_meta/index",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.footermetaManagement",
			Icon:               "ic:baseline-calendar-view-day",
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
		ParentId:  siteData.Id,
		Path:      "/sitemeta/homemeta",
		Name:      "HomemetaManagement",
		Redirect:  "",
		Component: "/pinkmoe/sitemeta/home_meta/index",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.homemetaManagement",
			Icon:               "ic:baseline-home",
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

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/version",
		Name:      "VersionManagement",
		Redirect:  "",
		Component: "/pinkmoe/version/index",
		Sort:      4,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.versionManagement",
			Icon:               "ic:baseline-settings-overscan",
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
		SetValue("{\"title\":\"PinkMoe Studio - 粉色计划\",\"keywords\":\"PinkMoe,PinkMoe程序,多元化网站应用程序,CMS程序,漫画程序,小说程序,音乐程序,社交程序,影视程序,视频程序\",\"desc\":\"可自由搭配的插件化多功能CMS程序! \",\"createTime\":\"2023-03-08T09:51:51.083Z\",\"url\":\"https://pinkmoe.com\",\"icon\":[\"/file_manager_0/image/2023-3-12/ba0e94b5-92eb-4d88-aad1-a568c32880ba.ico\"],\"logo\":[\"/file_manager_0/image/2023-3-12/418dce69-701d-4fe1-8463-f66e4e761122.png\"],\"background\":[\"/file_manager_0/image/2023-3-9/2295baa2-1c5b-48fe-adc4-1c8a86824a34.ico\"]}"),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_home_meta").
		SetValue("{\"title\":\"PinkMoe Studio - 粉色计划\",\"desc\":\"这一次,我们将倾尽所有!\",\"demoLink\":\"https://pinkmoe.com\",\"displayTitle\":\"本站产品展示\",\"displayDesc\":\"一站式内容管理平台\",\"displayService\":[\"0186d1b8-957f-7b25-9e73-870af059e148\"]}"),
	)

	sitemeta = append(sitemeta, l.svcCtx.DB.Sitemeta.Create().
		SetKey("website_footer_meta").
		SetValue("{\"product\":[{\"name\":\"程序更新记录\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"程序使用文档\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"一键部署\",\"url\":\"https://pinkmoe.com\"}],\"about\":[{\"name\":\"FAQ\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"联系我们\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"加入我们\",\"url\":\"https://pinkmoe.com\"}],\"site\":[{\"name\":\"帮助中心\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"捐助说明\",\"url\":\"https://pinkmoe.com\"},{\"name\":\"交流社区\",\"url\":\"https://pinkmoe.com\"}],\"friend\":[{\"name\":\"Coderzhaolu\",\"url\":\"https://coderzhaolu.com\"}],\"copy\":\"©2023 Pinkmoe Studio All Rights Reserved 独立工作室\",\"forRecord\":\"粤ICP备202300000000号\",\"desc\":\"这一次,我们将倾尽所有!\",\"title\":\"PinkMoe Studio Plan\"}"),
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
		SetName("微服务").SetDesc("后端微服务扩展").SetType(0).SetSort(1).SetIcon("ic:outline-home-repair-service").SetSlug("service").SetStatus(1),
	)

	category = append(category, l.svcCtx.DB.Category.Create().
		SetName("前端组件").SetDesc("前端组件插件").SetType(1).SetSort(2).SetIcon("ic:baseline-settings-input-component").SetSlug("component").SetStatus(1),
	)

	category = append(category, l.svcCtx.DB.Category.Create().
		SetName("接口插件").SetDesc("后端纯接口扩展").SetType(0).SetSort(3).SetIcon("ic:baseline-api").SetSlug("api").SetStatus(1),
	)

	category = append(category, l.svcCtx.DB.Category.Create().
		SetName("文档中心").SetDesc("服务文档").SetType(2).SetSort(4).SetIcon("ic:baseline-document-scanner").SetSlug("document").SetStatus(1),
	)

	err := l.svcCtx.DB.Category.CreateBulk(category...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}
