build-rest:
	@echo " >> building rest binaries"
	@go build -v -o productservice-rest cmd/rest/*.go

build-grpc:
	@echo " >> building grpc binaries"
	@go build -v -o productservice-grpc cmd/grpc/*.go

run-rest: build-rest
	@./productservice-rest

run-grpc: build-grpc
	@./productservice-grpc