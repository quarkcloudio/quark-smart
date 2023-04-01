package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
	"github.com/quarkcms/quark-go/pkg/lister"
	"github.com/quarkcms/quark-smart/internal/model"
)

type Category struct {
	adminresource.Template
}

// 初始化
func (p *Category) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "分类"

	// 模型
	p.Model = &model.Category{}

	// 默认排序
	p.IndexOrder = "sort asc"

	// 分页
	p.PerPage = false

	return p
}

func (p *Category) Fields(ctx *builder.Context) []interface{} {
	var tabPanes []interface{}

	// 基础字段
	basePane := (&tabs.TabPane{}).
		Init().
		SetTitle("基础").
		SetBody(p.BaseFields(ctx))
	tabPanes = append(tabPanes, basePane)

	// 扩展字段
	extendPane := (&tabs.TabPane{}).
		Init().
		SetTitle("扩展").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, extendPane)

	return tabPanes
}

// 基础字段
func (p *Category) BaseFields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	// 获取分类
	categorys, _ := (&model.Category{}).TreeSelect(true)

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Text("title", "标题").
			SetRules([]*rule.Rule{
				rule.Required(true, "标题必须填写"),
			}),

		field.Text("name", "缩略名").
			SetRules([]*rule.Rule{
				rule.Required(true, "缩略名必须填写"),
			}),

		field.TreeSelect("pid", "父节点").
			SetData(categorys).
			OnlyOnForms(),

		field.TextArea("description", "描述").
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetEditable(true),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Category) ExtendFields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.Image("cover_id", "封面图").
			SetMode("single").
			OnlyOnForms(),

		field.Text("index_tpl", "频道模板").
			OnlyOnForms(),

		field.Text("lists_tpl", "列表模板").
			OnlyOnForms(),

		field.Text("detail_tpl", "详情模板").
			OnlyOnForms(),

		field.Number("page_num", "分页数量").
			SetEditable(true),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true),
	}
}

// 搜索
func (p *Category) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("title", "标题"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("created_at", "创建时间"),
	}
}

// 行为
func (p *Category) Actions(ctx *builder.Context) []interface{} {

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

// 列表页面显示前回调
func (p *Category) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok == true && search != nil {
		result := []interface{}{}
		for _, v := range list {
			result = append(result, v)
		}

		return result
	}

	// 转换成树形表格
	tree, _ := lister.ListToTree(list, "id", "pid", "children", 0)

	return tree
}

// 创建页面显示前回调
func (p *Category) BeforeCreating(ctx *builder.Context) map[string]interface{} {

	// 表单初始化数据
	data := map[string]interface{}{
		"pid":      0,
		"sort":     0,
		"status":   true,
		"page_num": 10,
	}

	return data
}
