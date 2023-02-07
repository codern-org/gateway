BINARY_NAME=codern-gateway

build:
	GOARCH=amd64 GOOS=linux go build -o dist/${BINARY_NAME} cmd/main.go

clean:
	go clean
	rm -rf dist/${BINARY_NAME}

deps:
	go mod download
