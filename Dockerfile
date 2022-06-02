FROM golang:1.16

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o demo .

CMD ["./demo"]

