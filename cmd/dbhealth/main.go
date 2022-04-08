package main

import (
	"net/http"
	"os"

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
	tableName := os.Getenv("DYNAMODB_TABLE")
	getDbHealth := adapter.NewAdapter(svc)
	_, err := getDbHealth.Health(tableName)
	if err != nil {
		zap.L().Fatal("Test connection to dynamodb", zap.Any("error", err.Error()))
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
