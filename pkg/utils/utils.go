package utils

import (
	"github.com/quarkcms/quark-go/pkg/app/model"
)

// 获取文件路径
func GetFilePath(id interface{}) string {
	return (&model.File{}).GetPath(id)
}

// 获取多文件路径
func GetFilePaths(id interface{}) []string {
	return (&model.File{}).GetPaths(id)
}

// 获取图片路径
func GetPicturePath(id interface{}) string {
	return (&model.Picture{}).GetPath(id)
}

// 获取多图片路径
func GetPicturePaths(id interface{}) []string {
	return (&model.Picture{}).GetPaths(id)
}
