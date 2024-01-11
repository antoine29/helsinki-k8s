# Ping Pong serverless

## creates a cluster without traefik ingress   

```shell   
$ k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2 --registry-use k3d-myregistry.localhost:12345 --k3s-arg "--disable=traefik@server:0"
```

## installs knative   

```shell   
$ kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-crds.yaml

$ kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-core.yaml

```

## installs and sets contour as knative ingress   

```shell   
$ kubectl apply -f https://github.com/knative/net-contour/releases/download/knative-v1.0.0/contour.yaml \
-f https://github.com/knative/net-contour/releases/download/knative-v1.0.0/net-contour.yaml

$ kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress-class":"contour.ingress.networking.knative.dev"}}'
```

## make a serverless call

this step assumes that `ghcr.io/antoine29/ping-pong:latest` is an image available through internet. Since setting up a local registry for knative is not a straightforward process   

```shell   
$ kubectl apply -f ping-pong-ksvc.yaml
$ curl -H "Host: ping-pong.default.example.com" http://localhost:8081
```

