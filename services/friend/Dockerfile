FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/friend

FROM alpine:3.10
COPY --from=builder /go/bin/friend /go/bin/friend
EXPOSE 8080
ENTRYPOINT ["/go/bin/friend"]