package main

import (
	"fmt"
	"log"

	"github.com/FireExtin/exchange-from-first-principles/chapters/01-double-entry-ledger-go/internal/ledger"
)

func main() {
	book := ledger.New()
	if err := book.Apply("fund-alice", []ledger.Entry{
		{AccountID: "external", Asset: "USD", Delta: -100},
		{AccountID: "alice", Asset: "USD", Delta: 100},
	}); err != nil {
		log.Fatal(err)
	}
	if err := book.Apply("alice-pays-bob", []ledger.Entry{
		{AccountID: "alice", Asset: "USD", Delta: -100},
		{AccountID: "bob", Asset: "USD", Delta: 100},
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("alice USD=%d bob USD=%d\n", book.Balance("alice", "USD"), book.Balance("bob", "USD"))
}
