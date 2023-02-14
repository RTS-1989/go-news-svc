FROM golang:alpine as builder

WORKDIR /GoNewsSvc

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine

WORKDIR /GoNewsSvc

COPY --from=builder /GoNewsSvc/main /GoNewsSvc/main
COPY --from=builder /GoNewsSvc/pkg/config/envs/*.env /GoNewsSvc/

RUN chmod +x /GoNewsSvc/main

CMD ["./main"]
