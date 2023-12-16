package config

import (
	"github.com/quarkcms/quark-smart/pkg/env"
)

type AppConfig struct {
	Version        string // 应用版本
	Name           string // 应用名称
	Debug          bool   // 开启Debug模式
	Recover        bool   // 崩溃后自动恢复
	Env            string // 项目环境
	Host           string // 服务地址
	Key            string // 令牌加密key，如果设置绝对不可泄漏
	RootPath       string // Web根目录
	StaticPath     string // 静态文件路径
	TemplatePath   string // 模版文件路径
	Logger         bool   // 是否开启日志
	LoggerFilePath string // 日志文件路径
}

// APP配置信息
var App = &AppConfig{

	// 应用版本
	Version: "1.0.0",

	// 应用名称
	Name: env.Get("APP_NAME", "QuarkSmart").(string),

	// 开启Debug模式
	Debug: false,

	// 崩溃后自动恢复
	Recover: true,

	// 项目环境
	Env: env.Get("APP_ENV").(string),

	// 服务地址
	Host: env.Get("APP_HOST", "127.0.0.1:3000").(string),

	// 令牌加密key，如果设置绝对不可泄漏
	Key: env.Get("APP_KEY").(string),

	// Web根目录
	RootPath: env.Get("APP_ROOT_PATH", "./web/app").(string),

	// 静态文件路径
	StaticPath: env.Get("APP_STATIC_PATH", "./web/static").(string),

	// 模版文件路径
	TemplatePath: env.Get("APP_TEMPLATE_PATH", "web/template/*.html").(string),

	// 是否开启日志
	Logger: false,

	// 日志文件路径
	LoggerFilePath: "./app.log",
}
