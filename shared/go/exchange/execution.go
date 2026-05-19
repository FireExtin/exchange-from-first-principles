package exchange

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type TradeID string
type LiquidityRole string

const (
	RoleMaker LiquidityRole = "maker"
	RoleTaker LiquidityRole = "taker"
)

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
