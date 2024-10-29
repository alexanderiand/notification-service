.SILENT:

.PHONY: fmt lint test race build run

include local.env
export

fmt:
	go fmt ./...

lint: fmt # golangci-lint, but for KISS, choise other linter
	go vet ./...

test: lint
	go test -v -cover ./...

race: test
	go test -v -race ./...

rm_bin:
	rm ./notification_service

build_and_run: test
	go build -o notification_service cmd/notification_service/main.go && ./notification_service

run: race
	go run cmd/notification_service/main.go

.DEFAULT_GOAL := run