package adapter

import "github.com/FireExtin/exchange-from-first-principles/shared/go/funds"

// NewEngine returns a funds.Engine backed by wallet.Processor.
// Wire up wallet.Processor to implement funds.Engine:
//   - map each funds.Command kind to the corresponding Processor method
//   - translate Processor boolean results into funds.Event accept/reject
//   - enforce sequence continuity before delegating to the Processor
func NewEngine() funds.Engine {
	panic("TODO: wire up wallet.Processor to implement funds.Engine")
}
