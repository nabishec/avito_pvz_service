FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080 3000 9090

CMD [ "./main" ]