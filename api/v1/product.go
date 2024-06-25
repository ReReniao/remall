package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// CreateProductHandler 创建商品
func CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		var createProductService service.ProductService
		if err := c.ShouldBind(&createProductService); err == nil {
			res := createProductService.Create(c.Request.Context(), files)
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ListProductHandler 获取商品列表
func ListProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listProductService service.ProductService
		if err := c.ShouldBind(&listProductService); err == nil {
			res := listProductService.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}

// SearchProductHandler 搜索商品
func SearchProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchProductService service.ProductService
		if err := c.ShouldBind(&searchProductService); err == nil {
			res := searchProductService.Search(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}

// ShowProductHandler 展示商品信息
func ShowProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var showProductService service.ProductService
		if err := c.ShouldBind(&showProductService); err == nil {
			res := showProductService.Show(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}
