package controllers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowProductPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var product []models.Products
	if err := models.DB.Find(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Product"})
		return
	}

	c.HTML(http.StatusOK, "products.html", gin.H{"products": product})
}
