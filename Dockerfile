# Use the official Golang image as the base image
FROM golang:1.24.4


WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY *.go ./

RUN go build -o main .


EXPOSE 8080

CMD ["./main"]