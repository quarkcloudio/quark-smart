package resources

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-smart/internal/model"
	"gorm.io/gorm"
)

type Page struct {
	adminresource.Template
}

// 初始化
func (p *Page) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "单页"

	// 模型
	p.Model = &model.Post{}

	// 分页
	p.PerPage = 10

	return p
}

// 只查询单页类型
func (p *Page) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type", "PAGE")
}

// 字段
func (p *Page) Fields(ctx *builder.Context) []interface{} {
	field := &builder.AdminField{}

	// 获取分类
	pages, _ := (&model.Post{}).OrderedList(false)

	return []interface{}{
		field.ID("id", "ID"),

		field.Hidden("adminid", "AdminID"),

		field.Text("title", "标题").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "标题必须填写",
				},
			),

		field.Text("name", "缩略名").
			OnlyOnForms(),

		field.TextArea("description", "描述").
			SetRules(
				[]string{
					"max:200",
				},
				map[string]string{
					"max": "描述不能超过200个字符",
				},
			).
			OnlyOnForms(),

		field.Number("level", "排序").
			SetEditable(true),

		field.Select("pid", "根节点").
			SetOptions(pages).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Datetime("created_at", "创建时间"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			OnlyOnForms(),
	}
}

// 搜索
func (p *Page) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("title", "标题"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("created_at", "创建时间"),
	}
}

// 行为
func (p *Page) Actions(ctx *builder.Context) []interface{} {

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
