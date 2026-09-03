package main

import (
	"context"
	"encoding/json"
	"go-serverless/helpers"
	"go-serverless/internal/db/repositories"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pool, err := helpers.NewPool(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	defer pool.Close()

	page, errPage := strconv.Atoi(request.QueryStringParameters["page"])
	pageSize, errPageSize := strconv.Atoi(request.QueryStringParameters["pageSize"])

	if errPage != nil {
		page = 1
	}

	if errPageSize != nil {
		pageSize = 10
	}

	userRepo := repositories.NewUserRepository(pool)

	users, err := userRepo.List(ctx, &pageSize, &page)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: `{"error":"failed to get users", "message":"` + err.Error() + `"}`}, nil
	}

	if users == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: `{"error":"users not found"}`}, nil
	}

	body, err := json.Marshal(users)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: `{"error":"failed to marshal response"}`}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
