FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /go/bin/app

FROM haproxy:alpine
COPY --from=builder /go/bin/app /app

CMD /app && haproxy -f haproxy.cfg

