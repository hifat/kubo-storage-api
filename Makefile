run-grpc:
	go run ./cmd/r2 -envPath=./env/.env

pb-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./proto/storage/storage.proto

dck-build:
	docker build -t kubo-storage-api:latest .

dck-run:
	docker run -p 4000:4000 -p 9000:9000 \
  -v /path/to/.env:/app/env/.env \
  kubo-storage-api:latest