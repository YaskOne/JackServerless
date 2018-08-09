package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"os"
)

var database *dynamodb.DynamoDB = nil

func DB() *dynamodb.DynamoDB {
	if database == nil {
		initializeDB()
	}
	return database
}

func initializeDB() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println("ERROR: Failed opening aws session")
		fmt.Println(err.Error())
		os.Exit(0)
	}

	database = dynamodb.New(sess)

	initializeBusiness()
}