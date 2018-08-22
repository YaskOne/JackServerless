package main

import (
	"JackServerless/jack-api/db"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/kr/pretty"
)

type Section struct {
	db.Category
	Products []db.Product `json:"products"`
}

type BusinessProductsResponse struct {
	Stocks interface{} `json:"stocks"`
}

func GetBusinessStocks(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	println("YAYAYAYAYYAYAYAYAYAYA_____")
	id, err := strconv.ParseUint(request.QueryStringParameters["id"], 10,64)

	if err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	sections := []Section{}

	//business := db.GetBusiness(request.BusinessID)
	//if db.BusinessExists(business) == nil {
	//
	//}
	categories := []db.Category{}
	//categories := db.GetBusinessCategories(db.Business{db.Model{ID: 1}}).Related(&categories)
	db.DB().Where("business_id = ?", id).Find(&categories)

	i := 0
	pretty.Println(categories)
	for i < len(categories) {
		products := []db.Product {}
		db.DB().Where("category_id = ?", categories[i].ID).Find(&products)

		section := Section{}
		section.Category = categories[i]
		section.Products = products

		sections = append(sections, section)
		i++
	}
	pretty.Println(sections)

	return core.MakeHTTPResponse(http.StatusOK, BusinessProductsResponse{sections})
}

func main() {
	lambda.Start(GetBusinessStocks)
}