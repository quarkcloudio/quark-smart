package config

import (
	"github.com/quarkcms/quark-easy/pkg/env"
)

var Redis = map[string]interface{}{

	// 地址
	"host": env.Get("REDIS_HOST", "127.0.0.1"),

	// 密码
	"password": env.Get("REDIS_PASSWORD"),

	// 端口
	"port": env.Get("REDIS_PORT", "6379"),

	// 数据库
	"database": 0,
}
