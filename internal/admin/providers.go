package admin

import (
	"github.com/quarkcms/quark-simple/internal/admin/dashboards"
	"github.com/quarkcms/quark-simple/internal/admin/login"
	"github.com/quarkcms/quark-simple/internal/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
	&dashboards.Index{},
	&resources.Post{},
}
