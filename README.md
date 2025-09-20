For swagger:

go install github.com/swaggo/swag/cmd/swag@latest

export PATH=$PATH:$(go env GOPATH)/bin

swag init -g cmd/main.go