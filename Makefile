.PHONY: build clean deploy
environment = $(DEPLOY_ENV)
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/createAuction createAuction/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --stage=$(environment) --verbose

remove: 
	sls remove --stage=$(environment)
