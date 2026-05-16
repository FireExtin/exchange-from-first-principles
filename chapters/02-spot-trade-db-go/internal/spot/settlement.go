package spot

import (
	"errors"
	"fmt"

	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

type Trade struct {
	Ref         types.Ref
	Buyer       types.AccountID
	Seller      types.AccountID
	BaseAsset   types.Asset
	QuoteAsset  types.Asset
	BaseAmount  types.Amount
	QuoteAmount types.Amount
}

type Store struct {
	balances map[types.BalanceKey]types.Amount
}

func NewStore() *Store {
	return &Store{balances: make(map[types.BalanceKey]types.Amount)}
}

func (s *Store) Deposit(accountID types.AccountID, asset types.Asset, amount types.Amount) {
	s.balances[key(accountID, asset)] += amount
}

func (s *Store) Balance(accountID types.AccountID, asset types.Asset) types.Amount {
	return s.balances[key(accountID, asset)]
}

func (s *Store) Settle(trade Trade) error {
	_, err := s.SettleEvents(trade)
	return err
}

func (s *Store) SettleEvents(trade Trade) ([]funds.Event, error) {
	if err := validateTrade(trade); err != nil {
		return nil, err
	}

	next := make(map[types.BalanceKey]types.Amount, len(s.balances)+4)
	for k, v := range s.balances {
		next[k] = v
	}

	apply := func(accountID types.AccountID, asset types.Asset, delta types.Amount) error {
		k := key(accountID, asset)
		next[k] += delta
		if next[k] < 0 {
			return fmt.Errorf("insufficient funds: account=%s asset=%s", accountID, asset)
		}
		return nil
	}

	if err := apply(trade.Buyer, trade.QuoteAsset, -trade.QuoteAmount); err != nil {
		return nil, err
	}
	if err := apply(trade.Seller, trade.BaseAsset, -trade.BaseAmount); err != nil {
		return nil, err
	}
	if err := apply(trade.Buyer, trade.BaseAsset, trade.BaseAmount); err != nil {
		return nil, err
	}
	if err := apply(trade.Seller, trade.QuoteAsset, trade.QuoteAmount); err != nil {
		return nil, err
	}

	s.balances = next
	return []funds.Event{
		{
			Ref:    trade.Ref,
			Kind:   funds.EventTransferred,
			From:   trade.Buyer,
			To:     trade.Seller,
			Asset:  trade.QuoteAsset,
			Amount: trade.QuoteAmount,
		},
		{
			Ref:    trade.Ref,
			Kind:   funds.EventTransferred,
			From:   trade.Seller,
			To:     trade.Buyer,
			Asset:  trade.BaseAsset,
			Amount: trade.BaseAmount,
		},
	}, nil
}

func validateTrade(trade Trade) error {
	if trade.Ref == "" || trade.Buyer == "" || trade.Seller == "" {
		return errors.New("ref, buyer, and seller are required")
	}
	if trade.BaseAsset == "" || trade.QuoteAsset == "" {
		return errors.New("base_asset and quote_asset are required")
	}
	if trade.BaseAmount <= 0 || trade.QuoteAmount <= 0 {
		return errors.New("amounts must be positive")
	}
	return nil
}

func key(accountID types.AccountID, asset types.Asset) types.BalanceKey {
	return types.BalanceKey{AccountID: accountID, Asset: asset}
}
