default: testall

# Run all acceptance tests
.PHONY: testall
testall:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 120m ./internal/provider

# Run a subset of tests by specifying a regex (NAME).
# Usage: make test NAME=OrganizationAdmin
.PHONY: test
test:
	TF_ACC=1 go test ./... -v -run $(NAME) -timeout 120m

# Create new definition and build the code
# Usage: make gen SPEC_PATH="/organizations/{organizationId}/admins/{adminId}" NAME="Organization Admin"
# SPEC_PATH: The PUT API endpoint path
# NAME: The name of the definition, e.g. "Organization Admin"
.PHONY: gen
gen:
	go run gen/load_models.go
ifeq ($(SPEC_PATH),)
	go run ./gen/generator.go X "$(NAME)"
else
	go run ./gen/definition.go "$(SPEC_PATH)" "$(NAME)"
	go run ./gen/generator.go "$(SPEC_PATH)" "$(NAME)"
endif
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go

# Update all definitions from OpenAPI spec
.PHONY: genall
genall:
	go run gen/load_models.go
	go run ./gen/generator.go -a
	go run ./gen/generator.go
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go
