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

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (s *AddressService) Create(ctx context.Context) serializer.Response {
	var address *model.Address
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
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserId:  userInfo.Id,
		Name:    s.Name,
		Phone:   s.Phone,
		Address: s.Address,
	}
	err = addressDao.CreateAddress(address)
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

func (s *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(uint(addressId))
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
		Data:   serializer.BuildAddress(address),
		Msg:    e.GetMsg(code),
	}
}

func (s *AddressService) List(ctx context.Context) serializer.Response {
	var addresses []*model.Address
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
	addressDao := dao.NewAddressDao(ctx)
	addresses, err = addressDao.GetAddressByUserId(userInfo.Id)
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
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}

func (s *AddressService) Update(ctx context.Context, aId string) serializer.Response {
	var address *model.Address
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
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserId:  userInfo.Id,
		Name:    s.Name,
		Phone:   s.Phone,
		Address: s.Address,
	}
	addressId, _ := strconv.Atoi(aId)
	err = addressDao.UpdateAddressByUserId(address, uint(addressId), userInfo.Id)
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

func (s *AddressService) Delete(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
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
	addressDao := dao.NewAddressDao(ctx)
	err = addressDao.DeleteAddressById(userInfo.Id, uint(addressId))
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
