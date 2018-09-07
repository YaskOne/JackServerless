package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
)


func CreateUser(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := db.User{}

	if !(&user).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error in input parameters")
	}

	if valid, err := user.Valid(); !valid {
		return core.MakeHTTPError(400, err)
	}

	if !(&user).Create() {
		return core.MakeHTTPError(500, "Error creating user")
	}

	return core.MakeHTTPResponse(200, db.IdModel{user.ID})
}


func main() {
	lambda.Start(CreateUser)
}
