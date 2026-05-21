# Change Brief 05: Posted Facts And Derived State

Date: 2026-05-21

## Summary

Clarified a project-wide boundary inspired by Modern Treasury's lending ledger
article:

```text
posted financial facts
  != operational state
  != derived or prospective model state
```

The ledger owns posted historical financial facts. Orders, reservations,
positions, marks, margin requirements, risk views, and projections may be
important state, but they should not pretend to be posted ledger truth unless an
explicit business event posts journal entries.

## Key Changes

- Added a core principle for posted facts versus derived/prospective state.
- Split truth-source migration into posted facts, operational state, and
  derived/prospective state.
- Clarified that positions can be first-class state while unrealized PnL and
  mark-based risk remain derived views.
- Updated SQL/outbox/projection chapter notes so consumers preserve provenance
  and do not infer ledger truth from hidden row changes.
- Added TODO contract scenarios for settlement boundaries, mark/PnL behavior,
  risk/projection mutation boundaries, and projection provenance.

## Guardrails

- No business implementation was added.
- No SQL, memory, risk, or projection adapter was added.
- Unimplemented exchange behavior remains behind the `exchange_contract_todo`
  build tag.
- Older change briefs remain historical and were not rewritten.
