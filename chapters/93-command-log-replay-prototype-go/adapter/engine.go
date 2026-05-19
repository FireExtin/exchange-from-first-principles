package adapter

import "github.com/FireExtin/exchange-from-first-principles/shared/go/funds"

// NewEngine returns a funds.Engine backed by replay.FundsEngine.
// Implement replay.FundsEngine (decide/apply pattern) to make this work.
func NewEngine() funds.Engine {
	panic("TODO: wire up replay.FundsEngine to implement funds.Engine")
}
