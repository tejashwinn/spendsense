# Backend

For swagger:

    go install github.com/swaggo/swag/cmd/swag@latest
    
    export PATH=$PATH:$(go env GOPATH)/bin

    swag init -g main.go


To Run:

    go run main.go

To Build:
    
    go build -o bin/sependsense main.go

# Frontend

Dev: 
    pnpm dev

Generate API Client:
    // Picks the config from openapi-ts.config.ts
    pnpm dlx @hey-api/openapi-ts