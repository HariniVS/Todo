FROM golang:latest
WORKDIR /go/src/todo
COPY . .
RUN go build

CMD ["./todo"]
