# Go Serverless CRUD with MongoDB

This is a sample serverless application in Go that demonstrates CRUD operations with MongoDB. It uses the AWS Serverless Application Model (SAM) for deployment.

## Prerequisites

*   Go (version 1.18 or later)
*   Docker
*   AWS SAM CLI

## Build

To build the application, run the following command:

```bash
sam build
```

This command will build the Go binaries and package them for deployment.

## Run Locally

To run the application locally, you can use the SAM CLI to start a local API Gateway:

```bash
sam local start-api
```


or

you can invoke used this

```bash
sam local invoke HelloWorldFunction --event custom-event.json
```
This will start a local server at `http://127.0.0.1:3000`. You can then access the `/hello` endpoint to invoke the `HelloWorldFunction`.

### Invoke the function

To invoke the function, you can use `curl`:

```bash
curl http://127.0.0.1:3000/hello
```
