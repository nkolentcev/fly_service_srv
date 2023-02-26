FROM golang:1.20.1-alpine3.17
RUN mkdir /app
ADD ./cmd /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go build -o /cmd/main .
CMD ["/app/main"]