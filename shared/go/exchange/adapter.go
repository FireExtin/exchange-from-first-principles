package exchange

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

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

type Snapshot struct {
	Cursor    ProjectionCursor
	Orders    []OrderView
	Positions []PositionView
	Journal   []JournalTransaction
	Events    []Event
}

type Adapter interface {
	Submit(Command) ([]Event, error)
	AccountBalance(AccountRef) types.Amount
	Orders() []OrderView
	Positions() []PositionView
	Events() []Event
	Snapshot() Snapshot
	Restore(Snapshot) error
}
