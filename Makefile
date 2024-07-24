run:
	@go run main.go

dev:
	@air

service:
	@docker-compose up -d

service_stop:
	@docker-compose down

test:
	@go test ./... -v

.PHONY: run air service service_stop test
