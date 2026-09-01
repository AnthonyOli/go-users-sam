# Go Users SAM

A serverless Go application for managing user data, deployed on AWS Lambda with PostgreSQL as the backend database.

## Overview

This project demonstrates a clean, scalable architecture for building serverless APIs in Go using AWS Lambda and Amazon API Gateway. It includes a generic service and repository pattern for database operations.

## Tech Stack

- **Runtime**: Go 1.26.5
- **Serverless Framework**: AWS Serverless Application Model (SAM)
- **Compute**: AWS Lambda
- **API**: AWS API Gateway
- **Database**: PostgreSQL
- **Driver**: lib/pq, pgx/v5

## Project Structure

```
.
├── cmd/                      # Entry points
│   ├── get-user/            # Lambda handler for getting users
│   └── hello-world/         # Sample Lambda function
├── internal/
│   ├── base/                # Generic base repository and service
│   ├── db/
│   │   ├── entities/        # Data models (User, etc.)
│   │   └── repositories/    # Repository implementations
│   └── services/            # Business logic layer
├── helpers/                 # Utility functions (DB connection pooling, etc.)
└── template.yaml            # SAM template for infrastructure as code
```

## Features

- **Get User by ID**: Retrieve user information via REST API endpoint `/users/{id}`
- **Generic Repository Pattern**: Abstract base repository and service for code reuse
- **Database Connection Pooling**: Efficient database connection management
- **Lambda Runtime**: Optimized for minimal cold start times

## User Model

```go
type User struct {
    Id           string
    Name         string
    Email        string
    PasswordHash string
    Phone        string
    IsActive     bool
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

## API Endpoints

### GET /users/{id}
Retrieve a user by their ID.

**Response**: User object as JSON

## Prerequisites

- Go 1.26.5+
- AWS SAM CLI
- AWS credentials configured
- PostgreSQL database

## Environment Variables

The Lambda functions require the following environment variables:

- `DB_HOST`: PostgreSQL host
- `DB_USER`: PostgreSQL user
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `DB_PORT`: PostgreSQL port

## Building

```bash
sam build
```

## Deployment

```bash
sam deploy --guided
```

During deployment, you'll be prompted to provide the database connection parameters via CloudFormation parameters.

## Development

```bash
go mod download
go mod tidy

sam local start-api --env-vars env.json --warm-containers EAGER
```

## Migrations
```bash 
migrate create -ext sql -dir db/migrations -seq migration-name
```

## License

MIT
