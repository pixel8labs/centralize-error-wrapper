.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make run app   - Invoke InitApp method"
	@echo "  make run api   - Invoke InitApi method"

.PHONY: run
run:
	@go run example/cmd/main.go $(filter-out $@,$(MAKECMDGOALS))
%:
	@: