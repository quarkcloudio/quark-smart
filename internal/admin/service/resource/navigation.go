package resource

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-go/v2/pkg/utils/lister"
	"github.com/quarkcloudio/quark-smart/internal/model"
)

type Navigation struct {
	resource.Template
}

// 初始化
func (p *Navigation) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "导航"

	// 模型
	p.Model = &model.Navigation{}

	// 默认排序
	p.IndexQueryOrder = "sort asc"

	// 分页
	p.PerPage = false

	return p
}

func (p *Navigation) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	// 获取分类
	categorys, _ := (&model.Navigation{}).TreeSelect(true)

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Text("title", "标题").
			SetRules([]*rule.Rule{
				rule.Required(true, "标题必须填写"),
			}),

		field.TreeSelect("pid", "父节点").
			SetData(categorys).
			SetDefault(0).
			OnlyOnForms(),

		field.Image("cover_id", "封面图").
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Text("url", "链接"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true).
			OnlyOnForms(),
	}
}

// 搜索
func (p *Navigation) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Navigation) Actions(ctx *builder.Context) []interface{} {

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
func (p *Navigation) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
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
