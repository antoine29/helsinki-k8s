# Running log-output along ping-pong (communication between pods)
log-output pod is making a request to ping-pong pod (which is returning the ping/pong count)

## 1. Delete previous deployments, services and ingress
Check the deployment names with `$ kubectl get deployments`, the services with `$ kubectl get svc,ing` and delete them by using their names. (or by using the manifests files: `$ kubectl delete -f ingress.yml service.yml deployment.yml`)

## 2. (If you don't have the images builded and pushed) Build, tag and push the images
Check the previous steps/folders

## 3. Run the manifests 
Run the `namespace.yml` first
```shell
$ kubectl apply -f namespace.yml
$ kubectl apply -f .
```

## 4. test
Assuming you have your K3D LB on port 8081:

```shell
$ curl localhost:8081/status/http
62c7beb524773830a2939cbd71906c23
time: 2022-09-01 20:46:54.810321 +0000 UTC m=+422.393819896
Ping / Pongs: 5
```
