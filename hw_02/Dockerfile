FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o hw_02 .

FROM alpine:latest

COPY --from=builder /app/hw_02 .

CMD ["./hw_02"]