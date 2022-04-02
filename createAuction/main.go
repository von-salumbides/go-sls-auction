package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

type Auction struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type HTTPApiResponse events.APIGatewayV2HTTPResponse

func Handler(request events.APIGatewayV2HTTPRequest) (*HTTPApiResponse, error) {
	var auction Auction
	err := json.Unmarshal([]byte(request.Body), &auction)
	if err != nil {
		zap.L().Fatal("Error parsing:")
	}

	auctionBts, err := json.Marshal(auction)
	if err != nil {
		zap.L().Fatal("Error Marshal")
	}
	resp := &HTTPApiResponse{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            string(auctionBts),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	zap.L().Info("Event Received", zap.Any("body", request.Body))
	return resp, nil
}

func init() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}

func main() {
	lambda.Start(Handler)
}
