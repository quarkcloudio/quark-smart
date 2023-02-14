package main

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-simple/database"
	adminproviders "github.com/quarkcms/quark-simple/internal/admin"
	"github.com/quarkcms/quark-simple/internal/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 数据库配置信息
	dsn := "root:Bc5HQFJc4bLjZCcC@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{
		AppKey: "123456",
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		Providers: append(admin.Providers, adminproviders.Providers...),
		AdminLayout: &builder.AdminLayout{
			Title: "QuarkEasy",
		},
	}

	// 实例化对象
	b := builder.New(config)

	// 静态文件目录
	b.Static("/", "./website")

	// 构建quarkgo基础数据库、拉取静态文件
	b.Use(install.Handle)

	// 构建本项目数据库
	b.Use(database.Handle)

	// 后台中间件
	b.Use(middleware.Handle)

	// 路由
	b.GET("/", (&handlers.Home{}).Index)

	// 启动服务
	b.Run(":3000")
}
