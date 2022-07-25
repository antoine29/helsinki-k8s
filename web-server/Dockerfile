# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o web-server

# for some reason is not workig with the distroless image, anyways, using alpine we get an even littler image size (10 vs 20 mb)
#FROM gcr.io/distroless/base-debian10:debug
FROM alpine
WORKDIR /app
COPY --from=builder /app/web-server /app
#EXPOSE 8080
#USER nonroot:nonroot
#
ENTRYPOINT ["/app/web-server"]
#CMD [""]
#CMD ["5"]
#CMD: is the command being executed by the container on startup. You can override this on 'docker run'
#ENTRYPOINT: means the container will run this command, (im not sure if you can override it, but thats the mean)
#shell form: CMD ping 8.8.8.8 => will append the command (ping 8.8.8.8) to the container's shell (possible problems with PIDs n in the case of shell-less container)
#exec form: CMD ["/bin/ping", "8.8.8.8"], especifies the bin


#docker build . -t web-server
#docker run -it --rm -e "GO_PORT=80" web-server