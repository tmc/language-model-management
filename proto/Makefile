.PHONY: generate
generate: generate-deps ## Generate code.
	buf generate
	cd proto/llm/v1; d2 llm.erd.d2

.PHONY: generate-deps
generate-deps: ## Install dependencies for code generation.
	@command -v buf > /dev/null || brew install buf
