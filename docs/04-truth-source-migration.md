# Truth Source Migration

[English](#english) · [中文](#中文)

## English

This note owns the project spine:

```text
minimal business semantics
  -> custody/user liability
  -> balance states
  -> order reservation
  -> match and settlement
  -> ACID SQL truth
  -> SQL facts and outbox
  -> deterministic memory core
  -> replicated log core
  -> SQL projections and consumer views
```

Truth-source migration starts only after the smallest business semantics are
defined. The migration target is not only balances. It is the full exchange
semantic surface: ledger postings, reservations, orders, matching, executions,
positions, risk admission, projection cursors, caches, and push recovery.

## Truth Layers

The migration target has three layers:

| Layer | Examples | Rule |
| --- | --- | --- |
| Posted facts | ledger entries, execution facts, settlement facts | Historical financial truth; append or reverse, do not silently rewrite. |
| Operational state | orders, reservations, hot-path risk admission | Current state needed to accept, reject, sequence, and execute commands. |
| Prospective or derived state | marks, margin requirements, unrealized PnL, continuous risk views, projections | Model or view state; useful for decisions, but not posted ledger truth until an explicit event posts entries. |

Architecture migration may move where these layers live, but it should not blur
their meaning.

## v0: ACID SQL As First Truth Source

After chapters 01-04 define the component semantics, the first complete model
should be boring and familiar:

```text
request
  -> begin transaction
  -> validate business rules
  -> post double-entry ledger entries
  -> update reservations, orders, trades, positions, risk state
  -> commit
```

SQL is a good starting point because it gives durability, queryability,
transactional isolation, and a natural place to explain accounting. User
balances are platform liabilities. Custody accounts are platform assets. Every
journal transaction balances per asset.

This is the right place to teach:

- double-entry accounting;
- account reservation;
- atomic spot settlement;
- order acceptance and rejection;
- position updates;
- margin and risk admission;
- SQL constraints and transaction boundaries.

## v1: SQL Facts And Outbox

The first migration step still uses SQL as truth:

```text
request
  -> begin transaction
  -> insert command row
  -> insert event rows
  -> update current tables
  -> insert outbox rows
  -> commit
```

The conceptual shift is:

```text
facts = command_log + event_log + ledger_entries
views = balances + orders + positions + risk views + reports
```

Rows remain the committed truth, but they are no longer the only explanation.
Facts begin to explain why each row exists, and consumers gain a stable cursor
for projection, audit, and recovery.

## v2: Single-Node Memory Core

The hot path then moves into a deterministic state machine:

```text
sequenced command
  -> apply to private in-memory state
  -> emit ordered facts
  -> snapshot for recovery
  -> project facts to SQL consumers
```

The key change is where atomicity lives.

Database transaction atomicity:

```text
storage layer isolates concurrent effects
```

State-machine atomicity:

```text
commands are totally ordered, then applied one by one
```

The same exchange semantics must survive: reservations, matching, executions,
positions, risk decisions, and ledger-explainable facts.

## v3: Replicated Log Core

The replicated-log version changes the ordering and recovery substrate:

```text
client command
  -> replicated log / Raft-style ordering
  -> each node applies the same command stream
  -> each node reaches the same state
  -> snapshot + replay recover after failure
```

The replicated log does not invent different business meaning. It gives the
same deterministic transition a shared order, failover boundary, and recovery
contract.

## v4: SQL Projection And Consumers

SQL returns as the warm/cold-path store:

- OMS views;
- ledger reports;
- reconciliation records;
- compliance exports;
- risk projections;
- cache rebuild inputs;
- push recovery checkpoints.

At this stage SQL is not rejected. It is moved to the places where queryability,
auditability, reconciliation, and reporting matter most.

## Code Map

- `shared/go/exchange` defines the exchange semantic contract.
- `chapters/01-custody-and-user-ledger-go` starts the business semantic ramp.
- `chapters/02-balance-states-go` explains account state transitions.
- `chapters/03-order-reservation-go` explains reservation and release.
- `chapters/04-match-and-settlement-go` explains matching, fees, and settlement.
- `chapters/05-acid-sql-exchange-go` is the v0 SQL contract scaffold.
- `chapters/06-sql-facts-outbox-go` is the v1 facts/outbox scaffold.
- `chapters/07-single-node-memory-core-java` is the v2 memory-core scaffold.
- `chapters/08-replicated-log-core-aeron-java` is the v3 replicated-log
  skeleton.
- `chapters/09-sql-projection-consumers` is the v4 projection scaffold.
- `chapters/90-*` through `chapters/93-*` preserve earlier funds contract
  scaffolds.

## Future Work

1. Implement v0 SQL only after the contract tests are clear.
2. Connect v1 outbox tables to the same event and cursor contract.
3. Connect v2 memory state to the same adapter surface.
4. Connect v3 replicated log to the same command stream and snapshot contract.
5. Build v4 SQL projections from facts, not from hidden side effects.

---

## 中文

这份笔记负责项目脊柱：

```text
minimal business semantics
  -> custody/user liability
  -> balance states
  -> order reservation
  -> match and settlement
  -> ACID SQL truth
  -> SQL facts and outbox
  -> deterministic memory core
  -> replicated log core
  -> SQL projections and consumer views
```

真相源迁移要等最小业务语义定义清楚后再开始。迁移对象不只是余额，而是完整交易所
语义表面：ledger postings、reservations、orders、matching、executions、positions、
risk admission、projection cursors、caches 和 push recovery。

## 真相层次

迁移对象分成三层：

| 层次 | 例子 | 规则 |
| --- | --- | --- |
| 已过账事实 | ledger entries、execution facts、settlement facts | 历史财务真相；只能追加或冲正，不能偷偷重写。 |
| 运营状态 | orders、reservations、hot-path risk admission | 接受、拒绝、排序和执行命令所需的当前状态。 |
| 未来或派生状态 | marks、margin requirements、unrealized PnL、continuous risk views、projections | 模型或视图状态；可用于决策，但只有显式事件提交 entries 后才成为 ledger truth。 |

架构迁移可以改变这些层次在哪里运行，但不应该模糊它们的含义。

## v0：ACID SQL 作为第一真相源

第 01-04 章定义组件语义之后，第一版完整模型应该朴素且熟悉：

```text
request
  -> begin transaction
  -> validate business rules
  -> post double-entry ledger entries
  -> update reservations, orders, trades, positions, risk state
  -> commit
```

SQL 是好的起点，因为它提供持久性、可查询性、事务隔离，也自然适合解释会计。
用户余额是平台负债。custody 账户是平台资产。每笔 journal transaction 在每种
资产内分别平衡。

这里适合教学：

- double-entry accounting；
- account reservation；
- 原子现货结算；
- 订单接受与拒绝；
- 仓位更新；
- 保证金和风控准入；
- SQL 约束和事务边界。

## v1：SQL Facts And Outbox

第一步迁移仍然使用 SQL 作为真相：

```text
request
  -> begin transaction
  -> insert command row
  -> insert event rows
  -> update current tables
  -> insert outbox rows
  -> commit
```

概念转移是：

```text
facts = command_log + event_log + ledger_entries
views = balances + orders + positions + risk views + reports
```

数据库行仍然是已提交真相，但不再是唯一解释。事实开始解释每一行为什么存在，
consumer 也获得稳定 cursor，用于 projection、审计和恢复。

## v2：单机内存核心

热路径随后进入确定性状态机：

```text
sequenced command
  -> apply to private in-memory state
  -> emit ordered facts
  -> snapshot for recovery
  -> project facts to SQL consumers
```

关键变化是原子性在哪里承担。

数据库事务原子性：

```text
storage layer isolates concurrent effects
```

状态机原子性：

```text
commands are totally ordered, then applied one by one
```

同一套交易所语义必须存活：reservation、matching、execution、position、risk
decision，以及能被 ledger 解释的事实。

## v3：复制日志核心

复制日志版本改变的是排序和恢复底座：

```text
client command
  -> replicated log / Raft-style ordering
  -> each node applies the same command stream
  -> each node reaches the same state
  -> snapshot + replay recover after failure
```

复制日志不发明新的业务含义。它给同一个确定性转换提供共享顺序、failover 边界和
恢复契约。

## v4：SQL Projection And Consumers

SQL 回到 warm/cold-path store：

- OMS views；
- ledger reports；
- 对账记录；
- 合规导出；
- risk projections；
- cache rebuild inputs；
- push recovery checkpoints。

此时 SQL 不是被否定，而是被移动到最需要可查询性、可审计性、对账和报表的地方。

## 代码地图

- `shared/go/exchange` 定义交易所语义契约。
- `chapters/01-custody-and-user-ledger-go` 开始业务语义爬坡。
- `chapters/02-balance-states-go` 解释账户状态迁移。
- `chapters/03-order-reservation-go` 解释冻结和释放。
- `chapters/04-match-and-settlement-go` 解释撮合、手续费和结算。
- `chapters/05-acid-sql-exchange-go` 是 v0 SQL 契约脚手架。
- `chapters/06-sql-facts-outbox-go` 是 v1 facts/outbox 脚手架。
- `chapters/07-single-node-memory-core-java` 是 v2 内存核心脚手架。
- `chapters/08-replicated-log-core-aeron-java` 是 v3 复制日志骨架。
- `chapters/09-sql-projection-consumers` 是 v4 projection 脚手架。
- `chapters/90-*` 到 `chapters/93-*` 保留早期资金契约脚手架。

## 后续工作

1. 只有在契约测试清晰后，才实现 v0 SQL。
2. 将 v1 outbox 表接入同一事件和 cursor 契约。
3. 将 v2 内存状态接入同一 adapter 表面。
4. 将 v3 复制日志接入同一 command stream 和 snapshot 契约。
5. 从事实构建 v4 SQL projection，而不是依赖隐藏副作用。
