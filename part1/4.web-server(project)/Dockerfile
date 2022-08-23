# syntax=docker/dockerfile:1
FROM node:16.14.0-alpine as fe-builder 
WORKDIR /app
COPY ./todo-fe .
RUN yarn install && yarn run build

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
COPY --from=fe-builder /app/dist /app/todo-fe/dist

ENTRYPOINT ["/app/web-server"]
