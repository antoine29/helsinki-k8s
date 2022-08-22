# Running the temporised random string project as k3d deployment with a local k3d registry
The program can run in two modes:

- reader: exposes a port to query the status\
```shell
$ go run . -serverPort 8090 -reader
```

testing in other shell
```shell
$ curl localhost:8080/status/file (to check the status stored on file)
$ curl localhost:8080/status/memory (to check the status stored on memory)
```

- writer: to write the status (time stamp+randomStr) to the file and to memory
```shell
$ go run . -strLength 5 -secsInterval 5 -writer
```

- aditionally `mode` param can be omited to run in writer/reader mode
```shell
$ go run . -serverPort 8090 -strLength 5 -secsInterval 5
```

1. Build and test the project docker image
```shell
$ docker build . -t log-output
$ docker run -it --rm --name log-output -p 8090:8090 log-output -serverPort 8090 -strLength 5 -secsInterval 5
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
$ docker tag log-output k3d-myregistry.localhost:12345/log-output
$ docker push k3d-myregistry.localhost:12345/log-output
```
5. Create the deployment
```shell
$ kubectl create deployment log-output --image=k3d-myregistry.localhost:12345/log-output -- /app/exe 5
```

ToDo: we're passing the '/app/exe 5' param to override the entire image entrypoint and cmd, in the CLI mode of `create deployment there is no way to pass a cmd without a entrypoint`
https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-deployment-em-