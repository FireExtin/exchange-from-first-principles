# Change Brief 07: Credit Contract Lab

Date: 2026-05-21

## Summary

Added the first contract lab for the credit, margin, and funding extension:

```text
collateral / borrow / funding schedule / liquidation risk
  != posted ledger truth
```

The lab keeps the project focused on exchange semantics. It is not a Modern
Lending system.

## Key Changes

- Added `chapters/18-credit-and-collateral-accounts-go` as a contract scaffold.
- Added `shared/go/credit` as an optional extension contract package.
- Added TODO tests behind the `credit_contract_todo` build tag.
- Updated the roadmap, documentation map, root README, shared contract docs,
  and agent instructions.
- Kept chapters 19-20 as planned notes.

## Guardrails

- No credit, funding, liquidation, SQL, memory, or adapter implementation was
  added.
- The new contract remains separate from `shared/go/exchange` so the spot
  exchange contract stays focused.
- TODO tests are expected to fail until the project owner implements adapters
  and exercise logic.
- Older change briefs remain historical and were not rewritten.
