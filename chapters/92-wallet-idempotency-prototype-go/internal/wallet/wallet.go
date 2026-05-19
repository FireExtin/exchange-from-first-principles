package wallet

// WithdrawalStatus tracks whether a withdrawal is pending or complete.
type WithdrawalStatus string

const (
	WithdrawalRequested WithdrawalStatus = "requested"
	WithdrawalConfirmed WithdrawalStatus = "confirmed"
)

type Withdrawal struct {
	ID        string
	AccountID string
	Asset     string
	Amount    int64
	Status    WithdrawalStatus
}

// Processor is the in-memory wallet engine. It enforces idempotency on
// deposits (via callbackID) and withdrawal confirmations (via providerEventID),
// and enforces the two-phase withdrawal lifecycle (requested → confirmed).
type Processor struct{}

func NewProcessor() *Processor {
	panic("TODO: implement")
}

// HandleDeposit credits the account. Returns (false, nil) if callbackID was
// already seen (idempotent). Returns (false, err) on invalid input.
func (p *Processor) HandleDeposit(callbackID, accountID, asset string, amount int64) (bool, error) {
	panic("TODO: implement")
}

// RequestWithdrawal reserves funds and records the withdrawal as requested.
// Returns (false, nil) if withdrawalID already exists (idempotent dedup).
func (p *Processor) RequestWithdrawal(withdrawalID, accountID, asset string, amount int64) (bool, error) {
	panic("TODO: implement")
}

// ConfirmWithdrawal transitions the withdrawal to confirmed.
// Returns (false, nil) if providerEventID was already seen (idempotent).
func (p *Processor) ConfirmWithdrawal(providerEventID, withdrawalID string) (bool, error) {
	panic("TODO: implement")
}

// Transfer moves funds from one account to another atomically.
func (p *Processor) Transfer(from, to, asset string, amount int64) error {
	panic("TODO: implement")
}

func (p *Processor) Balance(accountID, asset string) int64 {
	panic("TODO: implement")
}

func (p *Processor) Withdrawal(id string) (Withdrawal, bool) {
	panic("TODO: implement")
}
