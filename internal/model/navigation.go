package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/lister"
	"gorm.io/gorm"
)

// 导航
type Navigation struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Pid       int            `json:"pid"`
	Title     string         `json:"title" gorm:"size:200;not null"`
	CoverId   string         `json:"cover_id" gorm:"size:500;default:null"`
	Sort      int            `json:"sort" gorm:"size:11;default:0;"`
	UrlType   int            `json:"url_type" gorm:"size:1;not null;default:1"`
	Url       string         `json:"url" gorm:"size:200;not null"`
	Status    int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Navigation) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(26) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 26, Name: "导航管理", GuardName: "admin", Icon: "", Type: "engine", Pid: 7, Sort: 0, Path: "/api/admin/navigation/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Navigation{
		{Title: "默认导航", Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取菜单的有序列表
func (model *Navigation) OrderedList(root bool) (list []map[string]interface{}, Error error) {
	var data []map[string]interface{}
	err := db.Client.
		Model(&model).
		Order("sort asc,id asc").
		Find(&data).Error
	if err != nil {
		return list, err
	}

	trees, err := lister.ListToTree(data, "id", "pid", "children", 0)
	if err != nil {
		return list, err
	}

	treeList, err := lister.TreeToOrderedList(trees, 0, "title", "children")
	if err != nil {
		return list, err
	}

	// 是否有跟节点
	if root {
		list = append(list, map[string]interface{}{
			"label": "根节点",
			"value": 0,
		})
	}

	for _, v := range treeList {
		option := map[string]interface{}{
			"label": v.((map[string]interface{}))["title"],
			"value": v.(map[string]interface{})["id"],
		}
		list = append(list, option)
	}

	return list, nil
}
