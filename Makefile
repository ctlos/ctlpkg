pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

run:
	docker compose down 
	docker compose up --build app -d


test:
	docker compose up -d
	go test -count 1 -cover ./...

check:
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest gofumpt -l -w .
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest golangci-lint run
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf lint

.PHONY: gen
gen:
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf format -w
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf generate
