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
}
