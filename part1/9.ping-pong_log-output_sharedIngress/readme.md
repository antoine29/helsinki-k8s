# Running the log-output along ping-pong with a shared ingress (1.log-output, PingPongApp)

## 1. Delete previous deployments, services and ingress
Check the deployment names with `$ kubectl get deployments`, the services with `$ kubectl get svc,ing` and delete them by using their names. (or by using the manifests files: `$ kubectl delete -f ingress.yml service.yml deployment.yml`)

## 2. (If you don't have the images builded and pushed) Build, tag and push the images
Check the previous steps/folders

## 3. Run the manifests 
Run the deployment
```shell
$ kubectl apply -f .
```

## 5. test
```shell
$ curl localhost:8081/current
2022-08-15 18:51:22.583066944 +0000 UTC m=+95.006844993 elwvx
$ curl localhost:8081/ping-pong
1
```
