# Todos liveness test

To set up the project images into a local k3d cluster:

```bash
$ docker build . -t todos-be
$ docker build . -t todos-be_migrations
$ docker build . -t todos-fe --build-arg VITE_API_URL=http://localhost:8081
$ docker build . -t todos-rproxy

$ docker tag todos-be_migrations k3d-myregistry.localhost:12345/todos-be_migrations
$ docker tag todos-be k3d-myregistry.localhost:12345/todos-be           
$ docker tag todos-fe k3d-myregistry.localhost:12345/todos-fe
$ docker tag todos-rproxy k3d-myregistry.localhost:12345/todos-rproxy
$ docker tag postgres:14.3-alpine3.16 k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16 

$ docker push k3d-myregistry.localhost:12345/todos-be_migrations 
$ docker push k3d-myregistry.localhost:12345/todos-be
$ docker push k3d-myregistry.localhost:12345/todos-fe
$ docker push k3d-myregistry.localhost:12345/todos-rproxy
$ docker push k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16 

$ kubectl apply -f ns.yml
$ kubectl apply -f .
```

At this point the application will be running ok. If you change the PG password provided in the cmap.yml file, the readinessProbe will start failing after some seconds (check be-deploy logs).
