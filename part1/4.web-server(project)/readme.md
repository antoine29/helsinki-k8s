# Running the web-server 

## 1. Running as a docker container
```shell
$ docker build . -t web-server
$ docker run -it --rm -e "GO_PORT=80" web-server
```

## 2. Running as a K3D deployment
Assuming you already have the K3D cluster, registry and the images.

Build, tag and push the image. 
```shell
$ docker build . -t wserver
$ docker tag wserver k3d-myregistry.localhost:12345/wserver
$ docker push k3d-myregistry.localhost:12345/wserver
```
Run the deployment
```shell
$ kubectl apply -f manifests/deployment.yaml
```
