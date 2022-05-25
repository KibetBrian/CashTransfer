#Builds binary file
FROM  golang:1.17.8-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Runs the binary file from the builder stage
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /app/main /app/  

EXPOSE 8080
CMD ["/app/main"]