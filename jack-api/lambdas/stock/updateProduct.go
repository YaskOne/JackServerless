package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

func updateProduct(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var params db.Product
	var product db.Product

	if !(&params).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	product.ID = params.ID

	if !(&product).Load() {
		return core.MakeHTTPError(400, "Category not found")
	}

	if params.Name != "" {
		product.Name = params.Name
	}
	if params.Url != "" {
		product.Url = params.Url
	}
	if params.Price != 0 {
		product.Price = params.Price
	}
	if params.CategoryID != 0 {
		product.CategoryID = params.CategoryID
	}

	db.DB().Save(&product)

	return core.MakeHTTPResponse(http.StatusOK, db.ProductResponse{product})
}

func main() {
	lambda.Start(updateProduct)
}
