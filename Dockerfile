FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o bin/main cmd/main.go

EXPOSE 8080

CMD ["bin/main"]