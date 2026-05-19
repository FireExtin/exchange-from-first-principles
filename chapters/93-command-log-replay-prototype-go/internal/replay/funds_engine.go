package replay

import (
	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

// FundsEngine implements funds.Engine using the decide/apply pattern:
//   decide(command) -> event   (pure function, no mutation)
//   apply(event)               (mutates state)
//
// This separation makes the business rules inspectable without side effects,
// and makes replay safe: replay events, not commands.
type FundsEngine struct{}

func NewFundsEngine() *FundsEngine {
	panic("TODO: implement")
}

func (e *FundsEngine) Handle(command funds.Command) ([]funds.Event, error) {
	panic("TODO: implement")
}

func (e *FundsEngine) Balance(accountID types.AccountID, asset types.Asset) types.Amount {
	panic("TODO: implement")
}

func (e *FundsEngine) Withdrawal(id types.WithdrawalID) (funds.Withdrawal, bool) {
	panic("TODO: implement")
}
