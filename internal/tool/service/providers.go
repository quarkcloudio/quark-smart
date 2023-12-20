package service

import "github.com/quarkcms/quark-smart/internal/tool/service/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
