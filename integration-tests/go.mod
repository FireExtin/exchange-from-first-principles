module github.com/FireExtin/exchange-from-first-principles/integration-tests

go 1.22

require (
	github.com/FireExtin/exchange-from-first-principles/chapters/03-wallet-deposit-withdrawal-go v0.0.0
	github.com/FireExtin/exchange-from-first-principles/chapters/04-command-log-replay-go v0.0.0
	github.com/FireExtin/exchange-from-first-principles/shared/go v0.0.0
)

replace github.com/FireExtin/exchange-from-first-principles/chapters/03-wallet-deposit-withdrawal-go => ../chapters/03-wallet-deposit-withdrawal-go

replace github.com/FireExtin/exchange-from-first-principles/chapters/04-command-log-replay-go => ../chapters/04-command-log-replay-go

replace github.com/FireExtin/exchange-from-first-principles/shared/go => ../shared/go
