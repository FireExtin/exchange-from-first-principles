package wallet

import (
	"errors"
	"fmt"
)

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

type Processor struct {
	balances          map[string]int64
	seenDepositIDs    map[string]struct{}
	withdrawalByID    map[string]Withdrawal
	seenConfirmations map[string]struct{}
}

func NewProcessor() *Processor {
	return &Processor{
		balances:          make(map[string]int64),
		seenDepositIDs:    make(map[string]struct{}),
		withdrawalByID:    make(map[string]Withdrawal),
		seenConfirmations: make(map[string]struct{}),
	}
}

func (p *Processor) HandleDeposit(callbackID, accountID, asset string, amount int64) (bool, error) {
	if callbackID == "" {
		return false, errors.New("callback_id is required")
	}
	if amount <= 0 {
		return false, errors.New("amount must be positive")
	}
	if _, ok := p.seenDepositIDs[callbackID]; ok {
		return false, nil
	}

	p.balances[key(accountID, asset)] += amount
	p.seenDepositIDs[callbackID] = struct{}{}
	return true, nil
}

func (p *Processor) RequestWithdrawal(withdrawalID, accountID, asset string, amount int64) (bool, error) {
	if withdrawalID == "" {
		return false, errors.New("withdrawal_id is required")
	}
	if amount <= 0 {
		return false, errors.New("amount must be positive")
	}
	if _, ok := p.withdrawalByID[withdrawalID]; ok {
		return false, nil
	}
	k := key(accountID, asset)
	if p.balances[k] < amount {
		return false, fmt.Errorf("insufficient funds: account=%s asset=%s", accountID, asset)
	}

	p.balances[k] -= amount
	p.withdrawalByID[withdrawalID] = Withdrawal{
		ID:        withdrawalID,
		AccountID: accountID,
		Asset:     asset,
		Amount:    amount,
		Status:    WithdrawalRequested,
	}
	return true, nil
}

func (p *Processor) ConfirmWithdrawal(providerEventID, withdrawalID string) (bool, error) {
	if providerEventID == "" {
		return false, errors.New("provider_event_id is required")
	}
	if _, ok := p.seenConfirmations[providerEventID]; ok {
		return false, nil
	}
	withdrawal, ok := p.withdrawalByID[withdrawalID]
	if !ok {
		return false, fmt.Errorf("unknown withdrawal: %s", withdrawalID)
	}

	withdrawal.Status = WithdrawalConfirmed
	p.withdrawalByID[withdrawalID] = withdrawal
	p.seenConfirmations[providerEventID] = struct{}{}
	return true, nil
}

func (p *Processor) Balance(accountID, asset string) int64 {
	return p.balances[key(accountID, asset)]
}

func (p *Processor) Withdrawal(id string) (Withdrawal, bool) {
	withdrawal, ok := p.withdrawalByID[id]
	return withdrawal, ok
}

func key(accountID, asset string) string {
	return accountID + ":" + asset
}
