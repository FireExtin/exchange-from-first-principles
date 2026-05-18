package replay

import (
	"fmt"

	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

type FundsEngine struct {
	lastSeq          types.Seq
	balances         map[types.BalanceKey]types.Amount
	seenCallbacks    map[types.CallbackID]struct{}
	seenProviderEvts map[types.ProviderEventID]struct{}
	withdrawals      map[types.WithdrawalID]funds.Withdrawal
	events           []funds.Event
}

func NewFundsEngine() *FundsEngine {
	return &FundsEngine{
		balances:         make(map[types.BalanceKey]types.Amount),
		seenCallbacks:    make(map[types.CallbackID]struct{}),
		seenProviderEvts: make(map[types.ProviderEventID]struct{}),
		withdrawals:      make(map[types.WithdrawalID]funds.Withdrawal),
	}
}

func (e *FundsEngine) Handle(command funds.Command) ([]funds.Event, error) {
	if command.Seq != e.lastSeq+1 {
		return nil, fmt.Errorf("sequence gap: got=%d want=%d", command.Seq, e.lastSeq+1)
	}
	event := e.decide(command)
	e.apply(event)
	e.lastSeq = command.Seq
	e.events = append(e.events, event)
	return []funds.Event{event}, nil
}

func (e *FundsEngine) Balance(accountID types.AccountID, asset types.Asset) types.Amount {
	return e.balances[types.BalanceKey{AccountID: accountID, Asset: asset}]
}

func (e *FundsEngine) Withdrawal(id types.WithdrawalID) (funds.Withdrawal, bool) {
	withdrawal, ok := e.withdrawals[id]
	return withdrawal, ok
}

func (e *FundsEngine) Events() []funds.Event {
	out := make([]funds.Event, len(e.events))
	copy(out, e.events)
	return out
}

func (e *FundsEngine) decide(command funds.Command) funds.Event {
	switch command.Kind {
	case funds.CommandDeposit:
		return e.decideDeposit(command)
	case funds.CommandRequestWithdrawal:
		return e.decideRequestWithdrawal(command)
	case funds.CommandConfirmWithdrawal:
		return e.decideConfirmWithdrawal(command)
	case funds.CommandTransfer:
		return e.decideTransfer(command)
	default:
		return rejectFunds(command, funds.RejectInvalidCommand)
	}
}

func (e *FundsEngine) decideDeposit(command funds.Command) funds.Event {
	if command.CallbackID == "" || command.AccountID == "" || command.Asset == "" || command.Amount <= 0 {
		return rejectFunds(command, funds.RejectInvalidAmount)
	}
	if _, ok := e.seenCallbacks[command.CallbackID]; ok {
		return rejectFunds(command, funds.RejectDuplicateCallback)
	}
	return funds.Event{
		Seq:        command.Seq,
		Ref:        command.Ref,
		Kind:       funds.EventDeposited,
		AccountID:  command.AccountID,
		Asset:      command.Asset,
		Amount:     command.Amount,
		CallbackID: command.CallbackID,
	}
}

func (e *FundsEngine) decideRequestWithdrawal(command funds.Command) funds.Event {
	if command.WithdrawalID == "" || command.AccountID == "" || command.Asset == "" || command.Amount <= 0 {
		return rejectFunds(command, funds.RejectInvalidAmount)
	}
	if _, ok := e.withdrawals[command.WithdrawalID]; ok {
		return rejectFunds(command, funds.RejectDuplicateWithdrawal)
	}
	if e.Balance(command.AccountID, command.Asset) < command.Amount {
		return rejectFunds(command, funds.RejectInsufficientFunds)
	}
	return funds.Event{
		Seq:          command.Seq,
		Ref:          command.Ref,
		Kind:         funds.EventWithdrawalRequested,
		WithdrawalID: command.WithdrawalID,
		AccountID:    command.AccountID,
		Asset:        command.Asset,
		Amount:       command.Amount,
	}
}

func (e *FundsEngine) decideConfirmWithdrawal(command funds.Command) funds.Event {
	if command.WithdrawalID == "" || command.ProviderEventID == "" {
		return rejectFunds(command, funds.RejectInvalidCommand)
	}
	if _, ok := e.seenProviderEvts[command.ProviderEventID]; ok {
		return rejectFunds(command, funds.RejectDuplicateProviderEvent)
	}
	if _, ok := e.withdrawals[command.WithdrawalID]; !ok {
		return rejectFunds(command, funds.RejectUnknownWithdrawal)
	}
	return funds.Event{
		Seq:             command.Seq,
		Ref:             command.Ref,
		Kind:            funds.EventWithdrawalConfirmed,
		WithdrawalID:    command.WithdrawalID,
		ProviderEventID: command.ProviderEventID,
	}
}

func (e *FundsEngine) decideTransfer(command funds.Command) funds.Event {
	if command.From == "" || command.To == "" || command.Asset == "" || command.Amount <= 0 {
		return rejectFunds(command, funds.RejectInvalidAmount)
	}
	if e.Balance(command.From, command.Asset) < command.Amount {
		return rejectFunds(command, funds.RejectInsufficientFunds)
	}
	return funds.Event{
		Seq:    command.Seq,
		Ref:    command.Ref,
		Kind:   funds.EventTransferred,
		From:   command.From,
		To:     command.To,
		Asset:  command.Asset,
		Amount: command.Amount,
	}
}

func (e *FundsEngine) apply(event funds.Event) {
	switch event.Kind {
	case funds.EventDeposited:
		e.balances[types.BalanceKey{AccountID: event.AccountID, Asset: event.Asset}] += event.Amount
		e.seenCallbacks[event.CallbackID] = struct{}{}
	case funds.EventWithdrawalRequested:
		key := types.BalanceKey{AccountID: event.AccountID, Asset: event.Asset}
		e.balances[key] -= event.Amount
		e.withdrawals[event.WithdrawalID] = funds.Withdrawal{
			ID:        event.WithdrawalID,
			AccountID: event.AccountID,
			Asset:     event.Asset,
			Amount:    event.Amount,
			Status:    funds.WithdrawalRequested,
		}
	case funds.EventWithdrawalConfirmed:
		withdrawal := e.withdrawals[event.WithdrawalID]
		withdrawal.Status = funds.WithdrawalConfirmed
		e.withdrawals[event.WithdrawalID] = withdrawal
		e.seenProviderEvts[event.ProviderEventID] = struct{}{}
	case funds.EventTransferred:
		e.balances[types.BalanceKey{AccountID: event.From, Asset: event.Asset}] -= event.Amount
		e.balances[types.BalanceKey{AccountID: event.To, Asset: event.Asset}] += event.Amount
	}
}

func rejectFunds(command funds.Command, reason funds.RejectReason) funds.Event {
	return funds.Event{Seq: command.Seq, Ref: command.Ref, Kind: funds.EventRejected, Reason: reason}
}
