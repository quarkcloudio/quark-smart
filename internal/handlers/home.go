package handlers

import "github.com/quarkcms/quark-go/pkg/builder"

// 结构体
type Home struct{}

// 首页
func (p *Home) Index(ctx *builder.Context) error {
	ctx.Write([]byte("Hello World!"))

	return nil
}
