package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
)

// Handler is the Lambda function handler
func createBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	business := db.Business{}

	if !(&business).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if valid, err := business.Valid(); !valid {
		return core.MakeHTTPError(400, err)
	}

	if !(&business).Create() {
		return core.MakeHTTPResponse(500, "Error creating business")
	}

	return core.MakeHTTPResponse(200, db.IdModel{business.ID})

}

func main() {
	lambda.Start(createBusiness)
}
