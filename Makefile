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
	docker compose down &
	flutter build web &
	# chromium --disable-web-security http://localhost:8080/ &
	go run . -w ${pwd}/build/web -u ${USER} -f pretty run

evans:
	evans --proto proto/v1/pacman.proto --web --host localhost -p 8080
