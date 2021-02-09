FROM golang:1.12.0-alpine3.9
RUN mkdir /app
WORKDIR /app
ADD . .
RUN go build -o main .
CMD ["/app/main"]