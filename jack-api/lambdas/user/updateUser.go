package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/kr/pretty"
)

func updateUser(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := db.User{}
	user := db.User{}

	if !(&params).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error parsing user")
	}

	user.ID = params.ID

	if !(&user).Load() {
		return core.MakeHTTPError(400, "User not found")
	}

	updateUserInfos(params, &user)
	updateUserFcmInfos(params, &user)
	updateUserStripeInfos(params, &user, request)
	db.DB().Save(&user)

	return core.MakeHTTPResponse(http.StatusOK, db.UserResponse{user})
}

func updateUserInfos(params db.User, user *db.User) bool {
	if params.Name != "" {
		user.Name = params.Name
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Password != "" {
		user.Password = params.Password
	}
	return true
}

func updateUserFcmInfos(params db.User, user *db.User) bool {
	if params.FcmToken != "" {
		user.FcmToken = params.FcmToken
	}
	return true
}

func updateUserStripeInfos(params db.User, user *db.User, request *events.APIGatewayProxyRequest) bool {
	stripe.Key = "sk_live_0FeoUv8ixC4nzyUndaRqyEC8"

	pretty.Println("11111")
	if user.Email == "" || (params.StripeKey == "" && params.StripeCustomerId == "") {
		return false
	}
	pretty.Println("2222")

	// Create a Customer:
	customerParams := &stripe.CustomerParams{
		Email: stripe.String(user.Email),
	}

	stripeKey := ""

	if params.StripeCustomerId != "" {
		stripeKey = params.StripeCustomerId
	} else {
		stripeKey = params.StripeKey
	}

	customerParams.SetSource(stripeKey)
	cus, err := customer.New(customerParams)
	println(cus.Sources.Data[0].Card.Country)
	println(cus.Sources.Data[0].Card.Brand)
	println(cus.Sources.Data[0].Card.ExpMonth)
	println(cus.Sources.Data[0].Card.Last4)
	pretty.Println("2222")

	if err != nil {
		pretty.Println(err)
		return false
	}
	pretty.Println("3333")
	pretty.Println(cus.ID)

	user.StripeCustomerId = cus.ID
	db.DB().Save(user)

	return true
}


func main() {
	lambda.Start(updateUser)
}
