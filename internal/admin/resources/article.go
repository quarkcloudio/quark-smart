package resources

import (
	"time"

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
			SetOptions([]map[string]interface{}{
				{"label": "首页推荐", "value": 1},
				{"label": "频道推荐", "value": 2},
				{"label": "列表推荐", "value": 3},
				{"label": "详情推荐", "value": 4},
			}),

		field.Radio("show_type", "展现形式").
			SetOptions([]map[string]interface{}{
				{"label": "无图", "value": 1},
				{"label": "单图", "value": 2},
				{"label": "多图", "value": 3},
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
			OnlyOnForms(),

		field.Select("pid", "分类目录").
			SetOptions(categorys).
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "请选择分类目录",
				},
			).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Datetime("created_at", "发布时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}),

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
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true),

		field.Datetime("created_at", "发布时间").
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

// 创建页面显示前回调
func (p *Article) BeforeCreating(ctx *builder.Context) map[string]interface{} {

	// 表单初始化数据
	data := map[string]interface{}{
		"level":      0,
		"view":       0,
		"show_type":  1,
		"comment":    0,
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
		"status":     true,
	}

	return data
}
