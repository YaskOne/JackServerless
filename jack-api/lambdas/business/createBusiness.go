package main

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"JackServerless/jack-api/db"
	"encoding/json"
	"JackServerless/jack-api/core"
	"github.com/kr/pretty"
)

// Handler is the Lambda function handler
func CreateBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	//pretty.Println(request.RequestContext)
	//pretty.Println(request.RequestContext.Authorizer)
	//pretty.Println(request.RequestContext.Authorizer["cognito:username"])
	log.Println("Lambda request", request.RequestContext.RequestID)

	business := db.Business{}

	if err := json.Unmarshal([]byte(request.Body), &business); err != nil {
		return core.MakeHTTPError(400, err.Error())
	}
	pretty.Println(business)
	if !db.ValidateCreateBusiness(business) {
		return core.MakeHTTPError(400, "insufficient or incomplete parameters")
	}

	if !db.CreateBusiness(&business) {
		return core.MakeHTTPResponse(500, "Error creating business")
	}

	//pretty.Println(*(&business))
	//if res := db.DB().Create(&business); res.Error != nil {
	//	println("errr not nil")
	//	pretty.Println(res.Error)
	//}
	//
	////business.Model = gorm.Model{}
	//pretty.Println(res)
	return core.MakeHTTPResponse(200, db.IDResponse{business.ID})

}

func main() {
	lambda.Start(CreateBusiness)
}
