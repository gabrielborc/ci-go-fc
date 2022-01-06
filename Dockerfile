FROM golang:latest

WORKDIR /app

COPY . . 

RUN go build math.go

CMD ["./math"]