package lambdas

import (
	"github.com/aws/aws-lambda-go/events"
	"JackServerless/jack-api/db"
	"JackServerless/jack-api/core"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"encoding/json"
	"net/http"
)

type deleteItemRequest struct {
	UserId        uint `json:"user_id"`
	BusinessId        uint `json:"business_id"`
	ProductId        uint `json:"product_id"`
	CategoryId        uint `json:"category_id"`
}

// Handler is the Lambda function handler
func deleteItem(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var params deleteItemRequest

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		return core.MakeHTTPError(http.StatusNotAcceptable, err.Error())
	}

	if params.UserId != 0 {
		return deleteUser(params.UserId)
	} else if params.BusinessId != 0 {
		return deleteBusiness(params.BusinessId)
	} else if params.ProductId != 0 {
		return deleteProduct(params.ProductId)
	} else if params.CategoryId != 0 {
		return deleteCategory(params.CategoryId)
	}

	return core.MakeHTTPError(http.StatusNotAcceptable, "no id given")
}

func deleteUser(id uint) (*events.APIGatewayProxyResponse, error) {
	object := db.User{ID: id}

	if !(&object).Load() {
		return core.MakeHTTPError(400, "User not found")
	}

	if !(&object).Delete() {
		return core.MakeHTTPError(500, "Error deleting user")
	}

	return core.MakeHTTPResponse(200, db.IdModel{object.ID})
}
func deleteBusiness(id uint) (*events.APIGatewayProxyResponse, error) {

	object := db.Business{ID: id}

	if !(&object).Load() {
		return core.MakeHTTPError(400, "Business not found")
	}

	if !(&object).Delete() {
		return core.MakeHTTPError(400, "Error deleting business")
	}

	return core.MakeHTTPResponse(200, db.IdModel{object.ID})
}
func deleteCategory(id uint) (*events.APIGatewayProxyResponse, error) {

	object := db.Category{ID: id}

	if !(&object).Load() {
		return core.MakeHTTPError(400, "Category not found")
	}

	if !(&object).Delete() {
		return core.MakeHTTPError(400, "Error deleting category")
	}

	return core.MakeHTTPResponse(200, db.IdModel{object.ID})
}
func deleteProduct(id uint) (*events.APIGatewayProxyResponse, error) {

	object := db.Product{ID: id}

	if !(&object).Load() {
		return core.MakeHTTPError(400, "Business not found")
	}

	if !(&object).Delete() {
		return core.MakeHTTPError(400, "Error deleting business")
	}

	return core.MakeHTTPResponse(200, db.IdModel{object.ID})
}


func main() {
	lambda.Start(deleteItem)
}

