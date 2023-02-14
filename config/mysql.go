package config

import (
	"github.com/quarkcms/quark-easy/pkg/env"
)

var Mysql = map[string]interface{}{
	// 地址
	"host": env.Get("DB_HOST", "127.0.0.1"),

	// 端口
	"port": env.Get("DB_PORT", "3306"),

	// 数据库
	"database": env.Get("DB_DATABASE", "quarkgo"),

	// 用户名
	"username": env.Get("DB_USERNAME", "root"),

	// 密码
	"password": env.Get("DB_PASSWORD", "root"),

	// 编码
	"charset": "utf8",
}
