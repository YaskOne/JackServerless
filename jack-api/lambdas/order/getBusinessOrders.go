package main

import (
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"strconv"
	"github.com/aws/aws-lambda-go/lambda"
)

type OrdersResponse struct {
	Order db.Order `json:"order"`
	Products []db.OrderProduct `json:"products"`
}

type GetOrdersResponse struct {
	Orders interface{} `json:"orders"`
}

func GetBusinessOrders(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.ParseUint(request.QueryStringParameters["id"], 10, 64)

	if err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	orders := []db.Order{}
	ordersResponse := []OrdersResponse{}
	business := db.Business{}

	business.ID = uint(id)
	db.DB().Model(business).Related(&orders)

	i := 0
	for i < len(orders) {
		ordersResponse = append(ordersResponse, OrdersResponse{})

		orderProducts := []db.OrderProduct{}
		db.DB().Model(orders[i]).Related(&orderProducts)

		ordersResponse[i].Order = orders[i]
		ordersResponse[i].Products = orderProducts

		i += 1
	}

	return core.MakeHTTPResponse(http.StatusOK, GetOrdersResponse{ordersResponse})
}


func main() {
	lambda.Start(GetBusinessOrders)
}