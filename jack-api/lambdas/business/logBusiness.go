package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
)


func logBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	business := db.Business{}

	business.Email = request.QueryStringParameters["email"]
	business.Password = request.QueryStringParameters["password"]

	if len(business.Email) == 0 || len(business.Password) == 0 {
		return core.MakeHTTPError(400, "Incorrect or empty parameters")
	}

	db.DB().Where(business).First(&business)

	if business.ID != 0 &&
		request.QueryStringParameters["password"] == business.Password {
		return core.MakeHTTPResponse(200, db.BusinessResponse{business})
	} else {
		return core.MakeHTTPError(400, "Wrong credidentials")
	}
}


func main() {
	lambda.Start(logBusiness)
}
