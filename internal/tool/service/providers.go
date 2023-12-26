package service

import "github.com/quarkcloudio/quark-smart/internal/tool/service/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
