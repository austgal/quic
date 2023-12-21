FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o broker

EXPOSE 6666/udp
EXPOSE 6667/udp

CMD ["./broker"]
