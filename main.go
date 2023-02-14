package main

import (
	appproviders "github.com/quarkcms/quark-go/pkg/app/handler/admin"
	appinstall "github.com/quarkcms/quark-go/pkg/app/install"
	appmiddleware "github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-smart/config"
	"github.com/quarkcms/quark-smart/database"
	"github.com/quarkcms/quark-smart/internal/admin"
	"github.com/quarkcms/quark-smart/internal/middleware"
	"github.com/quarkcms/quark-smart/internal/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

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

	// 配置资源
	getConfig := &builder.Config{
		AppKey: appKey,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		Providers: append(appproviders.Providers, admin.Providers...),
		AdminLayout: &builder.AdminLayout{
			Title:        config.Admin.Title,
			Logo:         config.Admin.Logo,
			Layout:       config.Admin.Layout,
			SplitMenus:   config.Admin.SplitMenus,
			ContentWidth: config.Admin.ContentWidth,
			PrimaryColor: config.Admin.PrimaryColor,
			FixedHeader:  config.Admin.FixedHeader,
			FixSiderbar:  config.Admin.FixSiderbar,
			IconfontUrl:  config.Admin.IconfontUrl,
			Locale:       config.Admin.Locale,
			SiderWidth:   config.Admin.SiderWidth,
			Copyright:    config.Admin.Copyright,
			Links:        config.Admin.Links,
		},
	}

	// 实例化对象
	b := builder.New(getConfig)

	// 静态文件目录
	b.Static("/", "./website")

	// 构建quarkgo基础数据库、拉取静态文件
	b.Use(appinstall.Handle)

	// 构建本项目数据库
	b.Use(database.Handle)

	// 后台中间件
	b.Use(appmiddleware.Handle)

	// 中间件
	b.Use(middleware.Handle)

	// 注册路由
	router.Register(b)

	// 启动服务
	b.Run(config.App.Host)
}
