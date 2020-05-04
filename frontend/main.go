/*
  AWS lambda front end service.
*/
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	//"github.com/nlopes/slack"
)


func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("HandleRequest: Processing request data for request=", request.RequestContext.RequestID, len(request.Body))

	// FIX TODO route message appropriately

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
