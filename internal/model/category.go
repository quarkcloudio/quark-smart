package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/lister"
	"gorm.io/gorm"
)

// 分类模型
type Category struct {
	Id          int            `json:"id" gorm:"autoIncrement"`
	Pid         int            `json:"pid"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Sort        int            `json:"sort" gorm:"size:11;default:0;"`
	CoverId     string         `json:"cover_id" gorm:"size:500;default:null"`
	Name        string         `json:"name" gorm:"size:100;default:null"`
	Description string         `json:"description" gorm:"size:500;default:null"`
	Count       int            `json:"count" gorm:"size:11;default:10;"`
	IndexTpl    string         `json:"index_tpl" gorm:"size:100;"`
	ListTpl     string         `json:"list_tpl" gorm:"size:100;"`
	DetailTpl   string         `json:"detail_tpl" gorm:"size:100;"`
	PageNum     int            `json:"page_num" gorm:"size:11;default:10;"`
	Type        string         `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	Status      int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Category) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(19) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 19, Name: "分类列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/category/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Category{
		{Title: "默认分类", Name: "default", Type: "ARTICLE", Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取SelectTree组件的数据
func (m *Category) SelectTreeData(root bool) (list []interface{}, Error error) {
	categories := []Category{}
	err := db.Client.Where("status = ?", 1).Select("title", "id", "pid").Find(&categories).Error
	if err != nil {
		return list, err
	}

	treeList := []map[string]interface{}{}

	// 是否有跟节点
	if root {
		list = append(list, map[string]interface{}{
			"label": "根节点",
			"value": 0,
		})
	}

	for _, v := range categories {
		item := map[string]interface{}{
			"value": v.Id,
			"pid":   v.Pid,
			"label": v.Title,
		}
		treeList = append(treeList, item)
	}

	tree, err := lister.ListToTree(treeList, "value", "pid", "children", 0)
	if err != nil {
		return list, err
	}
	list = append(list, tree...)

	return list, err
}

// 获取搜索框Select的属性
func (model *Category) Options() (list map[interface{}]interface{}, Error error) {
	options := map[interface{}]interface{}{}
	getList := []Category{}
	err := db.Client.Find(&getList).Error
	if err != nil {
		return options, err
	}
	for _, v := range getList {
		options[v.Id] = v.Title
	}

	return options, nil
}
