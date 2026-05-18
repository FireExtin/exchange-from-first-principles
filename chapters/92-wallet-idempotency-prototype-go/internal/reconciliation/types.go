package reconciliation

import "time"

type Source string

const (
	SourceInternalLedger     Source = "internal_ledger"
	SourceProviderCallback   Source = "provider_callback"
	SourceProviderSettlement Source = "provider_settlement"
	SourceBankOrChain        Source = "bank_or_chain"
)

type Status string

const (
	StatusSucceeded         Status = "succeeded"
	StatusFailed            Status = "failed"
	StatusSettled           Status = "settled"
	StatusRefunded          Status = "refunded"
	StatusPartiallyRefunded Status = "partially_refunded"
	StatusChargeback        Status = "chargeback"
	StatusPending           Status = "pending"
)

type RawRecord struct {
	Source         Source
	Provider       string
	ExternalID     string
	BusinessID     string
	AccountID      string
	Asset          string
	Amount         int64
	Fee            int64
	NetAmount      int64
	Status         Status
	BatchID        string
	TransactionAt  time.Time
	SettlementDate time.Time
	ReceivedAt     time.Time
	RawPayloadHash string
	RawPayload     string
}

type NormalizedRecord struct {
	Source         Source
	Provider       string
	ExternalID     string
	BusinessID     string
	AccountID      string
	Asset          string
	Amount         int64
	Fee            int64
	NetAmount      int64
	Status         Status
	BatchID        string
	TransactionAt  time.Time
	SettlementDate time.Time
	RawPayloadHash string
}

type DiscrepancyKind string

const (
	KindMatched           DiscrepancyKind = "matched"
	KindMissingInternal   DiscrepancyKind = "missing_internal"
	KindMissingExternal   DiscrepancyKind = "missing_external"
	KindAmountMismatch    DiscrepancyKind = "amount_mismatch"
	KindStatusMismatch    DiscrepancyKind = "status_mismatch"
	KindDuplicateExternal DiscrepancyKind = "duplicate_external"
	KindTimingDifference  DiscrepancyKind = "timing_difference"
	KindNeedsAdjustment   DiscrepancyKind = "needs_adjustment"
)

type MatchKind string

const (
	MatchExact            MatchKind = "exact"
	MatchBusinessID       MatchKind = "business_id"
	MatchBatchFeeAdjusted MatchKind = "batch_fee_adjusted"
	MatchTimingDifference MatchKind = "timing_difference"
	MatchManualReview     MatchKind = "manual_review"
)

type Match struct {
	Kind        MatchKind
	BusinessID  string
	Internal    NormalizedRecord
	External    []NormalizedRecord
	RuleID      string
	Confidence  int
	Explanation string
	AmountDelta int64
	FeeDelta    int64
}

type Discrepancy struct {
	Kind       DiscrepancyKind
	BusinessID string
	ExternalID string
	Reason     string
	Internal   *NormalizedRecord
	External   []NormalizedRecord
}

type AdjustmentReason string

const (
	AdjustmentPartialRefund AdjustmentReason = "partial_refund"
	AdjustmentProviderFee   AdjustmentReason = "provider_fee"
	AdjustmentUnknown       AdjustmentReason = "unknown"
)

type AdjustmentProposal struct {
	BusinessID string
	AccountID  string
	Asset      string
	Amount     int64
	Reason     AdjustmentReason
	Evidence   []string
}

type AdjustmentEntry struct {
	ID        string
	Proposal  AdjustmentProposal
	Operator  string
	Note      string
	Evidence  []string
	CreatedAt time.Time
}

type Report struct {
	RunID               string
	AsOf                time.Time
	Matches             []Match
	Discrepancies       []Discrepancy
	AdjustmentProposals []AdjustmentProposal
}
