package utils

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
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

// 获取配置
func GetConfig(key string) string {
	return (&model.Config{}).GetValue(key)
}

// 获取域名
func GetDomain() string {

	http := ""
	domain := (&model.Config{}).GetValue("WEB_SITE_DOMAIN")
	ssl := (&model.Config{}).GetValue("SSL_OPEN")
	if domain != "" {
		if ssl == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	return http + domain
}
