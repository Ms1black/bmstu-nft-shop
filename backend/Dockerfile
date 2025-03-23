FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o nft_backend main.go
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/nft_backend .
COPY certs/cert.pem /root/cert.pem
COPY certs/key.pem /root/key.pem
EXPOSE 443

CMD ["./nft_backend"]