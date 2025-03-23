# FROM golang:1.21 AS builder
# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build 0o nft_backend main.go
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/nft_backend .
# EXPOSE 8080

# CMD ["./nft_backend"]