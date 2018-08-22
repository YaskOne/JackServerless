package core

import (
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
)

const (
	// ContentType is the name of the HTTP header
	// specifying the MIME type of the body.
	ContentType = "Content-Type"

	// AuthenticationHeader is the name of the HTTP header
	// where the signature should belong.
	AuthenticationHeader = "Authorization"

	// JSON is the MIME type for json data.
	JSON = "application/json"
)

type HTTPErrorBody struct {
	Error interface{} `json:"error"`
}

func MakeHTTPResponse(status int, data interface{}) (resp *events.APIGatewayProxyResponse, err error) {
	body, _ := json.Marshal(data)

	return &events.APIGatewayProxyResponse{
		StatusCode:      status,
		Body:            string(body),
		Headers:         map[string]string{ContentType: JSON},
	}, nil
}

func MakeHTTPError(status int, data interface{}) (resp *events.APIGatewayProxyResponse, err error) {
	body, _ := json.Marshal(HTTPErrorBody{data})

	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(body),
		Headers:    map[string]string{ContentType: JSON},
	}, nil
}