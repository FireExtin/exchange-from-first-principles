# Minimal Exchange Design Paper

## Abstract

This project is a long-running exercise in deriving an exchange system from
small primitives. It starts with the most understandable money primitive:
database ACID transactions over ledger rows. It then moves, step by step, toward
explicit command logs, deterministic state machines, projections, and
replicated execution.

The migration is not about replacing databases because databases are bad. It is
about making the system's truth source, concurrency model, recovery model, and
service boundaries increasingly explicit.

The central thesis:

```text
Modern exchange architecture is not a bag of components.
It is a sequence of semantic migrations.
```

A database-first design starts with storage transactions as truth. Later designs
make ordered commands, events, snapshots, and projections explicit. The precise
technical hinge is that database transactions, in-memory state machines, and
replicated state machines can all present successful business mutations as an
explainable serial history. The database hides that history behind transactions,
locks, indexes, MVCC, and WAL. The state-machine model exposes it directly:
every command receives a sequence number, every state transition has one cause,
and every recovery path is snapshot plus replay.

This paper defines the model, boundaries, invariants, and implementation
chapters. It is intentionally not a complete implementation. The code should be
written chapter by chapter as practice.

## 1. Problem Statement

An exchange-like system has to answer several classes of questions at the same
time:

1. Money correctness:
   - Was value created accidentally?
   - Did a debit happen without a corresponding credit?
   - Are deposits, withdrawals, fees, and settlements explainable?

2. Trading correctness:
   - Why was an order accepted or rejected?
   - Which orders matched, in which order, and at what price?
   - Can the order book be rebuilt from facts?

3. Position and risk correctness:
   - What position does an account have now?
   - What exposure exists under current market prices?
   - Was an order allowed before admission?
   - Did risk keep watching after execution?

4. Recovery correctness:
   - Can we rebuild state after a crash?
   - Can another node apply the same facts and reach the same state?
   - Can incident review stop at a sequence number and inspect the system?

5. Latency correctness:
   - Can the hot path avoid general-purpose storage cost?
   - Can p99/p999 latency be explained instead of guessed?
   - Can backpressure be explicit rather than accidental?

The project should let these questions emerge naturally. Each chapter introduces
one new pressure and one new mechanism.

## 2. Design Principles

### 2.1 Truth Should Have A Shape

The first version treats database rows as truth:

```text
balances, ledger_entries, orders, trades, withdrawals, deposits
```

Later versions treat ordered facts as truth:

```text
command_log, event_log, snapshots
```

The final hot-path model treats the replicated ordered command log as truth:

```text
sequence -> command -> deterministic state transition -> event
```

The database remains important, but its role changes. It becomes a projection,
reporting store, reconciliation target, and compliance archive.

### 2.2 Separate Hot, Warm, And Cold Paths

Hot path:

```text
admitted command -> sequencer -> deterministic state machine -> emitted facts
```

Warm path:

```text
facts + market data -> positions -> exposure -> risk alerts/actions
```

Cold path:

```text
facts -> reports -> reconciliation -> audit -> compliance export
```

Cold path work is not less important. It is simply not allowed to block the hot
path. The system should be explainable without making every explanation
synchronous.

### 2.3 Prefer Explicit Order Over Hidden Contention

For a simple payment workflow, row locks are fine. For a trading hot path,
hidden lock contention and retry behavior can make behavior difficult to
explain.

The state-machine design pays a high infrastructure cost: sequencing, logs,
snapshots, backpressure, failover, and replay. In return, it gets a simple
semantic model:

```text
only one command mutates the core state at a time
```

This is not a claim that every subsystem should be single-threaded. It is a
claim that core business mutation should have a clear ordering boundary.

### 2.4 Ordering Mechanisms Are Different Places To Pay

Pessimistic concurrency control, optimistic concurrency control, and consensus
all answer the same business question: which successful command becomes true
first?

- Pessimistic locks and 2PL order conflicting work before or during execution:
  whoever gets the lock first runs first.
- MVCC, CAS, and optimistic validation order successful work at commit time:
  conflicting speculative work retries or rolls back.
- Raft and Paxos order work through replication: a command can be applied only
  after it occupies an accepted log position.

For an external observer, the successful financial commands should still be
explainable as a serial history. Serializable database isolation gives
equivalence to some serial history. Linearizability adds real-time ordering
constraints. Strict serializability combines both effects.

The project uses this as its core comparison language:

```text
different ordering substrate, same externally explainable business history
```

This is why full ordering can be worth its cost. Once the mutation boundary is
ordered, the upper layer does not reason about scheduling, interleavings, dirty
reads, non-repeatable reads, or phantom writes. It only asks:

```text
if command A runs before command B, what should happen?
```

### 2.5 Facts Before Views

Rows such as `available_balance`, `position_qty`, and `top_of_book` are views.
They are useful and often necessary, but they should be explainable from facts:

```text
deposit accepted
withdrawal requested
order accepted
cash reserved
trade executed
position updated
margin checked
```

When facts and views disagree, the system needs a reconciliation story. That is
why every chapter should prefer append-only facts first, then materialized
views.

## 3. Stage One: Database ACID As Source Of Truth

The earliest system can be built with SQL transactions:

```text
request
  -> begin transaction
  -> read rows
  -> validate
  -> write rows
  -> commit
```

This is the right starting point for:

- double-entry ledger;
- basic spot settlement;
- deposit callbacks;
- withdrawal state transitions;
- admin corrections;
- reconciliation reports.

The main invariant is conservation of value:

```text
sum(debits) == sum(credits)
```

External money movement is modeled explicitly:

```text
deposit: external source -> user account
withdrawal: user account -> external sink
fee: user account -> fee account
```

The database gives atomicity and durability. It is also easy to inspect. This is
why the project begins here.

### 3.1 Limits Of The DB-First Model

The DB-first model becomes harder when the hot path needs deterministic order
and low variance.

The database can tell us the committed result. It does not always give a clean
business story:

- Which command caused this final balance?
- Why did this request win the race?
- What exact state existed after operation `N`?
- Can another process replay history without trusting current rows?
- Can the matching engine avoid storage work during order admission?

These are not database defects. They are signals that the project has reached
the next stage.

## 4. Stage Two: Command Log Inside The Database

The bridge stage keeps SQL but introduces facts:

```text
request
  -> begin transaction
  -> insert command_log
  -> insert event_log
  -> update view rows
  -> commit
```

The source of explanation begins to move:

```text
truth candidate = command_log + event_log
view = balances/orders/positions/reports
```

This stage is important because it gives product engineers and trading engineers
a common language. A payment engineer can still rely on SQL. A trading engineer
can start thinking in replayable facts.

### 4.1 Outbox And Projection

A database-backed outbox can publish facts to downstream consumers:

```text
transaction commits rows and outbox event
background publisher reads unsent events
publisher sends to message bus
publisher marks event as sent
```

This is not the final hot-path architecture, but it teaches the right lessons:

- message publication must be tied to state change;
- duplicate messages are normal;
- consumers must be idempotent;
- projections may lag behind facts;
- reconciliation is part of the design, not an afterthought.

## 5. Stage Three: Deterministic State Machine

The trading hot path moves into a state machine:

```text
command
  -> sequence number
  -> deterministic apply()
  -> event
  -> snapshot/replay boundary
```

The state machine owns core mutable state:

- order book;
- account reservation state;
- accepted/rejected order state;
- execution events;
- minimal position hooks.

It should not call databases, remote services, or slow dependencies inside the
state transition.

### 5.1 Why This Can Be Equivalent To DB ACID

A serializable database history can be represented as:

```text
State_0 --tx_1--> State_1 --tx_2--> State_2 --tx_3--> State_3
```

A state-machine history can be represented as:

```text
State_0 --cmd_1--> State_1 --cmd_2--> State_2 --cmd_3--> State_3
```

If the same operations run in the same order and the transition function is
deterministic, both models produce the same final state.

The engineering difference is where the ordering is enforced:

- DB model: ordering is hidden inside storage concurrency control.
- State-machine model: ordering is explicit before business mutation.

The state-machine model looks more expensive because it requires sequencing,
logs, snapshots, and replica management. But its business semantics are simpler:
there is no race inside the core mutation boundary.

The strongest proof is to keep the transition contract stable across versions.
The surrounding substrate may move from SQL transactions to an in-memory writer
to a replicated log, but the business transition should still look like:

```text
ordered input + deterministic apply -> new state + emitted facts
```

If the same scenario suite passes against each version, the architecture has
changed while the external semantics have not.

### 5.2 Determinism Requirements

The state machine must avoid nondeterminism:

- no wall-clock reads inside mutation logic;
- no random number generation inside mutation logic;
- no remote calls inside mutation logic;
- no unordered map iteration that affects emitted facts;
- no floating-point money math;
- no dependency on local thread scheduling.

Inputs such as timestamps, ids, marks, and risk parameters must be part of the
command stream or configuration snapshot.

## 6. Stage Four: Replicated State Machine

Once one state machine is understandable, replication becomes the next pressure.

Target shape:

```text
client command
  -> replicated log / Aeron Cluster / Raft-like ordering
  -> node A apply(command)
  -> node B apply(command)
  -> node C apply(command)
  -> same state if same snapshot and same log
```

This stage is where Aeron Cluster or a Raft-like layer becomes relevant. The
project should not implement consensus from scratch. It should instead define
clear contracts at the boundary:

- command encoding;
- session identity;
- sequence position;
- snapshot format;
- event emission;
- backpressure handling;
- replay start point.

The learning goal is not to become a consensus implementer. The learning goal is
to understand what consistency service the trading system is asking from the
replicated log.

## 7. Domain Model

### 7.1 Accounts And Ledger

The ledger tracks value movement. It should support:

- account id;
- currency/asset;
- debit account;
- credit account;
- amount;
- reason;
- external reference;
- idempotency key;
- event sequence.

Important invariant:

```text
every internal movement has equal debit and credit
```

Deposits and withdrawals cross the system boundary, so their external references
must be explicit.

### 7.2 Orders And Executions

Orders are intents. Executions are facts.

An accepted order means:

```text
the system accepted an intent under known constraints
```

An execution means:

```text
the system created a trade fact
```

The system should not confuse the two. Positions are derived from executions,
not from open orders.

### 7.3 Position

Position state is the account's exposure by instrument:

- net quantity;
- average entry price;
- realized PnL;
- unrealized PnL;
- open notional;
- mark-price dependency;
- margin dependency.

Position management should be replayable from execution reports. If the trading
desk routes to external exchanges, position management may be more important
than local matching.

### 7.4 Desk Pre-Trade Risk

Pre-trade risk decides whether a command is allowed to enter the sequenced core.

Examples:

- account enabled;
- instrument enabled;
- notional below limit;
- price within band;
- order size below limit;
- enough available balance;
- enough margin;
- account or strategy not killed.

This component is close to order entry. It must be fast and deterministic.

### 7.5 Risk Cluster

The risk cluster watches continuous exposure after facts are emitted.

Inputs:

- executions;
- positions;
- mark prices;
- deposits;
- withdrawals;
- funding;
- manual adjustments;
- account configuration.

Outputs:

- alerts;
- kill-switch recommendations;
- margin pressure;
- exposure reports;
- reconciliation signals.

This is not the same as desk pre-trade risk. Pre-trade risk blocks bad commands.
The risk cluster detects changing exposure over time.

### 7.6 Margin

The first margin model should be deliberately small.

Spot:

```text
available_quote = quote_balance - frozen_quote
available_base = base_balance - frozen_base

buy_required_quote = price * quantity + fee_buffer
sell_required_base = quantity
```

Linear contract placeholder:

```text
notional = abs(position_qty) * mark_price

if position_qty > 0:
  unrealized_pnl = position_qty * (mark_price - entry_price)
else:
  unrealized_pnl = abs(position_qty) * (entry_price - mark_price)

initial_margin = notional * initial_margin_rate
maintenance_margin = notional * maintenance_margin_rate
equity = wallet_balance + unrealized_pnl
available = equity - initial_margin - frozen_margin - fee_buffer
```

Liquidation placeholder:

```text
equity <= maintenance_margin + liquidation_fee_buffer
```

The goal is not to model every exchange rule. The goal is to make the dependency
chain explicit:

```text
fills -> position -> mark dependent PnL -> equity -> margin -> risk action
```

## 8. Failure Model

### 8.1 DB Stage Failures

Expected failures:

- duplicate callback;
- transaction rollback;
- deadlock retry;
- idempotency conflict;
- partial external payment state;
- reconciliation mismatch.

The design response:

- idempotency keys;
- unique constraints;
- explicit status transitions;
- append-only ledger entries;
- reconciliation jobs.

### 8.2 Command-Log Stage Failures

Expected failures:

- command appended but projection lagged;
- event sent twice;
- consumer processed twice;
- snapshot too old;
- replay position unknown.

The design response:

- event ids;
- sequence numbers;
- idempotent consumers;
- replay checkpoints;
- snapshot metadata.

### 8.3 State-Machine Stage Failures

Expected failures:

- process crash;
- leader failover;
- backpressure;
- slow projection;
- divergent state due to nondeterminism;
- snapshot/replay bug.

The design response:

- deterministic transition tests;
- snapshot plus replay;
- no external calls in mutation;
- explicit backpressure;
- reproducible input streams;
- sequence-based incident review.

## 9. Implementation Chapters

The project should be implemented in small chapters.

1. Double-entry ledger in Go
   - prove conservation of value;
   - model external deposit/withdrawal boundaries.

2. Spot trade with DB-shaped transaction in Go
   - buyer/seller settlement;
   - rollback and insufficient funds;
   - simple rows as truth.

3. Wallet deposit/withdrawal in Go
   - idempotency;
   - duplicate callback;
   - status transition.

4. Command log and replay in Go
   - ordered commands;
   - replay;
   - snapshot;
   - equivalence with serial DB history.

5. In-memory single-writer in Java
   - remove DB lock from hot path;
   - introduce fixed event contracts;
   - observe allocation behavior.

6. Matching engine
   - price-time priority;
   - top-of-book;
   - execution facts.

7. Position manager
   - execution reports to position state;
   - account and instrument exposure.

8. Margin model
   - spot available balance;
   - linear contract placeholder;
   - mark-driven equity.

9. Desk pre-trade risk
   - synchronous admission checks.

10. Risk cluster projection
   - event and mark-price driven exposure.

11. Aeron/Raft replicated state-machine boundary in Java
   - transport and ordering;
   - backpressure;
   - replay and failover boundary.

12. OMS, ledger, compliance, and path split in Java/Go
   - local order state;
   - service edge APIs;
   - reporting, reconciliation, and compliance projections.

13. Cache coherence and market state
   - user/account/permission caches;
   - instrument and mark-price caches;
   - external exchange state caches.

14. Market-data and execution push
   - public snapshots and deltas;
   - private execution reports;
   - gap detection and recovery.

15. Rust hot path experiment
   - optional, later;
   - only after Java/Go chapters are useful.

16. Low-latency runtime and networking
   - zero-allocation work;
   - profiling;
   - OS and network tuning.

Each chapter should have one runnable demo or test, but the demo should be
small enough to rewrite from memory.

## Project Narrative

## 11. Non-Goals

This project should not start by implementing:

- full exchange liquidation rules;
- portfolio margin;
- production wallet custody;
- KYC/AML workflows;
- consensus from scratch;
- DPDK/XDP production networking;
- a full matching venue with all order types;
- a complete product backend.

Those can be researched later. The first goal is a clean chain of reasoning and
small executable demonstrations.

## 12. Success Criteria

The project succeeds if it can answer these questions with code and notes:

- Why does a ledger need double-entry?
- Why does settlement need an atomic boundary?
- Why do callbacks need idempotency?
- Why do facts become more important than current rows?
- Why can DB serial history and state-machine command history be equivalent?
- Why does total order simplify business semantics despite engineering cost?
- Why are matching, position, pre-trade risk, and risk cluster separate?
- Why does margin depend on mark prices and positions rather than just orders?
- Why should hot, warm, and cold paths have different responsibilities?
- Why does replay make incident review clearer?

If these questions can be answered, the repo becomes useful both as engineering
practice and as a working demonstration.

---

## 中文

### 摘要

本项目是一个长期的练习：从小型原语推导交易所系统。它从最易理解的钱
原语开始：账本行上的数据库 ACID 事务。然后逐步走向有序命令日志和确定性
复制状态机。

迁移不是关于因为数据库不好就替换数据库。它是将热路径真相源从隐式
存储层排序移动到显式业务排序。

核心论点：

```text
DB ACID 通过隔离并发存储效果给出正确性。
有序状态机通过从热路径移除并发业务效果给出正确性。
```

两者都可以产生串行历史。数据库将串行历史隐藏在事务、锁、索引、MVCC
和 WAL 之后。状态机模型直接暴露串行历史：每条命令收到一个序列号，
每个状态转换有一个原因，每个恢复路径是快照加重放。

本文定义了模型、边界、不变量和实现章节。它故意不是一个完整实现。
代码应该逐章作为练习来写。

### 1. 问题陈述

一个类交易所系统必须同时回答几类问题：

1. 资金正确性：
   - 价值是否被意外创造？
   - 借方是否有对应的贷方？
   - 入金、出金、费用和结算可解释吗？

2. 交易正确性：
   - 订单为什么被接受或拒绝？
   - 哪些订单成交，按什么顺序，以什么价格？
   - 订单簿可以从事实重建吗？

3. 仓位和风险正确性：
   - 账户现在有什么仓位？
   - 在当前市场价格下存在什么敞口？
   - 订单准入前被允许了吗？
   - 执行后风控持续监视了吗？

4. 恢复正确性：
   - 崩溃后能重建状态吗？
   - 另一个节点能应用相同事实达到相同状态吗？
   - 事件审查能在某个序列号停下检查系统吗？

5. 延迟正确性：
   - 热路径能避免通用存储成本吗？
   - p99/p999 延迟能解释而不是猜测吗？
   - 背压能是显式的而不是偶然的吗？

项目应该让这些问题自然浮现。每章引入一个新的压力和一个新的机制。

### 12. 成功标准

如果项目能用代码和笔记回答这些问题，它就成功了：

- 为什么账本需要双分录？
- 为什么结算需要原子边界？
- 为什么回调需要幂等性？
- 为什么事实变得比当前行更重要？
- 为什么 DB 串行历史和状态机命令历史可以等价？
- 为什么全排序能简化业务语义尽管有工程成本？
- 为什么撮合、仓位、下单前风控和风控集群是分开的？
- 为什么保证金依赖标记价格和仓位而不仅仅是订单？
- 为什么热、温、冷路径应该有不同职责？
- 为什么重放使事件审查更清晰？

如果这些问题能回答，仓库就既有工程练习也有工作演示的价值。
