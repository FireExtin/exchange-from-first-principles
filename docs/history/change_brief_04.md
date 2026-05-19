# Change Brief 04: Business Semantic Ramp

Date: 2026-05-18

## Summary

Moved the reader entry path from a large exchange semantic contract into a
business semantic ramp:

```text
custody / user liability
  -> available / locked / pending
  -> order reservation / cancel release
  -> minimal match / settlement
  -> ACID SQL exchange
  -> SQL facts/outbox
  -> memory core
  -> replicated log
  -> SQL projections
```

The full exchange contract still exists, but readers now meet it one business
idea at a time before the architecture migration begins.

## Key Changes

- Reordered main chapters into 01-04 business semantics, 05-09 architecture
  migration, and 10-17 domain/runtime deep dives.
- Moved the ACID SQL exchange chapter to chapter 05 so it composes earlier
  semantic pieces instead of introducing everything at once.
- Kept runnable Go funds examples in appendix chapters 90-93.
- Updated docs, README files, and run commands to the new paths.
- Regrouped TODO contract tests around the first four semantic chapters.

## Guardrails

- No business implementation was added.
- README-only chapters remain README-only.
- Existing skeleton/runnable chapters keep their existing code shape.
- Unimplemented exchange behavior remains behind the `exchange_contract_todo`
  build tag.
- Older change briefs remain historical and were not rewritten.
