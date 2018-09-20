OUTPUT = bin
LAMBDA_SUFFIX = .go

PACKAGED_TEMPLATE = packaged.yaml
S3_BUCKET := jack-world-app
STACK_NAME := JackStage
TEMPLATE = template.yaml

#JackTakewayDev
#JackTakewayProd
#JackTakewayStage

#JackBusinessDev
#JackBusinessProd
#JackBusinessStage



# Business
GET_BUSINESS_BY_ID_OUTPUT = getBusinessById
CREATE_BUSINESS = createBusiness
LOG_BUSINESS = logBusiness
DELETE_BUSINESS = deleteBusiness
UPDATE_BUSINESS = updateBusiness

# Stocks
CREATE_PRODUCT = createProduct
CREATE_CATEGORY = createCategory
GET_BUSINESS_STOCKS = getBusinessStocks
GET_PRODUCTS = getProducts
GET_CATEGORIES = getCategories
UPDATE_CATEGORY = updateCategory
UPDATE_PRODUCT = updateProduct
DELETE_ITEM = deleteItem
ORDERS_WATCHER = ordersWatcher

# Order
CREATE_ORDER = createOrder
GET_ORDERS = getOrders
UPDATE_ORDER = updateOrder

# User
LOG_USER = logUser
CREATE_USER = createUser
GET_USERS = getUsers
DELETE_USER = deleteUser
UPDATE_USER = updateUser

# Payment
EPHEMERAL_KEY = ephemeralKey

# Utilities
RESET_DB = resetDB
JACK_AUTHORIZER = jackAuthorizer

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(OUTPUT)/$(GET_BUSINESS_BY_ID_OUTPUT)
	rm -f $(OUTPUT)/$(CREATE_BUSINESS)
	rm -f $(OUTPUT)/$(CREATE_PRODUCT)
	rm -f $(OUTPUT)/$(CREATE_CATEGORY)
	rm -f $(OUTPUT)/$(GET_BUSINESS_STOCKS)
	rm -f $(OUTPUT)/$(GET_PRODUCTS)
	rm -f $(OUTPUT)/$(GET_CATEGORIES)
	rm -f $(OUTPUT)/$(CREATE_ORDER)
	rm -f $(OUTPUT)/$(GET_BUSINESS_ORDERS)
	rm -f $(OUTPUT)/$(RESET_DB)

.PHONY: install
install:
	go get ./...

main:
	go build -o $(OUTPUT)/$(GET_BUSINESS_BY_ID_OUTPUT) ./lambdas/business/$(GET_BUSINESS_BY_ID_OUTPUT).go
	go build -o $(OUTPUT)/$(CREATE_BUSINESS) ./lambdas/business/$(CREATE_BUSINESS).go
	go build -o $(OUTPUT)/$(LOG_BUSINESS) ./lambdas/business/$(LOG_BUSINESS).go
	go build -o $(OUTPUT)/$(DELETE_BUSINESS) ./lambdas/business/$(DELETE_BUSINESS).go
	go build -o $(OUTPUT)/$(UPDATE_BUSINESS) ./lambdas/business/$(UPDATE_BUSINESS).go

	go build -o $(OUTPUT)/$(CREATE_PRODUCT) ./lambdas/stock/$(CREATE_PRODUCT).go
	go build -o $(OUTPUT)/$(CREATE_CATEGORY) ./lambdas/stock/$(CREATE_CATEGORY).go
	go build -o $(OUTPUT)/$(GET_BUSINESS_STOCKS) ./lambdas/stock/$(GET_BUSINESS_STOCKS).go
	go build -o $(OUTPUT)/$(GET_PRODUCTS) ./lambdas/stock/$(GET_PRODUCTS).go
	go build -o $(OUTPUT)/$(GET_CATEGORIES) ./lambdas/stock/$(GET_CATEGORIES).go
	go build -o $(OUTPUT)/$(UPDATE_CATEGORY) ./lambdas/stock/$(UPDATE_CATEGORY).go
	go build -o $(OUTPUT)/$(UPDATE_PRODUCT) ./lambdas/stock/$(UPDATE_PRODUCT).go

	go build -o $(OUTPUT)/$(DELETE_ITEM) ./lambdas/$(DELETE_ITEM).go

	go build -o $(OUTPUT)/$(LOG_USER) ./lambdas/user/$(LOG_USER).go
	go build -o $(OUTPUT)/$(CREATE_USER) ./lambdas/user/$(CREATE_USER).go
	go build -o $(OUTPUT)/$(GET_USERS) ./lambdas/user/$(GET_USERS).go
	go build -o $(OUTPUT)/$(DELETE_USER) ./lambdas/user/$(DELETE_USER).go
	go build -o $(OUTPUT)/$(UPDATE_USER) ./lambdas/user/$(UPDATE_USER).go

	go build -o $(OUTPUT)/$(CREATE_ORDER) ./lambdas/order/$(CREATE_ORDER).go
	go build -o $(OUTPUT)/$(GET_ORDERS) ./lambdas/order/$(GET_ORDERS).go
	go build -o $(OUTPUT)/$(UPDATE_ORDER) ./lambdas/order/$(UPDATE_ORDER).go
	go build -o $(OUTPUT)/$(ORDERS_WATCHER) ./lambdas/order/$(ORDERS_WATCHER).go

	go build -o $(OUTPUT)/$(RESET_DB) ./lambdas/utils/$(RESET_DB).go
	go build -o $(OUTPUT)/$(JACK_AUTHORIZER) ./lambdas/auth/$(JACK_AUTHORIZER).go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

.PHONY: build
build: lambda

.PHONY: api
api: build
	sam local start-api

.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM
