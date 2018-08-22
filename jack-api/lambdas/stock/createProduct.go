package main

import (
	"net/http"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
	--- PRODUCTS
*/


/*
	 Create new stock
*/

func CreateProduct(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var params db.Product
	var place db.Business

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	if db.DB().Where(params.BusinessID).Find(&place).RecordNotFound() {
		return core.MakeHTTPError(http.StatusNotAcceptable, "Error: business not found")
	}
	db.CreateProduct(&params)

	return core.MakeHTTPResponse(http.StatusOK, db.IDResponse{params.ID})
}

func main() {
	lambda.Start(CreateProduct)
}


//type BusinessProductResponse struct {
//	products []db.Product
//}
///*
//	 Fetch all products from business
//*/
//
//func FetchBusinessProducts(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//	id, err := strconv.ParseUint(request.QueryStringParameters["id"], 10,64)
//
//	if err != nil {
//		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
//	}
//
//	var business db.Business
//
//	products := []db.Product {}
//
//	place := db.Model{}
//
//	place.ID = uint(id)
//
//	business.ID = uint(id)
//
//	db.DB().Model(business).Related(&products)
//
//	return core.MakeHTTPResponse(http.StatusOK, BusinessProductResponse{products})
//}
//
//func main() {
//	lambda.Start(FetchBusinessProducts)
//}

/*
	 Fetch products by ID
*/

//func FetchProductsByID(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//	products := []db.Product {}
//
//	ids, err := strconv.ParseUint(request.QueryStringParameters["id"], 10,64)
//
//	if err != nil {
//		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
//	}
//
//	db.DB().Where(ids).Find(&products)
//
//	return core.MakeHTTPResponse(http.StatusOK, BusinessProductResponse{products})
//}
//
//func main() {
//	lambda.Start(FetchProductsByID)
//}

/*
	 Fetch products by ID
*/

//func FetchProductsCategoriesByID(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//	categories := []db.Category {}
//
//	ids := request.QueryStringParameters["ids"]
//
//	db.DB().Where(utils.SplitArrayString(ids)).Find(&categories)
//
//	return core.MakeHTTPResponse(http.StatusOK, BusinessCategoriesResponse{categories})
//}
//
//func main() {
//	lambda.Start(FetchProductsCategoriesByID)
//}