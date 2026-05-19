package exchange

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type OwnerType string

const (
	OwnerPlatform OwnerType = "platform"
	OwnerUser     OwnerType = "user"
	OwnerClearing OwnerType = "clearing"
	OwnerExternal OwnerType = "external"
)

type AccountPurpose string

const (
	PurposeCustody           AccountPurpose = "custody"
	PurposeAvailable         AccountPurpose = "available"
	PurposeLocked            AccountPurpose = "locked"
	PurposePendingWithdrawal AccountPurpose = "pending_withdrawal"
	PurposeFeeRevenue        AccountPurpose = "fee_revenue"
	PurposeEquity            AccountPurpose = "equity"
)

type NormalSide string

const (
	NormalDebit  NormalSide = "debit"
	NormalCredit NormalSide = "credit"
)

type EntrySide string

const (
	EntryDebit  EntrySide = "debit"
	EntryCredit EntrySide = "credit"
)

type AccountRef struct {
	OwnerType  OwnerType
	OwnerID    types.AccountID
	Asset      types.Asset
	Purpose    AccountPurpose
	NormalSide NormalSide
}

type Entry struct {
	Account AccountRef
	Side    EntrySide
	Amount  types.Amount
	Asset   types.Asset
}

type JournalTransaction struct {
	Ref     types.Ref
	Kind    TransactionKind
	Entries []Entry
	Facts   []Event
}

type TransactionKind string

const (
	TransactionDeposit              TransactionKind = "deposit"
	TransactionWithdrawalRequest    TransactionKind = "withdrawal_request"
	TransactionWithdrawalConfirm    TransactionKind = "withdrawal_confirm"
	TransactionTransfer             TransactionKind = "transfer"
	TransactionOrderReservation     TransactionKind = "order_reservation"
	TransactionOrderCancelRelease   TransactionKind = "order_cancel_release"
	TransactionTradeExecution       TransactionKind = "trade_execution"
	TransactionPartialFillRelease   TransactionKind = "partial_fill_release"
	TransactionProjectionCheckpoint TransactionKind = "projection_checkpoint"
)
