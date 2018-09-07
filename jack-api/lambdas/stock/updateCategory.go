package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

func updateCategory(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var params db.Category
	var category db.Category

	if !(&params).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	category.ID = params.ID

	if !(&category).Load() {
		return core.MakeHTTPError(400, "Category not found")
	}

	if params.Name != "" {
		category.Name = params.Name
	}

	db.DB().Save(&category)

	return core.MakeHTTPResponse(http.StatusOK, db.CategoryResponse{category})
}

func main() {
	lambda.Start(updateCategory)
}
