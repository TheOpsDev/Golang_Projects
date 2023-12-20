SHELL=/bin/bash
APP_NAME := top-ten

build:
	echo "Staring build..."
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -o ${APP_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${APP_NAME}-linux main.go
	echo "Build completed!"
test:
	echo "Initiating tests..."
	go test
	echo "Tests completed!"
run: build test
	./generate_sample.sh -n 10000
	go run main.go -filename ./tmp_sample.txt
clean:
	rm -rf ./tmp_sample.txt *-darwin *-linux
	go clean
all: build test run clean
	echo "All tasks completed!"
