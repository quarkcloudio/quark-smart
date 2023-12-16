package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	admininstall "github.com/quarkcms/quark-go/v2/pkg/app/admin/install"
	adminmiddleware "github.com/quarkcms/quark-go/v2/pkg/app/admin/middleware"
	adminservice "github.com/quarkcms/quark-go/v2/pkg/app/admin/service"
	miniappinstall "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/install"
	miniappmiddleware "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/middleware"
	miniappservice "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/service"
	toolservice "github.com/quarkcms/quark-go/v2/pkg/app/tool/service"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-smart/config"
	"github.com/quarkcms/quark-smart/database"
	"github.com/quarkcms/quark-smart/internal/admin/service"
	"github.com/quarkcms/quark-smart/internal/middleware"
	"github.com/quarkcms/quark-smart/internal/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 模板结构体
type Template struct {
	templates *template.Template
}

// 模板渲染方法
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// 服务
	var providers []interface{}

	// 配置信息
	var (
		appKey     = config.App.Key
		dbUser     = config.Mysql.Username
		dbPassword = config.Mysql.Password
		dbHost     = config.Mysql.Host
		dbPort     = config.Mysql.Port
		dbName     = config.Mysql.Database
		dbCharset  = config.Mysql.Charset
	)

	// 数据库配置信息
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local"

	// 加载后台服务
	providers = append(providers, adminservice.Providers...)

	// 加载自定义后台服务
	providers = append(providers, service.Provider...)

	// 加载MiniApp服务
	providers = append(providers, miniappservice.Providers...)

	// 加载工具服务
	providers = append(providers, toolservice.Providers...)

	// 配置资源
	getConfig := &builder.Config{
		AppKey: appKey,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		Providers: providers,
	}

	// 实例化对象
	b := builder.New(getConfig)

	// WEB根目录
	b.Static("/", config.App.RootPath)

	// 静态文件目录
	b.Static("/static/", config.App.StaticPath)

	// 构建quarkgo基础数据库、拉取静态文件
	admininstall.Handle()

	// 构建本项目数据库
	database.Handle()

	// 后台中间件
	b.Use(adminmiddleware.Handle)

	// 构建MiniApp数据库
	miniappinstall.Handle()

	// MiniApp中间件
	b.Use(miniappmiddleware.Handle)

	// 中间件
	b.Use((&middleware.AppMiddleware{}).Handle)

	// 开启Debug模式
	b.Echo().Debug = config.App.Debug

	// 加载Html模板
	b.Echo().Renderer = &Template{
		templates: template.Must(template.ParseGlob(config.App.TemplatePath)),
	}

	// 日志中间件
	if config.App.Logger {
		b.Echo().Use(echomiddleware.Logger())
	}

	// 日志文件位置
	if config.App.LoggerFilePath != "" {
		f, _ := os.OpenFile(config.App.LoggerFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

		// 记录日志
		b.Echo().Logger.SetOutput(io.MultiWriter(f, os.Stdout))
	}

	// 崩溃后自动恢复
	if config.App.Recover {
		b.Echo().Use(echomiddleware.Recover())
	}

	// 注册后台路由
	router.AdminRegister(b)

	// 注册Web路由
	router.WebRegister(b)

	// 启动服务
	b.Run(config.App.Host)
}
