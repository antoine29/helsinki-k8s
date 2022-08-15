# Running the wev-server application (4.web-server)

## 1. Delete the previous deployments
Check the deployment names with `$ kubectl get deployments` then delete them

## 2. (If you don't have the images builded and pushed) Build, tag and push the image. 
```shell
1.web-server $ docker build . -t web-server
$ docker tag web-server k3d-myregistry.localhost:12345/web-server
$ docker push k3d-myregistry.localhost:12345/web-server
```

## 3. Run the manifests 
Run the deployment
```shell
$ kubectl apply -f deployment.yml
$ kubectl apply -f service.yml
$ kubectl apply -f ingress.yml
```

## 5. test
```shell
$ curl localhost:8081/api/health
```
