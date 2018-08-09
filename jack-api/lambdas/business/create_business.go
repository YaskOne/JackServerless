package business

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"JackServerless/jack-api/models"
)

// Handler is the Lambda function handler
func CreateBusiness(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda request", request.RequestContext.RequestID)

	latLng := models.LatLng{12, 10}
	item := models.Business{
		latLng,
		"Arthur",
		"19 route des gardes",
		"Resto",
		"yayayayaya",
		"htttp:dwinfwnefnewnfwennjwe",
	}
	av, err := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("Movies"),
	}

	_, err = models.DB().PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added 'The Big New Movie' (2015) to Movies table")

	return events.APIGatewayProxyResponse{
		Body:       string("CreateBusiness"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(GetBusinessById)
}
