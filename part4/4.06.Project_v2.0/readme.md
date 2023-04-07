# Project v2.0 Todos NATS queue

## 0. Initial set-up: Building images

6 images need to be build.

```
project/
--- be                      todos-be
--- be/dbMigrations         todos-be_migrations
--- fe                      todos-fe
--- rproxy                  todos-rproxy
--- todos-queue/consumer    todos-q-consumer
--- todos-queue/publisher   todos-q-publisher
```

## 1. Create the namespace, install nats, decrypt (or manually set) the secret and run the manifests

```shell
$ k apply -f ns.yml
$ helm install my-nats nats/nats
$ echo "decript (sops) the .secret.enc.yaml or create a secret.yaml manually"
$ k apply -f . 
```

![alt](demo.gif)
