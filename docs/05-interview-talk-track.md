# Interview Talk Track

The project is not about using a language for its own sake.

The design point is:

> A trading system needs a narrow truth boundary. I put command sequencing,
> append-only logging, deterministic state transition, and replay in one core.
> Everything else is an adapter or projection.

For trading-system interviews, I show that core in Java because the market
expects Java low-latency literacy: GC behavior, off-heap buffers, SBE/Agrona,
Aeron boundaries, profiling, and live coding under familiar Java collections.

For product/asset-system interviews, I show the service edges in Go: gateway,
ledger simulation, idempotency, reconciliation, and operational tooling.

The same commands, events, and scenarios are reused in both tracks. That keeps
the discussion about system behavior rather than language taste.

## Galaxy / Citi / Bullish / BTSE Framing

I would frame the Java side as:

> The core is a deterministic event-driven state machine. Commands are ordered,
> logged, applied once, and replayed into the same state. The performance work is
> about controlling allocation, avoiding avoidable GC pressure, using stable
> data structures, and keeping the hot path observable through JFR or
> async-profiler.

Focus points:

- order book and top-of-book aggregation;
- position updates from execution reports;
- replay from a command/event offset;
- SBE/Agrona buffer handling;
- Aeron Cluster as ordering/replication infrastructure, not business logic;
- GC and allocation tradeoffs.

## Infini Framing

I would frame the Go side as:

> The service layer handles business correctness around money movement:
> idempotent callbacks, ledger entries, reconciliation, frozen/in-flight funds,
> and clear operational recovery. I would keep concurrency explicit and avoid
> scattering goroutines across business ordering boundaries.

Focus points:

- idempotency keys and duplicate callback handling;
- ledger entry model and reconciliation reports;
- context cancellation and timeout boundaries;
- pprof CPU/heap/block/mutex basics;
- Kafka offset handling and exactly-once effects through idempotent writes.
