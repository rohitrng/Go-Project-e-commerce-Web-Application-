package controllers

import (
	"ecommerce/models"
	"net/http"
	"strconv"

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

func AddToCart(c *gin.Context) {
	session := sessions.Default(c)
	userid, ok := session.Get("user_id").(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
		return
	}

	productDstr := c.Param("id")
	productID, err := strconv.ParseUint(productDstr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var product models.Products
	if err := models.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
		return
	}

	var cart models.Cart
	if err := models.DB.Where("user_id = ? AND product_id = ?", userid, uint(productID)).First(&cart).Error; err != nil {
		if err.Error() == "record not found" {
			cart = models.Cart{
				UserID:    userid,
				ProductID: uint(productID),
				Quantity:  1,
				Price:     product.Price,
			}
			models.DB.Create(&cart)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
			return
		}
	} else {
		cart.Quantity++
		cart.Price = product.Price
		models.DB.Save(&cart)
	}
	c.JSON(http.StatusAccepted, gin.H{"message": cart})
}
