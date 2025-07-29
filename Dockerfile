FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]
