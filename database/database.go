package database

import (
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-smart/internal/model"
)

// 执行数据库操作
func Handle() {

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Post{},
		&model.Category{},
	)

	// 数据填充
	(&model.Post{}).Seeder()
	(&model.Category{}).Seeder()
}
