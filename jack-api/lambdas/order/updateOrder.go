package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"encoding/json"
)

type UpdateOrderRequest struct {
	OrderId uint `json:"order_id"`
	Status string `json:"status"`
	State string `json:"state"`
	Canceled bool `json:"canceled"`
}

func updateOrder(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := UpdateOrderRequest{}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(400, err.Error())
	}

	order := db.Order{}
	db.DB().First(&order, params.OrderId)

	if len(params.Status) != 0 {
		db.DB().Model(&order).Update("status", db.OrderStatus(params.Status))
	}
	if len(params.State) != 0 {
		db.DB().Model(&order).Update("state", db.OrderStatus(params.State))
	}
	if params.Canceled {
		db.DB().Model(&order).Update("canceled", true)
	}

	products := []db.OrderProduct{}
	db.DB().Model(order).Related(&products)

	return core.MakeHTTPResponse(http.StatusOK, db.OrderResponse{order, products})
}


func main() {
	lambda.Start(updateOrder)
}
