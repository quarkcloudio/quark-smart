package upload

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/upload"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/storage"
	"github.com/quarkcms/quark-smart/config"
)

type File struct {
	upload.Template
}

// 初始化
func (p *File) Init(ctx *builder.Context) interface{} {

	// 限制文件大小
	p.LimitSize = config.App.UploadFileSize

	// 限制文件类型
	p.LimitType = config.App.UploadFileType

	// 设置文件上传路径
	p.SavePath = config.App.UploadFileSavePath

	return p
}

// 上传前回调
func (p *File) BeforeHandle(ctx *builder.Context, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	getFileInfo, _ := (&model.File{}).GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if getFileInfo.Id != 0 {
		fileInfo := &storage.FileInfo{
			Name: getFileInfo.Name,
			Size: getFileInfo.Size,
			Ext:  getFileInfo.Ext,
			Path: getFileInfo.Path,
			Url:  getFileInfo.Url,
			Hash: getFileInfo.Hash,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *File) AfterHandle(ctx *builder.Context, result *storage.FileInfo) error {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").
		String()

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = (&model.File{}).GetPath(result.Url)
	}

	adminInfo, err := (&model.Admin{}).GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 插入数据库
	id, err := (&model.File{}).InsertGetId(&model.File{
		ObjType: "ADMINID",
		ObjId:   adminInfo.Id,
		Name:    result.Name,
		Size:    result.Size,
		Ext:     result.Ext,
		Path:    result.Path,
		Url:     result.Url,
		Hash:    result.Hash,
		Status:  1,
	})
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success(
		"上传成功",
		"",
		map[string]interface{}{
			"id":          id,
			"contentType": result.ContentType,
			"ext":         result.Ext,
			"hash":        result.Hash,
			"name":        result.Name,
			"path":        result.Path,
			"size":        result.Size,
			"url":         result.Url,
		},
	))
}
