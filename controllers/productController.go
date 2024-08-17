package controllers

import (
	"ecommerce/models"
	"log"
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
	// c.JSON(http.StatusAccepted, gin.H{"message": cart})
	c.Redirect(http.StatusFound, "/user/cart")
}

func CetCart(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please Login first"})
		return
	}

	var cartitems []models.Cart
	if err := models.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartitems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faied to fetch cart items"})
		return
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{"cartitems": cartitems})
}

func PlaceOrder(c *gin.Context) {
	session := sessions.Default(c)
	userId, ok := session.Get("user_id").(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please Login First"})
		return
	}

	var cartitems []models.CartItem
	if err := models.DB.Table("carts").Select("carts.id as cart_id, carts.user_id,carts.product_id,carts.quantity,products.name as product_name , products.price").Joins("left join products on products.id = carts.product_id").Where("user_id = ?", userId).Scan(&cartitems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart item"})
		return
	}

	log.Println("cart Items", cartitems)

	var total int
	for _, item := range cartitems {
		if item.ProductID == 0 || item.Price == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Product details"})
			return
		}
		total += int(item.Quantity) * item.Price
	}

	order := models.Order{
		UserID: userId,
		Total:  total,
	}
	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Order"})
		return
	}

}
