# Running the web-server 

## 1. Running as CLI (debug) mode

```shell
$ GO_PORT=8090 go run main.go
```

## 2. Running as a docker container
```shell
$ docker build . -t web-server
$ docker run -it --rm -e "GO_PORT=80" web-server
```

## 3. Running as a K3D deployment
Build the image
```shell
$ docker build . -t web-server
```

Create the K3D registry and cluster
```shell
$ k3d registry create myregistry.localhost --port 12345
$ k3d cluster create mycluster -a 2 --registry-use k3d-myregistry.localhost:12345
```

Tag and push the docker image to the K3D registry
```shell
$ docker tag web-server k3d-myregistry.localhost:12345/web-server
$ docker push k3d-myregistry.localhost:12345/web-server
```

Create the deployment

* notice there is not an option to set env vars in the CLI way of `create deployment`
```shell
$ kubectl create deployment web-server --image=k3d-myregistry.localhost:12345/web-server
```

Get logs
```shell
$ kubectl get pods
$ kubectl logs -f web-server-76bd4c6485-2szjs
Error: 'GO_PORT' environment variable not set, using 8080 as defult.
Server started in port: 8080
```

Delete deployment
```shell
$ kubectl delete deployment web-server
```

## 4. Running as a single image into K3D
```shell
$ kubectl run web-server --image=k3d-myregistry.localhost:12345/web-server --env="GO_PORT=80"
$ kubectl logs -f web-server
Server started in port: 80
$ kubectl delete pod web-server
```