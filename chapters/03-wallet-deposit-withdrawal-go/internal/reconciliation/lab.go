package reconciliation

import "time"

type ProviderCallbackSimulator struct {
	provider string
	baseTime time.Time
}

func NewProviderCallbackSimulator(provider string, baseTime time.Time) ProviderCallbackSimulator {
	return ProviderCallbackSimulator{provider: provider, baseTime: baseTime}
}

func (s ProviderCallbackSimulator) Succeeded(eventID, businessID, accountID, asset string, amount int64) RawRecord {
	panic("TODO: build a provider callback raw record for a successful charge")
}

func (s ProviderCallbackSimulator) Failed(eventID, businessID, accountID, asset string, amount int64) RawRecord {
	panic("TODO: build a provider callback raw record for a failed charge")
}

func (s ProviderCallbackSimulator) PartialRefund(eventID, businessID, accountID, asset string, amount int64) RawRecord {
	panic("TODO: build a provider callback raw record for a partial refund")
}

func (s ProviderCallbackSimulator) Chargeback(eventID, businessID, accountID, asset string, amount int64) RawRecord {
	panic("TODO: build a provider callback raw record for a chargeback")
}

func (s ProviderCallbackSimulator) Duplicate(original RawRecord, receivedAt time.Time) RawRecord {
	panic("TODO: preserve the external event id but change receipt metadata")
}

func (s ProviderCallbackSimulator) OutOfOrderPartialRefund(refundEventID, successEventID, businessID, accountID, asset string, refundAmount, chargeAmount int64) (RawRecord, RawRecord) {
	panic("TODO: return a refund callback that arrives before its success callback")
}

type ProviderReportBuilder struct {
	provider string
}

func ProviderReportMock(provider string) ProviderReportBuilder {
	return ProviderReportBuilder{provider: provider}
}

func (b ProviderReportBuilder) SettledCharge(batchID, externalID, businessID, accountID, asset string, grossAmount, fee int64, settlementDate time.Time) RawRecord {
	panic("TODO: build a provider settlement report row with gross, fee, and net")
}

type BankOrChainRecordBuilder struct {
	sourceName string
}

func BankOrChainRecordMock(sourceName string) BankOrChainRecordBuilder {
	return BankOrChainRecordBuilder{sourceName: sourceName}
}

func (b BankOrChainRecordBuilder) Payout(externalID, batchID, asset string, netAmount int64, settlementTime time.Time) RawRecord {
	panic("TODO: build a bank/custody/on-chain final cash movement record")
}

func InternalLedgerRecord(externalID, businessID, accountID, asset string, amount int64, status Status, settlementDate time.Time) NormalizedRecord {
	panic("TODO: build a normalized internal ledger/funds fact")
}

func Normalize(raw RawRecord) (NormalizedRecord, error) {
	panic("TODO: convert callback, provider report, and bank/chain raw records into one normalized shape")
}

func Reconcile(internal []NormalizedRecord, external []NormalizedRecord, asOf time.Time) Report {
	panic("TODO: match records, classify discrepancies, and produce adjustment proposals")
}

type AdjustmentJournal struct {
	entries []AdjustmentEntry
}

func NewAdjustmentJournal() *AdjustmentJournal {
	return &AdjustmentJournal{}
}

func (j *AdjustmentJournal) Append(proposal AdjustmentProposal, operator, note string, createdAt time.Time) AdjustmentEntry {
	panic("TODO: append audit evidence without mutating wallet or ledger state")
}

func (j *AdjustmentJournal) Entries() []AdjustmentEntry {
	return append([]AdjustmentEntry(nil), j.entries...)
}
