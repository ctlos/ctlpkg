pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

run:
	docker compose down 
	docker compose up --build app -d

test:
	docker compose up -d
	go test -count 1 -cover ./...

check:
	gofumpt -l -w .
	golangci-lint run
	buf lint

.PHONY: gen
gen:
	sudo chmod a+rwx -R .
	buf format -w
	buf generate
