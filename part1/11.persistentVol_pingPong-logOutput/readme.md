# ping-pong and log-output applications with node shared volume
The ping-pong exposes an endpoint (actually any /) to increase a hit counter and return it along the current time-stamp. Aditionally, this response is being saved in `/tmp/status` file.
Log-output is run in reader mode, so on a `/staus/file` call is returning the content of the `/tmp/status` file.
Finally, `/tmp` is being shared as a persistent volume (persisting in the node) across the containers.

## 1. Delete the previous deployments, services and ingress
Delete by using `kubectl delete -f {manifiests}` or check the deployment names with `$ kubectl get deployments` and services with `ubectl get svc,ing` then delete them

## 2. (If you don't have the images builded and pushed) Build, tag and push the images. 

## 3. Run the manifests
Run the deployment
```shell
$ kubectl apply -f .
```

## 5. test
Assuming you have the k3d cluster set with a load balancer in port 8081 
```shell
$ curl localhost:8081/ping 
time: 2022-08-23 21:20:14.987956738 +0000 UTC m=+113.349354995 	 counter: 1

$ curl localhost:8081/status/file
MD5 file hash: 7ae9b2b6b87e1e926f2d17e8f7e17523 
time: 2022-08-23 21:20:18.168047804 +0000 UTC m=+116.529446143 	 counter: 3
```
