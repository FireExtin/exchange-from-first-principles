package credit

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type CommandKind string

const (
	CommandPledgeCollateral  CommandKind = "pledge_collateral"
	CommandReleaseCollateral CommandKind = "release_collateral"
	CommandBorrow            CommandKind = "borrow"
	CommandRepay             CommandKind = "repay"
	CommandAccrueFunding     CommandKind = "accrue_funding"
	CommandLiquidate         CommandKind = "liquidate"
)

type EventKind string

const (
	EventCollateralPledged   EventKind = "collateral_pledged"
	EventCollateralReleased  EventKind = "collateral_released"
	EventBorrowOpened        EventKind = "borrow_opened"
	EventRepaid              EventKind = "repaid"
	EventFundingAccrued      EventKind = "funding_accrued"
	EventLiquidationSettled  EventKind = "liquidation_settled"
	EventCreditRejected      EventKind = "credit_rejected"
	EventCreditViewProjected EventKind = "credit_view_projected"
)

type RejectReason string

const (
	RejectInvalidCommand         RejectReason = "invalid_command"
	RejectInvalidAmount          RejectReason = "invalid_amount"
	RejectInsufficientCollateral RejectReason = "insufficient_collateral"
	RejectInsufficientLiquidity  RejectReason = "insufficient_liquidity"
	RejectNoBorrowLiability      RejectReason = "no_borrow_liability"
	RejectMarginHealthy          RejectReason = "margin_healthy"
)

type Command struct {
	Seq             types.Seq
	Ref             types.Ref
	Kind            CommandKind
	AccountID       types.AccountID
	Asset           types.Asset
	Amount          types.Amount
	CollateralAsset types.Asset
	BorrowAsset     types.Asset
	Instrument      types.Asset
	MarkPrice       types.Amount
}

type Event struct {
	Seq         types.Seq
	Ref         types.Ref
	Kind        EventKind
	AccountID   types.AccountID
	Asset       types.Asset
	Amount      types.Amount
	JournalRefs []types.Ref
	Reason      RejectReason
}

type Balance struct {
	AccountID types.AccountID
	Asset     types.Asset
	Amount    types.Amount
}

type CreditView struct {
	AccountID          types.AccountID
	Collateral         []Balance
	BorrowLiabilities  []Balance
	AccruedFunding     []Balance
	MarginRequirement  []Balance
	LiquidationPending bool
}

type Snapshot struct {
	Seq    types.Seq
	Views  []CreditView
	Events []Event
}

type Adapter interface {
	Submit(Command) ([]Event, error)
	View(types.AccountID) CreditView
	Events() []Event
	Snapshot() Snapshot
	Restore(Snapshot) error
}
