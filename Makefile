.PHONY: build clean
build:
	@./build.sh build
	@go mod tidy
clean:
	@./build.sh clean