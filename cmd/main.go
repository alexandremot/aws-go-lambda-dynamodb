package main

import (
	"context"
	"net/http"

	"github.com/alexandremot/aws-go-lambda-dynamodb/myaws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := request.PathParameters["uuid"]

	responseBody := myaws.ConsultaDynamo(id)

	return events.APIGatewayProxyResponse{
		Body:       string(responseBody),
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
