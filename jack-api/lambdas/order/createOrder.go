package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"net/http"
	"time"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kr/pretty"
	"encoding/json"
)

/*
	 Create new order route
*/

func createOrder(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var order db.Order
	var params db.OrderRequest

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	retrieveDate, _ := time.Parse(time.RFC3339, params.RetrieveDate)

	order.BusinessID = params.BusinessID
	order.UserID = params.UserID
	order.RetrieveDate = retrieveDate
	order.Canceled = false
	order.Status = db.OrderStatus("PENDING")
	order.State = db.OrderState("WAITING")

	pretty.Println(order)

	// creates place
	if valid, err := order.Valid(); !valid {
		return core.MakeHTTPError(400, err)
	}

	if !(&order).Create() {
		return core.MakeHTTPError(http.StatusInternalServerError, "Error: creating order")
	}

	products := db.GetProductsById(params.ProductIds)

	price := 0.0
	i := 0
	for i < len(products) {
		orderProduct := db.OrderProduct{}
		orderProduct.OrderID = order.ID
		orderProduct.ProductID = params.ProductIds[i]
		db.DB().Create(&orderProduct)
		price += products[i].Price
		i += 1
	}

	order.Price = price
	db.DB().Save(&order)

	return core.MakeHTTPResponse(http.StatusOK, db.IdModel{order.ID})
}

func main() {
	lambda.Start(createOrder)
}