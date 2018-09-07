package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
)

// Handler is the Lambda function handler
func ResetDB(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	if request.QueryStringParameters["business"] == "true" {
		db.DB().DropTableIfExists(&db.Business{})
	}
	if request.QueryStringParameters["category"] == "true" {
		db.DB().DropTableIfExists(&db.Category{})
	}
	if request.QueryStringParameters["product"] == "true" {
		db.DB().DropTableIfExists(&db.Product{})
	}
	if request.QueryStringParameters["order"] == "true" {
		db.DB().DropTableIfExists(&db.Order{})
		db.DB().DropTableIfExists(&db.OrderProduct{})
	}
	if request.QueryStringParameters["user"] == "true" {
		db.DB().DropTableIfExists(&db.User{})
	}

	db.DB().AutoMigrate(&db.Business{})
	db.DB().AutoMigrate(&db.Category{})
	db.DB().AutoMigrate(&db.LatLng{})
	db.DB().AutoMigrate(&db.Product{})
	db.DB().AutoMigrate(&db.Order{})
	db.DB().AutoMigrate(&db.OrderProduct{})
	db.DB().AutoMigrate(&db.User{})

	return core.MakeHTTPResponse(200, "yeah")
}

func main() {
	lambda.Start(ResetDB)
}
