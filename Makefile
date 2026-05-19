.PHONY: test test-go test-java tree

test: test-go test-java

# Chapters 90-93 compile and are expected to panic at TODO boundaries.
# The tests define the semantic contract; implement the functions to make them pass.
test-go:
	cd shared/go && go test ./...
	cd chapters/90-funds-double-entry-prototype-go && go test ./...
	cd chapters/91-spot-settlement-transaction-prototype-go && go test ./...
	cd chapters/92-wallet-idempotency-prototype-go && go test ./...
	cd chapters/93-command-log-replay-prototype-go && go test ./...
	go test ./integration-tests/...
	cd tools/go && go test ./...

test-java:
	@if command -v gradle >/dev/null 2>&1; then \
		cd chapters/08-replicated-log-core-aeron-java && gradle --no-daemon clean test; \
	else \
		echo "SKIP: gradle not found. Install Gradle locally or run this target on the remote dev box."; \
	fi

tree:
	find . -maxdepth 3 -type f | sort
