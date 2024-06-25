package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) ListProductImg(pId uint) (productImg []*model.ProductImg, total int64, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", pId).Find(&productImg).Count(&total).Error
	return
}
