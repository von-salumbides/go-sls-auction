package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/von-salumbides/auction/internal/repository/adapter"
	"github.com/von-salumbides/auction/internal/server"
	httpApi "github.com/von-salumbides/auction/utils/http"
	"github.com/von-salumbides/auction/utils/logger"
)

func Handler(request events.APIGatewayV2HTTPRequest) (*httpApi.HTTPApiResponse, error) {
	svc := server.GetConnection()
	// tableName for DyanmoDb
	tableName := os.Getenv("DYNAMODB_TABLE")
	getDbHealth := adapter.NewAdapter(svc)
	_, err := getDbHealth.Health(tableName)
	if err != nil {
		logger.ERROR("Test connection to dynamodb", err.Error())
		return httpApi.ERRORInternalServer(), err
	}
	return httpApi.OKResponse("{\"healthy\":\"ok\"}"), nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
