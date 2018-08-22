package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"encoding/json"
	"JackServerless/jack-api/core"
	"net/http"
	"time"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kr/pretty"
)

/*
	 Create new order route
*/

func CreateOrder(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	println("1111")
	var order db.Order
	var params db.OrderRequest

	println("1111")
	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	println("2222")
	pretty.Println(params)
	retrieveDate, _ := time.Parse(time.RFC3339, params.RetrieveDate)

	//order.Products = products
	order.BusinessID = params.BusinessID
	order.UserID = params.UserID
	order.RetrieveDate = retrieveDate
	println("3333")
	pretty.Println(order)

	// creates place
	if !db.CreateOrder(&order) {
		return core.MakeHTTPError(http.StatusInternalServerError, "Error: creating order")
	}
	println("4444")

	pretty.Println(params)

	i := 0
	for i < len(params.ProductIds) {
		orderProduct := db.OrderProduct{}
		//orderProduct.Model = db.Model{}
		orderProduct.OrderID = order.ID
		orderProduct.ProductID = params.ProductIds[i]
		db.DB().Create(&orderProduct)
		println("---------------")
		i += 1
	}

	println(order.BusinessID)

	//"orders_id": ordersResponse,
	return core.MakeHTTPResponse(http.StatusOK, order)
}

func main() {
	lambda.Start(CreateOrder)
}