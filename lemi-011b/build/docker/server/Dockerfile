FROM arm64v8/golang:1.17.2-alpine AS builder

WORKDIR /build
COPY go.mod .
COPY go.sum .

COPY pkg pkg/
COPY cmd/server .

RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
    -ldflags="-w -s" -o app .

FROM scratch

ENV LEMI011B_SERVER_TIMESCALEDB_URL="postgres://postgres:password@10.160.11.142:5432/lemi011b"
ENV LEMI011B_SERVER_REST_PORT=8080

EXPOSE 8080

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]
