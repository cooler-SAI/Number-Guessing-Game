FROM golang:latest
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM ubuntu:latest

LABEL authors="cooler-SAI"

WORKDIR /root/

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
