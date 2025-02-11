FROM golang:1.23-alpine AS builder
WORKDIR /go/src/app

RUN apk update && apk upgrade && \
  apk --update add git gcc make libc-dev openssh linux-headers

COPY . .
RUN go mod tidy
RUN go build -o build/main cmd/*.go

FROM alpine:latest as release
WORKDIR /app

COPY --from=builder /go/src/app/build/main main

ENTRYPOINT [ "./main", "-env=production" ]
