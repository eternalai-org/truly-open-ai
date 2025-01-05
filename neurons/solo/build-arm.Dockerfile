FROM --platform=linux/arm64 golang:1.21-alpine3.17 AS build

RUN apk update && apk add gcc musl-dev gcompat libc-dev linux-headers
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -ldflags "-linkmode external -extldflags -static" -o workersv

RUN chmod +x /app/workersv

FROM scratch AS export-stage

COPY --from=build /app/workersv .
