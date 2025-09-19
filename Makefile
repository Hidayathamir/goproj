run:
	go run main.go

swag-init:
	swag fmt && swag init --generalInfo ./cmd/server/server/swagger.go
