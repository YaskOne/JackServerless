package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"encoding/json"
	"JackServerless/jack-api/utils"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
	"time"
	"fmt"
)

type UpdateOrderRequest struct {
	OrderId uint `json:"order_id"`
	OrderStatus db.OrderStatus `json:"status"`
	State int `json:"state"`
	Canceled bool `json:"canceled"`
	RetrieveDate string `form:"retrieve_date" json:"retrieve_date" gorm:"not null"`
}

func updateOrder(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := UpdateOrderRequest{}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(400, err.Error())
	}

	order := db.Order{}
	db.DB().First(&order, params.OrderId)

	data := map[string]string{
		"type": utils.OrderUpdated,
		"id": string(order.ID),
	}

	retrieveDate, err := time.Parse(time.RFC3339, params.RetrieveDate)

	if err == nil {
		order.RetrieveDate = retrieveDate

		utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", fmt.Sprintf(string(utils.JKOrderAccepted), order.RetrieveDate.Format("15:04")), data)
		db.DB().Save(&order)
	}

	if err := updateStatus(params, &order, data); err != "" {
		core.MakeHTTPError(400, err)
	}

	products := []db.OrderProduct{}
	db.DB().Model(order).Related(&products)

	return core.MakeHTTPResponse(http.StatusOK, db.OrderResponse{order, products})
}

func updateStatus(params UpdateOrderRequest, order *db.Order, data map[string]string) string {
	if params.OrderStatus > order.OrderStatus {
		db.DB().Model(order).Update("status", params.OrderStatus)
		order.OrderStatus = params.OrderStatus
		db.DB().Save(order)
		if order.OrderStatus == db.REJECTED {
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", string(utils.JKOrderRejected), data)

			return refundOrder(order)
		} else if order.OrderStatus == db.ACCEPTED {
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", fmt.Sprintf(string(utils.JKOrderAccepted), order.RetrieveDate.Format("15:04")), data)
		} else if order.OrderStatus == db.PREPARING {
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", fmt.Sprintf(string(utils.JKOrderPreparing), order.RetrieveDate.Format("15:04")), data)
		} else if order.OrderStatus == db.READY {
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", string(utils.JKOrderReady), data)
		} else if order.OrderStatus == db.CLIENT_CANCELED {
			utils.SendPushToClient("Business", order.Business().FcmToken, "Jack Restaurants", string(utils.JKClientCanceledOrder), data)
			return refundOrder(order)
		} else if order.OrderStatus == db.BUSINESS_CANCELED {
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", string(utils.JKBusinessOrderCanceled), data)

			return refundOrder(order)
		}
	}
	return ""
}

func refundOrder(order *db.Order) string {

	if core.Develop {
		return ""
	}
	stripe.Key = utils.StripKey

	params := &stripe.RefundParams{
		Charge: stripe.String(order.ChargeId),
	}
	ref, err := refund.New(params)

	transaction := db.Transaction{OrderId: order.ID}
	if err != nil {
		transaction.Status = db.REFUND_FAIDED
		return err.Error()
	}
	transaction.Status = db.REFUNDED
	order.RefundId = ref.ID

	db.DB().Save(order)
	(&transaction).Create()

	return ""
}

func main() {
	lambda.Start(updateOrder)
}
