package main

import (
	"github.com/quarkcms/quark-easy/config"
	"github.com/quarkcms/quark-easy/database"
	"github.com/quarkcms/quark-easy/internal/admin"
	"github.com/quarkcms/quark-easy/internal/middleware"
	"github.com/quarkcms/quark-easy/internal/router"
	appproviders "github.com/quarkcms/quark-go/pkg/app/handler/admin"
	appinstall "github.com/quarkcms/quark-go/pkg/app/install"
	appmiddleware "github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 配置信息
	var (
		appName    = config.App["name"].(string)
		appKey     = config.App["key"].(string)
		dbUser     = config.Mysql["username"].(string)
		dbPassword = config.Mysql["password"].(string)
		dbHost     = config.Mysql["host"].(string)
		dbPort     = config.Mysql["port"].(string)
		dbName     = config.Mysql["database"].(string)
		dbCharset  = config.Mysql["charset"].(string)
	)

	// 数据库配置信息
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{
		AppKey: appKey,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		Providers: append(appproviders.Providers, admin.Providers...),
		AdminLayout: &builder.AdminLayout{
			Title: appName,
		},
	}

	// 实例化对象
	b := builder.New(config)

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
	b.Run(":3000")
}
