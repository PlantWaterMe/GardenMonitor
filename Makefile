setup:
	go mod download

run:
	go run main.go

run-api:
	go run cmd/api/main.go

run-queue-consumer:
	go run cmd/queue-consumer/main.go

run-garden-monitor:
	go run cmd/garden-monitor/main.go