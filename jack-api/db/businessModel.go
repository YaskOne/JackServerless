package db

import (
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
	"context"
	"log"
	"encoding/json"
	"time"
)

type LatLng struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DisponibilityStatus int

const (
	AVAILABLE    DisponibilityStatus = 0
	UNAVAILABLE   DisponibilityStatus = 1
	TEMPORARILY_UNAVAILABLE    DisponibilityStatus = 2
)

type Business struct {
	//Model

	ID        uint `json:"id" gorm:"primary_key"`
	LatLng
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique" binding:"required"`
	Password string `json:"password" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Description string `json:"description"`
	Url string `json:"url"`

	DisponibilityStatus DisponibilityStatus `json:"disponibility_status"`
	DefaultPreparationDuration time.Duration `json:"default_preparation_duration"`

	Token string `json:"token"`
	FcmToken string `json:"fcm_token"`
}

type BusinessesResponse struct {
	Businesses interface{} `json:"businesses"`
}

type BusinessResponse struct {
	Business interface{} `json:"business"`
}

func (model *Business) Parse(data string) bool {
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return false
	}
	return true
}

func (model Business) Exists() bool {
	if model.ID == 0 {
		return DB().Where(&Business{Name: model.Email, Address: model.Address}).First(&model).Error == nil
	}
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model *Business) Load() bool {
	return DB().Where(model.ID).First(&model).Error == nil
}

func (model Business) Valid() (bool, string) {
	if len(model.Name) < 1 {
		return false, "name too short, must be at leat 1 characters"
	} else if len(model.Password) < 6 {
		return false, "password too short, must be at least 6 characters"
	} else if model.Address == "" {
		return false, "address must be specified"
	} else if model.Email == "" {
		return false, "email must be specified"
	} else if model.Type == "" {
		return false, "type must be specified"
	} else if model.Exists() {
		return false, "business already exists"
	}
	return true, ""
}

func (model *Business) Create() bool {
	res, pos := GeocodeAddress(model.Address)

	if !res {
		return false
	}

	model.Latitude = pos.Latitude
	model.Longitude = pos.Longitude

	model.DisponibilityStatus = AVAILABLE
	model.DefaultPreparationDuration = time.Duration(10 * time.Minute)

	return DB().Create(model).Error == nil
}

func (model *Business) Delete() bool {
	orders := []Order{}
	categories := []Category{}

	DB().Model(model).Related(&orders)
	DB().Model(model).Related(&categories)

	i := 0
	for i < len(orders) {
		(&orders[i]).Delete()
		i += 1
	}

	i = 0
	for i < len(categories) {
		(&categories[i]).Delete()
		i += 1
	}

	return DB().Delete(model).Error == nil
}

func GeocodeAddress(address string) (success bool, pos LatLng) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCPC2asFTIIc7ysBdaDe78dg6aWsjUDLxY"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// extrenalize
	r := &maps.GeocodingRequest{
		Address: address,
	}
	latLng := LatLng{}
	res, err := c.Geocode(context.Background(), r)

	if err == nil {
		pretty.Println(res[0].Geometry.Location)

		latLng.Latitude = res[0].Geometry.Location.Lat
		latLng.Longitude = res[0].Geometry.Location.Lng
	}
	return err == nil, latLng
}