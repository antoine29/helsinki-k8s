# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o temporised-string

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/temporised-string /app
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/temporised-string"]
#CMD [""]
#CMD ["5"]
#CMD: is the command being executed by the container on startup. You can override this on 'docker run'
#ENTRYPOINT: means the container will run this command, (im not sure if you can override it, but thats the mean)
#shell form: CMD ping 8.8.8.8 => will append the command (ping 8.8.8.8) to the container's shell (possible problems with PIDs n in the case of shell-less container)
#exec form: CMD ["/bin/ping", "8.8.8.8"], especifies the bin



