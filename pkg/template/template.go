package template

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// 模板结构体
type Template struct {
	templates *template.Template
}

func New(templatePath string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(templatePath)),
	}
}

// 模板方法：输出Html标签
func html(x string) interface{} {
	return template.HTML(x)
}

// 模板渲染方法
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// 注入上下文
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	// 注册模板方法
	t.templates = t.templates.Funcs(template.FuncMap{"html": html})

	return t.templates.ExecuteTemplate(w, name, data)
}
