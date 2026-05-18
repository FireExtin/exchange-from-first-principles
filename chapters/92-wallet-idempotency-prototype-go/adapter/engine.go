package adapter

import (
	"fmt"

	"github.com/FireExtin/exchange-from-first-principles/chapters/92-wallet-idempotency-prototype-go/internal/wallet"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/funds"
	"github.com/FireExtin/exchange-from-first-principles/shared/go/types"
)

type Engine struct {
	processor *wallet.Processor
	lastSeq   types.Seq
}

func NewEngine() funds.Engine {
	return &Engine{processor: wallet.NewProcessor()}
}

func (e *Engine) Handle(command funds.Command) ([]funds.Event, error) {
	if command.Seq != e.lastSeq+1 {
		return nil, fmt.Errorf("sequence gap: got=%d want=%d", command.Seq, e.lastSeq+1)
	}
	event := e.handleInOrder(command)
	e.lastSeq = command.Seq
	return []funds.Event{event}, nil
}

func (e *Engine) Balance(accountID types.AccountID, asset types.Asset) types.Amount {
	return types.Amount(e.processor.Balance(string(accountID), string(asset)))
}

func (e *Engine) Withdrawal(id types.WithdrawalID) (funds.Withdrawal, bool) {
	withdrawal, ok := e.processor.Withdrawal(string(id))
	if !ok {
		return funds.Withdrawal{}, false
	}
	return funds.Withdrawal{
		ID:        types.WithdrawalID(withdrawal.ID),
		AccountID: types.AccountID(withdrawal.AccountID),
		Asset:     types.Asset(withdrawal.Asset),
		Amount:    types.Amount(withdrawal.Amount),
		Status:    funds.WithdrawalStatus(withdrawal.Status),
	}, true
}

func (e *Engine) handleInOrder(command funds.Command) funds.Event {
	switch command.Kind {
	case funds.CommandDeposit:
		return e.handleDeposit(command)
	case funds.CommandRequestWithdrawal:
		return e.handleRequestWithdrawal(command)
	case funds.CommandConfirmWithdrawal:
		return e.handleConfirmWithdrawal(command)
	case funds.CommandTransfer:
		return e.handleTransfer(command)
	default:
		return reject(command, funds.RejectInvalidCommand)
	}
}

func (e *Engine) handleDeposit(command funds.Command) funds.Event {
	if command.CallbackID == "" || command.AccountID == "" || command.Asset == "" || command.Amount <= 0 {
		return reject(command, funds.RejectInvalidAmount)
	}
	applied, err := e.processor.HandleDeposit(
		string(command.CallbackID),
		string(command.AccountID),
		string(command.Asset),
		int64(command.Amount),
	)
	if err != nil {
		return reject(command, funds.RejectInvalidCommand)
	}
	if !applied {
		return reject(command, funds.RejectDuplicateCallback)
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

func (e *Engine) handleRequestWithdrawal(command funds.Command) funds.Event {
	if command.WithdrawalID == "" || command.AccountID == "" || command.Asset == "" || command.Amount <= 0 {
		return reject(command, funds.RejectInvalidAmount)
	}
	if _, ok := e.processor.Withdrawal(string(command.WithdrawalID)); ok {
		_, _ = e.processor.RequestWithdrawal(
			string(command.WithdrawalID),
			string(command.AccountID),
			string(command.Asset),
			int64(command.Amount),
		)
		return reject(command, funds.RejectDuplicateWithdrawal)
	}
	if e.Balance(command.AccountID, command.Asset) < command.Amount {
		return reject(command, funds.RejectInsufficientFunds)
	}
	applied, err := e.processor.RequestWithdrawal(
		string(command.WithdrawalID),
		string(command.AccountID),
		string(command.Asset),
		int64(command.Amount),
	)
	if err != nil {
		return reject(command, funds.RejectInvalidCommand)
	}
	if !applied {
		return reject(command, funds.RejectDuplicateWithdrawal)
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

func (e *Engine) handleConfirmWithdrawal(command funds.Command) funds.Event {
	if command.WithdrawalID == "" || command.ProviderEventID == "" {
		return reject(command, funds.RejectInvalidCommand)
	}
	applied, err := e.processor.ConfirmWithdrawal(string(command.ProviderEventID), string(command.WithdrawalID))
	if err != nil {
		return reject(command, funds.RejectUnknownWithdrawal)
	}
	if !applied {
		return reject(command, funds.RejectDuplicateProviderEvent)
	}
	return funds.Event{
		Seq:             command.Seq,
		Ref:             command.Ref,
		Kind:            funds.EventWithdrawalConfirmed,
		WithdrawalID:    command.WithdrawalID,
		ProviderEventID: command.ProviderEventID,
	}
}

func (e *Engine) handleTransfer(command funds.Command) funds.Event {
	if command.From == "" || command.To == "" || command.Asset == "" || command.Amount <= 0 {
		return reject(command, funds.RejectInvalidAmount)
	}
	if e.Balance(command.From, command.Asset) < command.Amount {
		return reject(command, funds.RejectInsufficientFunds)
	}
	if err := e.processor.Transfer(string(command.From), string(command.To), string(command.Asset), int64(command.Amount)); err != nil {
		return reject(command, funds.RejectInvalidCommand)
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

func reject(command funds.Command, reason funds.RejectReason) funds.Event {
	return funds.Event{Seq: command.Seq, Ref: command.Ref, Kind: funds.EventRejected, Reason: reason}
}
