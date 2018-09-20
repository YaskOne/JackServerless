package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/utils"
	"time"
)

func ordersWatcher(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	orders := []db.Order{}

	db.DB().Find(&orders)

	println(len(orders))
	i := 0
	for i < len(orders) {
		order := orders[i]

		data := map[string]string{
			"type": utils.OrderUpdated,
			"id": string(order.ID),
		}

		currentTime := time.Now().Add(2 * time.Hour)

		println(currentTime.String())
		println(order.StartPreparationTime().String())
		println(order.EndPreparationTime().String())
		println(order.DeliveredTime().String())

		println(currentTime.After(order.StartPreparationTime()))
		println(currentTime.After(order.EndPreparationTime()))
		println(currentTime.After(order.DeliveredTime()))
		if currentTime.After(order.StartPreparationTime()) && order.OrderStatus == db.ACCEPTED {

			order.OrderStatus = db.PREPARING
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKOrderPreparing], data)

		} else if currentTime.After(order.EndPreparationTime()) && order.OrderStatus == db.PREPARING {

			order.OrderStatus = db.READY
			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKOrderReady], data)

		} else if currentTime.After(order.DeliveredTime()) && order.OrderStatus == db.READY {

			order.OrderStatus = db.DELIVERED

		}
		i += 1
		db.DB().Save(&order)
	}

	return core.MakeHTTPResponse(http.StatusOK, "ok")
}
//
//func updateStatus(params UpdateOrderRequest, order *db.Order, data map[string]string) string {
//	if params.OrderStatus > order.OrderStatus {
//		db.DB().Model(order).Update("status", params.OrderStatus)
//
//		if order.OrderStatus == db.REJECTED {
//			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKOrderRejected], data)
//
//			return refundOrder(order)
//		} else if order.OrderStatus == db.PREPARING {
//			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKOrderPreparing], data)
//		} else if order.OrderStatus == db.READY {
//			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKOrderReady], data)
//		} else if order.OrderStatus == db.CLIENT_CANCELED {
//			utils.SendPushToClient("Business", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKClientCanceledOrder], data)
//			return refundOrder(order)
//		} else if order.OrderStatus == db.BUSINESS_CANCELED {
//			utils.SendPushToClient("Takeway", order.User().FcmToken, "Jack Restaurants", utils.NotificationTexts[utils.JKBusinessOrderCanceled], data)
//
//			return refundOrder(order)
//		}
//	}
//	return ""
//}
//
//func refundOrder(order *db.Order) string {
//
//	if core.Develop {
//		return ""
//	}
//	stripe.Key = utils.StripKey
//
//	params := &stripe.RefundParams{
//		Charge: stripe.String(order.ChargeId),
//	}
//	ref, err := refund.New(params)
//
//	transaction := db.Transaction{OrderId: order.ID}
//	if err != nil {
//		transaction.Status = db.REFUND_FAIDED
//		return err.Error()
//	}
//	transaction.Status = db.REFUNDED
//	order.RefundId = ref.ID
//
//	db.DB().Save(order)
//	(&transaction).Create()
//
//	return ""
//}

func main() {
	lambda.Start(ordersWatcher)
}
