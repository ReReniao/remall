package service

import (
	"context"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

func (s *CartService) Create(ctx context.Context) serializer.Response {
	var cart *model.Cart
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(s.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(s.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    userInfo.Id,
		ProductId: s.ProductId,
		BossId:    s.BossId,
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCart(cart, product, boss),
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

func (s *CartService) List(ctx context.Context) serializer.Response {
	var carts []*model.Cart
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(ctx)
	carts, err = cartDao.ListCartByUserId(userInfo.Id)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCarts(ctx, carts),
		Msg:    e.GetMsg(code),
	}
}

func (s *CartService) Update(ctx context.Context, cId string) serializer.Response {
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(ctx)
	CartId, _ := strconv.Atoi(cId)
	err = cartDao.UpdateCartNumByUserId(uint(CartId), userInfo.Id, s.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
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

func (s *CartService) Delete(ctx context.Context, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(ctx)
	err = cartDao.DeleteCartById(userInfo.Id, uint(cartId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
