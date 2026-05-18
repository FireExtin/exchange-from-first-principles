//go:build exchange_contract_todo

package exchange

import "testing"

func TestDepositCreatesCustodyAssetAndUserLiability(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement the engine here")
}

func TestOrderPlacementMovesAvailableToLocked(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement reservation here")
}

func TestCancellationReleasesLockedToAvailable(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement cancellation here")
}

func TestTradeExecutionBalancesUSDAndBTCSeparately(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement matching or posting here")
}

func TestBuyerFeePostsToFeeRevenue(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement fee posting here")
}

func TestPartialFillReleasesUnusedLockedFunds(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement release logic here")
}

func TestExecutionFactsUpdatePositions(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement position logic here")
}

func TestMarginAndRiskRejectDangerousOrders(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement risk logic here")
}

func TestReplayPreservesBalancesOrdersPositionsAndFacts(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement replay logic here")
}

func TestReplicatedNodesReachSameStateFromSameCommandStream(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires adapters; do not implement Raft or replication here")
}

func TestSQLProjectionRebuildsReadModelFromSnapshotAndEvents(t *testing.T) {
	t.Fatal("TODO(exchange_contract_todo): contract scenario requires an adapter; do not implement projection rebuild here")
}
