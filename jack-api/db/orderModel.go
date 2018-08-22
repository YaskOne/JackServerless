package db

import (
	"time"
)

/*
	DB db binding:"required"
*/

type Order struct {
	ID        uint `json:"id" gorm:"primary_key"`

	RetrieveDate time.Time `form:"retrieve_date" json:"retrieve_date"`
	Products []Product `form:"products" json:"products"`

	UserID uint `json:"user_id" gorm:"not null"`
	BusinessID uint `json:"business_id" gorm:"not null"`
}

type OrderRequest struct {
	ID        uint `json:"id" gorm:"primary_key"`

	RetrieveDate string `form:"retrieve_date" json:"retrieve_date" gorm:"not null"`
	ProductIds []uint `form:"product_ids" json:"product_ids"`
	UserID uint `form:"user_id" json:"user_id" gorm:"not null"`
	BusinessID uint `form:"business_id" json:"business_id" gorm:"not null"`
}

type OrderProduct struct {
	ID        uint `json:"id" gorm:"primary_key"`

	ProductID uint `form:"product_id" json:"product_id" gorm:"not null"`
	OrderID uint `form:"order_id" json:"order_id" gorm:"not null"`
}

/*
	Model mofiers
*/

func CreateOrder(request *Order) (success bool) {
	//request.Model = Model{}

	return DB().Create(request).Error == nil
}
