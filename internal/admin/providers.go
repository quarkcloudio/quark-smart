package admin

import (
	"github.com/quarkcms/quark-smart/internal/admin/dashboards"
	"github.com/quarkcms/quark-smart/internal/admin/login"
	"github.com/quarkcms/quark-smart/internal/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
	&dashboards.Index{},
	&resources.Article{},
	&resources.Page{},
	&resources.Category{},
	&resources.Banner{},
	&resources.BannerCategory{},
	&resources.Navigation{},
}
