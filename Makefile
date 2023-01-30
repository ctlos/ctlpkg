pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: run
run:
	docker compose down 
	docker compose up --build app -d

.PHONY: gen
gen:
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf generate .

.PHONY: test
test:
	docker compose up -d
	go test -count 1 -cover ./...

.PHONY: lint
lint:
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest gofumpt -l -w .
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest golangci-lint run
