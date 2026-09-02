package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-serverless/helpers"
	"go-serverless/internal/db/entities"
	"go-serverless/internal/db/repositories"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func castUpdateUserRequestToEntity(request UpdateUserRequest) entities.User {
	return entities.User{
		Id:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		IsActive: request.IsActive,
	}
}

func validateUpdateUserRequest(request UpdateUserRequest) error {
	if request.Id == "" {
		return fmt.Errorf("id is required")
	}
	if request.Name == "" {
		return fmt.Errorf("name is required")
	}
	if request.Email == "" {
		return fmt.Errorf("email is required")
	}
	if request.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if request.IsActive != true && request.IsActive != false {
		return fmt.Errorf("is_active must be a boolean")
	}
	return nil
}

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

	var updateUserRequest UpdateUserRequest

	err = json.Unmarshal([]byte(request.Body), &updateUserRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	updateUserRequest.Id = userId

	err = validateUpdateUserRequest(updateUserRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	userRepo := repositories.NewUserRepository(pool)
	userEntity := castUpdateUserRequestToEntity(updateUserRequest)

	savedUser, err := userRepo.Save(ctx, &userEntity)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error":"failed to save user", "message":"` + err.Error() + `"}`,
		}, nil
	}

	body, err := json.Marshal(savedUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error":"failed to marshal response"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
