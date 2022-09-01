# syntax=docker/dockerfile:1
FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go get && \
  go install github.com/swaggo/swag/cmd/swag && \
  go build -o web-server

# for some reason is not workig with the distroless image, anyways, using alpine we get an even littler image size (10 vs 20 mb)
#FROM gcr.io/distroless/base-debian10:debug
FROM alpine
WORKDIR /app
COPY --from=builder /app/web-server /app

ENTRYPOINT ["/app/web-server"]
