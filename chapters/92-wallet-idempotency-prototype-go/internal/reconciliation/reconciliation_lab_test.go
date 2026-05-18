//go:build reconciliation_lab_todo

package reconciliation

import (
	"testing"
	"time"

	"github.com/FireExtin/exchange-from-first-principles/chapters/92-wallet-idempotency-prototype-go/internal/wallet"
)

func TestProviderCallbackSimulatorProducesDuplicateOutOfOrderAndPartialRefund(t *testing.T) {
	base := time.Date(2026, 5, 17, 10, 0, 0, 0, time.UTC)
	sim := NewProviderCallbackSimulator("stripe", base)

	success := sim.Succeeded("evt-success-1", "charge-1", "alice", "USD", 10000)
	duplicate := sim.Duplicate(success, base.Add(time.Second))
	refund, delayedSuccess := sim.OutOfOrderPartialRefund("evt-refund-1", "evt-success-2", "charge-2", "bob", "USD", 3000, 10000)

	if success.Source != SourceProviderCallback || success.ExternalID != "evt-success-1" || success.Status != StatusSucceeded {
		t.Fatalf("unexpected success callback: %+v", success)
	}
	if duplicate.ExternalID != success.ExternalID || duplicate.ReceivedAt.Equal(success.ReceivedAt) {
		t.Fatalf("duplicate should preserve external id but have a later receipt time: %+v", duplicate)
	}
	if refund.Status != StatusPartiallyRefunded || delayedSuccess.Status != StatusSucceeded {
		t.Fatalf("unexpected out-of-order callbacks: refund=%+v success=%+v", refund, delayedSuccess)
	}
	if !refund.ReceivedAt.Before(delayedSuccess.ReceivedAt) {
		t.Fatal("refund callback should arrive before success in this fixture")
	}
}

func TestNormalizerConvertsProviderReportAndBankRecords(t *testing.T) {
	settlementDate := time.Date(2026, 5, 19, 0, 0, 0, 0, time.UTC)

	report := ProviderReportMock("stripe").SettledCharge("batch-1", "report-charge-1", "charge-1", "alice", "USD", 10000, 150, settlementDate)
	bank := BankOrChainRecordMock("hsbc").Payout("bank-payout-1", "batch-1", "USD", 9850, settlementDate.Add(12*time.Hour))

	reportRecord, err := Normalize(report)
	if err != nil {
		t.Fatalf("normalize provider report: %v", err)
	}
	bankRecord, err := Normalize(bank)
	if err != nil {
		t.Fatalf("normalize bank record: %v", err)
	}

	if reportRecord.Amount != 10000 || reportRecord.Fee != 150 || reportRecord.NetAmount != 9850 {
		t.Fatalf("provider report amounts not normalized: %+v", reportRecord)
	}
	if bankRecord.Source != SourceBankOrChain || bankRecord.BatchID != "batch-1" || bankRecord.NetAmount != 9850 {
		t.Fatalf("bank payout not normalized: %+v", bankRecord)
	}
}

func TestReconcileMatchesNormalDepositAcrossProviderReportAndBankPayout(t *testing.T) {
	settlementDate := time.Date(2026, 5, 19, 0, 0, 0, 0, time.UTC)
	asOf := settlementDate.Add(24 * time.Hour)

	internal := []NormalizedRecord{
		InternalLedgerRecord("ledger-1", "charge-1", "alice", "USD", 10000, StatusSucceeded, settlementDate),
	}
	external := normalizeAll(t,
		NewProviderCallbackSimulator("stripe", settlementDate.Add(-48*time.Hour)).Succeeded("evt-success-1", "charge-1", "alice", "USD", 10000),
		ProviderReportMock("stripe").SettledCharge("batch-1", "report-charge-1", "charge-1", "alice", "USD", 10000, 150, settlementDate),
		BankOrChainRecordMock("hsbc").Payout("bank-payout-1", "batch-1", "USD", 9850, settlementDate.Add(12*time.Hour)),
	)

	report := Reconcile(internal, external, asOf)

	if len(report.Matches) != 1 {
		t.Fatalf("matches = %d, want 1: %+v", len(report.Matches), report)
	}
	match := report.Matches[0]
	if match.Kind != MatchBatchFeeAdjusted || match.AmountDelta != 150 || match.FeeDelta != 150 {
		t.Fatalf("expected a fee-adjusted record match, got %+v", match)
	}
	if match.RuleID == "" || match.Confidence == 0 || match.Explanation == "" {
		t.Fatalf("match should explain the rule and confidence used: %+v", match)
	}
	if len(report.Discrepancies) != 0 {
		t.Fatalf("unexpected discrepancies: %+v", report.Discrepancies)
	}
	if len(report.AdjustmentProposals) != 0 {
		t.Fatalf("unexpected adjustment proposals: %+v", report.AdjustmentProposals)
	}
}

func TestReconcileRepresentsBatchLevelProviderChargesAgainstOnePayout(t *testing.T) {
	settlementDate := time.Date(2026, 5, 19, 0, 0, 0, 0, time.UTC)
	asOf := settlementDate.Add(24 * time.Hour)

	internal := []NormalizedRecord{
		InternalLedgerRecord("ledger-1", "deposit-1", "alice", "USD", 4000, StatusSucceeded, settlementDate),
		InternalLedgerRecord("ledger-2", "deposit-2", "bob", "USD", 3000, StatusSucceeded, settlementDate),
		InternalLedgerRecord("ledger-3", "deposit-3", "carol", "USD", 3000, StatusSucceeded, settlementDate),
	}
	external := normalizeAll(t,
		ProviderReportMock("stripe").SettledCharge("batch-1", "report-charge-1", "deposit-1", "alice", "USD", 4000, 60, settlementDate),
		ProviderReportMock("stripe").SettledCharge("batch-1", "report-charge-2", "deposit-2", "bob", "USD", 3000, 45, settlementDate),
		ProviderReportMock("stripe").SettledCharge("batch-1", "report-charge-3", "deposit-3", "carol", "USD", 3000, 45, settlementDate),
		BankOrChainRecordMock("hsbc").Payout("bank-payout-1", "batch-1", "USD", 9850, settlementDate.Add(12*time.Hour)),
	)

	report := Reconcile(internal, external, asOf)

	match := findMatch(t, report, MatchBatchFeeAdjusted)
	if match.AmountDelta != 150 || match.FeeDelta != 150 {
		t.Fatalf("batch match should explain gross-to-net fee delta, got %+v", match)
	}
	if len(match.External) != 4 {
		t.Fatalf("batch match should keep provider rows and bank payout together, got %d external records", len(match.External))
	}
	if len(report.Discrepancies) != 0 {
		t.Fatalf("batch fee match should not become an exception: %+v", report.Discrepancies)
	}
}

func TestReconcileClassifiesDuplicatesTimingAndAmountMismatch(t *testing.T) {
	now := time.Date(2026, 5, 17, 10, 0, 0, 0, time.UTC)
	sim := NewProviderCallbackSimulator("stripe", now)
	first := sim.Succeeded("evt-success-1", "charge-1", "alice", "USD", 10000)
	duplicate := sim.Duplicate(first, now.Add(time.Second))

	internal := []NormalizedRecord{
		InternalLedgerRecord("ledger-1", "charge-1", "alice", "USD", 9000, StatusSucceeded, now),
		InternalLedgerRecord("ledger-2", "charge-pending", "alice", "USD", 5000, StatusSucceeded, now.Add(48*time.Hour)),
		InternalLedgerRecord("ledger-3", "charge-missing", "alice", "USD", 7000, StatusSucceeded, now.Add(-48*time.Hour)),
	}
	external := normalizeAll(t, first, duplicate)

	report := Reconcile(internal, external, now)

	assertDiscrepancy(t, report, KindDuplicateExternal, "evt-success-1")
	assertDiscrepancy(t, report, KindAmountMismatch, "charge-1")
	assertDiscrepancy(t, report, KindTimingDifference, "charge-pending")
	assertDiscrepancy(t, report, KindMissingExternal, "charge-missing")
}

func TestPartialRefundProducesAdjustmentProposal(t *testing.T) {
	now := time.Date(2026, 5, 17, 10, 0, 0, 0, time.UTC)
	sim := NewProviderCallbackSimulator("stripe", now)

	internal := []NormalizedRecord{
		InternalLedgerRecord("ledger-1", "charge-1", "alice", "USD", 10000, StatusSucceeded, now),
	}
	external := normalizeAll(t,
		sim.Succeeded("evt-success-1", "charge-1", "alice", "USD", 10000),
		sim.PartialRefund("evt-refund-1", "charge-1", "alice", "USD", 3000),
	)

	report := Reconcile(internal, external, now.Add(time.Hour))

	assertDiscrepancy(t, report, KindNeedsAdjustment, "charge-1")
	if len(report.AdjustmentProposals) != 1 {
		t.Fatalf("adjustment proposals = %d, want 1: %+v", len(report.AdjustmentProposals), report.AdjustmentProposals)
	}
	proposal := report.AdjustmentProposals[0]
	if proposal.BusinessID != "charge-1" || proposal.Amount != -3000 || proposal.Reason != AdjustmentPartialRefund {
		t.Fatalf("unexpected proposal: %+v", proposal)
	}
}

func TestAdjustmentJournalRecordsEvidenceWithoutMutatingWallet(t *testing.T) {
	p := wallet.NewProcessor()
	if _, err := p.HandleDeposit("funding-1", "alice", "USD", 10000); err != nil {
		t.Fatalf("funding: %v", err)
	}

	journal := NewAdjustmentJournal()
	entry := journal.Append(AdjustmentProposal{
		BusinessID: "charge-1",
		AccountID:  "alice",
		Asset:      "USD",
		Amount:     -3000,
		Reason:     AdjustmentPartialRefund,
		Evidence:   []string{"evt-refund-1", "provider-report-1"},
	}, "ops-user-1", "partial refund reviewed", time.Date(2026, 5, 17, 12, 0, 0, 0, time.UTC))

	if entry.Operator != "ops-user-1" || len(entry.Evidence) != 2 {
		t.Fatalf("journal entry did not preserve audit evidence: %+v", entry)
	}
	if got := p.Balance("alice", "USD"); got != 10000 {
		t.Fatalf("adjustment journal must not mutate wallet balance, got %d", got)
	}
}

func normalizeAll(t *testing.T, raws ...RawRecord) []NormalizedRecord {
	t.Helper()
	records := make([]NormalizedRecord, 0, len(raws))
	for _, raw := range raws {
		record, err := Normalize(raw)
		if err != nil {
			t.Fatalf("normalize %+v: %v", raw, err)
		}
		records = append(records, record)
	}
	return records
}

func assertDiscrepancy(t *testing.T, report Report, kind DiscrepancyKind, key string) {
	t.Helper()
	for _, discrepancy := range report.Discrepancies {
		if discrepancy.Kind != kind {
			continue
		}
		if discrepancy.BusinessID == key || discrepancy.ExternalID == key {
			return
		}
	}
	t.Fatalf("missing discrepancy kind=%s key=%s in %+v", kind, key, report.Discrepancies)
}

func findMatch(t *testing.T, report Report, kind MatchKind) Match {
	t.Helper()
	for _, match := range report.Matches {
		if match.Kind == kind {
			return match
		}
	}
	t.Fatalf("missing match kind=%s in %+v", kind, report.Matches)
	return Match{}
}
