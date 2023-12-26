package router

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-smart/internal/admin/handler"
)

// 注册Admin路由
func AdminRegister(b *builder.Engine) {
	g := b.Group("/api/admin")
	g.GET("/index/index", (&handler.Index{}).Index)
}
