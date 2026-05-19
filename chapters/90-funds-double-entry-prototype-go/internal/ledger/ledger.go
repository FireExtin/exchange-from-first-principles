package ledger

type Entry struct {
	AccountID string
	Asset     string
	Delta     int64
}

type Ledger struct{}

func New() *Ledger {
	panic("TODO: implement")
}

// Apply posts a balanced journal transaction. Every asset must sum to zero
// across all entries. Returns an error if the transaction is unbalanced or
// would produce a negative balance in any non-external account.
func (l *Ledger) Apply(ref string, entries []Entry) error {
	panic("TODO: implement")
}

func (l *Ledger) Balance(accountID, asset string) int64 {
	panic("TODO: implement")
}
