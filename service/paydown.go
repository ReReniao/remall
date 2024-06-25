package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strconv"
)

type OrderPay struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"`
}

func (s *OrderPay) PayDown(ctx context.Context) serializer.Response {
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
	err = dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {
		util.Encrypt.SetKey(s.Key)
		uId := userInfo.Id
		orderDao := dao.NewOrderDaoByDB(tx)

		order, err := dao.NewOrderDao(ctx).GetOrderById(s.OrderId, uId)
		if err != nil {
			util.LogrusObj.Info(err)
			return err
		}
		money := order.Money
		num := order.Num
		money = money * float64(num)

		userDao := dao.NewUserDaoByDB(tx)
		user, err := userDao.GetUserById(uId)
		if err != nil {
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		// 对钱进行解密。减去订单。再进行加密。
		moneyStr := util.Encrypt.AesDecoding(user.Money)
		moneyFloat, _ := strconv.ParseFloat(user.Money, 64)
		if moneyFloat-money < 0.0 {
			util.LogrusObj.Info(err)
			code = e.Error
			return errors.New("金额不足")
		}

		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(uId, user)
		if err != nil { // 更新用户金额失败，回滚
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		boss := new(model.User)
		boss, err = userDao.GetUserById(s.BossId)
		moneyStr = util.Encrypt.AesDecoding(boss.Money)
		moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
		finMoney = fmt.Sprintf("%f", moneyFloat+money)
		boss.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(s.BossId, boss)
		if err != nil { // 更新boss金额失败，回滚
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		product := new(model.Product)
		productDao := dao.NewProductDaoByDB(tx)
		product, err = productDao.GetProductById(s.ProductId)
		if err != nil {
			return err
		}

		product.Num -= num
		err = productDao.UpdateProductById(s.ProductId, product)
		if err != nil {
			// 更新商品数量减少失败，回滚
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		order.Type = 2
		err = orderDao.UpdateOrderById(s.OrderId, order)
		if err != nil { // 更新订单失败，回滚
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		productUser := model.Product{
			Name:          product.Name,
			CategoryId:    product.CategoryId,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Num:           num,
			OnSale:        false,
			BossId:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}

		err = productDao.CreateProduct(&productUser)
		if err != nil { // 买完商品后创建成了自己的商品失败。订单失败，回滚
			util.LogrusObj.Info(err)
			code = e.Error
			return err
		}

		return nil
	})

	if err != nil {
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
