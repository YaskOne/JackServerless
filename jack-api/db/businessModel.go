package db

import (
	"googlemaps.github.io/maps"
	"github.com/kr/pretty"
	"context"
	"log"
)

type LatLng struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Business struct {
	//Model
	LatLng
	ID        uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Description string `json:"description"`
	Url string `json:"url"`
}

type BusinessResponse struct {
	Businesses interface{} `json:"businesses"`
}

func ValidateCreateBusiness(business Business) (bool) {
	return len(business.Name) > 1 &&
		business.Address != "" &&
		business.Type != "" &&
		!BusinessExists(business)
}

func BusinessExists(business Business) (bool) {
	request := Business{}

	res := DB().Where(&Business{Name: business.Name, Address: business.Address}).First(&request)
	return res.Error == nil
}

func CreateBusiness(business *Business) (success bool) {
	//business.Model = Model{}

	res, pos := GeocodeAddress(business.Address)

	if !res {
		return false
	}

	business.Latitude = pos.Latitude
	business.Longitude = pos.Longitude

	pretty.Println(business)
	return DB().Create(business).Error == nil
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