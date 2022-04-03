package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/von-salumbides/auction/utils/logger"
	"go.uber.org/zap"
)

type Auction struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

type HTTPApiResponse events.APIGatewayV2HTTPResponse

func Handler(request events.APIGatewayV2HTTPRequest) (*HTTPApiResponse, error) {
	var auction Auction
	err := json.Unmarshal([]byte(request.Body), &auction)
	if err != nil {
		zap.L().Fatal("Error parsing:")
	}
	auction.DateCreated = time.Now().Format("01-02-2006 15:04:05 Monday")
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
	logger.InitZap()
}

func main() {
	lambda.Start(Handler)
}
