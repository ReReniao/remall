package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (s *ProductService) Create(ctx context.Context, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	bossInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	useDao := dao.NewUserDao(ctx)
	boss, _ = useDao.GetUserById(bossInfo.Id)

	// 以第一张图作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, bossInfo.Id, s.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("ProductService.Create:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          s.Name,
		CategoryId:    s.CategoryId,
		Title:         s.Title,
		Info:          s.Info,
		ImgPath:       path,
		Price:         s.Price,
		DiscountPrice: s.DiscountPrice,
		OnSale:        false,
		Num:           s.Num,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, boss.ID, s.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Data:   nil,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Data:   nil,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (s *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success
	if s.PageSize == 0 {
		s.PageSize = 15
	}
	condition := make(map[string]interface{})
	if s.CategoryId != 0 {
		condition["category_id"] = s.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, s.BasePage)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (s *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if s.PageSize == 0 {
		s.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, total, err := productDao.SearchProduct(s.Info, s.BasePage)
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
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))

}

func (s *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pId, _ := strconv.Atoi(id)
	fmt.Println(pId)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
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
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}
