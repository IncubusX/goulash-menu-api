.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print $$3}'

.PHONY: build
build:
	go mod download && go build -o ./bin/goulash-menu-api$(shell go env GOEXE) ./cmd/api/main.go
