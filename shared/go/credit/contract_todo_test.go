//go:build credit_contract_todo

package credit

import "testing"

func TestCollateralPledgeMovesAvailableToCollateralAccount(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): collateral pledge must move available funds into an explicit collateral account")
}

func TestBorrowDrawdownCreatesBorrowLiabilityAndUsableFunds(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): borrow drawdown must create borrow liability and usable funds through explicit entries")
}

func TestMarkPriceChangeDoesNotCreateLedgerEntries(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): mark price changes are model inputs and must not create ledger entries by themselves")
}

func TestFundingScheduleDoesNotPostUntilAccrualEvent(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): funding schedules are prospective state until an explicit accrual event posts entries")
}

func TestRepaymentReducesBorrowLiabilityThroughExplicitEntries(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): repayment must reduce borrow liability through explicit journal entries")
}

func TestLiquidationSettlesCollateralLiabilityAndFeeThroughEntries(t *testing.T) {
	t.Fatal("TODO(credit_contract_todo): liquidation must settle collateral, borrow liability, and fees through explicit entries")
}
