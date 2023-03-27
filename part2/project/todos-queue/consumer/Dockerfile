# syntax=docker/dockerfile:1
FROM golang:1.17-alpine as builder
WORKDIR /app
COPY ./consumer/ /app/consumer/
COPY ./telegram-client/ /app/telegram-client/
WORKDIR /app/consumer/cmd
RUN go get && go build -o exe

# for some reason is not workig with the distroless image, anyways, using alpine we get an even littler image size (10 vs 20 mb)
#FROM gcr.io/distroless/base-debian10:debug
FROM alpine
WORKDIR /app
COPY --from=builder /app/consumer/cmd/exe /app

ENTRYPOINT ["/app/exe"]

# build this image from parent folder:
# $ docker build . -f ./consumer/Dockerfile -t todos-q-consumer

