module github.com/FireExtin/exchange-from-first-principles/integration-tests

go 1.22

require (
	github.com/FireExtin/exchange-from-first-principles/chapters/92-wallet-idempotency-prototype-go v0.0.0
	github.com/FireExtin/exchange-from-first-principles/chapters/93-command-log-replay-prototype-go v0.0.0
	github.com/FireExtin/exchange-from-first-principles/shared/go v0.0.0
)

replace github.com/FireExtin/exchange-from-first-principles/chapters/92-wallet-idempotency-prototype-go => ../chapters/92-wallet-idempotency-prototype-go

replace github.com/FireExtin/exchange-from-first-principles/chapters/93-command-log-replay-prototype-go => ../chapters/93-command-log-replay-prototype-go

replace github.com/FireExtin/exchange-from-first-principles/shared/go => ../shared/go
