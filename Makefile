go-test:
	go test ./...

format:
	gofumpt -l -w .

golint:
	golangci-lint -c ../.golangci.yaml run ./...

gomod:
	go mod tidy
	go mod vendor

install-hooks:
	cp githooks/* .git/hooks

.FORCE:
	go test ./...