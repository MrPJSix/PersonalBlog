package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategorys(pageSize int, pageNum int) []Category {
	var categorys []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categorys
}

// 编辑分类信息
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err := db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
