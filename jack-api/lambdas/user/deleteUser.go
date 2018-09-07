package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

// Handler is the Lambda function handler
func deleteUser(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := db.User{}

	if !(&user).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if !(&user).Load() {
		return core.MakeHTTPError(400, "User not found")
	}

	if !(&user).Delete() {
		return core.MakeHTTPError(500, "Error deleting user")
	}

	return core.MakeHTTPResponse(200, db.IdModel{user.ID})

}

func main() {
	lambda.Start(deleteUser)
}

