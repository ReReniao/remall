package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// CreateCartHandler 创建订单
func CreateCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createCartService := service.CartService{}
		if err := c.ShouldBind(&createCartService); err == nil {
			res := createCartService.Create(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// UpdateCartHandler 更新订单
func UpdateCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createCartService := service.CartService{}
		if err := c.ShouldBind(&createCartService); err == nil {
			res := createCartService.Update(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ListCartHandler 获取订单列表
func ListCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listCartService service.CartService
		if err := c.ShouldBind(&listCartService); err == nil {
			res := listCartService.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}

// DeleteCartHandler 删除订单
func DeleteCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteCartService service.CartService
		if err := c.ShouldBind(&deleteCartService); err == nil {
			res := deleteCartService.Delete(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}
