# Change Brief 03: Exchange Semantic Version Line

Date: 2026-05-18

## Summary

Reorganized the project around a full exchange semantic version line:

```text
ACID SQL exchange
  -> SQL facts/outbox
  -> single-node memory core
  -> replicated log core
  -> SQL projection consumers
```

The runnable Go funds chapters were moved to appendix prototype chapters
90-93. The main chapter sequence now starts with an exchange-level semantic
contract and keeps business implementation intentionally absent where the
chapter is only a scaffold.

## Key Changes

- Added `shared/go/exchange` as an interface-first exchange semantic contract.
- Added `exchange_contract_todo` tests that intentionally fail until adapters
  and implementations exist.
- Reordered chapter directories to the v0-v4 version line plus domain deep
  dives.
- Preserved existing runnable funds examples as appendix prototypes.
- Updated active docs to describe accounting, reservation, matching,
  executions, positions, risk, projections, caches, and push as one semantic
  surface.

## Guardrails

- No business implementation was added.
- Existing runnable prototype tests remain normal tests.
- Unimplemented exchange behavior is behind an explicit build tag.
- Older change briefs remain historical and were not rewritten.
