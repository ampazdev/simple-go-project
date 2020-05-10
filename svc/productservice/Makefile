build-rest:
	@echo " >> building rest binaries"
	@go build -v -o productservice-rest cmd/productservice/*.go

run-rest: build-rest
	@./productservice-rest