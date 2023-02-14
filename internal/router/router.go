package router

import (
	"github.com/quarkcms/quark-easy/internal/handler"
	"github.com/quarkcms/quark-go/pkg/builder"
)

// 注册路由
func Register(b *builder.Engine) {
	b.GET("/", (&handler.Home{}).Index)
}
