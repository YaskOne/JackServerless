package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

/*
	 Create new category
*/

func createCategory(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	category := db.Category{}

	if !(&category).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if valid, err := category.Valid(); !valid {
		return core.MakeHTTPError(400, err)
	}

	if !(&category).Create() {
		return core.MakeHTTPError(400, "Error creating catgeory")
	}

	return core.MakeHTTPResponse(http.StatusOK, db.IdModel{category.ID})
}

func main() {
	lambda.Start(createCategory)
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
