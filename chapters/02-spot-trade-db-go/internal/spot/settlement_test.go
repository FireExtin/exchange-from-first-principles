package spot

import "testing"

func TestSettleSpotTradeAtomically(t *testing.T) {
	store := NewStore()
	store.Deposit("alice", "USD", 10000)
	store.Deposit("bob", "BTC", 2)

	err := store.Settle(Trade{
		Ref:         "trade-1",
		Buyer:       "alice",
		Seller:      "bob",
		BaseAsset:   "BTC",
		QuoteAsset:  "USD",
		BaseAmount:  1,
		QuoteAmount: 5000,
	})
	if err != nil {
		t.Fatalf("settle: %v", err)
	}

	if got := store.Balance("alice", "USD"); got != 5000 {
		t.Fatalf("alice USD = %d, want 5000", got)
	}
	if got := store.Balance("alice", "BTC"); got != 1 {
		t.Fatalf("alice BTC = %d, want 1", got)
	}
	if got := store.Balance("bob", "USD"); got != 5000 {
		t.Fatalf("bob USD = %d, want 5000", got)
	}
	if got := store.Balance("bob", "BTC"); got != 1 {
		t.Fatalf("bob BTC = %d, want 1", got)
	}
}

func TestFailedSettlementDoesNotPartiallyUpdate(t *testing.T) {
	store := NewStore()
	store.Deposit("alice", "USD", 100)
	store.Deposit("bob", "BTC", 2)

	err := store.Settle(Trade{
		Ref:         "trade-2",
		Buyer:       "alice",
		Seller:      "bob",
		BaseAsset:   "BTC",
		QuoteAsset:  "USD",
		BaseAmount:  1,
		QuoteAmount: 5000,
	})
	if err == nil {
		t.Fatal("expected insufficient funds")
	}

	if got := store.Balance("alice", "USD"); got != 100 {
		t.Fatalf("alice USD changed after failed settlement: %d", got)
	}
	if got := store.Balance("bob", "BTC"); got != 2 {
		t.Fatalf("bob BTC changed after failed settlement: %d", got)
	}
}
