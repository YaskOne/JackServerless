package main


import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"log"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/ephemeralkey"
)

func ephemeralKey(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	stripeVersion := request.QueryStringParameters["api_version"]
	id := request.QueryStringParameters["id"]

	if stripeVersion == "" {
		return core.MakeHTTPError(400, "Stripe-Version not found")
	}
	params := &stripe.EphemeralKeyParams{
		Customer: stripe.String(id),
		StripeVersion: stripe.String(stripeVersion),
	}
	key, err := ephemeralkey.New(params)
	if err != nil {
		log.Printf("Stripe bindings call failed, %v\n", err)
		return core.MakeHTTPError(500, "Stripe bindings call failed")
	}
	return core.MakeHTTPResponse(http.StatusOK, key.RawJSON)
}


func main() {
	lambda.Start(ephemeralKey)
}
