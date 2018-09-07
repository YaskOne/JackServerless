package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/core"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/utils"
)


func getBusinessById(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	businesses := []db.Business {}

	ids := request.QueryStringParameters["ids"]

	if len(ids) == 0 {
		db.DB().Find(&businesses)
	} else {
		db.DB().
			Where(utils.SplitArrayString(ids)).
			Find(&businesses)
	}

	return core.MakeHTTPResponse(200, db.BusinessesResponse{businesses})
}


func main() {
	lambda.Start(getBusinessById)
}
