package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/utils"
)

func getUsers(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	users := []db.User{}
	userObjects := []db.UserObject{}
	ids := request.QueryStringParameters["ids"]

	if len(ids) == 0 {
		db.DB().Find(&users)
	} else {
		db.DB().
			Where(utils.SplitArrayString(ids)).
			Find(&users)
	}
	var i = 0
	for i < len(users) {
		userObjects = append(userObjects, db.GetUserObject(users[i]))
		i += 1
	}

	return core.MakeHTTPResponse(200, db.UsersResponse{userObjects})
}

func main() {
	lambda.Start(getUsers)
}

