package db

import (
	"encoding/json"
)

/*
	DB db binding:"required"
*/

type TransactionStatus uint

const (
	PAYED    TransactionStatus = 0
	REFUNDED   TransactionStatus = 1
	PAY_FAIDED   TransactionStatus = 2
	REFUND_FAIDED   TransactionStatus = 3
)

type Transaction struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Status TransactionStatus `json:"status"`

	OrderId uint `json:"order_id"`
}

func (model Transaction) Order() Order {
	order := Order{}

	DB().Where(model.OrderId).First(&order)
	return order
}

func (model *Transaction) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}

func (model Transaction) Exists() bool {
	return false
}

func (model *Transaction) Load() bool {
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model Transaction) Valid() (bool, string) {
	return true, ""
}

func (model *Transaction) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *Transaction) Delete() bool {
	return DB().Delete(model).Error == nil
}

