FROM golang:1.23-buster as deps

WORKDIR /app

RUN apt-get update && apt-get install -y build-essential git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM deps as builder
COPY . .
RUN go build -o imagine-backend .

FROM debian:buster as runner

ENV GIN_MODE=release

# Install Google Chrome and dependencies
RUN apt-get update && apt-get install -y wget gnupg
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list
RUN apt-get update && apt-get install -y google-chrome-stable

WORKDIR /app
COPY --from=builder /app/imagine-backend /app/imagine-backend

RUN chmod +x /app/imagine-backend
CMD ["/app/imagine-backend"]
