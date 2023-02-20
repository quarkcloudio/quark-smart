package admin

import (
	"github.com/quarkcms/quark-smart/internal/admin/dashboard"
	"github.com/quarkcms/quark-smart/internal/admin/login"
	"github.com/quarkcms/quark-smart/internal/admin/resource"
)

// 注册服务
var Provider = []interface{}{
	&login.Index{},
	&dashboard.Index{},
	&resource.Article{},
	&resource.Page{},
	&resource.Category{},
	&resource.Banner{},
	&resource.BannerCategory{},
	&resource.Navigation{},
}
