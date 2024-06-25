package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// CreateOrderHandler 创建订单
func CreateOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createOrderService := service.OrderService{}
		if err := c.ShouldBind(&createOrderService); err == nil {
			res := createOrderService.Create(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ShowOrderHandler 更新订单
func ShowOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		showOrderService := service.OrderService{}
		if err := c.ShouldBind(&showOrderService); err == nil {
			res := showOrderService.Show(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ListOrderHandler 获取订单列表
func ListOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listOrderService service.OrderService
		if err := c.ShouldBind(&listOrderService); err == nil {
			res := listOrderService.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}

// DeleteOrderHandler 删除订单
func DeleteOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteOrderService service.OrderService
		if err := c.ShouldBind(&deleteOrderService); err == nil {
			res := deleteOrderService.Delete(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}
