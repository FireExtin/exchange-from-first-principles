package exchange

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type PositionUpdate struct {
	AccountID     types.AccountID
	Instrument    Instrument
	Quantity      types.Amount
	AverageEntry  types.Amount
	RealizedPnL   types.Amount
	UnrealizedPnL types.Amount
}

type PositionView struct {
	AccountID    types.AccountID
	Instrument   Instrument
	Quantity     types.Amount
	AverageEntry types.Amount
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
