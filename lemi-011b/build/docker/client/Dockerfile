FROM arm64v8/golang:1.17.2-alpine AS builder

WORKDIR /build
COPY go.mod .
COPY go.sum .

COPY pkg pkg/
COPY cmd/client .

RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
    -ldflags="-w -s" -o app .

FROM scratch

ENV LEMI011B_CLIENT_REST_URL="http://localhost:8081"
ENV LEMI011B_CLIENT_SERIAL_PORT="/dev/ttyUSB0"
ENV LEMI011B_CLIENT_SERIAL_BAUD=19200

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]
