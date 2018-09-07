package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
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

	if params.Name != "" {
		user.Name = params.Name
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Password != "" {
		user.Password = params.Password
	}

	db.DB().Save(&user)

	return core.MakeHTTPResponse(http.StatusOK, db.UserResponse{user})
}


func main() {
	lambda.Start(updateUser)
}
