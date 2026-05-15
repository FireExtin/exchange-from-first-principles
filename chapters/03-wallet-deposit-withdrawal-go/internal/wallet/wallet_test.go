package wallet

import "testing"

func TestDepositCallbackIsIdempotent(t *testing.T) {
	p := NewProcessor()

	applied, err := p.HandleDeposit("chain-tx-1", "alice", "USDT", 100)
	if err != nil || !applied {
		t.Fatalf("first deposit applied=%v err=%v", applied, err)
	}
	applied, err = p.HandleDeposit("chain-tx-1", "alice", "USDT", 100)
	if err != nil {
		t.Fatalf("duplicate deposit: %v", err)
	}
	if applied {
		t.Fatal("duplicate callback should not apply twice")
	}
	if got := p.Balance("alice", "USDT"); got != 100 {
		t.Fatalf("balance = %d, want 100", got)
	}
}

func TestWithdrawalRequestAndConfirmationAreIdempotent(t *testing.T) {
	p := NewProcessor()
	if _, err := p.HandleDeposit("funding-1", "alice", "USDT", 100); err != nil {
		t.Fatalf("funding: %v", err)
	}

	applied, err := p.RequestWithdrawal("wd-1", "alice", "USDT", 40)
	if err != nil || !applied {
		t.Fatalf("first withdrawal applied=%v err=%v", applied, err)
	}
	applied, err = p.RequestWithdrawal("wd-1", "alice", "USDT", 40)
	if err != nil {
		t.Fatalf("duplicate withdrawal: %v", err)
	}
	if applied {
		t.Fatal("duplicate withdrawal should not debit twice")
	}
	if got := p.Balance("alice", "USDT"); got != 60 {
		t.Fatalf("balance = %d, want 60", got)
	}

	applied, err = p.ConfirmWithdrawal("provider-event-1", "wd-1")
	if err != nil || !applied {
		t.Fatalf("confirm applied=%v err=%v", applied, err)
	}
	applied, err = p.ConfirmWithdrawal("provider-event-1", "wd-1")
	if err != nil {
		t.Fatalf("duplicate confirmation: %v", err)
	}
	if applied {
		t.Fatal("duplicate confirmation should not apply twice")
	}
}
