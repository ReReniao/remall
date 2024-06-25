package service

import (
	"context"
	"fmt"
	"math/rand"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"order_id"`
	OrderNum  int     `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (s *OrderService) Create(ctx context.Context) serializer.Response {
	var order *model.Order
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
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserId:    userInfo.Id,
		ProductId: s.ProductId,
		BossId:    s.BossId,
		Num:       s.Num,
		Money:     s.Money,
		Type:      1, // 未支付
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(s.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressID = address.ID

	// 订单号创建
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(s.ProductId))
	userNum := strconv.Itoa(int(s.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	err = orderDao.CreateOrder(order)
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

func (s *OrderService) Show(ctx context.Context, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
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
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uint(orderId), userInfo.Id)
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
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(order.ProductId)
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
		Data:   serializer.BuildOrder(order, product, address),
		Msg:    e.GetMsg(code),
	}
}

func (s *OrderService) List(ctx context.Context) serializer.Response {
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
	if s.PageSize == 0 {
		s.PageSize = 15
	}
	orderDao := dao.NewOrderDao(ctx)

	condition := make(map[string]interface{})
	if s.Type != 0 {
		condition["type"] = s.Type
	}
	condition["user_id"] = userInfo.Id
	orders, total, err := orderDao.GetOrderByCondition(condition, s.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))
}

func (s *OrderService) Delete(ctx context.Context, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
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
	OrderDao := dao.NewOrderDao(ctx)
	err = OrderDao.DeleteOrderById(userInfo.Id, uint(orderId))
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
