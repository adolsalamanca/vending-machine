FROM golang:1.13.4-alpine as builder

ENV CGO_ENABLED=0
ENV GO111MODULE="on"

WORKDIR /go/src
COPY . .

RUN go test ./... && go install ./...

FROM alpine as runner

COPY --from=builder /go/bin/* /app/

WORKDIR /app
CMD ["./vending-machine"]
