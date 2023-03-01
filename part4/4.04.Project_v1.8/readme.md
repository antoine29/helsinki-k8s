# Todos canary release  

## 0. Initial set-up

First install prometheus and grafana and open ports for its web interfaces. (Useful to test the prometheus queries and to visually check the pods cpu usage)

```shell
$ helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus

$ k get po -n prometheus
NAME                                                              READY   STATUS    RESTARTS       AGE
kube-prometheus-stack-1675-operator-78567977b-559wd               1/1     Running   0              4h10m
kube-prometheus-stack-1675116557-prometheus-node-exporter-rf5n4   1/1     Running   0              4h10m
kube-prometheus-stack-1675116557-prometheus-node-exporter-lhc62   1/1     Running   0              4h10m
kube-prometheus-stack-1675116557-kube-state-metrics-565fb95f6bs   1/1     Running   0              4h10m
kube-prometheus-stack-1675116557-prometheus-node-exporter-lh5vq   1/1     Running   0              4h10m
alertmanager-kube-prometheus-stack-1675-alertmanager-0            2/2     Running   2 (4h9m ago)   4h10m
prometheus-kube-prometheus-stack-1675-prometheus-0                2/2     Running   0              4h10m
kube-prometheus-stack-1675116557-grafana-976b6468b-jtd4s          3/3     Running   0              4h10m

$ kubectl -n prometheus port-forward prometheus-kube-prometheus-stack-1675-prometheus-0 9090:9090
Forwarding from 127.0.0.1:9090 -> 9090
Forwarding from [::1]:9090 -> 9090
Handling connection for 9090
Handling connection for 9090

$ kubectl -n prometheus port-forward kube-prometheus-stack-1675116557-grafana-976b6468b-jtd4s 3000:3000
```

Grafana credentials:
```
admin
prom-operator
```

Then install argo-rollouts to have a "Rollout" crd (similar to canary).

```
$ kubectl create namespace argo-rollouts
$ kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml
```

Prepare the images. We'll need multiple tags for backend image just for demonstration purposes. 


```shell
$ docker tag todos-be todos-be:V0-init                  
$ docker tag todos-be todos-be:V1-ok  
$ docker tag todos-be todos-be:V2-failing

$ docker tag todos-be:V0-init k3d-myregistry.localhost:12345/todos-be:V0-init
$ docker tag todos-be:V1-ok k3d-myregistry.localhost:12345/todos-be:V1-ok  
$ docker tag todos-be:V2-failing k3d-myregistry.localhost:12345/todos-be:V2-failing

$ docker push k3d-myregistry.localhost:12345/todos-be:V0-init 
$ docker push k3d-myregistry.localhost:12345/todos-be:V1-ok 
$ docker push k3d-myregistry.localhost:12345/todos-be:V2-failing
```

## 1. First run

- Set the `todos-ns-pods-cpu-usage-analysis-temp.yml` condition to pass
  ```
  ...
  successCondition: result < 3
  ...
  ```

- Set the `be-deploy.yml` be container image to 'V0-init'
  ```
  ...
  containers:
    - name: todos-be
      image: k3d-myregistry.localhost:12345/todos-be:V0-init
  ...
  ```

- Apply the manifests
  ```bash
  $ kubectl apply -f ns.yml
  $ kubectl apply -f todos-ns-pods-cpu-usage-analysis-temp.yml
  $ kubectl apply -f .
  ```

- Two pods/replicas should be created for 'todos-be', both should be running up with 'V0-init' image tag

  * `k` is an alias for `kubectl`

  ```bash
  $ k get po
  NAME                           READY   STATUS    RESTARTS         AGE
  pg-ss-0                        1/1     Running   0                16m
  todos-fe-57fd4dcf9d-j5lcq      1/1     Running   0                16m
  todos-rproxy-94f44f97c-sg2n9   1/1     Running   0                16m
  todos-be-6547b7d988-9c598      1/1     Running   0                16m
  todos-be-6547b7d988-5pd6w      1/1     Running   0                16m

  $ k describe pod todos-be-6547b7d988-9c598 | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V0-init
  $ k describe pod todos-be-6547b7d988-5pd6w | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V0-init
  ```

## 2. Applying a failing deployment:

- change the `todos-ns-pods-cpu-usage-analysis-temp.yml` condition to fail
  ```
  ...
  successCondition: result < 0.01 
  ...
  ```
 
  change the `be-deploy.yml` backend image tag to 'V2-failing' (just for observability purposes)
  ```
  ...
  containers:
    - name: todos-be
      image: k3d-myregistry.localhost:12345/todos-be:V2-failing
  ...
  ```

- apply the changes
  ```shell
  $ k apply -f todos-ns-pods-cpu-usage-analysis-temp.yml
  $ k apply -f be-deploy.yml
  ```

- A third pod should have been created using 'V2-failing' tag. At some point this new pod will replace one of the two replicas previously created (following the 50% deployment canary deployment strategy)
  ```shell 
  $ k get po
  NAME                           READY   STATUS    RESTARTS         AGE
  pg-ss-0                        1/1     Running   7                1d
  todos-fe-57fd4dcf9d-j5lcq      1/1     Running   7                1d
  todos-rproxy-94f44f97c-sg2n9   1/1     Running   14               1d
  todos-be-6547b7d988-9c598      1/1     Running   0                20m
  todos-be-6547b7d988-5pd6w      1/1     Running   0                20m
  todos-be-bfb5fccf6-8m77j       1/1     Running   0                14s

  $ k describe pod todos-be-bfb5fccf6-8m77j | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V2-failing
  ```

- After some time and since the analysisTemplate condition was set up to fail, the new 'V2-failing' pod will be replaced by a new pod using the 'V0-init' tag. Meaning that the deployment has failed and hence a rollback to the previous version was applied

  ```shell
  $ k get po
  NAME                           READY   STATUS    RESTARTS         AGE
  pg-ss-0                        1/1     Running   7                1d
  todos-fe-57fd4dcf9d-j5lcq      1/1     Running   7                1d
  todos-rproxy-94f44f97c-sg2n9   1/1     Running   14               1d
  todos-be-6547b7d988-5pd6w      1/1     Running   0                27m
  todos-be-6547b7d988-6gztx      1/1     Running   0                2m15s

  $ k describe pod todos-be-6547b7d988-6gztx  | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V0-init
  ```

## 3. Applying a passing deployment

- Change/apply the analysisTemplate condition to passing

- Change/apply the be deployment to use 'V1-ok' image tag

- After some time, both previously created pods will be replaced with new pods using 'V1-ok' image tag. Meaning the deployment was successful

  ```bash
  $ k get po
  NAME                           READY   STATUS    RESTARTS         AGE
  pg-ss-0                        1/1     Running   0                1d
  todos-fe-57fd4dcf9d-j5lcq      1/1     Running   7                1d
  todos-rproxy-94f44f97c-sg2n9   1/1     Running   14               1d
  todos-be-5cd4bd856c-n7mms      1/1     Running   0                8m13s
  todos-be-5cd4bd856c-pqswh      1/1     Running   0                3m

  $ k describe pod todos-be-5cd4bd856c-n7mms  | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V1-ok
 
  $ k describe pod todos-be-5cd4bd856c-pqswh | grep -i Image:
      Image:         k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16
      Image:          k3d-myregistry.localhost:12345/todos-be_migrations
      Image:          k3d-myregistry.localhost:12345/todos-be:V1-ok
  ```

