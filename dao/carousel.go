package dao

import (
	"context"
	"gorm.io/gorm"
	"re-mall/model"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// ListCarousel 根据Id获取carousel
func (dao *CarouselDao) ListCarousel() (carousel []*model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
