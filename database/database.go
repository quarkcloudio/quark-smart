package database

import (
	"os"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-smart/internal/model"
)

// 判断路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

// 执行数据库操作
func Handle(ctx *builder.Context) error {

	// 如果锁定文件存在则不执行操作
	if PathExist("install.lock") {
		return nil
	}

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Post{},
	)

	// 如果超级管理员不存在，初始化数据库数据
	adminInfo, err := (&appmodel.Admin{}).GetInfoById(1)
	if err != nil && err.Error() != "record not found" {
		panic(err)
	}
	if adminInfo.Id == 0 {
		// 数据填充
		(&model.Post{}).Seeder()
	}

	return nil
}
