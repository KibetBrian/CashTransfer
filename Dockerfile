FROM  golang:1.17.8-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /app/main /app/
COPY .env.example /app/

EXPOSE 8080
CMD ["/app/main"]