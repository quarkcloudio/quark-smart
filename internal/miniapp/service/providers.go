package service

import (
	"github.com/quarkcloudio/quark-smart/internal/miniapp/service/forms"
	"github.com/quarkcloudio/quark-smart/internal/miniapp/service/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
	&forms.Demo{},
}
