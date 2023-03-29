pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: gen
gen:
	buf format -w
	buf generate
	protoc --dart_out=grpc:lib/generated -Iproto proto/v1/pacman.proto

check:
	gofumpt -l -w .
	golangci-lint run
	buf lint

run:
	docker compose down
	docker compose up --build app -d

gorun:
	docker compose down &
	flutter build web &
	chromium --disable-web-security &
	go run . -w ${pwd}/build/web -u ${USER} -l pretty run
