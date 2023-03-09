# Publisher

This golang project exposes an http endpoint and publishes the payloads being sent to a nats queue.

```shell
$ export SUBJECT=foo
$ go run .

$ curl -X 'POST' \
  'http://localhost:8080/api/publish' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "id",
    "todo": "todo", "status": "status"
}'

```
