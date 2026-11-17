FROM golang:1.27-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /api .

FROM alpine:3.24
COPY --from=builder /api /api
USER 65534:65534
EXPOSE 8080
ENTRYPOINT ["/api"]
