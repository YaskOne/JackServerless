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
	"JackServerless/jack-api/utils"
	"github.com/stripe/stripe-go"
	"fmt"
)

/*
	 Create new order route
*/

func createOrder(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var order db.Order
	var params db.OrderRequest
	//utils.SendSMS("+33687918380", "Ca roule ? Yoyoyoyoyoyoyo woulouwlou jiwjfbwe niwfe nwf wfji wejfwb hnweijd iwjbfwfjkwenf")

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	//println(params.RetrieveDate)
	retrieveDate, _ := time.Parse(time.RFC3339, params.RetrieveDate)
	//println(retrieveDate.String())

	order.BusinessID = params.BusinessID
	order.UserID = params.UserID
	order.RetrieveDate = retrieveDate
	order.Canceled = false
	order.OrderStatus = db.PENDING
	order.State = 0

	pretty.Println(retrieveDate)
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

	data := map[string]string{
		"type": utils.NewOrder,
		"id": string(order.ID),
	}

	pretty.Println(order.Business().ID)
	pretty.Println(order.Business().FcmToken)

	if business := order.Business(); business.ID != 0 {
		utils.SendPushToClient("Business", business.FcmToken, "Jack Restaurants", fmt.Sprintf(string(utils.JKOrderPreparing), order.RetrieveDate.Format("15:04")), data)
	}

	if !core.Develop {
		charge, err := utils.ChargeCustomer(int(order.Price * 100),
			stripe.CurrencyEUR,
			order.User().StripeCustomerId,
			order.Business().Name + "|" + fmt.Sprint(order.UserID) + "|" + order.RetrieveDate.String())

		transaction := db.Transaction{OrderId: order.ID}

		if err != nil {
			transaction.Status = db.PAY_FAIDED
			return core.MakeHTTPError(400, err)
		}

		transaction.Status = db.PAYED
		order.ChargeId = charge.ID

		(&transaction).Create()
		db.DB().Save(&order)
	}

	return core.MakeHTTPResponse(http.StatusOK, db.IdModel{order.ID})
}

func main() {
	lambda.Start(createOrder)
}