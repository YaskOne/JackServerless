OUTPUT = bin
LAMBDA_SUFFIX = .go
MAIN_OUTPUT = main
GET_BUSINESS_BY_ID_OUTPUT = getBusinessById
CREATE_BUSINESS = createBusiness
PACKAGED_TEMPLATE = packaged.yaml
S3_BUCKET := jack-world-app
STACK_NAME := JackStack3
TEMPLATE = template.yaml

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(MAIN_OUTPUT) $(GET_BUSINESS_BY_ID_OUTPUT) $(CREATE_BUSINESS) $(PACKAGED_TEMPLATE)

.PHONY: install
install:
	go get ./...

main: ./main.go
	go build -o $(OUTPUT)/$(MAIN_OUTPUT) ./$(MAIN_OUTPUT).go
	go build -o $(OUTPUT)/$(GET_BUSINESS_BY_ID_OUTPUT) ./lambdas/business/$(GET_BUSINESS_BY_ID_OUTPUT).go
	go build -o $(OUTPUT)/$(CREATE_BUSINESS) ./lambdas/business/$(CREATE_BUSINESS).go

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
