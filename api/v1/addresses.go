package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// CreateAddressHandler 创建地址
func CreateAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createAddressService := service.AddressService{}
		if err := c.ShouldBind(&createAddressService); err == nil {
			res := createAddressService.Create(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// UpdateAddressHandler 创建地址
func UpdateAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createAddressService := service.AddressService{}
		if err := c.ShouldBind(&createAddressService); err == nil {
			res := createAddressService.Update(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ListAddressHandler 获取地址列表
func ListAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listAddressService service.AddressService
		if err := c.ShouldBind(&listAddressService); err == nil {
			res := listAddressService.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}

// DeleteAddressHandler 删除地址
func DeleteAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteAddressService service.AddressService
		if err := c.ShouldBind(&deleteAddressService); err == nil {
			res := deleteAddressService.Delete(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}

// ShowAddressHandler 展示地址信息
func ShowAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var showAddressService service.AddressService
		if err := c.ShouldBind(&showAddressService); err == nil {
			res := showAddressService.Show(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}
