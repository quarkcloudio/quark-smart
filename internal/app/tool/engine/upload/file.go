package upload

import (
	"github.com/quarkcloudio/quark-go/v3"
	appupload "github.com/quarkcloudio/quark-go/v3/app/tool/upload"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

type File struct {
	appupload.File
}

// 初始化
func (p *File) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = config.App.UploadFileSize

	// 限制文件类型
	p.LimitType = config.App.UploadFileType

	// 设置文件上传路径
	p.SavePath = config.App.UploadFileSavePath

	return p
}
