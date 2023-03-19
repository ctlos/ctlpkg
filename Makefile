pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: gen
gen:
	buf generate
	flutter pub run build_runner build

check:
	gofumpt -l -w .
	golangci-lint run
	buf lint

run:
	docker build -t ctlpkg:latest .
	docker run -p 8080:8080 -p 9080:9080 ctlpkg:latest run
