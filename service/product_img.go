package service

import (
	"context"
	"re-mall/dao"
	"re-mall/serializer"
	"strconv"
)

type ProductImgService struct {
}

func (s *ProductImgService) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productImgId, _ := strconv.Atoi(pId)
	productImgs, total, _ := productImgDao.ListProductImg(uint(productImgId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(total))
}
