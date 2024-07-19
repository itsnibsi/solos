.PHONY: run-dev lint-backend lint-frontend test-backend generate-swagger build-dev

run-dev:
	docker-compose -f docker-compose.dev.yml up --build

lint-backend:
	golangci-lint run ./...

lint-frontend:
	cd client && npm run lint

test-backend:
	go test ./...

generate-swagger:
	swag init -g cmd/api/main.go -o ./docs --parseInternal

build-dev:
	docker-compose -f docker-compose.dev.yml build