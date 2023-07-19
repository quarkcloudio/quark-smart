package handler

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *builder.Context) error {
	return ctx.JSON(200, message.Success("Hello, world!"))
}
