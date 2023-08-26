.PHONY: run
run: fmt
	go run cmd/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: fmt
	go test ./... -v -cover

.PHONY: build
build: test
	go build -o main cmd/main.go