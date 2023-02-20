package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-smart/internal/admin/search"
	"github.com/quarkcms/quark-smart/internal/model"
)

type Banner struct {
	adminresource.Template
}

// 初始化
func (p *Banner) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "广告"

	// 模型
	p.Model = &model.Banner{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Banner) Fields(ctx *builder.Context) []interface{} {
	field := &builder.AdminField{}

	// 获取分类
	categorys, _ := (&model.BannerCategory{}).List()

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("title", "标题").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "标题必须填写",
				},
			),

		field.Image("cover_id", "图片").
			OnlyOnForms(),

		field.Select("category_id", "广告位").
			SetOptions(categorys).
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "请选择广告位",
				},
			).
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Text("url", "链接"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Banner) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("title", "标题"),
		(&search.BannerCategory{}).Init("category_id", "广告位"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("created_at", "创建时间"),
	}
}

// 行为
func (p *Banner) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&actions.CreateLink{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.EditLink{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
