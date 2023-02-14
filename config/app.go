package config

import (
	"github.com/quarkcms/quark-easy/pkg/env"
)

var App = map[string]interface{}{

	// 应用名称
	"name": env.Get("APP_NAME", "QuarkEasy"),

	// 服务地址
	"host": env.Get("APP_HOST", "127.0.0.1:3000"),

	// 令牌加密key，如果设置绝对不可泄漏
	"key": env.Get("APP_KEY"),
}
