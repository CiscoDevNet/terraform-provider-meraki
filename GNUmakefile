default: test

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Clean up leftover test resources from previous runs
# Usage: make clean
.PHONY: clean
clean:
	@echo "========================================="
	@echo "Cleaning up test resources..."
	@echo "========================================="
	@if [ -z "$(MERAKI_API_KEY)" ]; then \
		echo "SKIPPED: MERAKI_API_KEY is not configured"; \
	else \
		go run internal/cleanup/main.go; \
	fi

# Run acceptance tests
# Usage: make test [NAME="Definition Name"] [DEBUG=1]
# NAME must be a definition name (e.g., "Organization Admin") matching the name: field in gen/definitions/*.yaml
.PHONY: test
test:
	@echo "========================================="
	@echo "Running acceptance tests..."
	@echo "========================================="
	@if [ -z "$(MERAKI_API_KEY)" ]; then \
		echo "SKIPPED: MERAKI_API_KEY is not configured"; \
		echo "To enable tests, configure MERAKI_API_KEY and other variables in your .env file"; \
	else \
		TEST_NAME=""; \
		CLEAN_NAME=$$(echo "$(NAME)" | sed 's/\$$$$//'); \
		ANCHOR=""; \
		if [ "$(NAME)" != "$${CLEAN_NAME}" ]; then ANCHOR='$$'; fi; \
		if [ -n "$(NAME)" ]; then \
			MATCH=$$(grep -rl "^name: $${CLEAN_NAME}$$" gen/definitions/*.yaml | head -1); \
			if [ -n "$${MATCH}" ]; then \
				CAMEL=$$(echo "$${CLEAN_NAME}" | awk '{for(i=1;i<=NF;i++) $$i=toupper(substr($$i,1,1)) substr($$i,2)}1' | tr -d ' '); \
				TEST_NAME="TestAcc.*Meraki$${CAMEL}$${ANCHOR}"; \
				echo "Resolved definition '$${CLEAN_NAME}' to test pattern: $${TEST_NAME}"; \
			else \
				echo "ERROR: No definition found matching '$${CLEAN_NAME}' in gen/definitions/*.yaml"; \
				exit 1; \
			fi; \
		fi; \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output.log";) \
		if [ -n "$${TEST_NAME}" ]; then echo "Running tests matching: $${TEST_NAME}"; fi; \
		TF_ACC=1 \
		MERAKI_RETRY_ON_ERROR_CODES="308,401,404" \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $${TEST_NAME:+-run "$${TEST_NAME}"} $(TESTARGS) -count 1 -timeout 120m ./internal/provider $(if $(DEBUG),2>&1 | tee test-output.log); \
	fi

# Generate code from definitions
# Usage:
#   make gen                                    - Regenerate all code from all definitions
#   make gen NAME="Organization Admin"          - Regenerate code for an existing definition
#   make gen SPEC_PATH="/orgs/{orgId}/..." NAME="Organization Admin" - Create new definition and generate code
.PHONY: gen
gen:
	go run gen/load_models.go
ifeq ($(NAME),)
	go run ./gen/generator.go -a
	go run ./gen/generator.go
else ifeq ($(SPEC_PATH),)
	go run ./gen/generator.go X "$(NAME)"
else
	go run ./gen/definition.go "$(SPEC_PATH)" "$(NAME)"
	go run ./gen/generator.go "$(SPEC_PATH)" "$(NAME)"
endif
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go
