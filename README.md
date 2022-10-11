# Advance Golang (day 1/2)

> online class @TDG with Somkiat Puisungnoen

## run /app

```bash
PORT=8080 go run app/main.go
```

or if use flag

```bash
go run app/main.go -port=8080
```

## run /another-app

```bash
go run another-app/main.go
```

## run e2e test for /app

```bash
go test --tags=e2e ./app/... -v
```

> use rest client for rest-client-test dir (VSCode extension)