package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
)


func LogUser(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := db.User{}

	user.Email = request.QueryStringParameters["email"]
	user.Password = request.QueryStringParameters["password"]

	if len(user.Email) == 0 || len(user.Password) == 0 {
		return core.MakeHTTPError(400, "Incorrect or empty parameters")
	}

	db.DB().Where(user).First(&user)

	if user.ID != 0 &&
		request.QueryStringParameters["email"] == user.Email &&
		request.QueryStringParameters["password"] == user.Password {
		return core.MakeHTTPResponse(200, db.UserResponse{db.GetUserObject(user)})
	} else {
		return core.MakeHTTPError(400, "Wrong credidentials")
	}
}


func main() {
	lambda.Start(LogUser)
}
