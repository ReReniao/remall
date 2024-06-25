package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"re-mall/api/v1"
	middleware "re-mall/middlewave"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 轮播图
		v1.GET("carousels", api.ListCarouselHandler())

		// 商品操作
		v1.GET("products", api.ListProductHandler())
		v1.GET("products/:id", api.ShowProductHandler())
		v1.GET("imgs/:id", api.ListProductImgHandler())
		v1.GET("categories", api.ListCategoryHandler())

		authed := v1.Group("/") // 需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdateHandler())
			authed.POST("avatar", api.UploadAvatarHandler())
			authed.POST("user/sending-email", api.SendEmailHandler())
			authed.POST("user/valid-email", api.ValidEmailHandler())

			// 显示金额
			authed.POST("money", api.ShowMoneyHandler())

			// 商品操作
			authed.POST("product", api.CreateProductHandler())
			authed.POST("products", api.SearchProductHandler())

			// 收藏夹操作
			authed.GET("favorites", api.ListFavoriteHandler())
			authed.POST("favorites", api.CreateFavoriteHandler())
			authed.DELETE("favorites/:id", api.DeleteFavoriteHandler())

			// 地址操作
			authed.POST("addresses", api.CreateAddressHandler())
			authed.GET("addresses/:id", api.ShowAddressHandler())
			authed.GET("addresses", api.ListAddressHandler())
			authed.PUT("addresses/:id", api.UpdateAddressHandler())
			authed.DELETE("addresses/:id", api.DeleteAddressHandler())

			// 购物车操作
			authed.POST("carts", api.CreateCartHandler())
			authed.GET("carts", api.ListCartHandler())
			authed.PUT("carts/:id", api.UpdateCartHandler())
			authed.DELETE("carts/:id", api.DeleteCartHandler())

			// 订单操作
			authed.POST("orders", api.CreateOrderHandler())
			authed.GET("orders", api.ListOrderHandler())
			authed.GET("orders/:id", api.ShowOrderHandler())
			authed.DELETE("orders/:id", api.DeleteOrderHandler())

			// 支付功能
			authed.POST("paydown", api.OrderPayHandler())
		}
	}
	return r
}
