# SplitwiseClone Backend

## Overview
SplitwiseClone is a backend web application built using the Gin framework and AWS DynamoDB. It provides a platform for users to manage expenses within groups, track balances, and settle debts.

## Features
- User registration and authentication
- Group creation and management
- Expense tracking and splitting
- Balance calculation and settlement

## Technologies Used
- Go (Golang)
- Gin Framework
- AWS DynamoDB
- JWT for authentication

## Project Structure
```
splitwiseclone-backend
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── user.go        # User-related API handlers
│   │   │   ├── group.go       # Group-related API handlers
│   │   │   ├── expense.go     # Expense-related API handlers
│   │   │   └── auth.go        # Authentication-related API handlers
│   │   └── routes.go          # API route definitions
│   ├── db
│   │   ├── dynamo.go          # DynamoDB client setup
│   │   └── models
│   │       ├── user.go        # User model
│   │       ├── group.go       # Group model
│   │       ├── group_member.go # GroupMember model
│   │       ├── expense.go      # Expense model
│   │       └── expense_split.go # ExpenseSplit model
│   ├── logic
│   │   └── balances.go        # Logic for calculating balances
│   ├── auth
│   │   └── jwt.go             # JWT token handling
│   ├── config
│   │   └── config.go          # Configuration settings
│   └── utils
│       ├── hash.go            # Password hashing utilities
│       └── response.go        # JSON response utilities
├── go.mod                      # Go module definition
├── go.sum                      # Module dependency checksums
└── README.md                   # Project documentation
```

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd splitwiseclone-backend
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Configure AWS credentials:**
   Ensure your AWS credentials are set up in your environment or through the AWS CLI.

4. **Run the application:**
   ```
   go run cmd/main.go
   ```

5. **API Documentation:**
   Refer to the API routes defined in `internal/api/routes.go` for available endpoints and their usage.

## Usage
- **User Registration:** POST `/users`
- **User Login:** POST `/auth/login`
- **Create Group:** POST `/groups`
- **Add Expense:** POST `/groups/{id}/expenses`
- **Get Group Balances:** GET `/groups/{id}/balances`

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.