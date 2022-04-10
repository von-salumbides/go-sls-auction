package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/von-salumbides/auction/internal/models"
	"github.com/von-salumbides/auction/internal/repository/adapter"
	"github.com/von-salumbides/auction/internal/server"
	httpApi "github.com/von-salumbides/auction/utils/http"
	"github.com/von-salumbides/auction/utils/logger"
	"go.uber.org/zap"
)

func Handler(request events.APIGatewayV2HTTPRequest) (*httpApi.HTTPApiResponse, error) {
	svc := server.GetConnection()
	// itemUuid is a unique id for item
	itemUuid := uuid.New().String()

	// itemString unmarshal to Auction to access object properties
	itemString := request.Body
	itemStruct := models.Auction{}
	err := json.Unmarshal([]byte(itemString), &itemStruct)
	if err != nil {
		logger.ERROR("Error parsing request body", err.Error())
		return httpApi.ERRORInternalServer(), err
	}

	// create of new item of type Auction
	itemTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	item := models.Auction{
		Id:          itemUuid,
		Title:       itemStruct.Title,
		Status:      itemStruct.Status,
		DateCreated: itemTime,
	}
	// tableName for DyanmoDb
	tableName := os.Getenv("DYNAMODB_TABLE")

	// PutItem createOrupdate data to Dynamodb
	putItem := adapter.NewAdapter(svc)
	_, err = putItem.CreateOrUpdate(item, tableName)
	if err != nil {
		logger.ERROR("Got error calling put item", err.Error())
		return httpApi.ERRORInternalServer(), err
	}
	// marshal item
	av, err := json.Marshal(item)
	logger.INFO("Response", zap.Any("value", av))
	if err != nil {
		logger.ERROR("Error marshalling item", err.Error())
		return httpApi.ERRORInternalServer(), err
	}

	// resp build the response to client
	resp := httpApi.OKResponse(string(av))
	return resp, nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
