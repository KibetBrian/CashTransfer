FROM alpine:1.17-alpine3.15
WORKDIR /app
COPY . /app/
RUN go build -o main main.go

EXPOSE 8080
CMD ["/app/main"]