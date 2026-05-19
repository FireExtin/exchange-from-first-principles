package replay

// Command is a generic single-asset ledger command used in this prototype.
// Type is "credit" or "debit". Seq must be strictly sequential starting at 1.
type Command struct {
	Seq       int64
	Type      string
	AccountID string
	Asset     string
	Amount    int64
	Ref       string
}

// Event is the fact produced by applying a Command.
// Type is "account_credited" or "account_debited".
type Event struct {
	Seq       int64
	Type      string
	AccountID string
	Asset     string
	Amount    int64
	Ref       string
}

// CommandLog is an append-only in-memory command store.
type CommandLog struct{}

func (l *CommandLog) Append(command Command) {
	panic("TODO: implement")
}

func (l *CommandLog) Commands() []Command {
	panic("TODO: implement")
}

// Engine applies commands one at a time and tracks resulting balances and
// events. Apply must be called in strict sequence order (1, 2, 3...).
type Engine struct{}

// Snapshot captures the full state at a point in time, allowing replay
// to resume from a known position rather than from command 1.
type Snapshot struct {
	LastSeq  int64
	Balances map[string]int64
}

func NewEngine() *Engine {
	panic("TODO: implement")
}

func (e *Engine) Apply(command Command) (Event, error) {
	panic("TODO: implement")
}

// Replay rebuilds an Engine by applying all commands from the log in order.
// Demonstrates that deterministic replay reaches the same state as serial
// execution.
func Replay(commands []Command) (*Engine, error) {
	panic("TODO: implement")
}

func (e *Engine) Balance(accountID, asset string) int64 {
	panic("TODO: implement")
}

func (e *Engine) Events() []Event {
	panic("TODO: implement")
}

func (e *Engine) Snapshot() Snapshot {
	panic("TODO: implement")
}
