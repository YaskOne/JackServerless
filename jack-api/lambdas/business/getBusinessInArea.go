package main

import (
	"github.com/aws/aws-lambda-go/events"
	"context"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

/*
	 Fetch businesses in area
*/

type FetchBusinessInAreaRequest struct {
	NearLeftLatitude float64 `json:"near_left_latitude"`
	NearLeftLongitude float64 `json:"near_left_longitude"`
	FarRightLatitude float64 `json:"far_right_latitude"`
	FarRightLongitude float64 `json:"far_right_longitude"`
}

func FetchBusinessInArea(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	businesses := []db.Business {}

	nearLeftLat, nearLeftLatErr := strconv.ParseFloat(request.QueryStringParameters["near_left_latitude"], 64)
	nearLeftLng, nearLeftLngErr := strconv.ParseFloat(request.QueryStringParameters["near_left_longitude"], 64)
	farRightLat, farRightLatErr := strconv.ParseFloat(request.QueryStringParameters["far_right_latitude"], 64)
	farRightLng, farRightLngErr := strconv.ParseFloat(request.QueryStringParameters["far_right_longitude"], 64)

	if nearLeftLatErr != nil || nearLeftLngErr != nil || farRightLatErr != nil || farRightLngErr != nil {
		return core.MakeHTTPError(400, "Error in request parameters")
	}

	println("YAYAYAYA")

	db.DB().Where(
			"latitude >= ? AND latitude <= ? AND longitude >= ? AND longitude <= ?",
			nearLeftLat,
			farRightLat,
			nearLeftLng,
			farRightLng,
		).Find(&businesses)

	return core.MakeHTTPResponse(200, db.BusinessResponse{businesses})
	//return core.MakeHTTPError(200, businesses)

	//var i interface{}
	//i = businesses
	//
	//body, _ := json.Marshal(core.HTTPErrorBody{i})
	//
	//return &events.APIGatewayProxyResponse{
	//	StatusCode: 200,
	//	Body:       string(body),
	//	Headers:    map[string]string{core.ContentType: core.JSON},
	//}, nil
}

func main() {
	lambda.Start(FetchBusinessInArea)
}
