build-rest:
	@echo " >> building rest binaries"
	@go build -v -o userservice-rest cmd/userservice/*.go

run-rest: build-rest
	@./userservice-rest