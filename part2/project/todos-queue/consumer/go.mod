module github.com/antoine29/todos-queue-consumer

go 1.19

require (
	github.com/antoine29/todos-queue-telegram-client v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/nats-io/nats.go v1.24.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/nats-io/nats-server/v2 v2.9.15 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/antoine29/todos-queue-telegram-client => ../telegram-client
