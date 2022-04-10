package http

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type HTTPApiResponse events.APIGatewayV2HTTPResponse

func ERRORInternalServer() *HTTPApiResponse {
	return &HTTPApiResponse{
		StatusCode: http.StatusInternalServerError,
	}
}

func OKResponse(body string) *HTTPApiResponse {
	return &HTTPApiResponse{
		StatusCode:      http.StatusOK,
		Body:            body,
		IsBase64Encoded: false,
		Headers:         map[string]string{"Content-Type": "application/json"},
	}
}
