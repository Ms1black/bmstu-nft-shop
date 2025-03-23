FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o nft_backend main.go
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/nft_backend .
EXPOSE 8080

CMD ["./nft_backend"]