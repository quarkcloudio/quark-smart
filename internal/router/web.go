package router

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-smart/internal/handler"
)

// 注册Web路由
func WebRegister(b *builder.Engine) {
	b.GET("/", (&handler.Home{}).Index)
}
