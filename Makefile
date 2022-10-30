.PHONY: build
build:
	go build -v ./cmd/main
.PHONY: run
run:
	go run -v ./cmd/main.go
