FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -v -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /crud_app

COPY --from=builder /app ./

CMD ["./app"]
