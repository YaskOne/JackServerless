package db

/*
	DB models
*/

type Category struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `binding:"required" json:"name" gorm:"not null"`
	BusinessID uint `json:"business_id" gorm:"not null" binding:"required"`
}

type Product struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `form:"name" json:"name" gorm:"not null" binding:"required"`
	Price float32 `form:"price" json:"price" gorm:"not null" binding:"required"`
	Url string `form:"url" json:"url" gorm:"not null"`

	CategoryID uint `form:"category_id" json:"category_id" gorm:"not null" binding:"required"`
	BusinessID uint `form:"business_id" json:"business_id" gorm:"not null" binding:"required"`
}

type ProductResponse struct {
	Products interface{} `json:"products"`
}

/*
	Model mofiers
*/

func CreateProduct(request *Product) (success bool) {
	// request.BusinessID = request.BusinessID

	return DB().Create(request).Error == nil
}

func CreateCategory(request *Category) (success bool) {
	// request.BusinessID = request.BusinessID

	return DB().Create(request).Error == nil
}

/*
	Accessors
*/

func GetCategoryProducts(category Category) ([]Product) {
	products := []Product{}

	DB().Model(category).Find(&products)

	return products
}

func GetBusinessCategories(business Business) ([]Category) {
	categories := []Category{}

	DB().Model(business).Find(&categories)

	return categories
}
