run-grpc:
	go run ./cmd/r2 -envPath=./env/.env

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./proto/storage/storage.proto