package main

import (
	"context"
	"encoding/json"
	"go-serverless/helpers"
	"go-serverless/internal/db/repositories"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pool, err := helpers.NewPool(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	defer pool.Close()

	queryParams := request.PathParameters["id"]

	userRepo := repositories.NewUserRepository(pool)

	user, err := userRepo.GetById(ctx, queryParams)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
