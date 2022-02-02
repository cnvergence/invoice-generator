
FROM golang:1.17.2 AS builder

WORKDIR /go/src/invoice-generator
COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/invoice-generator


FROM alpine:3.13
COPY --from=builder /go/bin/invoice-generator /usr/local/bin/invoice-generator
USER 1001