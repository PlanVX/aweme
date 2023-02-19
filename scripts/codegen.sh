#!/usr/bin/env bash
# generate swagger documentation
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/main.go -p snakecase

# optional: generate mocks for testing
# mockery code generation
# for mocking interfaces to test logic
if [ -n "$MOCKERY" ]; then
  go install github.com/vektra/mockery/v2@v2.20.0
  go generate ./...
fi

# generate gorm queries
go mod tidy
go run scripts/db/gen.go
