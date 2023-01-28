pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: run
run:
	docker compose down 
	docker compose up --build app -d

.PHONY: gen
gen:
	docker run --rm -v ${pwd}:/src -w /src unibeautify/sqlformat -r -s -o sqlc.sql sqlc.sql
	docker run --rm -v ${pwd}:/src -w /src rvolosatovs/protoc --proto_path=/src --go_out=. --go-sqlc_out=require_unimplemented_servers=false:. api.proto
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate -f sqlc.yml

.PHONY: test
test:
	docker compose up -d
	go test -count 1 -cover ./...

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run
