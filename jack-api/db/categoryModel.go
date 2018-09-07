package db

import "encoding/json"

type Category struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `binding:"required" json:"name" gorm:"not null"`
	BusinessID uint `json:"business_id" gorm:"not null" binding:"required"`
}

type CategoriesResponse struct {
	Categories interface{} `json:"categories"`
}
type CategoryResponse struct {
	Category interface{} `json:"category"`
}

func (model *Category) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}


func (model Category) Exists() bool {
	if model.ID == 0 {
		return DB().Where(&Category{Name: model.Name}).First(&model).Error == nil
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model *Category) Load() bool {
	if model.ID == 0 {
		return false
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model Category) Valid() (bool, string) {
	if len(model.Name) < 1 {
		return false, "name too short, must be at leat 1 characters"
	} else if model.Exists() {
		return false, "category already exists"
	}
	return true, ""
}

func (model *Category) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *Category) Delete() bool {
	products := []Product{}

	DB().Model(model).Related(&products)

	i := 0
	for i < len(products) {
		(&products[i]).Delete()
		i += 1
	}

	return DB().Delete(model).Error == nil
}