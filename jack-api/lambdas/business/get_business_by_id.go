package business

import (
        "log"

        "github.com/aws/aws-lambda-go/events"
        "github.com/aws/aws-lambda-go/lambda"
        "context"
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
        "github.com/aws/aws-sdk-go/service/dynamodb"
        "github.com/aws/aws-sdk-go/aws"
        "fmt"
        "JackServerless/jack-api/models"
)

// Handler is the Lambda function handler
func GetBusinessById(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        log.Println("Lambda request", request.RequestContext.RequestID)

        result, err := models.DB().GetItem(&dynamodb.GetItemInput{
                TableName: aws.String("Business"),
                Key: map[string]*dynamodb.AttributeValue{
                        "id": {
                                N: aws.String("0"),
                        },
                },
        })

        if err != nil {
                fmt.Println(err.Error())
        }

        item := models.Business{}

        err = dynamodbattribute.UnmarshalMap(result.Item, &item)

        if err != nil {
                panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
        }

        fmt.Println("Found item:")
        fmt.Println("Year:  ", result)
        fmt.Println("Year:  ", result.Item)
        fmt.Println("Year:  ", item.Name)
        fmt.Println("Title: ", item.Latitude)
        fmt.Println("Plot:  ", item.Address)
        fmt.Println("Rating:", item.Type)


        return events.APIGatewayProxyResponse{
                Body:       string("GetBusiness"),
                StatusCode: 200,
        }, nil
}

func main() {
        lambda.Start(GetBusinessById)
}
