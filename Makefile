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

flutter-docker:
	docker run --rm -it -v ${pwd}:/src -w /src -p 80:80 cirrusci/flutter bash -c "flutter pub get && flutter build web && cd build/web && python3 -m http.server 80"

.PHONY: gen
gen:
	buf format -w
	buf generate
