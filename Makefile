pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: run
run:
	docker compose down 
	docker compose up --build app -d

.PHONY: gen
gen:
	docker run --rm -v ${pwd}:/src -w /src rvolosatovs/protoc --proto_path=/src --go_out=. --go-grpc_out=require_unimplemented_servers=false:. api.proto

.PHONY: test
test:
	docker compose up -d
	go test -count 1 -cover ./...

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run
