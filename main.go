package main

import (
	"ecommerce/models"
	"ecommerce/routes"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func mul(a, b int) int {
	return a * b
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	models.ConnectDatabase()
	r.SetFuncMap(template.FuncMap{
		"mul": mul,
	})
	r.LoadHTMLGlob("template/*")
	routes.SetupRoutes(r)
	r.Run(":8081")
}
