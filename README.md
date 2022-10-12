# Advance Golang (day 1/2)

> online class @TDG with Somkiat Puisungnoen

## Day1

### run /app

```bash
PORT=8080 go run day1/app/main.go
```

or if use flag

```bash
go run day1/app/main.go -port=8080
```

### run /another-app

```bash
go run day1/another-app/main.go
```

### run e2e test for /app

```bash
go test --tags=e2e .day1//app/... -v
```

> use rest client for rest-client-test dir (VSCode extension)

## Day2

> ...
