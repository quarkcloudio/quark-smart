package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/lister"
	"github.com/quarkcms/quark-smart/internal/model"
)

type Navigation struct {
	adminresource.Template
}

// 初始化
func (p *Navigation) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "导航"

	// 模型
	p.Model = &model.Navigation{}

	// 默认排序
	p.IndexOrder = "sort asc"

	// 分页
	p.PerPage = false

	return p
}

func (p *Navigation) Fields(ctx *builder.Context) []interface{} {
	field := &builder.AdminField{}

	// 获取分类
	categorys, _ := (&model.Navigation{}).OrderedList(true)

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Text("title", "标题").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "标题必须填写",
				},
			),

		field.Select("pid", "父节点").
			SetOptions(categorys).
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
		(&searches.Input{}).Init("title", "标题"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("created_at", "创建时间"),
	}
}

// 行为
func (p *Navigation) Actions(ctx *builder.Context) []interface{} {

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
func (p *Navigation) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
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
