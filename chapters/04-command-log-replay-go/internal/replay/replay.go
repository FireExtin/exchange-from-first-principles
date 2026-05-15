package replay

import (
	"errors"
	"fmt"
)

type Command struct {
	Seq       int64
	Type      string
	AccountID string
	Asset     string
	Amount    int64
	Ref       string
}

type Event struct {
	Seq       int64
	Type      string
	AccountID string
	Asset     string
	Amount    int64
	Ref       string
}

type CommandLog struct {
	commands []Command
}

func (l *CommandLog) Append(command Command) {
	l.commands = append(l.commands, command)
}

func (l *CommandLog) Commands() []Command {
	out := make([]Command, len(l.commands))
	copy(out, l.commands)
	return out
}

type Engine struct {
	lastSeq  int64
	balances map[string]int64
	events   []Event
}

type Snapshot struct {
	LastSeq  int64
	Balances map[string]int64
}

func NewEngine() *Engine {
	return &Engine{balances: make(map[string]int64)}
}

func (e *Engine) Apply(command Command) (Event, error) {
	if command.Seq != e.lastSeq+1 {
		return Event{}, fmt.Errorf("sequence gap: got=%d want=%d", command.Seq, e.lastSeq+1)
	}
	if command.Ref == "" || command.AccountID == "" || command.Asset == "" {
		return Event{}, errors.New("ref, account_id, and asset are required")
	}
	if command.Amount <= 0 {
		return Event{}, errors.New("amount must be positive")
	}

	var event Event
	switch command.Type {
	case "credit":
		e.balances[key(command.AccountID, command.Asset)] += command.Amount
		event = Event{Seq: command.Seq, Type: "account_credited", AccountID: command.AccountID, Asset: command.Asset, Amount: command.Amount, Ref: command.Ref}
	case "debit":
		k := key(command.AccountID, command.Asset)
		if e.balances[k] < command.Amount {
			return Event{}, fmt.Errorf("insufficient funds: account=%s asset=%s", command.AccountID, command.Asset)
		}
		e.balances[k] -= command.Amount
		event = Event{Seq: command.Seq, Type: "account_debited", AccountID: command.AccountID, Asset: command.Asset, Amount: command.Amount, Ref: command.Ref}
	default:
		return Event{}, fmt.Errorf("unknown command type: %s", command.Type)
	}

	e.lastSeq = command.Seq
	e.events = append(e.events, event)
	return event, nil
}

func Replay(commands []Command) (*Engine, error) {
	engine := NewEngine()
	for _, command := range commands {
		if _, err := engine.Apply(command); err != nil {
			return nil, err
		}
	}
	return engine, nil
}

func (e *Engine) Balance(accountID, asset string) int64 {
	return e.balances[key(accountID, asset)]
}

func (e *Engine) Events() []Event {
	out := make([]Event, len(e.events))
	copy(out, e.events)
	return out
}

func (e *Engine) Snapshot() Snapshot {
	balances := make(map[string]int64, len(e.balances))
	for key, value := range e.balances {
		balances[key] = value
	}
	return Snapshot{LastSeq: e.lastSeq, Balances: balances}
}

func key(accountID, asset string) string {
	return accountID + ":" + asset
}
