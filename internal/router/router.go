package router

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-smart/internal/handler"
)

// 注册路由
func Register(b *builder.Engine) {
	b.GET("/", (&handler.Home{}).Index)
}
