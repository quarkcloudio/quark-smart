package config

import (
	"github.com/quarkcms/quark-easy/pkg/env"
)

type AppConfig struct {
	Name    string // 应用名称
	Version string // 应用版本
	Host    string // 服务地址
	Key     string // 令牌加密key，如果设置绝对不可泄漏
}

// APP配置信息
var App = &AppConfig{

	// 应用名称
	Name: env.Get("APP_NAME", "QuarkEasy").(string),

	// 应用版本
	Version: "1.0.0",

	// 服务地址
	Host: env.Get("APP_HOST", "127.0.0.1:3000").(string),

	// 令牌加密key，如果设置绝对不可泄漏
	Key: env.Get("APP_KEY").(string),
}