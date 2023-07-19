package search

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-smart/internal/model"
	"gorm.io/gorm"
)

type Category struct {
	searches.TreeSelect
}

// 初始化
func (p *Category) Init(column string, name string) *Category {
	p.ParentInit()
	p.Column = column
	p.Name = name

	return p
}

// 执行查询
func (p *Category) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	if value == "" || value == nil {
		return query
	}

	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *Category) Options(ctx *builder.Context) interface{} {
	options, _ := (&model.Category{}).TreeSelect(false)

	return options
}
