package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"context"
	"JackServerless/jack-api/db"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

type body struct {
	Message string `json:"message"`
}

// Handler is the Lambda function handler
func Handler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda request", request.RequestContext.RequestID)

	//result, err := db.DB().ListTables(&dynamodb.ListTablesInput{})
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return events.APIGatewayProxyResponse{
	//		Body:       string(err.Error()),
	//		StatusCode: 200,
	//	}, nil
	//}

	fmt.Println("Tables:")
	fmt.Println("")

	//for _, n := range result.TableNames {
	//	fmt.Println(*n)
	//}
	//
	//params := &dynamodb.ScanInput{
	//	TableName: aws.String("BusinessTable"),
	//}
	//result2, err2 := db.DB().Scan(params)
	//if err2 != nil {
	//	fmt.Errorf("failed to make Query API call, %v", err)
	//	return events.APIGatewayProxyResponse{
	//		Body:       string(err2.Error()),
	//		StatusCode: 200,
	//	}, nil
	//}
	//obj := []db.Business{}
	//err = dynamodbattribute.UnmarshalListOfMaps(result2.Items, &obj)
	//if err != nil {
	//	fmt.Errorf("failed to unmarshal Query result items, %v", err)
	//	return events.APIGatewayProxyResponse{
	//		Body:       string(err.Error()),
	//		StatusCode: 200,
	//	}, nil
	//}
	//pretty.Println(obj[0])
	//if len(obj) > 0 {
	//	println(obj[0].Address)
	//	b, _ := json.Marshal(body{Message: obj[0].Address})
	//
	//	return events.APIGatewayProxyResponse{
	//		Body:       string(b),
	//		StatusCode: 200,
	//	}, nil
	//} else {

		b, _ := json.Marshal(body{Message: "main"})

		return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 200,
		}, nil
	//}
}


func main() {
	db.DB()
	lambda.Start(Handler)
}
