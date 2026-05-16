package types

type Seq int64
type Ref string
type AccountID string
type Asset string
type Amount int64
type CallbackID string
type WithdrawalID string
type ProviderEventID string

type BalanceKey struct {
	AccountID AccountID
	Asset     Asset
}
