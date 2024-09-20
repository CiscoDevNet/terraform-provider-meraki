default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 120m ./internal/provider

# Update all definitions from OpenAPI spec
.PHONY: update
update:
	go run ./gen/generator.go -a

# Generate the code and docs from definitions
.PHONY: gen
gen:
	go generate

# Update all definitions from spec and generate code and docs
.PHONY: build
build: update gen
