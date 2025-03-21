package upload

import (
	"github.com/quarkcloudio/quark-go/v3"
	appupload "github.com/quarkcloudio/quark-go/v3/app/tool/upload"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

type Image struct {
	appupload.Image
}

// 初始化
func (p *Image) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = config.App.UploadImageSize

	// 限制文件类型
	p.LimitType = config.App.UploadImageType

	// 设置文件上传路径
	p.SavePath = config.App.UploadImageSavePath

	return p
}
