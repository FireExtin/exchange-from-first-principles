package ledger

import (
	"errors"
	"fmt"
)

type Entry struct {
	AccountID string
	Asset     string
	Delta     int64
}

type Ledger struct {
	balances map[string]int64
}

func New() *Ledger {
	return &Ledger{balances: make(map[string]int64)}
}

func (l *Ledger) Apply(ref string, entries []Entry) error {
	if ref == "" {
		return errors.New("ref is required")
	}
	if len(entries) < 2 {
		return errors.New("a balanced transaction needs at least two entries")
	}

	sumByAsset := make(map[string]int64)
	next := make(map[string]int64, len(l.balances)+len(entries))
	for key, value := range l.balances {
		next[key] = value
	}

	for _, entry := range entries {
		if entry.AccountID == "" || entry.Asset == "" {
			return errors.New("account_id and asset are required")
		}
		sumByAsset[entry.Asset] += entry.Delta
		key := balanceKey(entry.AccountID, entry.Asset)
		next[key] += entry.Delta
		if entry.AccountID != "external" && next[key] < 0 {
			return fmt.Errorf("negative balance: account=%s asset=%s", entry.AccountID, entry.Asset)
		}
	}

	for asset, sum := range sumByAsset {
		if sum != 0 {
			return fmt.Errorf("unbalanced transaction for %s: sum=%d", asset, sum)
		}
	}

	l.balances = next
	return nil
}

func (l *Ledger) Balance(accountID, asset string) int64 {
	return l.balances[balanceKey(accountID, asset)]
}

func balanceKey(accountID, asset string) string {
	return accountID + ":" + asset
}
