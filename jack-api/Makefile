OUTPUT = bin
LAMBDA_SUFFIX = .go

MAIN_OUTPUT = main

# Business
GET_BUSINESS_BY_ID_OUTPUT = getBusinessById
CREATE_BUSINESS = createBusiness
GET_BUSINESS_IN_AREA = getBusinessInArea

# Stocks
CREATE_PRODUCT = createProduct
CREATE_CATEGORY = createCategory
GET_BUSINESS_STOCKS = getBusinessStocks
GET_PRODUCTS_BY_ID = getProductsById

# Order
CREATE_ORDER = createOrder
GET_BUSINESS_ORDERS = getBusinessOrders

# Utilities
RESET_DB = resetDB

PACKAGED_TEMPLATE = packaged.yaml
S3_BUCKET := jack-world-app
STACK_NAME := JackStack3
TEMPLATE = template.yaml

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(MAIN_OUTPUT) $(GET_BUSINESS_BY_ID_OUTPUT) $(CREATE_BUSINESS) $(GET_BUSINESS_IN_AREA) $(RESET_DB) $(PACKAGED_TEMPLATE)

.PHONY: install
install:
	go get ./...

main:
	go build -o $(OUTPUT)/$(MAIN_OUTPUT) ./$(MAIN_OUTPUT).go

	go build -o $(OUTPUT)/$(GET_BUSINESS_BY_ID_OUTPUT) ./lambdas/business/$(GET_BUSINESS_BY_ID_OUTPUT).go
	go build -o $(OUTPUT)/$(CREATE_BUSINESS) ./lambdas/business/$(CREATE_BUSINESS).go
	go build -o $(OUTPUT)/$(GET_BUSINESS_IN_AREA) ./lambdas/business/$(GET_BUSINESS_IN_AREA).go

	go build -o $(OUTPUT)/$(CREATE_PRODUCT) ./lambdas/stock/$(CREATE_PRODUCT).go
	go build -o $(OUTPUT)/$(CREATE_CATEGORY) ./lambdas/stock/$(CREATE_CATEGORY).go
	go build -o $(OUTPUT)/$(GET_BUSINESS_STOCKS) ./lambdas/stock/$(GET_BUSINESS_STOCKS).go
	go build -o $(OUTPUT)/$(GET_PRODUCTS_BY_ID) ./lambdas/stock/$(GET_PRODUCTS_BY_ID).go

	go build -o $(OUTPUT)/$(CREATE_ORDER) ./lambdas/order/$(CREATE_ORDER).go
	go build -o $(OUTPUT)/$(GET_BUSINESS_ORDERS) ./lambdas/order/$(GET_BUSINESS_ORDERS).go

	go build -o $(OUTPUT)/$(RESET_DB) ./lambdas/utils/$(RESET_DB).go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

.PHONY: build
build: clean lambda

.PHONY: api
api: build
	sam local start-api

.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM
