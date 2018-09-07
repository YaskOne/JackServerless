package main

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/kr/pretty"
)

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func jackAuthorizer(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	println("Authorizer ----------")
	pretty.Println(event.Type)
	token := event.AuthorizationToken

	switch strings.ToLower(token) {
		case "allow":
			return generatePolicy("user", "Allow", event.MethodArn), nil
		case "deny":
			return generatePolicy("user", "Deny", event.MethodArn), nil
		case "unauthorized":
			return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
		default:
			return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Error: Invalid token")
	}
}

func main() {
	lambda.Start(jackAuthorizer)
}