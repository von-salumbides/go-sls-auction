package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/von-salumbides/auction/internal/repository/adapter"
	"github.com/von-salumbides/auction/internal/server"
	httpApi "github.com/von-salumbides/auction/utils/http"
	"github.com/von-salumbides/auction/utils/logger"
	"go.uber.org/zap"
)

func Handler(request events.APIGatewayV2HTTPRequest) (*httpApi.HTTPApiResponse, error) {
	svc := server.GetConnection()
	// tableName for DyanmoDb
	getDbHealth := adapter.NewAdapter(svc)
	ok := getDbHealth.Health()
	if !ok {
		zap.L().Fatal("Test connection to dynamodb")
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return &httpApi.HTTPApiResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
