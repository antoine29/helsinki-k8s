# Running the Todo_s backend 

## Running locally
```shell
$ export PG_HOST=localhost
$ export PG_PORT=5432
$ export PG_USER=postgres
$ export PG_PASSWORD=postgres
$ export PG_DBNAME=postgres
$ export PG_SCHEMA=todo
$ GO_PORT=8080 go run main.go
```

## 1. Running as a docker container
```shell
$ docker build . -t todos_be
$ docker run -it --rm \
  -e "GO_PORT=8080" \
  -e "PG_HOST=localhost" \
  -e "PG_PORT=5432" \
  -e "PG_USER=postgres" \
  -e "PG_PASSWORD=postgres" \
  -e "PG_DBNAME=postgres" \
  -e "PG_SCHEMA=todo" \
  -p 8080:8080 todos_be
```

## 2. To build and push the image to a local K3D image registry
Assuming you already have the K3D registry, build, tag and push the image. 

```shell
$ docker build . -t todos_be
$ docker tag todos_be k3d-myregistry.localhost:12345/todos_be
$ docker push k3d-myregistry.localhost:12345/todos_be
```

