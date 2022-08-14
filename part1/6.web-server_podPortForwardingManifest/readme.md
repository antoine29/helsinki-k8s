# Running the web-server (4.web-server application) with pod port forwarding as a manifest

## 1. Delete the previous web-server deployment
```shell
5.web-server_podPortForwarding $ kubectl delete -f deployment.yml        
deployment.apps "web-server" deleted
```
or just the deployment name by looking at the deployment manifest, or by listing deployments with `$ kubectl get deploymments` and delete it with `$ kubectl delete deployment web-server` (`web-server is the deployment name`) 

## 2. Delete the K3D cluster
We need to recreate the cluster opening some ports for agents. Assuming your cluster is named as `mycluster` (check it with `$ k3d cluster list`)
```shell
$ k3d cluster stop mycluster
$ k3d cluster delete mycluster
```

Assuming you're still using the local k3d registry
```shell
$ k3d cluster create mycluster --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2 --registry-use k3d-myregistry.localhost:12345
```

## 3. (If you don't have the images builded and pushed) Build, tag and push the image. 
```shell
$ docker build . -t web-server
$ docker tag web-server k3d-myregistry.localhost:12345/web-server
$ docker push k3d-myregistry.localhost:12345/web-server
```

## 4. Run the manifests 
Run the deployment
```shell
$ kubectl apply -f deployment.yml
$ kubectl apply -f service.yml
```

## 5. test
```shell
$ curl localhost:8082/api/health
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
