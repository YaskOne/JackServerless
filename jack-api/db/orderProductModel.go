package db

import "encoding/json"

type OrderProduct struct {
	ID        uint `json:"id" gorm:"primary_key"`

	ProductID uint `form:"product_id" json:"product_id" gorm:"not null"`
	OrderID uint `form:"order_id" json:"order_id" gorm:"not null"`
}

func (model *OrderProduct) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}


func (model OrderProduct) Exists() bool {
	return false
}

func (model *OrderProduct) Load() bool {
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model OrderProduct) Valid() (bool, string) {
	return true, ""
}

func (model *OrderProduct) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *OrderProduct) Delete() bool {
	return DB().Delete(model).Error == nil
}