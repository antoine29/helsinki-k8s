# Ping-Pong

## 1. Running as CLI (debug) mode

```shell
$ GO_PORT=8090 go run main.go
```

## 2. Running as a docker container
```shell
$ docker build . -t ping-pong
$ docker run -it --rm -p 80:80 -e "GO_PORT=80" ping-pong
```

## 3. Running as a K3D deployment
Build the image
```shell
$ docker build . -t ping-pong
```

Create the K3D registry and cluster
```shell
$ k3d registry create myregistry.localhost --port 12345
$ k3d cluster create mycluster -a 2 --registry-use k3d-myregistry.localhost:12345
```

Tag and push the docker image to the K3D registry
```shell
$ docker tag ping-pong k3d-myregistry.localhost:12345/ping-pong
$ docker push k3d-myregistry.localhost:12345/ping-pong
```

Create the deployment

* notice there is not an option to set env vars in the CLI way of `create deployment`
```shell
$ kubectl create deployment ping-pong --image=k3d-myregistry.localhost:12345/ping-pong
```

Get logs
```shell
$ kubectl get pods
$ kubectl logs -f ping-pong-76bd4c6485-2szjs
Error: 'GO_PORT' environment variable not set, using 8080 as defult.
Server started in port: 8080
```

Delete deployment
```shell
$ kubectl delete deployment ping-pong
```

## 4. Running as a single image into K3D
```shell
$ kubectl run ping-pong --image=k3d-myregistry.localhost:12345/ping-pong --env="GO_PORT=80"
$ kubectl logs -f ping-pong
Server started in port: 80
$ kubectl delete pod ping-pong
```

## PG DB mode
To save the ping-pong counts into a PG database the following env vars should be set:

```
GO_RUNMODE=db
PG_HOST=pg-svc
PG_PORT=5432
PG_DBNAME=postgres
PG_USER=postgres
PG_PASSWORD=postgres
PG_SCHEMA=pingpong
PG_TABLE=counts
```


## Health endpoint
Set `.env` file and then:

```shell
$ go run .
Reading .env file
Running in db mode
Listening on: 8090

$ curl http://localhost:8090/health
DB is ready
```
