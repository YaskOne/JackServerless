package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"encoding/json"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

/*
	 Create new category
*/

func CreateCategory(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var params db.Category

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	var business db.Business
	if db.DB().Where(params.BusinessID).Find(&business).RecordNotFound() {
		return core.MakeHTTPError(http.StatusNotAcceptable, "Business not found")
	}

	db.CreateCategory(&params)

	return core.MakeHTTPResponse(http.StatusOK, db.IDResponse{params.ID})
}

func main() {
	lambda.Start(CreateCategory)
}

/*
	 Fetch all products categories from business
*/

//type BusinessCategoriesResponse struct {
//	categories []db.Category
//}
//
//func FetchBusinessProductsCategories(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//	var place db.Business
//
//	categories := []db.Category {}
//
//	id, err := strconv.ParseUint(request.QueryStringParameters["id"], 10,64)
//
//	if err != nil {
//		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
//	}
//	place.ID = uint(id)
//
//	db.DB().Model(place).Related(&categories)
//
//	return core.MakeHTTPResponse(http.StatusOK, BusinessCategoriesResponse{categories})
//}
//
//func main() {
//	lambda.Start(FetchBusinessProductsCategories)
//}
