# **`ARTIFACTS_DIR` is an environment variable provided by the AWS SAM CLI that tells your Makefile where to place the compiled build artifacts
build-HelloWorldFunction:
	env GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap cmd/function/main.go

build-HttpHelloWorldFunction:
	env GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap cmd/http/main.go

.PHONY: build-HelloWorldFunction
.PHONY: build-HttpHelloWorldFunction
