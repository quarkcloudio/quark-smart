package resource

import (
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/checkbox"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tabs"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-smart/internal/admin/search"
	"github.com/quarkcms/quark-smart/internal/model"
	"gorm.io/gorm"
)

type Article struct {
	resource.Template
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

// 只查询文章类型
func (p *Article) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type", "ARTICLE")
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
	field := &resource.Field{}

	// 获取分类
	categorys, _ := (&model.Category{}).TreeSelect(false)

	return []interface{}{
		field.ID("id", "ID"),

		field.Hidden("adminid", "AdminID"),

		field.Hidden("cover_ids", "封面图"),

		field.Text("title", "标题").
			SetRules([]*rule.Rule{
				rule.Required(true, "标题必须填写"),
			}),

		field.TextArea("description", "描述").
			SetRules([]*rule.Rule{
				rule.Max(200, "描述不能超过200个字符"),
			}).
			OnlyOnForms(),

		field.Text("author", "作者"),

		field.Number("level", "排序").
			SetEditable(true),

		field.Text("source", "来源").
			OnlyOnForms(),

		field.Checkbox("position", "推荐位").
			SetOptions([]*checkbox.Option{
				{Value: 1, Label: "首页推荐"},
				{Value: 2, Label: "频道推荐"},
				{Value: 3, Label: "列表推荐"},
				{Value: 4, Label: "详情推荐"},
			}),

		field.Radio("show_type", "展现形式").
			SetOptions([]*radio.Option{
				{Value: 1, Label: "无图"},
				{Value: 2, Label: "单图"},
				{Value: 3, Label: "多图"},
			}).
			SetWhen(2, func() interface{} {
				return []interface{}{
					field.Image("single_cover_ids", "封面图").
						SetMode("multiple").
						SetLimitNum(1).
						OnlyOnForms(),
				}
			}).
			SetWhen(3, func() interface{} {
				return []interface{}{
					field.Image("multiple_cover_ids", "封面图").
						SetMode("multiple").
						OnlyOnForms(),
				}
			}).
			SetDefault(1).
			OnlyOnForms(),

		field.TreeSelect("category_id", "分类目录").
			SetData(categorys).
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择分类目录"),
			}).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Datetime("created_at", "发布时间").
			SetDefault(time.Now().Format("2006-01-02 15:04:05")),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Article) ExtendFields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Text("name", "缩略名").
			OnlyOnForms(),

		field.Number("level", "排序").
			OnlyOnForms(),

		field.Number("view", "浏览量").
			OnlyOnForms(),

		field.Number("comment", "评论量").
			OnlyOnForms(),

		field.Text("password", "访问密码").
			OnlyOnForms(),

		field.File("file_ids", "附件").
			OnlyOnForms(),

		field.Switch("comment_status", "允许评论").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),

		field.Datetime("created_at", "发布时间").
			OnlyOnForms(),

		field.Switch("status", "状态").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),
	}
}

// 搜索
func (p *Article) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("title", "标题"),
		(&search.Category{}).Init("category_id", "分类目录"),
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

// 编辑页面显示前回调
func (p *Article) BeforeEditing(request *builder.Context, data map[string]interface{}) map[string]interface{} {
	if data["show_type"] == 2 {
		data["single_cover_ids"] = data["cover_ids"]
	}

	if data["show_type"] == 3 {
		data["multiple_cover_ids"] = data["cover_ids"]
	}

	return data
}

// 保存数据前回调
func (p *Article) BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	if int(submitData["show_type"].(float64)) == 2 {
		submitData["cover_ids"] = submitData["single_cover_ids"]
	}

	if int(submitData["show_type"].(float64)) == 3 {
		submitData["cover_ids"] = submitData["multiple_cover_ids"]
	}

	return submitData, nil
}
