package model

import (
	"myblog/utils/errcode"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory  查询分类是否存在
func CheckCategory(name string) *errcode.Error {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	//大于0  用户已存在 返回状态码
	if cate.ID > 0 {
		return errcode.ErrorCateNameUsed
	}
	return errcode.Success
}

// CreateCategory 新增分类
func CreateCategory(data *Category) *errcode.Error {
	err := db.Create(data).Error
	if err != nil {
		return errcode.ServerError // 500
	}
	return errcode.Success
}

// GetCate  查询分类列表
func GetCate(pageSize, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// EditCategory  编辑分类信息
func EditCategory(id int, data *Category) *errcode.Error {
	var cate Category
	var maps = make(map[string]any)
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errcode.ServerError
	}
	return errcode.Success
}

// DeleteCategory  删除分类
func DeleteCategory(id int) *errcode.Error {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errcode.ServerError
	}
	return errcode.Success
}
