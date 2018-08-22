package main

import (
	"JackServerless/jack-api/db"
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/utils"
)

func GetProductsById(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products := []db.Product {}

	ids := request.QueryStringParameters["ids"]

	if len(ids) == 0 {
		db.DB().Find(&products)
	} else {
		i := 0
		uintIds := utils.SplitArrayString(ids)
		product := db.Product {}

		for i < len(uintIds) {
			db.DB().Where(uintIds[i]).Find(&product)
			products[i] = product
		}
	}

	return core.MakeHTTPResponse(200, db.ProductResponse{products})
}

func main() {
	lambda.Start(GetProductsById)
}