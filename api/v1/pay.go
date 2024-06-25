package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

func OrderPayHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderPay := service.OrderPay{}
		if err := c.ShouldBind(&orderPay); err == nil {
			res := orderPay.PayDown(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("List Product:", err)
		}
	}
}
