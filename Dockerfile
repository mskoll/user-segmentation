FROM golang:1.20.1-alpine3.17

WORKDIR /src/app

COPY . .
RUN go mod tidy
RUN go build -o app ./cmd/main.go

CMD ["./app"]


