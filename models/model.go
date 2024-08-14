package models

type User struct {
	ID       uint   `form:"id" gorm:"primaryKey"`
	Username string `form:"username" gorm:"size:100;not null"`
	Email    string `form:"email" gorm:"size:100;null,unique"`
	Password string `form:"password" gorm:"size:100;not null"`
}

type Products struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Description string
	Price       int `gorm:"not null"`
}

type Cart struct {
	ID        uint     `gorm:"primary_key"`
	UserID    uint     `gorm:"not null"`
	ProductID uint     `gorm:"not null"`
	Quantity  int      `gorm:"default:1"`
	Product   Products `gorm:"foreignKey:ProductID"`
	Price     int
}
