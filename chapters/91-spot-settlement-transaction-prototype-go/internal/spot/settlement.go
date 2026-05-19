package spot

import (
	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

// Trade describes a spot match: buyer pays QuoteAsset, seller delivers
// BaseAsset. Settlement is atomic: both legs apply or neither does.
type Trade struct {
	Ref         types.Ref
	Buyer       types.AccountID
	Seller      types.AccountID
	BaseAsset   types.Asset
	QuoteAsset  types.Asset
	BaseAmount  types.Amount
	QuoteAmount types.Amount
}

type Store struct{}

func NewStore() *Store {
	panic("TODO: implement")
}

func (s *Store) Deposit(accountID types.AccountID, asset types.Asset, amount types.Amount) {
	panic("TODO: implement")
}

func (s *Store) Balance(accountID types.AccountID, asset types.Asset) types.Amount {
	panic("TODO: implement")
}

// Settle applies the trade atomically. Either both legs succeed or neither
// does. Returns an error if either account has insufficient funds.
func (s *Store) Settle(trade Trade) error {
	panic("TODO: implement")
}

// SettleEvents applies the trade and returns two funds.EventTransferred
// facts: QuoteAsset buyer→seller and BaseAsset seller→buyer.
func (s *Store) SettleEvents(trade Trade) ([]funds.Event, error) {
	panic("TODO: implement")
}
