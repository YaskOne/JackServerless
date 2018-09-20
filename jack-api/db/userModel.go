package db

import "encoding/json"

/*
	DB models
*/

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `json:"name" gorm:"not null" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Email string `json:"email" gorm:"not null;unique" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`

	DeviceToken string `json:"devie_token"`
	Token string `json:"token"`
	FcmToken string `json:"fcm_token"`

	StripeKey string `json:"stripe_key"`
	StripeCustomerId string `json:"stripe_customer_id"`
}

type UserObject struct {
	ID        uint `json:"id" gorm:"primary_key"`

	Name string `json:"name" gorm:"not null" binding:"required"`
	Email string `json:"email" gorm:"not null;unique" binding:"required"`

	StripeCustomerId string `json:"stripe_customer_id"`
}

type UserResponse struct {
	User interface{} `json:"user"`
}

type UsersResponse struct {
	Users interface{} `json:"users"`
}

func GetUserObject(user User) UserObject {
	userObject := UserObject{}

	userObject.ID = user.ID
	userObject.Name = user.Name
	userObject.Email = user.Email
	userObject.StripeCustomerId = user.StripeCustomerId

	return userObject
}

func (model *User) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}

func (model User) Exists() bool {
	if model.ID == 0 {
		return DB().Where(&User{Email: model.Email}).First(&model).Error == nil
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model *User) Load() bool {
	if model.ID == 0 {
		return false
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model User) Valid() (bool, string) {
	if len(model.Name) < 2 {
		return false, "name too short, must be at leat 2 characters"
	} else if model.Email == "" {
		return false, "address must be specified"
	} else if len(model.Password) < 6 {
		return false, "password too short, must be at least 6 characters"
	} else if model.Exists() {
		return false, "user already exists"
	}
	return true, ""
}

func (model *User) Create() bool {
	return DB().Create(model).Error == nil
}

func (model *User) Delete() bool {
	orders := []Order{}

	DB().Model(model).Related(&orders)

	i := 0
	for i < len(orders) {
		orders[i].Delete()
		i += 1
	}
	return DB().Delete(model).Error == nil
}