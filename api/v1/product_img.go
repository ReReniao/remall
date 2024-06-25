package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// ListProductImgHandler 获取图片地址
func ListProductImgHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listProductImg service.ProductImgService
		if err := c.ShouldBind(&listProductImg); err == nil {
			res := listProductImg.List(c.Request.Context(), c.Param("id"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List ProductImg:", err)
		}
	}
}
