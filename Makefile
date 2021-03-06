.PHONY: build clean deploy
environment = $(DEPLOY_ENV)
function_name = $(FUNCTION_NAME)
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/createAuction cmd/createAuction/*.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getAuctions cmd/getAuctions/*.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getAuction cmd/getAuction/*.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/dbHealth cmd/dbhealth/*.go
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --stage=$(environment) --verbose

deployFunc: clean build
	sls deploy -f $(function_name) --stage=$(environment) --verbose

remove: 
	sls remove --stage=$(environment) --verbose
