.DEFAULT_GOAL := help

.PHONY: help
help: ## This help menu
	@grep -E '^\S+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

ALL_TESTS = go test ./...
.PHONY: test
test: ## Run tests
	@$(call COLOR_COMMAND, $(ALL_TESTS))

ALL_TESTS_VERBOSE = go test -v ./...
.PHONY: test-verb
test-verb: ## Run tests with verbose output
	@$(call COLOR_COMMAND, $(ALL_TESTS_VERBOSE))

.PHONY: test-watch
test-watch: ## Run tests in watch mode
	@$(call COLOR_COMMAND, $(ALL_TESTS))
	@while true; do \
		inotifywait -qq -r -e create,modify,move,delete ./; \
		printf "\n[ . . . Re-running command . . . ]\n"; \
		$(call COLOR_COMMAND, $(ALL_TESTS)); \
	done

# Fuzz tests
TIME=10
FUZZ_TEST_FILES=$(shell grep -r --include='**_test.go' --files-with-matches 'func Fuzz' .)
.PHONY: test-fuzz
test-fuzz: ## Run fuzz tests for [TIME] seconds each
	@for file in $(FUZZ_TEST_FILES); do \
		echo Fuzzing $${file} for $(TIME) seconds; \
		go test -fuzz=Fuzz $${file} -fuzztime=$(TIME)s; \
	done \

# Color Command
PASS_COLOR=$(shell echo -e "\e[1;32m")
FAIL_COLOR=$(shell echo -e "\e[1;31m")
RESET_COLOR=$(shell echo -e "\e[0m")

COLORED_PASS_TERMS=✅ $(PASS_COLOR)&$(RESET_COLOR)
COLORED_FAIL_TERMS=❌ $(FAIL_COLOR)&$(RESET_COLOR)
COLOR_COMMAND = $(1) | sed -Ee "s/\<pass\>|\<ok\>/$(COLORED_PASS_TERMS)/I" -Ee "s/\<fail\>|\<failed\>/$(COLORED_FAIL_TERMS)/I"
