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

	////
	db.DB().DropTableIfExists(&db.Business{})
	db.DB().DropTableIfExists(&db.Category{})
	db.DB().DropTableIfExists(&db.LatLng{})
	db.DB().DropTableIfExists(&db.Product{})
	db.DB().DropTableIfExists(&db.Order{})
	db.DB().DropTableIfExists(&db.OrderProduct{})

	db.DB().AutoMigrate(&db.Business{})
	db.DB().AutoMigrate(&db.Category{})
	db.DB().AutoMigrate(&db.LatLng{})
	db.DB().AutoMigrate(&db.Product{})
	db.DB().AutoMigrate(&db.Order{})
	db.DB().AutoMigrate(&db.OrderProduct{})

	return core.MakeHTTPResponse(200, "yeah")
}

func main() {
	lambda.Start(ResetDB)
}
