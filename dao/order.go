package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(DB *gorm.DB) *OrderDao {
	return &OrderDao{DB}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(oId, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oId, uId).First(&order).
		Error
	return
}

func (dao *OrderDao) GetOrderByUserId(uId uint) (orders []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id = ?", uId).Find(&orders).Error
	return
}

func (dao *OrderDao) ShowOrderByUserId(order *model.Order, oId uint, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oId, uId).First(&order).Error
}

func (dao *OrderDao) DeleteOrderById(uId, oId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oId, uId).Delete(&model.Order{}).Error
}

func (dao *OrderDao) GetOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, count int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Count(&count).
		Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&orders).
		Error
	return
}

func (dao *OrderDao) UpdateOrderById(oId uint, order *model.Order) error {
	return dao.Where("id = ? and user_id").Updates(order).Error
}
