.PHONY: generate
generate: generate-deps ## Generate code.
	buf generate

.PHONY: generate-deps
generate-deps: ## Install dependencies for code generation.
	@command -v buf > /dev/null || brew install buf
