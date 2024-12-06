create-proto:
	protoc --proto_path=./3_1/ --proto_path=. --go_out=./pkg/api/order --go_opt=paths=source_relative --go-grpc_out=./pkg/api/order --go-grpc_opt=paths=source_relative --grpc-gateway_out=./pkg/api/order --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true ./3_1/order.proto
start-server:
	go run ./cmd/main/main.go
create-binary:
	go build -o main ./cmd/main/
# create-table:
# 	migrate -database postgres://Dahre:Daniil2007@localhost:8000/lyceum?sslmode=disable -path migrations up
