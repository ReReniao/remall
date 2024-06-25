package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// CreateFavoriteHandler 创建收藏夹
func CreateFavoriteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		createFavoriteService := service.FavoriteService{}
		if err := c.ShouldBind(&createFavoriteService); err == nil {
			fmt.Println(createFavoriteService.BossId, createFavoriteService.ProductId)
			res := createFavoriteService.Create(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Create Product:", err)
		}
	}
}

// ListFavoriteHandler 获取收藏夹列表
func ListFavoriteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listFavoriteService service.FavoriteService
		if err := c.ShouldBind(&listFavoriteService); err == nil {
			res := listFavoriteService.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}

// DeleteFavoriteHandler 删除收藏夹
func DeleteFavoriteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteFavoriteService service.FavoriteService
		if err := c.ShouldBind(&deleteFavoriteService); err == nil {
			res := deleteFavoriteService.Delete(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("Search Product:", err)
		}
	}
}
