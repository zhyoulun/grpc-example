gen_grpc:
	protoc --go_out=./proto/sdk --go_opt=paths=source_relative \
		--go-grpc_out=./proto/sdk --go-grpc_opt=paths=source_relative \
		./proto/hello.proto

gen_grpc_gateway:
	protoc --grpc-gateway_out ./proto/sdk \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./proto/hello.proto
