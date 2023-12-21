FROM golang:alpine3.18

WORKDIR /bin

COPY . .

RUN go build -o cmd/app/main ./cmd/app/main.go

WORKDIR /bin/cmd/app

CMD ["./main"]