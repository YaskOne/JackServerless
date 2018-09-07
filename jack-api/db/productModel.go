package db

import "encoding/json"

/*
	DB models
*/

type Product struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `form:"name" json:"name" gorm:"not null" binding:"required"`
	Price float64 `form:"price" json:"price" gorm:"not null" binding:"required"`
	Url string `form:"url" json:"url" gorm:"not null"`

	CategoryID uint `form:"category_id" json:"category_id" gorm:"not null" binding:"required"`
	BusinessID uint `form:"business_id" json:"business_id" gorm:"not null" binding:"required"`
}

type ProductsResponse struct {
	Products interface{} `json:"products"`
}
type ProductResponse struct {
	Product interface{} `json:"product"`
}

/*
	Accessors
*/

func GetProductsById(ids []uint) ([]Product) {
	products := []Product{}

	var i = 0
	for i < len(ids) {
		product := Product{}
		DB().Where(ids[i]).First(&product)
		products = append(products, product)
		i += 1
	}

	return products
}

func (model *Product) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}

func (model Product) Exists() bool {
	if model.ID == 0 {
		return DB().Where(&Product{Name: model.Name, BusinessID: model.BusinessID}).First(&model).Error == nil
	}
	return DB().Where(model.ID).Find(&model).Error == nil
}

func (model *Product) Load() bool {
	if model.ID == 0 {
		return false
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model Product) Valid() (bool, string) {
	if len(model.Name) < 1 {
		return false, "name too short, must be at leat 1 characters"
	} else if model.Exists() {
		return false, "product already exists"
	}
	return true, ""
}

func (model *Product) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *Product) Delete() bool {
	return DB().Delete(model).Error == nil
}