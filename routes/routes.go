package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/register", controllers.ShowRegisterPage)
	r.POST("/register", controllers.Registerdata)
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.Logindata)
	r.GET("/logout", controllers.LogoutUser)
	auth := r.Group("user")
	auth.Use()
	{
		auth.GET("products", controllers.ShowProductPage)
		auth.POST("/cart/add/:id", controllers.AddToCart)
		auth.GET("/cart", controllers.CetCart)
		auth.GET("/order", controllers.PlaceOrder)
	}
}
