package db

import (
	"time"
	"encoding/json"
)

/*
	DB db binding:"required"
*/

type OrderStatus uint

const (
	PENDING    OrderStatus = 0
	ACCEPTED   OrderStatus = 1
	PREPARING    OrderStatus = 2
	READY   OrderStatus = 3
	DELIVERED   OrderStatus = 4
	REJECTED    OrderStatus = 5
	CLIENT_CANCELED   OrderStatus = 6
	BUSINESS_CANCELED   OrderStatus = 7
)

type Order struct {
	ID        uint `json:"id" gorm:"primary_key"`

	RetrieveDate time.Time `json:"retrieve_date"`
	Products []Product `json:"products"`
	Price float64 `json:"price"`

	Canceled bool `json:"canceled"`
	ChargeId string `json:"charge_id"`
	RefundId string `json:"refund_id"`

	OrderStatus OrderStatus `json:"status"`
	State int `json:"state"`

	UserID uint `json:"user_id" gorm:"not null"`
	BusinessID uint `json:"business_id" gorm:"not null"`

	Position uint `json:"position" gorm:"not null"`
	CellPhone uint `json:"cell_phone" gorm:"not null"`
}

type OrderRequest struct {
	ID        uint `json:"id" gorm:"primary_key"`

	RetrieveDate string `form:"retrieve_date" json:"retrieve_date" gorm:"not null"`
	ProductIds []uint `form:"product_ids" json:"product_ids"`
	UserID uint `form:"user_id" json:"user_id" gorm:"not null"`
	BusinessID uint `form:"business_id" json:"business_id" gorm:"not null"`
}

type OrderResponse struct {
	Order Order `json:"order"`
	Products []OrderProduct `json:"products"`
}

type GetOrdersResponse struct {
	Orders interface{} `json:"orders"`
}

func (model Order) StartPreparationTime() time.Time {
	business := model.Business()
	return model.RetrieveDate.Add(-business.DefaultPreparationDuration / 1000)
}

func (model Order) EndPreparationTime() time.Time {
	return model.RetrieveDate
}

func (model Order) DeliveredTime() time.Time {
	return model.RetrieveDate.Add(time.Minute * 5)
}

func (model Order) User() User {
	user := User{}

	DB().Where(model.UserID).First(&user)
	return user
}

func (model Order) Business() Business {
	business := Business{}

	DB().Where(model.BusinessID).First(&business)
	return business
}

/*
	Model mofiers
*/

func (model *Order) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}


func (model Order) Exists() bool {
	return false
}

func (model *Order) Load() bool {
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model Order) Valid() (bool, string) {
	if len(model.Products) > 0 {
		return false, "must have at least one product"
	} else if !(Business{ID: model.BusinessID}).Exists() {
		return false, "business doesn't exist"
	} else if !(User{ID: model.UserID}).Exists() {
		return false, "user doesn't exist"
	}
	return true, ""
}

func (model *Order) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *Order) Delete() bool {
	orderProducts := []OrderProduct{}

	DB().Model(model).Related(&orderProducts)

	i := 0
	for i < len(orderProducts) {
		orderProducts[i].Delete()
		i += 1
	}

	return DB().Delete(model).Error == nil
}