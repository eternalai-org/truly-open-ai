FROM golang:1.23-alpine3.19 as deps

WORKDIR /app

RUN apk add --no-cache build-base git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM deps as builder
COPY . .
RUN go build -o imagine-backend .

FROM alpine:3.19 as runner

ENV GIN_MODE=release

WORKDIR /app
COPY --from=builder /app/imagine-backend /app/imagine-backend

RUN chmod +x /app/imagine-backend
CMD ["/app/imagine-backend"]
