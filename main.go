package main

import (
	"ecommerce/models"
	"ecommerce/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	models.ConnectDatabase()
	r.LoadHTMLGlob("template/*")
	routes.SetupRoutes(r)
	r.Run(":8081")
}
