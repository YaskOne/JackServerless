package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

type LatLng struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Business struct {
	LatLng
	Name string `json:"name"`
	Address string `json:"address"`
	Type string `json:"type"`
	Description string `json:"description"`
	Url string `json:"url"`
}

func initializeBusiness() {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("latitude"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("longitude"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("address"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("type"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("description"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("url"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("latitude"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("longitude"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("address"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("type"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("description"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("url"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Business"),
	}

	_, err := DB().CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}

}