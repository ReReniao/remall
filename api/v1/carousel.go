package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

// ListCarouselHandler 轮播图
func ListCarouselHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var listCarousel service.CarouselService
		if err := c.ShouldBind(&listCarousel); err == nil {
			res := listCarousel.List(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("ListCarouselHandler:", err)

		}
	}
}
