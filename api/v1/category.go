package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// ListCategoryHandler 轮播图
func ListCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listCategory service.CategoryService
		if err := c.ShouldBind(&listCategory); err == nil {
			res := listCategory.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("ListCarouselHandler:", err)
		}
	}
}
