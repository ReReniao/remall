package service

import (
	"context"
	"re-mall/dao"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
)

type CategoryService struct {
}

func (s *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(category), uint(len(category)))
}
