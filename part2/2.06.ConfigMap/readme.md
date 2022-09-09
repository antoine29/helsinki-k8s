# Running log-output along ping-pong (communication between pods)
- log-output pod is making a request to ping-pong pod (which is returning the ping/pong count)
- `confimap.yml` is used to set env vars into the pod

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
Hello
81a6bfc0811d146265f4d27dfb2791d4
time: 2022-09-09 00:59:27.988490644 +0000 UTC m=+17991.240595868
Ping / Pongs: 42
```
