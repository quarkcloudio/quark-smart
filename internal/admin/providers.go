package admin

import (
	"github.com/quarkcms/quark-easy/internal/admin/dashboards"
	"github.com/quarkcms/quark-easy/internal/admin/login"
	"github.com/quarkcms/quark-easy/internal/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
	&dashboards.Index{},
	&resources.Post{},
}
