//go:build exchange_contract_todo

package exchange

import "testing"

func TestChapter01DepositPostsCustodyAssetAndUserLiability(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 01 scenario requires an adapter; do not implement posting here")
}

func TestChapter02AvailableMovesToLocked(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 scenario requires an adapter; do not implement balance-state transitions here")
}

func TestChapter02WithdrawalRequestMovesAvailableToPending(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 scenario requires an adapter; do not implement withdrawal state here")
}

func TestChapter02FeeRevenueIsPlatformAccountState(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 02 scenario requires an adapter; do not implement fee account state here")
}

func TestChapter03PlaceOrderReservesFunds(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 03 scenario requires an adapter; do not implement reservation here")
}

func TestChapter03CancelOrderReleasesFunds(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 03 scenario requires an adapter; do not implement cancellation or release here")
}

func TestChapter04MatchEmitsExecutionFacts(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 scenario requires an adapter; do not implement matching here")
}

func TestChapter04TradeBalancesUSDAndBTCSeparately(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 scenario requires an adapter; do not implement settlement posting here")
}

func TestChapter04BuyerFeePostsToPlatformRevenue(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 scenario requires an adapter; do not implement fee posting here")
}

func TestChapter04PartialFillReleasesSurplus(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapter 04 scenario requires an adapter; do not implement partial-fill release here")
}

func TestChapter05PlusSameScenariosRunAgainstArchitectureAdapters(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): chapters 05+ require adapters; do not implement SQL, memory, or projection engines here")
}

func TestExecutionFactsUpdatePositions(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): later contract scenario requires an adapter; do not implement position logic here")
}

func TestMarginAndRiskRejectDangerousOrders(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): later contract scenario requires an adapter; do not implement risk logic here")
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
