FROM golang:1.15 AS builder

ENV CGO_ENABLED=0
ADD . /app
WORKDIR /app
RUN go build .

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/jwt-decode /app/jwt-decode
USER nonroot:nonroot

ENTRYPOINT ["/app/jwt-decode"]
