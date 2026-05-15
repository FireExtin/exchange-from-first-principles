package spot

import (
	"errors"
	"fmt"
)

type Trade struct {
	Ref         string
	Buyer       string
	Seller      string
	BaseAsset   string
	QuoteAsset  string
	BaseAmount  int64
	QuoteAmount int64
}

type Store struct {
	balances map[string]int64
}

func NewStore() *Store {
	return &Store{balances: make(map[string]int64)}
}

func (s *Store) Deposit(accountID, asset string, amount int64) {
	s.balances[key(accountID, asset)] += amount
}

func (s *Store) Balance(accountID, asset string) int64 {
	return s.balances[key(accountID, asset)]
}

func (s *Store) Settle(trade Trade) error {
	if err := validateTrade(trade); err != nil {
		return err
	}

	next := make(map[string]int64, len(s.balances)+4)
	for k, v := range s.balances {
		next[k] = v
	}

	apply := func(accountID, asset string, delta int64) error {
		k := key(accountID, asset)
		next[k] += delta
		if next[k] < 0 {
			return fmt.Errorf("insufficient funds: account=%s asset=%s", accountID, asset)
		}
		return nil
	}

	if err := apply(trade.Buyer, trade.QuoteAsset, -trade.QuoteAmount); err != nil {
		return err
	}
	if err := apply(trade.Seller, trade.BaseAsset, -trade.BaseAmount); err != nil {
		return err
	}
	if err := apply(trade.Buyer, trade.BaseAsset, trade.BaseAmount); err != nil {
		return err
	}
	if err := apply(trade.Seller, trade.QuoteAsset, trade.QuoteAmount); err != nil {
		return err
	}

	s.balances = next
	return nil
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

func key(accountID, asset string) string {
	return accountID + ":" + asset
}
