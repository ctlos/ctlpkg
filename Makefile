pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

check:
	gofumpt -l -w .
	golangci-lint run
	buf lint

run:
	docker build -t ctlpkg:latest  -f dockerfiles/main.dockerfile .

run-backend:
	docker build -t ctlpkg:backend  -f dockerfiles/backend.dockerfile .
	docker run -p 8080:8080 -p 9080:9080 ctlpkg:backend

run-frontend:
	docker build -t ctlpkg:frontend  -f dockerfiles/frontend.dockerfile .
	docker run -p 80:80 ctlpkg:frontend
