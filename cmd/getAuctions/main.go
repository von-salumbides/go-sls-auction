package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/von-salumbides/auction/internal/models"
	"github.com/von-salumbides/auction/internal/repository/adapter"
	"github.com/von-salumbides/auction/internal/server"
	httpApi "github.com/von-salumbides/auction/utils/http"
	"github.com/von-salumbides/auction/utils/logger"
	"go.uber.org/zap"
)

func Handler(request events.APIGatewayV2HTTPRequest) (*httpApi.HTTPApiResponse, error) {
	svc := server.GetConnection()
	tableName := os.Getenv("DYNAMODB_TABLE")
	getAll := adapter.NewAdapter(svc)
	tableData, err := getAll.GetAll(tableName)
	if err != nil {
		zap.L().Fatal("Scan API call failed", zap.Any("error", err.Error()))
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	itemArray := []models.Auction{}
	for _, i := range tableData.Items {
		item := models.Auction{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			zap.L().Fatal("Got error unmarshalling", zap.Any("error", err.Error()))
			return &httpApi.HTTPApiResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}
		itemArray = append(itemArray, item)
	}
	zap.L().Info("Succesfully unmarshal item", zap.Any("itemArray", itemArray))
	itemArrayString, err := json.Marshal(itemArray)
	if err != nil {
		zap.L().Fatal("Got error marshalling result", zap.Any("error", err.Error()))
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return &httpApi.HTTPApiResponse{
		StatusCode: http.StatusOK,
		Body:       string(itemArrayString),
	}, nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
