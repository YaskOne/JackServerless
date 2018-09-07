package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/utils"
)

func GetCategory(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	ids := request.QueryStringParameters["ids"]
	business_id := request.QueryStringParameters["business_id"]

	if len(ids) > 0 {
		return getCategoriesById(ids, ctx, request)
	} else if len(business_id) > 0 {
		return getCategoriesByBusiness(business_id, ctx, request)
	}

	return core.MakeHTTPError(401, "Invalid parameters")
}

func getCategoriesByBusiness(params string, ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.ParseUint(params, 10, 64)

	if err != nil {
		return core.MakeHTTPError(401, "Invalid parameters")
	}

	categories := []db.Category {}
	business := db.Business{}
	business.ID = uint(id)

	db.DB().Model(business).Related(&categories)

	return core.MakeHTTPResponse(200, db.CategoriesResponse{categories})
}
func getCategoriesById(params string, ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	categories := []db.Category {}

	db.DB().Where(utils.SplitArrayString(params)).Find(&categories)

	return core.MakeHTTPResponse(200, db.CategoriesResponse{categories})
}

func main() {
	lambda.Start(GetCategory)
}