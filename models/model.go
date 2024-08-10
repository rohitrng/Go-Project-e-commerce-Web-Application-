package models

type User struct {
	ID       uint   `form:"id" gorm:"primaryKey"`
	Username string `form:"username" gorm:"size:100;not null"`
	Email    string `form:"email" gorm:"size:100;null,unique"`
	Password string `form:"password" gorm:"size:100;not null"`
}
