.PHONY: build clean deploy remove
environment = $(DEPLOY_ENV)
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --stage=$(environment) --verbose

remove: 
	sls remove --stage=$(environment)
