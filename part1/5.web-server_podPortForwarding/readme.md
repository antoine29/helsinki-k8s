# Running the web-server (4.web-server application) with pod port forwarding

## 1. Running as a K3D deployment
Assuming you already have the K3D cluster and the its registry.

Build, tag and push the image. 
```shell
$ docker build . -t web-server
$ docker tag web-server k3d-myregistry.localhost:12345/web-server
$ docker push k3d-myregistry.localhost:12345/web-server
```
Run the deployment
```shell
$ kubectl apply -f deployment.yaml
```

Get the deployment pod
```shell
$ kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
web-server-64cd79d554-f7tsl   1/1     Running   0          8m22s
```

Pod port forwarding
```shell
$ kubectl port-forward web-server-64cd79d554-f7tsl 8090:8090
```

Check going to localhost:8090/swagger/indexhtml or by curl
```shell
$ curl localhost:8090/api/health
```





