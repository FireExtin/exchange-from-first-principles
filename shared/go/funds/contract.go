package funds

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type CommandKind string

const (
	CommandDeposit           CommandKind = "deposit"
	CommandRequestWithdrawal CommandKind = "request_withdrawal"
	CommandConfirmWithdrawal CommandKind = "confirm_withdrawal"
	CommandTransfer          CommandKind = "transfer"
)

type Command struct {
	Seq             types.Seq
	Ref             types.Ref
	Kind            CommandKind
	AccountID       types.AccountID
	Asset           types.Asset
	Amount          types.Amount
	CallbackID      types.CallbackID
	WithdrawalID    types.WithdrawalID
	ProviderEventID types.ProviderEventID
	From            types.AccountID
	To              types.AccountID
}

type EventKind string

const (
	EventDeposited           EventKind = "deposited"
	EventWithdrawalRequested EventKind = "withdrawal_requested"
	EventWithdrawalConfirmed EventKind = "withdrawal_confirmed"
	EventTransferred         EventKind = "transferred"
	EventRejected            EventKind = "rejected"
)

type RejectReason string

const (
	RejectInvalidAmount          RejectReason = "invalid_amount"
	RejectDuplicateCallback      RejectReason = "duplicate_callback"
	RejectDuplicateWithdrawal    RejectReason = "duplicate_withdrawal"
	RejectDuplicateProviderEvent RejectReason = "duplicate_provider_event"
	RejectUnknownWithdrawal      RejectReason = "unknown_withdrawal"
	RejectInsufficientFunds      RejectReason = "insufficient_funds"
	RejectSequenceGap            RejectReason = "sequence_gap"
	RejectInvalidCommand         RejectReason = "invalid_command"
)

type Event struct {
	Seq             types.Seq
	Ref             types.Ref
	Kind            EventKind
	AccountID       types.AccountID
	Asset           types.Asset
	Amount          types.Amount
	CallbackID      types.CallbackID
	WithdrawalID    types.WithdrawalID
	ProviderEventID types.ProviderEventID
	From            types.AccountID
	To              types.AccountID
	Reason          RejectReason
}

type WithdrawalStatus string

const (
	WithdrawalRequested WithdrawalStatus = "requested"
	WithdrawalConfirmed WithdrawalStatus = "confirmed"
)

type Withdrawal struct {
	ID        types.WithdrawalID
	AccountID types.AccountID
	Asset     types.Asset
	Amount    types.Amount
	Status    WithdrawalStatus
}

type Engine interface {
	Handle(Command) ([]Event, error)
	Balance(types.AccountID, types.Asset) types.Amount
	Withdrawal(types.WithdrawalID) (Withdrawal, bool)
}
