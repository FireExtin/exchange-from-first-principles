package integrationtests_test

import (
	"reflect"
	"testing"

	walletadapter "github.com/FireExtin/exchange-from-first-principles/chapters/92-wallet-idempotency-prototype-go/adapter"
	replayadapter "github.com/FireExtin/exchange-from-first-principles/chapters/93-command-log-replay-prototype-go/adapter"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

type engineCase struct {
	name string
	new  func() funds.Engine
}

func TestFundsEnginesShareBusinessSemantics(t *testing.T) {
	engines := []engineCase{
		{name: "wallet-workflow", new: walletadapter.NewEngine},
		{name: "command-replay", new: replayadapter.NewEngine},
	}

	scenarios := []struct {
		name       string
		commands   []funds.Command
		events     []funds.Event
		balances   map[types.BalanceKey]types.Amount
		withdrawal types.WithdrawalID
		status     funds.WithdrawalStatus
	}{
		{
			name: "deposit callback is idempotent",
			commands: []funds.Command{
				{Seq: 1, Ref: "deposit-1", Kind: funds.CommandDeposit, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "deposit-1-duplicate", Kind: funds.CommandDeposit, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
			},
			events: []funds.Event{
				{Seq: 1, Ref: "deposit-1", Kind: funds.EventDeposited, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "deposit-1-duplicate", Kind: funds.EventRejected, Reason: funds.RejectDuplicateCallback},
			},
			balances: map[types.BalanceKey]types.Amount{
				{AccountID: "alice", Asset: "USDT"}: 100,
			},
		},
		{
			name: "withdrawal cannot overdraft",
			commands: []funds.Command{
				{Seq: 1, Ref: "deposit-1", Kind: funds.CommandDeposit, AccountID: "alice", Asset: "USDT", Amount: 50, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "withdrawal-1", Kind: funds.CommandRequestWithdrawal, WithdrawalID: "wd-1", AccountID: "alice", Asset: "USDT", Amount: 100},
			},
			events: []funds.Event{
				{Seq: 1, Ref: "deposit-1", Kind: funds.EventDeposited, AccountID: "alice", Asset: "USDT", Amount: 50, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "withdrawal-1", Kind: funds.EventRejected, Reason: funds.RejectInsufficientFunds},
			},
			balances: map[types.BalanceKey]types.Amount{
				{AccountID: "alice", Asset: "USDT"}: 50,
			},
		},
		{
			name: "withdrawal confirmation is idempotent",
			commands: []funds.Command{
				{Seq: 1, Ref: "deposit-1", Kind: funds.CommandDeposit, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "withdrawal-1", Kind: funds.CommandRequestWithdrawal, WithdrawalID: "wd-1", AccountID: "alice", Asset: "USDT", Amount: 40},
				{Seq: 3, Ref: "confirm-1", Kind: funds.CommandConfirmWithdrawal, WithdrawalID: "wd-1", ProviderEventID: "provider-event-1"},
				{Seq: 4, Ref: "confirm-1-duplicate", Kind: funds.CommandConfirmWithdrawal, WithdrawalID: "wd-1", ProviderEventID: "provider-event-1"},
			},
			events: []funds.Event{
				{Seq: 1, Ref: "deposit-1", Kind: funds.EventDeposited, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "withdrawal-1", Kind: funds.EventWithdrawalRequested, WithdrawalID: "wd-1", AccountID: "alice", Asset: "USDT", Amount: 40},
				{Seq: 3, Ref: "confirm-1", Kind: funds.EventWithdrawalConfirmed, WithdrawalID: "wd-1", ProviderEventID: "provider-event-1"},
				{Seq: 4, Ref: "confirm-1-duplicate", Kind: funds.EventRejected, Reason: funds.RejectDuplicateProviderEvent},
			},
			balances: map[types.BalanceKey]types.Amount{
				{AccountID: "alice", Asset: "USDT"}: 60,
			},
			withdrawal: "wd-1",
			status:     funds.WithdrawalConfirmed,
		},
		{
			name: "transfer moves funds with replayable facts",
			commands: []funds.Command{
				{Seq: 1, Ref: "deposit-1", Kind: funds.CommandDeposit, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "transfer-1", Kind: funds.CommandTransfer, From: "alice", To: "bob", Asset: "USDT", Amount: 30},
			},
			events: []funds.Event{
				{Seq: 1, Ref: "deposit-1", Kind: funds.EventDeposited, AccountID: "alice", Asset: "USDT", Amount: 100, CallbackID: "chain-tx-1"},
				{Seq: 2, Ref: "transfer-1", Kind: funds.EventTransferred, From: "alice", To: "bob", Asset: "USDT", Amount: 30},
			},
			balances: map[types.BalanceKey]types.Amount{
				{AccountID: "alice", Asset: "USDT"}: 70,
				{AccountID: "bob", Asset: "USDT"}:   30,
			},
		},
	}

	for _, scenario := range scenarios {
		for _, engineCase := range engines {
			t.Run(scenario.name+"/"+engineCase.name, func(t *testing.T) {
				engine := engineCase.new()
				var gotEvents []funds.Event
				for _, command := range scenario.commands {
					events, err := engine.Handle(command)
					if err != nil {
						t.Fatalf("handle command %+v: %v", command, err)
					}
					gotEvents = append(gotEvents, events...)
				}

				if !reflect.DeepEqual(gotEvents, scenario.events) {
					t.Fatalf("events mismatch\n got: %#v\nwant: %#v", gotEvents, scenario.events)
				}
				for key, want := range scenario.balances {
					if got := engine.Balance(key.AccountID, key.Asset); got != want {
						t.Fatalf("balance %+v = %d, want %d", key, got, want)
					}
				}
				if scenario.withdrawal != "" {
					withdrawal, ok := engine.Withdrawal(scenario.withdrawal)
					if !ok {
						t.Fatalf("withdrawal %s not found", scenario.withdrawal)
					}
					if withdrawal.Status != scenario.status {
						t.Fatalf("withdrawal status = %s, want %s", withdrawal.Status, scenario.status)
					}
				}
			})
		}
	}
}
