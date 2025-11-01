# **`ARTIFACTS_DIR` is an environment variable provided by the AWS SAM CLI that tells your Makefile where to place the compiled build artifacts
build-HelloWorldFunction:
	env GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap cmd/function/main.go

build-HttpHelloWorldFunction:
	env GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap cmd/http/main.go

# NOTE:
# -o $(ARTIFACTS_DIR)/bootstrap is location where the compiled binary will be placed.
# The name "bootstrap" is required for custom runtime Lambda functions.
# The cmd/function/main.go and cmd/http/main.go are the entry points for the respective Lambda functions.

.PHONY: build-HelloWorldFunction
.PHONY: build-HttpHelloWorldFunction
