

.PHONY: clean all init generate generate_mocks

all: build/main

build/main: cmd/api/main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf generated

init: clean generate
	go mod tidy
	go mod vendor

test:
	go clean -testcache
	go test -short -coverprofile coverage.out -short -v ./...

test_api:
	go clean -testcache
	go test ./tests/...

generate: generated generate_mocks

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package generated -generate types,server,spec $< > generated/api.gen.go

INTERFACES_GO_FILES := $(shell find internal/entity/interfaces -name "estate.go")

INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:internal/entity/interfaces/%.go=internal/entity/mocks/%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
internal/entity/mocks/%.mock.gen.go: internal/entity/interfaces/%.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=mocks
