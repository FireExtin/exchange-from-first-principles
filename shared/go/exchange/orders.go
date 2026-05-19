package exchange

import "github.com/FireExtin/exchange-from-first-principles/shared/go/types"

type OrderID string
type Instrument string
type Side string
type OrderStatus string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"

	OrderAccepted        OrderStatus = "accepted"
	OrderRejected        OrderStatus = "rejected"
	OrderPartiallyFilled OrderStatus = "partially_filled"
	OrderFilled          OrderStatus = "filled"
	OrderCancelled       OrderStatus = "cancelled"
)

type OrderCommand struct {
	OrderID    OrderID
	Instrument Instrument
	Side       Side
	Price      types.Amount
	Quantity   types.Amount
	MaxFee     types.Amount
}

type OrderEvent struct {
	OrderID         OrderID
	Instrument      Instrument
	Status          OrderStatus
	ReservedAsset   types.Asset
	ReservedAmount  types.Amount
	ReleasedAsset   types.Asset
	ReleasedAmount  types.Amount
	RemainingAmount types.Amount
}

type OrderView struct {
	OrderID    OrderID
	AccountID  types.AccountID
	Instrument Instrument
	Status     OrderStatus
	Remaining  types.Amount
}
