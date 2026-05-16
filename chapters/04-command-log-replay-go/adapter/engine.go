package adapter

import (
	"github.com/FireExtin/exchange-from-first-principles/chapters/04-command-log-replay-go/internal/replay"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
)

func NewEngine() funds.Engine {
	return replay.NewFundsEngine()
}
