# log-output application with simple shared volume (1.log-output)

## 1. Delete the previous deployments
Check the deployment names with `$ kubectl get deployments` then delete them

## 2. (If you don't have the images builded and pushed) Build, tag and push the image. 
```shell
1.log-output $ docker build . -t log-output
$ docker tag log-output k3d-myregistry.localhost:12345/log-output
$ docker push k3d-myregistry.localhost:12345/log-output
```

## 3. Run the manifests 
Run the deployment
```shell
$ kubectl apply -f .
```

## 5. test
```shell
$ curl localhost:8081/status/file
```
