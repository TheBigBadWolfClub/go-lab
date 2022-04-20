.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down:
	docker-compose -f docker-compose.yaml down
	docker system prune 

coverage:
	go test -v -short -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

golint:
	golangci-lint -c .golangci.yaml run ./...

go-test:
	go test -v -short ./...

go-test-integration:
	go test -v  ./...