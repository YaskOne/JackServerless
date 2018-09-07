package main

import (
	"JackServerless/jack-api/db"
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/utils"
	"strconv"
)

func GetProducts(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	ids := request.QueryStringParameters["ids"]
	business_id := request.QueryStringParameters["business_id"]
	category_id := request.QueryStringParameters["category_id"]

	if len(ids) > 0 {
		return getProductsById(ids, ctx, request)
	} else if len(business_id) > 0 {
		return getProductsByBusiness(business_id, ctx, request)
	} else if len(category_id) > 0 {
		return getProductsByCategory(category_id, ctx, request)
	}

	return core.MakeHTTPError(401, "Invalid input parameters, try 'ids', 'business_id' or 'category_id'")
}

func getProductsById(params string, ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	uintIds := utils.SplitArrayString(params)

	products := db.GetProductsById(uintIds)

	return core.MakeHTTPResponse(200, db.ProductsResponse{products})
}
func getProductsByBusiness(params string, ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.ParseUint(params, 10, 64)

	if err != nil {
		return core.MakeHTTPError(401, "Invalid business_id parameters")
	}

	products := []db.Product {}
	business := db.Business{}
	business.ID = uint(id)

	db.DB().Model(business).Related(&products)

	return core.MakeHTTPResponse(200, db.ProductsResponse{products})
}
func getProductsByCategory(params string, ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.ParseUint(params, 10, 64)

	if err != nil {
		return core.MakeHTTPError(401, "Invalid category_id parameters")
	}

	products := []db.Product {}
	category := db.Category{ID: uint(id)}

	db.DB().Model(category).Related(&products)

	return core.MakeHTTPResponse(200, db.ProductsResponse{products})
}

func main() {
	lambda.Start(GetProducts)
}