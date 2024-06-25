package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/pkg/util"
	"re-mall/service"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRegister service.UserService
		if err := c.ShouldBind(&userRegister); err == nil {
			res := userRegister.Register(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Register:", err)
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLogin service.UserService
		if err := c.ShouldBind(&userLogin); err == nil {
			res := userLogin.Login(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Login:", err)

		}
	}
}

func UserUpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userUpdate service.UserService
		if err := c.ShouldBind(&userUpdate); err == nil {
			res := userUpdate.Update(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user update：", err)

		}
	}
}

func UploadAvatarHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, fileHeader, _ := c.Request.FormFile("file")
		filesize := fileHeader.Size
		var uploadAvatar service.UserService
		if err := c.ShouldBind(&uploadAvatar); err == nil {
			res := uploadAvatar.Post(c.Request.Context(), file, filesize)
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Upload Avatar：", err)

		}
	}
}

func SendEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sendEmail service.SendEmailService
		if err := c.ShouldBind(&sendEmail); err == nil {
			res := sendEmail.Send(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Send Email：", err)

		}
	}
}

func ValidEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var validEmail service.ValidEmailService
		if err := c.ShouldBind(&validEmail); err == nil {
			res := validEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Valid Email：", err)

		}
	}
}

func ShowMoneyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var showMoney service.ShowMoneyService
		if err := c.ShouldBind(&showMoney); err == nil {
			res := showMoney.Show(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.LogrusObj.Infoln("user Show Money:", err)

		}
	}
}
