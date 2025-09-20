start:
	go run ./cmd/gateway/main.go 

build:
	docker build -t api-gateway .

up:
	docker run --rm -p 8080:8080 api-gateway
