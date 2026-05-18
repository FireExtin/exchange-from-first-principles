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
	OwnerType OwnerType
	OwnerID   types.AccountID
	Asset     types.Asset
	Purpose   AccountPurpose
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

type CommandKind string

const (
	CommandDeposit           CommandKind = "deposit"
	CommandRequestWithdrawal CommandKind = "request_withdrawal"
	CommandConfirmWithdrawal CommandKind = "confirm_withdrawal"
	CommandTransfer          CommandKind = "transfer"
	CommandPlaceOrder        CommandKind = "place_order"
	CommandCancelOrder       CommandKind = "cancel_order"
	CommandApplyExecution    CommandKind = "apply_execution"
	CommandUpdateMark        CommandKind = "update_mark"
	CommandKillSwitch        CommandKind = "kill_switch"
)

type Command struct {
	Seq          types.Seq
	Ref          types.Ref
	Kind         CommandKind
	AccountID    types.AccountID
	Asset        types.Asset
	Amount       types.Amount
	From         types.AccountID
	To           types.AccountID
	Order        OrderCommand
	Execution    ExecutionReport
	RiskOverride RiskOverride
}

type EventKind string

const (
	EventDeposited             EventKind = "deposited"
	EventWithdrawalRequested   EventKind = "withdrawal_requested"
	EventWithdrawalConfirmed   EventKind = "withdrawal_confirmed"
	EventTransferred           EventKind = "transferred"
	EventOrderAccepted         EventKind = "order_accepted"
	EventOrderRejected         EventKind = "order_rejected"
	EventOrderReserved         EventKind = "order_reserved"
	EventOrderCancelled        EventKind = "order_cancelled"
	EventOrderReleased         EventKind = "order_released"
	EventTradeExecuted         EventKind = "trade_executed"
	EventPositionUpdated       EventKind = "position_updated"
	EventPreTradeRiskAccepted  EventKind = "pre_trade_risk_accepted"
	EventPreTradeRiskRejected  EventKind = "pre_trade_risk_rejected"
	EventMarginAccepted        EventKind = "margin_accepted"
	EventMarginRejected        EventKind = "margin_rejected"
	EventKillSwitchActivated   EventKind = "kill_switch_activated"
	EventProjectionCheckpoint  EventKind = "projection_checkpoint"
	EventProjectionGapDetected EventKind = "projection_gap_detected"
	EventRejected              EventKind = "rejected"
)

type Event struct {
	Seq         types.Seq
	Ref         types.Ref
	Kind        EventKind
	AccountID   types.AccountID
	Asset       types.Asset
	Amount      types.Amount
	Order       OrderEvent
	Execution   ExecutionReport
	Position    PositionUpdate
	Risk        RiskDecision
	Projection  ProjectionCursor
	JournalRefs []types.Ref
	Reason      RejectReason
}

type RejectReason string

const (
	RejectInvalidCommand    RejectReason = "invalid_command"
	RejectInvalidAmount     RejectReason = "invalid_amount"
	RejectInsufficientFunds RejectReason = "insufficient_funds"
	RejectUnknownOrder      RejectReason = "unknown_order"
	RejectInvalidOrderState RejectReason = "invalid_order_state"
	RejectRiskLimit         RejectReason = "risk_limit"
	RejectMarginLimit       RejectReason = "margin_limit"
	RejectKillSwitch        RejectReason = "kill_switch"
	RejectProjectionGap     RejectReason = "projection_gap"
)

type OrderID string
type TradeID string
type Instrument string
type Side string
type OrderStatus string
type LiquidityRole string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"

	OrderAccepted        OrderStatus = "accepted"
	OrderRejected        OrderStatus = "rejected"
	OrderPartiallyFilled OrderStatus = "partially_filled"
	OrderFilled          OrderStatus = "filled"
	OrderCancelled       OrderStatus = "cancelled"

	RoleMaker LiquidityRole = "maker"
	RoleTaker LiquidityRole = "taker"
)

type OrderCommand struct {
	OrderID    OrderID
	Instrument Instrument
	Side       Side
	Price      types.Amount
	Quantity   types.Amount
	MaxFee     types.Amount
}

type OrderEvent struct {
	OrderID         OrderID
	Instrument      Instrument
	Status          OrderStatus
	ReservedAsset   types.Asset
	ReservedAmount  types.Amount
	ReleasedAsset   types.Asset
	ReleasedAmount  types.Amount
	RemainingAmount types.Amount
}

type ExecutionReport struct {
	TradeID       TradeID
	MakerOrderID  OrderID
	TakerOrderID  OrderID
	Instrument    Instrument
	Price         types.Amount
	Quantity      types.Amount
	MakerAccount  types.AccountID
	TakerAccount  types.AccountID
	FeeAccount    types.AccountID
	FeeAsset      types.Asset
	FeeAmount     types.Amount
	LiquidityRole LiquidityRole
}

type PositionUpdate struct {
	AccountID     types.AccountID
	Instrument    Instrument
	Quantity      types.Amount
	AverageEntry  types.Amount
	RealizedPnL   types.Amount
	UnrealizedPnL types.Amount
}

type RiskDecision struct {
	AccountID  types.AccountID
	Instrument Instrument
	Accepted   bool
	Reason     RejectReason
}

type RiskOverride struct {
	AccountID types.AccountID
	Enabled   bool
	Reason    string
}

type ProjectionCursor struct {
	SnapshotID string
	EventSeq   types.Seq
	HasGap     bool
}

type ProjectionRebuildResult struct {
	Cursor        ProjectionCursor
	RowsRebuilt   int
	GapDetected   bool
	RebuildStatus string
}

type OrderView struct {
	OrderID    OrderID
	AccountID  types.AccountID
	Instrument Instrument
	Status     OrderStatus
	Remaining  types.Amount
}

type PositionView struct {
	AccountID    types.AccountID
	Instrument   Instrument
	Quantity     types.Amount
	AverageEntry types.Amount
}

type Snapshot struct {
	Cursor ProjectionCursor
	Bytes  []byte
}

type Adapter interface {
	Submit(Command) ([]Event, error)
	Balance(AccountRef) types.Amount
	Order(OrderID) (OrderView, bool)
	Position(types.AccountID, Instrument) (PositionView, bool)
	Facts() []Event
	Snapshot() (Snapshot, error)
	Restore(Snapshot) error
}
