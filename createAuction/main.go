package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayV2HTTPResponse
type Auction struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

func ParseResponse(respString string) []byte {
	b := []byte(respString)
	var auction Auction
	err := json.Unmarshal(b, &auction)
	if err == nil {
		zap.L().Info("Unmarshall successfully:", zap.String("Title:", auction.Title))
	} else {
		zap.L().Fatal("Could not unmarshall JSON string", zap.Error(err))
	}
	return b
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (Response, error) {
	zap.L().Info("event received", zap.Any("method", event.RequestContext.HTTP.Method), zap.Any("path", event.RequestContext.HTTP.Path), zap.Any("body", event.Body), zap.Any("ctx", ctx))
	respBody := ParseResponse(event.Body)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(respBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Func-Reply": "createAuction",
		},
	}

	return resp, nil
}

func main() {
	logger := zap.L()
	logger.Sugar().Infow("Starting Lambda Service")
	lambda.Start(Handler)
}
