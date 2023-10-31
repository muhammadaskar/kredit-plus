FROM golang:latest

WORKDIR /app

COPY .env.example .env

COPY . .

RUN go get -d -v ./...

RUN go test ./app/transaction/delivery/http -v

RUN go build -o main .

CMD ["./main"]
