# PingPong and LogOutput readiness probes

1. Build, tag and push the images

```shell
part1/PingPongApp $ docker build . -t pingpong
part1/PingPongApp/dbMigrations $ docker build . -t pingpong-dbmigrations
part1/1.log-output $ docker build . -t logoutut

$ docker tag pingpong k3d-myregistry.localhost:12345/pingpong
$ docker tag pingpong-dbmigrations k3d-myregistry.localhost:12345/pingpong-dbmigrations
$ docker tag logoutput k3d-myregistry.localhost:12345/logoutput
$ docker tag postgres:14.3-alpine3.16 k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16 

$ docker push k3d-myregistry.localhost:12345/pingpong
$ docker push k3d-myregistry.localhost:12345/pingpong-dbmigrations
$ docker push k3d-myregistry.localhost:12345/logoutput
$ docker push k3d-myregistry.localhost:12345/postgres:14.3-alpine3.16 



k apply -f ns.yml
k aplpy -f secret.yml
k aplpy pg_ss.yml
k aplpy pg_ss-svc.yml
k apply -f pingpong-deploy.yml
k apply -f pingpong-svc.yml
k apply -f logoutput-deploy.yml
k apply -f logoutput-svc.yml
k apply -f ing.yml

# assuming cluster has LB on 8081 port, ask ping pong status with 
$ curl localhost:8081/status/http
```



