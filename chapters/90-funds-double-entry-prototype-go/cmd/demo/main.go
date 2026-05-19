package main

import (
	"github.com/FireExtin/exchange-from-first-principles/chapters/90-funds-double-entry-prototype-go/internal/ledger"
)

// Demo shows two balanced transactions:
//   1. external funds alice's account by 100 USD
//   2. alice pays bob 100 USD
// After both, alice USD = 0 and bob USD = 100.
func main() {
	_ = ledger.New()
	panic("TODO: implement ledger to run this demo")
}
