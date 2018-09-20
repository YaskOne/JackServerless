package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/utils"
	"strconv"
)


func getBusinessById(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	businesses := []db.Business {}

	nearLeftLat, nearLeftLatErr := strconv.ParseFloat(request.QueryStringParameters["near_left_latitude"], 64)
	nearLeftLng, nearLeftLngErr := strconv.ParseFloat(request.QueryStringParameters["near_left_longitude"], 64)
	farRightLat, farRightLatErr := strconv.ParseFloat(request.QueryStringParameters["far_right_latitude"], 64)
	farRightLng, farRightLngErr := strconv.ParseFloat(request.QueryStringParameters["far_right_longitude"], 64)

	if nearLeftLatErr != nil || nearLeftLngErr != nil || farRightLatErr != nil || farRightLngErr != nil {
		ids := request.QueryStringParameters["ids"]

		if len(ids) == 0 {
			db.DB().Find(&businesses)
		} else {
			db.DB().
				Where(utils.SplitArrayString(ids)).
				Find(&businesses)
		}
	} else {
		db.DB().Where(
			"latitude >= ? AND latitude <= ? AND longitude >= ? AND longitude <= ?",
			nearLeftLat,
			farRightLat,
			nearLeftLng,
			farRightLng,
		).Find(&businesses)
	}

	return core.MakeHTTPResponse(200, db.BusinessesResponse{businesses})
}

type fetchBusinessInAreaRequest struct {
	NearLeftLatitude float64 `json:"near_left_latitude"`
	NearLeftLongitude float64 `json:"near_left_longitude"`
	FarRightLatitude float64 `json:"far_right_latitude"`
	FarRightLongitude float64 `json:"far_right_longitude"`
}

func main() {
	lambda.Start(getBusinessById)
}
