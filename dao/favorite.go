package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// ListFavorite 展示收藏夹
func (dao *FavoriteDao) ListFavorite(uId uint) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uId).Find(&favorites).Error
	return
}

// FavoriteExistOrNot 是否已经收藏
func (dao *FavoriteDao) FavoriteExistOrNot(uId, pId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND product_id = ?", uId, pId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, err
}

// CreateFavorite 创建收藏夹
func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) (err error) {
	err = dao.DB.Where(&model.Favorite{}).Create(&in).Error
	return
}

func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) (err error) {
	return dao.DB.Model(&model.Favorite{}).Where("id = ? AND user_id = ?", fId, uId).
		Delete(&model.Favorite{}).Error
}
