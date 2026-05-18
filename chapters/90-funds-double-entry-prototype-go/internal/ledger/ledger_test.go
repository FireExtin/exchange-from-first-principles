package ledger

import (
	"strings"
	"testing"
)

func TestApplyBalancedTransfer(t *testing.T) {
	book := New()
	err := book.Apply("fund-alice", []Entry{
		{AccountID: "external", Asset: "USD", Delta: -1000},
		{AccountID: "alice", Asset: "USD", Delta: 1000},
	})
	if err != nil {
		t.Fatalf("fund alice: %v", err)
	}

	err = book.Apply("pay-bob", []Entry{
		{AccountID: "alice", Asset: "USD", Delta: -250},
		{AccountID: "bob", Asset: "USD", Delta: 250},
	})
	if err != nil {
		t.Fatalf("pay bob: %v", err)
	}

	if got := book.Balance("alice", "USD"); got != 750 {
		t.Fatalf("alice balance = %d, want 750", got)
	}
	if got := book.Balance("bob", "USD"); got != 250 {
		t.Fatalf("bob balance = %d, want 250", got)
	}
}

func TestRejectsUnbalancedTransaction(t *testing.T) {
	book := New()
	err := book.Apply("bad-credit", []Entry{
		{AccountID: "alice", Asset: "USD", Delta: 100},
		{AccountID: "bob", Asset: "USD", Delta: 1},
	})
	if err == nil || !strings.Contains(err.Error(), "unbalanced") {
		t.Fatalf("expected unbalanced error, got %v", err)
	}
}
