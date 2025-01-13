FROM golang:1.23.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o decentralized-inference .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/decentralized-inference /app/config.json /app/

CMD ["./decentralized-inference", "server"]


