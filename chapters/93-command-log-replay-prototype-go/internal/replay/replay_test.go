package replay

import "testing"

func TestReplayRestoresStateFromCommandLog(t *testing.T) {
	var log CommandLog
	log.Append(Command{Seq: 1, Type: "credit", AccountID: "alice", Asset: "USD", Amount: 1000, Ref: "funding"})
	log.Append(Command{Seq: 2, Type: "debit", AccountID: "alice", Asset: "USD", Amount: 250, Ref: "withdrawal"})

	engine, err := Replay(log.Commands())
	if err != nil {
		t.Fatalf("replay: %v", err)
	}
	if got := engine.Balance("alice", "USD"); got != 750 {
		t.Fatalf("balance = %d, want 750", got)
	}
	if got := len(engine.Events()); got != 2 {
		t.Fatalf("events = %d, want 2", got)
	}
}

func TestReplayRejectsSequenceGaps(t *testing.T) {
	_, err := Replay([]Command{
		{Seq: 1, Type: "credit", AccountID: "alice", Asset: "USD", Amount: 1000, Ref: "funding"},
		{Seq: 3, Type: "debit", AccountID: "alice", Asset: "USD", Amount: 250, Ref: "withdrawal"},
	})
	if err == nil {
		t.Fatal("expected sequence gap")
	}
}

func TestSequencedStateMachineMatchesSerializableHistory(t *testing.T) {
	commands := []Command{
		{Seq: 1, Type: "credit", AccountID: "alice", Asset: "USD", Amount: 1000, Ref: "funding"},
		{Seq: 2, Type: "debit", AccountID: "alice", Asset: "USD", Amount: 250, Ref: "withdrawal"},
		{Seq: 3, Type: "credit", AccountID: "bob", Asset: "USD", Amount: 250, Ref: "settlement"},
	}

	stateMachine, err := Replay(commands)
	if err != nil {
		t.Fatalf("state machine replay: %v", err)
	}
	dbModel, err := applyAsSerializableTransactions(commands)
	if err != nil {
		t.Fatalf("serializable DB model: %v", err)
	}

	for key, want := range dbModel {
		if got := stateMachine.Snapshot().Balances[key]; got != want {
			t.Fatalf("balance %s = %d, want %d", key, got, want)
		}
	}
}

func applyAsSerializableTransactions(commands []Command) (map[string]int64, error) {
	balances := make(map[string]int64)
	for _, command := range commands {
		next := make(map[string]int64, len(balances)+1)
		for key, value := range balances {
			next[key] = value
		}

		key := command.AccountID + ":" + command.Asset
		switch command.Type {
		case "credit":
			next[key] += command.Amount
		case "debit":
			if next[key] < command.Amount {
				return nil, errInsufficientFunds
			}
			next[key] -= command.Amount
		default:
			return nil, errUnknownCommand
		}
		balances = next
	}
	return balances, nil
}

var (
	errInsufficientFunds = testError("insufficient funds")
	errUnknownCommand    = testError("unknown command")
)

type testError string

func (e testError) Error() string {
	return string(e)
}
