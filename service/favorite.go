package service

import (
	"context"
	"fmt"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (s *FavoriteService) List(ctx context.Context) serializer.Response {
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favoriteDao := dao.NewFavoriteDao(ctx)
	favorites, err := favoriteDao.ListFavorite(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}

func (s *FavoriteService) Create(ctx context.Context) serializer.Response {
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	fmt.Println(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(userInfo.Id, s.ProductId)
	if exist {
		code = e.ErrorFavoriteExist
		util.LogrusObj.Infoln("err", e.GetMsg(code))
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userInfo.Id)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(s.BossId)
	if err != nil {
		code = e.Error
		fmt.Println(s.BossId)
		util.LogrusObj.Infoln("err", err)

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(s.ProductId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserId:    userInfo.Id,
		Product:   *product,
		ProductId: s.ProductId,
		Boss:      *boss,
		BossId:    s.BossId,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (s *FavoriteService) Delete(ctx context.Context, fId string) serializer.Response {
	code := e.Success
	favoriteId, _ := strconv.Atoi(fId)
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favoriteDao := dao.NewFavoriteDao(ctx)
	err = favoriteDao.DeleteFavorite(userInfo.Id, uint(favoriteId))
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   nil,
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}
