pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: gen
gen:
	buf format -w
	buf generate
	protoc --dart_out=grpc:gen/dart -Iproto proto/v1/pacman.proto

check:
	gofumpt -l -w .
	golangci-lint run
	buf lint

run:
	docker build -t ctlpkg:latest .
	docker run -p 8080:8080 -p 9080:9080 ctlpkg:latest run
