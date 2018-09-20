package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/utils"
	"strconv"
)

func getUsers(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	users := []db.User{}
	userObjects := []interface{}{
	}
	ids := request.QueryStringParameters["ids"]

	userId64, _ := strconv.ParseUint(request.QueryStringParameters["user_id"], 10,64)
	userId := uint(userId64)

	if len(ids) == 0 {
		db.DB().Find(&users)
	} else {
		db.DB().
			Where(utils.SplitArrayString(ids)).
			Find(&users)
	}
	var i = 0
	for i < len(users) {
		if users[i].ID == userId {
			userObjects = append(userObjects, users[i])
		} else {
			userObjects = append(userObjects, db.GetUserObject(users[i]))
		}
		i += 1
	}

	return core.MakeHTTPResponse(200, db.UsersResponse{userObjects})
}

func main() {
	lambda.Start(getUsers)
}

