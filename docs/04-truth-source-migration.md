# Truth Source Migration

[English](#english) · [中文](#中文)

## English

This note is the project spine:

> Start with database ACID as the source of truth, then migrate truth into an
> ordered fact log and a deterministic replicated state machine.

The point is not to argue that databases are bad. The database model is the
right starting point because it is simple, durable, queryable, and familiar. The
migration happens only when the hot path needs properties that row-state
transactions provide indirectly or expensively: explicit ordering, replay,
deterministic recovery, lower tail latency, and multi-replica consistency.

## The Starting Point: Database Rows Are Truth

The first version should be boring:

```text
request
  -> begin transaction
  -> read balances/orders
  -> validate business rules
  -> update balances/orders/trades/ledger_entries
  -> commit
```

In this model, the source of truth is the current database state:

- `balances`
- `orders`
- `trades`
- `ledger_entries`
- `withdrawals`
- `deposits`

Atomicity comes from the storage layer. The application asks the database to
protect invariants while multiple requests are running at the same time. The
database uses locks, MVCC, WAL, indexes, and commit ordering to make concurrent
transactions appear as if they happened in a safe order.

This model is excellent for the earliest chapters:

- double-entry ledger;
- simple spot settlement;
- deposit callback idempotency;
- withdrawal workflow;
- reconciliation reports;
- admin operations.

It is also close to real payment systems: for many money workflows, a good SQL
transaction and a clear ledger schema beat an overbuilt event engine.

## The Pressure: Why Row State Becomes Hard To Explain

The first pressure is concurrency. If the same account, order, instrument, or
position is touched by many requests, the database must serialize conflicts
through locks or retries. That can be correct, but the business ordering is not
always obvious from the outside.

The second pressure is recovery. A database can recover its own storage state,
but business recovery asks different questions:

- Which command caused this balance?
- Which event caused this order to become filled?
- What did the system believe after sequence `N`?
- Can we rebuild the state on another machine and get the same result?
- Can risk, reporting, and audit all consume the same facts?

The third pressure is latency. Database transactions pay for generality:
locking, indexes, WAL, query planning, buffer management, and cross-row
coordination. For a trading hot path, those costs often create p99/p999 latency
variance that is hard to reason about.

These pressures do not make the DB model wrong. They create the need for a
second model.

## Bridge Stage: Facts Begin To Compete With Rows

The first migration step should still use the database:

```text
request
  -> begin transaction
  -> insert command_log row
  -> insert event_log rows
  -> update materialized tables
  -> commit
```

The conceptual shift is:

```text
facts = command_log + event_log
views = balances + orders + reports
```

At this stage, the database still stores the facts and the views. But current
rows are no longer the only explanation of truth. A balance row becomes a
materialized view of an ordered fact history.

This is the bridge between product engineering and trading-system engineering.
It lets us keep SQL durability while introducing replay, audit, and event
contracts.

## The State-Machine Stage

The next stage moves the hot path into a deterministic state machine:

```text
request
  -> sequencer assigns seq
  -> append command log
  -> state machine applies command
  -> emit events
  -> async projections update databases
```

The key change is where atomicity lives.

Database transaction atomicity:

```text
storage layer isolates concurrent effects
```

State-machine atomicity:

```text
business commands are totally ordered, then applied one by one
```

From the model's point of view, both can produce a serial history:

```text
State_0 --op_1--> State_1 --op_2--> State_2 --op_3--> State_3
```

The database creates a serializable effect behind a transaction API. The state
machine exposes the serial history directly as a business primitive.

This is why the apparently expensive model can be easier to understand. Total
ordering has high infrastructure cost, but simple semantic shape:

- every command has a sequence number;
- every state transition has one cause;
- recovery is snapshot plus replay;
- replicas agree by applying the same ordered commands;
- incidents can be debugged by stopping at a sequence and inspecting state.

## Replicated State Machine

The final hot-path model is a replicated state machine:

```text
client command
  -> replicated log / Aeron Cluster / Raft-like ordering
  -> every node applies the same command stream
  -> every node reaches the same state
  -> snapshot + replay recover after failure
```

At this stage, the source of truth is:

```text
replicated ordered command log
```

The database is still important, but it is no longer the hot-path truth source.
It becomes:

- read model;
- reporting store;
- audit store;
- reconciliation target;
- compliance export path;
- cold-path archive.

## Code Map

- `shared/go` defines the first shared funds command/event contract.
- `integration-tests/funds_conformance_test.go` compares implementations
  through that contract instead of through their internal storage shape.
- `chapters/02-spot-trade-db-go` models the DB-transaction-shaped starting
  point and emits shared funding events for spot settlement.
- `chapters/03-wallet-deposit-withdrawal-go` models direct wallet workflow:
  deposits, withdrawal requests, withdrawal confirmations, and transfers.
- `chapters/04-command-log-replay-go` models ordered commands and deterministic
  replay over the same funds contract.
- `chapters/04-command-log-replay-go/internal/replay/replay_test.go` contains a
  tiny equivalence test: serial DB transactions and a sequenced state machine
  reach the same state from the same ordered operations.
- `chapters/11-replicated-state-machine-aeron-java` is the future replicated-log
  boundary.

## Future Work

1. Extend chapter 2 with SQL-backed command/event tables.
2. Make chapter 4 replay from a file-backed log and add snapshot cut points.
3. Expand the shared contract from funds into orders, executions, and positions.
4. Connect the Java single-writer state machine to the same command/event shape.
5. Connect the Aeron replicated-log boundary to the same semantic contract.
6. Add cold-path projectors that consume events into query tables.

---

## 中文

这份笔记是项目的脊柱：

> 从数据库 ACID 作为真相源开始，然后将真相迁移到有序事实日志
> 和确定性复制状态机。

要点的不是论证数据库不好。数据库模型是正确的起点，因为它简单、持久、
可查询且熟悉。迁移只在热路径需要行状态事务间接或昂贵提供的属性时才发生：
显式排序、重放、确定性恢复、更低尾延迟和多副本一致性。

## 起点：数据库行即真相

第一个版本应该很无聊：

```text
request
  -> begin transaction
  -> read balances/orders
  -> validate business rules
  -> update balances/orders/trades/ledger_entries
  -> commit
```

在这个模型中，真相源是当前数据库状态：

- `balances`
- `orders`
- `trades`
- `ledger_entries`
- `withdrawals`
- `deposits`

原子性来自存储层。应用程序请求数据库在多个请求同时运行时保护不变量。
数据库使用锁、MVCC、WAL、索引和提交排序使并发事务看起来像是以安全顺序发生的。

这个模型最适合早期章节：

- 双分录账本；
- 简单现货结算；
- 入金回调幂等性；
- 出金工作流；
- 对账报告；
- 管理操作。

它也接近真实支付系统：对于许多资金工作流，一个好的 SQL 事务加上一份清晰的
账本 schema 胜过过度建设的事件引擎。

## 压力：为什么行状态变得难以解释

第一个压力是并发。如果同一个账户、订单、合约或仓位被许多请求触及，
数据库必须通过锁或重试序列化冲突。这可以是正确的，但业务排序并不总是
从外部显而易见的。

第二个压力是恢复。数据库可以恢复自己的存储状态，但业务恢复提出不同的问题：

- 哪条命令导致了这个余额？
- 哪个事件导致这个订单变为成交？
- 系统在序列 `N` 后相信什么？
- 我们能在另一台机器上重建相同的状态吗？
- 风控、报告和审计都能消费相同的事实吗？

第三个压力是延迟。数据库事务为通用性付出代价：锁、索引、WAL、查询规划、
缓冲管理和跨行协调。对于交易热路径，这些成本通常会产生难以推理的 p99/p999
延迟方差。

这些压力并不使 DB 模型错误。它们创造了对第二种模型的需求。

## 桥阶段：事实开始与行竞争

第一个迁移步骤仍应使用数据库：

```text
request
  -> begin transaction
  -> insert command_log row
  -> insert event_log rows
  -> update materialized tables
  -> commit
```

概念转变是：

```text
facts = command_log + event_log
views = balances + orders + reports
```

在这个阶段，数据库仍然存储事实和视图。但当前行不再是真相的唯一解释。
余额行变成有序事实历史的物化视图。

这是产品工程和交易系统工程的桥梁。它让我们保留 SQL 持久性，同时引入
重放、审计和事件契约。

## 状态机阶段

下一阶段将热路径移入确定性状态机：

```text
request
  -> sequencer assigns seq
  -> append command log
  -> state machine applies command
  -> emit events
  -> async projections update databases
```

关键变化是原子性所在的位置。

数据库事务原子性：

```text
storage layer isolates concurrent effects
```

状态机原子性：

```text
business commands are totally ordered, then applied one by one
```

从模型的角度，两者都可以产生串行历史：

```text
State_0 --op_1--> State_1 --op_2--> State_2 --op_3--> State_3
```

数据库在事务 API 背后创建可串行化效果。状态机将串行历史直接作为业务原语暴露。

这就是为什么看起来昂贵的模型反而更容易理解。全排序有高基础设施成本，
但语义形状简单：

- 每条命令都有一个序列号；
- 每个状态转换都有一个原因；
- 恢复是快照加重放；
- 副本通过应用相同的排序命令达成一致；
- 事件可以通过停在某个序列并检查状态来调试。

## 复制状态机

最终热路径模型是复制状态机：

```text
client command
  -> replicated log / Aeron Cluster / Raft-like ordering
  -> every node applies the same command stream
  -> every node reaches the same state
  -> snapshot + replay recover after failure
```

在这个阶段，真相源是：

```text
replicated ordered command log
```

数据库仍然重要，但它不再是热路径真相源。它变成：

- 读取模型；
- 报告存储；
- 审计存储；
- 对账目标；
- 合规导出路径；
- 冷路径归档。

## 代码地图

- `shared/go` 定义第一份共享资金命令/事件契约。
- `integration-tests/funds_conformance_test.go` 通过这份契约比较不同实现，
  而不是比较它们内部如何存储状态。
- `chapters/02-spot-trade-db-go` 建模 DB 事务形状的起点，并为现货结算输出
  共享资金事件。
- `chapters/03-wallet-deposit-withdrawal-go` 建模直接钱包工作流：入金、出金请求、
  出金确认和转账。
- `chapters/04-command-log-replay-go` 基于同一份资金契约建模有序命令和确定性重放。
- `chapters/04-command-log-replay-go/internal/replay/replay_test.go` 包含一个
  微小等价测试：串行 DB 事务和排序状态机从相同的排序操作达到相同状态。
- `chapters/11-replicated-state-machine-aeron-java` 是未来的复制日志边界。

## 未来工作

1. 在第 2 章扩展 SQL 支持的命令/事件表。
2. 让第 4 章从文件支持的日志重放，并添加快照切分点。
3. 将共享契约从资金扩展到订单、成交和仓位。
4. 将 Java 单写者状态机连接到相同的命令/事件形状。
5. 将 Aeron 复制日志边界连接到相同的语义契约。
6. 添加消费事件到查询表的冷路径投影器。
