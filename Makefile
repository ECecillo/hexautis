.PHONY: run build-image run-image

build:
	@go build -v -o bin/hexautis ./cmd/main.go

run:
	@go run cmd/main.go

# Docker

build-image:
	@docker build -f Dockerfile --tag hexautis .

run-image:
	@docker run --rm -i hexautis
