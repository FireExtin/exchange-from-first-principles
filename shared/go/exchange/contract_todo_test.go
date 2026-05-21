//go:build exchange_contract_todo

package exchange

import "testing"

func TestChapter01DepositPostsDebitCustodyAndCreditUserAvailable(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 01 deposit must post debit custody / credit user available through an adapter")
}

func TestChapter02AvailableMovesToLockedByJournalEntries(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 available-to-locked movement must be explained by journal entries")
}

func TestChapter02WithdrawalRequestMovesAvailableToPending(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 withdrawal request must move available liability into pending withdrawal")
}

func TestChapter02WithdrawalConfirmationMovesPendingToCustody(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 withdrawal confirmation must debit pending withdrawal and credit custody")
}

func TestChapter02FeeRevenueIsPlatformAccountState(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 fee revenue must be a platform account movement, not a user balance edit")
}

func TestChapter03BuyOrderReservesQuoteAsset(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 03 buy order must reserve quote asset before entering the book")
}

func TestChapter03SellOrderReservesBaseAsset(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 03 sell order must reserve base asset before entering the book")
}

func TestChapter03CancelOrderReleasesFunds(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 03 cancellation must release only the remaining locked amount")
}

func TestChapter04MatchEmitsExecutionFacts(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 match must emit execution facts before settlement and projections consume them")
}

func TestChapter04TradeBalancesUSDAndBTCSeparately(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 settlement must balance USD and BTC separately")
}

func TestChapter04BuyerFeePostsToPlatformRevenue(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 buyer fee must post to platform fee revenue")
}

func TestChapter04PartialFillReleasesSurplus(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 partial fill must release unused locked funds")
}

func TestChapter05ACIDAdapterCommitsLedgerOrderExecutionAndPositionAtomically(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 05 ACID adapter must commit ledger, order/reservation, execution, and position facts atomically")
}

func TestExecutionFactsTranslateToLedgerPostingsAtSettlementBoundary(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): execution facts must become ledger postings through an explicit settlement boundary")
}

func TestExecutionFactsUpdatePositions(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): later contract scenario requires an adapter; do not implement position logic here")
}

func TestMarksAndUnrealizedPnLDoNotCreateLedgerEntries(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): marks and unrealized PnL are derived views and must not create ledger entries by themselves")
}

func TestMarginAndRiskRejectDangerousOrders(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): later contract scenario requires an adapter; do not implement risk logic here")
}

func TestRiskProjectionDoesNotSilentlyMutatePostedBalances(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): risk and projection state can reject or alert but must not silently mutate posted balances")
}

func TestReplayPreservesBalancesOrdersPositionsAndFacts(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): later contract scenario requires an adapter; do not implement replay logic here")
}

func TestReplicatedNodesReachSameStateFromSameCommandStream(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 08 scenario requires adapters; do not implement Raft or replication here")
}

func TestSQLProjectionRebuildsReadModelFromSnapshotAndEvents(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 09 scenario requires an adapter; do not implement projection rebuild here")
}

func TestProjectionRebuildPreservesPostedAndDerivedProvenance(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): projection rebuild may combine posted facts and model inputs but must preserve provenance")
}
