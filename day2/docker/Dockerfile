# Step 1 :: Compile go -> binary file
FROM golang:1.19.2-alpine3.16 as step_1
WORKDIR /app
COPY go.* ./
# Download dependencies
RUN go mod tidy 
COPY . .
RUN go build -o demo-service main.go

# Step 2 :: Copy binary file from step 1
FROM alpine:3.16
WORKDIR /demo
RUN apk --no-cache add ca-certificates
COPY --from=step_1 /app/demo-service .
EXPOSE 1323
ENTRYPOINT [ "/demo/demo-service" ]