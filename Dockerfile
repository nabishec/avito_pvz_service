FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go ./cmd/grpc.go ./cmd/http.go

EXPOSE 8080 3000

CMD [ "./main" ]