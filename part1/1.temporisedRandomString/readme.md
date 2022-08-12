# Running the temporised random string project as k3d deployment with a local k3d registry

required program parameters:
- -secsInterval
- -strLength
- -serverPort

`GET: host:{serverPort}/current` returns the current stamp/random str

1. Build and test the project docker image
```shell
$ docker build . -t temporised-random-string
$ docker run temporised-random-string 3
random string: usgwbdo interval to print:  3s
Tick:  2022-07-27 01:37:36.165060436 +0000 UTC m=+3.001884546 usgwbdo
Tick:  2022-07-27 01:37:39.16509235 +0000 UTC m=+6.001916430 usgwbdo
Tick:  2022-07-27 01:37:42.16504319 +0000 UTC m=+9.001867236 usgwbdo
...
```

2. Create a local k3d registry on port 12345
```shell
$ k3d registry create myregistry.localhost --port 12345
```

3. Create the k3d cluster with 2 agents and linking it to the created k3d registry
```shell
$ k3d cluster create mycluster -a 2 --registry-use k3d-myregistry.localhost:12345
```

4. Tag and push the created docker images
```shell
$ docker tag temporised-random-string k3d-myregistry.localhost:12345/temporised-random-string
$ docker push k3d-myregistry.localhost:12345/temporised-random-string
```
5. Create the deployment
```shell
$ kubectl create deployment temporiser --image=k3d-myregistry.localhost:12345/temporised-random-string -- /app/exe 5
```

ToDo: we're passing the '/app/exe 5' param to override the entire image entrypoint and cmd, in the CLI mode of `create deployment there is no way to pass a cmd without a entrypoint`
https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-deployment-em-