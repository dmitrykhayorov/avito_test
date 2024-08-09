FROM golang:1.22 AS builder
WORKDIR /avito_test
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service ./cmd/app/main.go
FROM golang:1.22
WORKDIR /avito_test
COPY --from=builder ./avito_test/service .
EXPOSE 8080
CMD ["./service"]
