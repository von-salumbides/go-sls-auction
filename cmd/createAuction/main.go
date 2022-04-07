package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/von-salumbides/auction/internal/models"
	httpApi "github.com/von-salumbides/auction/utils/http"
	"github.com/von-salumbides/auction/utils/logger"
	"go.uber.org/zap"
)

func Handler(request events.APIGatewayV2HTTPRequest) (*httpApi.HTTPApiResponse, error) {

	// itemString unmarshal to Auction to access object properties
	itemString := request.Body
	itemStruct := models.Auction{}
	err := json.Unmarshal([]byte(itemString), &itemStruct)
	if err != nil {
		zap.L().Fatal("Error parsing", zap.Any("error", err.Error()))
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	// create of new item of type Auction
	itemTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	item := models.Auction{
		Title:       itemStruct.Title,
		Status:      itemStruct.Status,
		DateCreated: itemTime,
	}

	// marshal item
	av, err := json.Marshal(item)
	if err != nil {
		zap.L().Fatal("Error marshalling item", zap.Any("error", err.Error()))
		return &httpApi.HTTPApiResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	resp := &httpApi.HTTPApiResponse{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            string(av),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func init() {
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
