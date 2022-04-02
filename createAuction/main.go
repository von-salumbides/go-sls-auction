package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
		log.Fatalf("Error parsing: %v", auction)
	}

	auctionBts, err := json.Marshal(auction)
	if err != nil {
		log.Fatalf("Error marshalling: %v", auctionBts)
	}
	resp := &HTTPApiResponse{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            string(auctionBts),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
