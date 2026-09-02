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

	userId := request.PathParameters["id"]
	if helpers.IsValidUUID(userId) == false {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: `{"error":"invalid user ID"}`}, nil
	}

	userRepo := repositories.NewUserRepository(pool)

	user, err := userRepo.GetById(ctx, userId)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: `{"error":"failed to get user", "message":"` + err.Error() + `"}`}, nil
	}

	if user == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: `{"error":"user not found"}`}, nil
	}

	body, err := json.Marshal(user)
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
