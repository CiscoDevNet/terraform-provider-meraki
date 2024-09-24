default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 120m ./internal/provider

# Create new definition and build the code
.PHONY: def
def:
	go run gen/load_models.go
	go run ./gen/definition.go "$(DEF_PATH)" "$(DEF_NAME)"
	go run ./gen/generator.go "$(DEF_PATH)" "$(DEF_NAME)"
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go

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
