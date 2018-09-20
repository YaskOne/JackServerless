package main

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
)

func updateBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := db.Business{DisponibilityStatus: -1}
	business := db.Business{}

	if !(&params).Parse(request.Body) {
		return core.MakeHTTPError(400, "Error parsing business")
	}

	business.ID = params.ID

	if !(&business).Load() {
		return core.MakeHTTPError(400, "Business not found")
	}

	if params.Name != "" {
		business.Name = params.Name
	}
	if params.Url != "" {
		business.Url = params.Url
	}
	if params.Type != "" {
		business.Type = params.Type
	}
	if params.Address != "" {
		business.Address = params.Address
	}
	if params.Description != "" {
		business.Description = params.Description
	}
	if params.Email != "" {
		business.Email = params.Email
	}
	if params.Password != "" {
		business.Password = params.Password
	}
	if params.FcmToken != "" {
		business.FcmToken = params.FcmToken
	}
	if params.DefaultPreparationDuration != 0 {
		business.DefaultPreparationDuration = params.DefaultPreparationDuration
	}
	if params.DisponibilityStatus != -1 {
		business.DefaultPreparationDuration = params.DefaultPreparationDuration
	}

	db.DB().Save(&business)

	return core.MakeHTTPResponse(http.StatusOK, db.BusinessResponse{business})
}


func main() {
	lambda.Start(updateBusiness)
}
