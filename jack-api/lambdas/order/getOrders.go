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

func getOrders(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	businessId, err1 := strconv.ParseUint(request.QueryStringParameters["business_id"], 10, 64)
	userId, err2 := strconv.ParseUint(request.QueryStringParameters["user_id"], 10, 64)

	orders := []db.Order{}
	ordersResponse := []db.OrderResponse{}

	if err1 == nil {
		business := db.Business{}
		business.ID = uint(businessId)

		db.DB().Model(business).Related(&orders)
	} else if err2 == nil {
		user := db.User{}
		user.ID = uint(userId)

		db.DB().Model(user).Related(&orders)
	}

	i := 0
	for i < len(orders) {
		ordersResponse = append(ordersResponse, db.OrderResponse{})

		orderProducts := []db.OrderProduct{}
		db.DB().Model(orders[i]).Related(&orderProducts)

		ordersResponse[i].Order = orders[i]
		ordersResponse[i].Products = orderProducts

		i += 1
	}

	return core.MakeHTTPResponse(http.StatusOK, db.GetOrdersResponse{ordersResponse})
}


func main() {
	lambda.Start(getOrders)
}