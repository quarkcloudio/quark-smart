package utils

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/model"
)

// 获取文件路径
func GetFilePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return (&model.File{}).GetPath(id)
}

// 获取多文件路径
func GetFilePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return (&model.File{}).GetPaths(id)
}

// 获取图片路径
func GetPicturePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return (&model.Picture{}).GetPath(id)
}

// 获取多图片路径
func GetPicturePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return (&model.Picture{}).GetPaths(id)
}
