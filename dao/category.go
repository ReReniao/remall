package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 展示轮播图
func (dao *CategoryDao) ListCategory() (category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return
}
