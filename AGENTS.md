# AGENTS.md

This file gives working rules for AI coding agents in this repository.

The project is an explanation-first systems repo. Preserve that shape: small
runnable exercises, explicit contracts, and documentation that explains why the
next architecture step becomes necessary.

## Source Of Truth

- `README.md` is the project front door.
- `docs/README.md` is the canonical documentation map.
- `docs/00-goal.md` defines the project goal and documentation hierarchy.
- `docs/01-core-principles.md` owns semantic contracts, command/event rules,
  replay, recovery, and test principles.
- `docs/07-chapter-roadmap.md` owns chapter order and chapter status.
- `docs/02-design-paper.md` is the full design narrative.
- Chapter READMEs own chapter-local behavior, pressure, and run commands.
- `shared/README.md` and `integration-tests/README.md` own cross-chapter
  semantic contracts.
- `docs/history/change_brief_*.md` files are append-only historical records.
  Add a new brief for a meaningful documentation or architecture
  reorganization; do not rewrite old briefs to describe the latest state.

When documents disagree, prefer the more specific source:

1. chapter README for chapter behavior;
2. `docs/07-chapter-roadmap.md` for chapter status;
3. `docs/README.md` for document organization;
4. `docs/02-design-paper.md` for project-level design narrative.

## Language

- Active project documentation should use mirrored English/Chinese sections.
- This `AGENTS.md` is English-only.
- Code, identifiers, package names, and test names should remain English.

## Toolchain

Runnable parts currently expect:

- Go 1.22 or newer;
- Java 21, tested with Azul Zulu 21;
- Gradle for the Java/Aeron chapter;
- Rust stable for chapter 15.

## Test Commands

Use the smallest relevant test first, then widen if the change crosses
contracts or chapter boundaries.

From the repo root:

```bash
make test-go
make test-rust
make test-java
```

Useful focused commands:

```bash
cd chapters/01-double-entry-ledger-go && go run ./cmd/demo
cd chapters/03-wallet-deposit-withdrawal-go && go test ./...
cd chapters/03-wallet-deposit-withdrawal-go && go test -tags reconciliation_lab_todo ./internal/reconciliation
cd chapters/11-replicated-state-machine-aeron-java && gradle --no-daemon clean test
cd chapters/15-rust-hot-path && cargo test
go test ./integration-tests/...
```

The reconciliation lab tests are intentionally behind the
`reconciliation_lab_todo` build tag and are expected to fail until the lab is
implemented.

For exercise and lab chapters, agents may add or refine interfaces, contracts,
fixtures, and tests, but should leave the core exercise implementation for the
project owner to write. Keep TODO boundaries explicit and avoid filling in the
solution unless the user specifically asks for it.

## Architecture Boundaries

- Go owns service edges: funds workflows, idempotency, callbacks,
  reconciliation, adapters, tools, and integration tests.
- Java owns the primary trading hot path: deterministic state application,
  matching, sequencing boundaries, snapshots, and replay rules.
- Aeron owns replicated-log integration concerns, not business rules.
- Rust is exploratory: useful for measuring a clean hot-path contract and FFI
  boundary, but not the active main implementation track.

Do not add broad frameworks or infrastructure before a chapter explains the
pressure that makes them necessary.

## Documentation Rules

- Keep root `README.md` short and navigational.
- Put the full documentation catalog in `docs/README.md`.
- Put chapter-specific run commands in the chapter README.
- Update `docs/07-chapter-roadmap.md` when a chapter status changes.
- Update `docs/README.md` when adding, removing, or reclassifying Markdown docs.
- Prefer adding a new `docs/history/change_brief_NN.md` for substantial
  reorganizations.
- Do not duplicate long design explanations across many files; link to the
  source-of-truth document instead.

## Chapter README Contract

Chapter READMEs must stay practical and local. They should answer:

1. what pressure this chapter introduces;
2. what status the chapter is in;
3. how to run it, or explicitly say that no runnable implementation exists yet;
4. what belongs in this chapter and what is intentionally out of scope.

Use mirrored English/Chinese sections. Keep the wording concise. Link to
`docs/01-core-principles.md`, `docs/02-design-paper.md`, or
`docs/07-chapter-roadmap.md` for broad theory instead of repeating it.

## Code Rules

- Preserve chapter independence unless a shared contract is deliberately being
  introduced.
- Shared contracts may be consumed by chapters, but shared code must not import
  chapter code.
- Integration tests should assert externally visible semantics, not internal
  data structures.
- Keep exercises small. A chapter should prove one pressure clearly before it
  grows another subsystem.
- Avoid unrelated refactors while changing docs or a single chapter.

## Git And Generated Files

- Do not commit build outputs, caches, or local environment files.
- Do not edit unrelated user changes.
- Before finishing, run `git diff --check`.
- For docs-only changes, link checks and `git diff --check` are usually enough.
  For code or contract changes, run the relevant test target.
