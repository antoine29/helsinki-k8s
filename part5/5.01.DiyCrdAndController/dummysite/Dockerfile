# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o dummysite

FROM alpine
WORKDIR /app
COPY --from=builder /app/dummysite /app

ENTRYPOINT ["/app/dummysite"]

# docker run -e PORT=8083 -e URL=https://example.com/ -p 8083:8083 dummysite
