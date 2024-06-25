package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartById(aId uint) (Cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", aId).First(&Cart).
		Error
	return
}

func (dao *CartDao) GetCartByUserId(uId uint) (Carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&Carts).Error
	return
}

func (dao *CartDao) UpdateCartNumByUserId(cId, uId uint, num uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id  = ?", cId, uId).Update("num", num).Error
}

func (dao *CartDao) DeleteCartById(uId, cId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cId, uId).Delete(&model.Cart{}).Error
}

func (dao *CartDao) ListCartByUserId(uId uint) (Carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&Carts).Error
	return
}
