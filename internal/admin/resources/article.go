package resources

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
	"github.com/quarkcms/quark-smart/internal/model"
)

type Article struct {
	adminresource.Template
}

// 初始化
func (p *Article) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "文章"

	// 模型
	p.Model = &model.Post{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Article) Fields(ctx *builder.Context) []interface{} {
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
func (p *Article) BaseFields(ctx *builder.Context) []interface{} {
	field := &builder.AdminField{}

	// 获取分类
	categorys, _ := (&model.Category{}).OrderedList()

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

		field.Text("author", "作者"),

		field.Number("level", "排序").
			SetEditable(true),

		field.Text("source", "来源").
			OnlyOnForms(),

		field.Checkbox("position", "推荐位").
			SetOptions(map[interface{}]interface{}{
				1: "首页推荐",
				2: "频道推荐",
				3: "列表推荐",
				4: "详情推荐",
			}),

		field.Select("pid", "分类目录").
			SetOptions(categorys).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Article) ExtendFields(ctx *builder.Context) []interface{} {
	field := &builder.AdminField{}

	return []interface{}{
		field.Image("cover_id", "封面图").
			SetMode("single").
			OnlyOnForms(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true),
	}
}

// 搜索
func (p *Article) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("title", "标题"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("created_at", "创建时间"),
	}
}

// 行为
func (p *Article) Actions(ctx *builder.Context) []interface{} {

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
