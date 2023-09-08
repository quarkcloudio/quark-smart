package resource

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tabs"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/utils/lister"
	"github.com/quarkcms/quark-smart/internal/model"
)

type Category struct {
	resource.Template
}

// 初始化
func (p *Category) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "分类"

	// 模型
	p.Model = &model.Category{}

	// 默认排序
	p.IndexQueryOrder = "sort asc"

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
	field := &resource.Field{}

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
			SetEditable(true),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true).
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Category) ExtendFields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

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
			SetEditable(true).
			SetDefault(10),

		field.Switch("status", "状态").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),
	}
}

// 搜索
func (p *Category) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Category) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.EditLink(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 列表页面显示前回调
func (p *Category) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok && search != nil {
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
