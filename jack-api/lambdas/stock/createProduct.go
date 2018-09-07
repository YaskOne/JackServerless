package main

import (
	"net/http"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
	 Create new stock
*/

func createProduct(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var product db.Product

	if !(&product).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if valid, err := product.Valid(); !valid {
		return core.MakeHTTPError(400, err)
	}

	if !(&product).Create() {
		return core.MakeHTTPError(http.StatusInternalServerError, "Error creating product")
	}

	return core.MakeHTTPResponse(http.StatusOK, db.IdModel{product.ID})
}

func main() {
	lambda.Start(createProduct)
}