#Builds binary file
FROM  golang:1.17.8-alpine AS builder
WORKDIR /app
COPY . /app/
RUN go build -o main main.go

#Copy binary file from builder stage
FROM golang:1.17.8-alpine3.15
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]