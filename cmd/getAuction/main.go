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
	pathParamId := request.PathParameters["id"]
	getAuction := adapter.NewAdapter(svc)
	auction, err := getAuction.Get(pathParamId, tableName)
	if err != nil {
		logger.ERROR("GET request failed", err.Error())
		return httpApi.ERRORInternalServer(), err
	}
	if len(auction.Item) == 0 {
		logger.INFO("0 item")
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusNoContent,
			Body:       "{\"items\":\"0\"}",
		}, nil
	}
	item := models.Auction{}
	err = dynamodbattribute.UnmarshalMap(auction.Item, &item)
	if err != nil {
		logger.ERROR("Failed to UnmarshalMap", err.Error())
		return httpApi.ERRORInternalServer(), err
	}
	// marshal to type bytes
	marshalledItem, err := json.Marshal(item)
	if err != nil {
		logger.ERROR("Failed to Marshal", err.Error())
		return httpApi.ERRORInternalServer(), err
	}
	logger.INFO("Success", zap.Any("item", string(marshalledItem)))
	return httpApi.OKResponse(string(marshalledItem)), nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
