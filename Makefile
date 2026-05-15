.PHONY: test test-go test-java test-rust run-ledger tree

test: test-go test-java

test-go:
	cd chapters/01-double-entry-ledger-go && go test ./...
	cd chapters/02-spot-trade-db-go && go test ./...
	cd chapters/03-wallet-deposit-withdrawal-go && go test ./...
	cd chapters/04-command-log-replay-go && go test ./...
	cd tools/go && go test ./...

test-java:
	@if command -v gradle >/dev/null 2>&1; then \
		cd chapters/11-replicated-state-machine-aeron-java && gradle --no-daemon clean test; \
	else \
		echo "SKIP: gradle not found. Install Gradle locally or run this target on the remote dev box."; \
	fi

test-rust:
	cd chapters/15-rust-hot-path && cargo test

run-ledger:
	cd chapters/01-double-entry-ledger-go && go run ./cmd/demo

tree:
	find . -maxdepth 3 -type f | sort
