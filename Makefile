pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: run
run:
	docker compose down 
	docker compose up --build app -d

.PHONY: gen
gen:
	docker run --rm -v ${pwd}:/src -w /src bufbuild/buf generate .

.PHONY: test
test:
	docker compose up -d
	go test -count 1 -cover ./...

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run

