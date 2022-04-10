package main

import (
	"encoding/json"
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
		logger.ERROR("Scan API call failed", err.Error())
		return httpApi.ERRORInternalServer(), err
	}

	itemArray := []models.Auction{}
	for _, i := range tableData.Items {
		item := models.Auction{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			logger.ERROR("Got error unmarshalling", err.Error())
			return httpApi.ERRORInternalServer(), err
		}
		itemArray = append(itemArray, item)
	}
	logger.INFO("Succesfully unmarshal item", zap.Any("items", itemArray))
	itemArrayString, err := json.Marshal(itemArray)
	if err != nil {
		logger.ERROR("Got error marshalling result", err.Error())
		return httpApi.ERRORInternalServer(), err
	}

	return httpApi.OKResponse(string(itemArrayString)), nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
