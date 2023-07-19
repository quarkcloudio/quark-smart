package database

import (
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/quarkcms/quark-smart/internal/model"
)

// 执行数据库操作
func Handle() {

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Post{},
		&model.Category{},
		&model.Banner{},
		&model.BannerCategory{},
		&model.Navigation{},
	)

	// 数据填充
	(&model.Post{}).Seeder()
	(&model.Category{}).Seeder()
	(&model.Banner{}).Seeder()
	(&model.BannerCategory{}).Seeder()
	(&model.Navigation{}).Seeder()
}
