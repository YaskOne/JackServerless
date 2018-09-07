package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

// Handler is the Lambda function handler
func deleteBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	business := db.Business{}

	if !(&business).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if !(&business).Load() {
		return core.MakeHTTPError(400, "Business not found")
	}

	if !(&business).Delete() {
		return core.MakeHTTPError(400, "Error deleting business")
	}

	return core.MakeHTTPResponse(200, db.IdModel{business.ID})
}

func main() {
	lambda.Start(deleteBusiness)
}
