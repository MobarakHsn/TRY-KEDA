FROM golang:alpine AS builder
RUN mkdir /go-http-server
COPY . /go-http-server
WORKDIR /go-http-server
RUN go build -o go-http-server

FROM alpine
WORKDIR /go-http-server
COPY --from=builder /go-http-server/ /go-http-server/

# Command to run the executable
ENTRYPOINT ["./go-http-server"]
