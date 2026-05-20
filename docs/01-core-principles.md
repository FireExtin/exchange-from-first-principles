# Core Principles

[English](#english) · [中文](#中文)

## English

This document owns the stable engineering principles for the project. It
combines the original project principles, the command/event model, and the
minimal replay/recovery checks.

## Project Shape

This repository is an education-oriented engineering project. It does not try
to build a full exchange first. It teaches the business semantics in small
steps, then shows how the same semantics survive architecture changes:

```text
custody/user liability
  -> balance states
  -> order reservation
  -> match and settlement
  -> ACID SQL exchange
  -> SQL facts and outbox
  -> deterministic memory core
  -> replicated log core
  -> SQL projections and consumer views
```

Architecture changes are not goals by themselves. They are trade-offs among
Correctness, Performance, and Reliability. Each stage chooses where to pay for
ordering, durability, recovery, latency, and operational complexity.

The project uses four Architecture Lenses to make those trade-offs visible:

- source of truth: what must stay correct;
- ordering model: where latency and contention are paid;
- recovery model: how the system survives failure;
- ownership and publication boundary: who owns state and emits facts.

## Semantic Contract First

Business semantics should stay stable while implementation substrates change.
The earlier Go funds contract lives in `shared/go/funds`. The active target
contract is the exchange-level surface in `shared/go/exchange`, presented
progressively from accounting and funds to orders, executions, positions, risk,
and projections.

The core shape is:

```text
old_state + command -> new_state + events
```

An implementation may be an ACID SQL transaction, a SQL facts/outbox bridge, a
single-node memory core, a replicated log core, or a SQL projection consumer.
It should still expose explainable commands, facts, balances, orders,
positions, rejection reasons, and replay behavior.

## Balances Are Explained, Not Assigned

Balances are not the primary story. Journal entries and emitted facts explain
why a balance changed. Implementations may materialize balances for speed, but
the materialized value should be reproducible from entries or from a documented
projection of entries.

Early chapters therefore introduce each business action in this order:

```text
user action -> account map -> journal template -> contract checks
```

Modern Treasury's wallet ledger article is a useful external reference for this
accounting teaching shape, but this repository's source of truth remains its
own docs and contracts:
[Accounting For Developers, Part II](https://www.moderntreasury.com/journal/accounting-for-developers-part-ii).

## Commands And Events

Commands express intent. Events express facts produced by the system.

Current and near-term command examples:

- `Deposit`
- `RequestWithdrawal`
- `ConfirmWithdrawal`
- `Transfer`
- `PlaceLimit`
- `Cancel`
- `ApplyExecution`

Current and near-term event examples:

- `Deposited`
- `WithdrawalRequested`
- `WithdrawalConfirmed`
- `Transferred`
- `OrderAccepted`
- `OrderRejected`
- `TradeExecuted`
- `OrderRested`
- `OrderCancelled`
- `CancelRejected`
- `PositionUpdated`
- `MarginRejected`
- `Rejected`

Only events mutate downstream views. Commands are inputs. Events are facts.

## Semantic Errors

Rejection reasons should describe business semantics, not implementation
details.

Examples:

- `InvalidAmount`
- `DuplicateCallback`
- `DuplicateWithdrawal`
- `DuplicateProviderEvent`
- `UnknownWithdrawal`
- `InsufficientFunds`
- `SequenceGap`
- `InvalidCommand`

SQL errors, map lookup failures, driver errors, or transport failures should be
wrapped into explicit semantic or operational errors before they cross a stable
contract boundary.

## Replay And Recovery

Replay rebuilds state by applying ordered input again. Speed matters later; the
first requirement is that the same command stream produces the same facts and
final state.

Minimal recovery checks:

- command sequence numbers are contiguous;
- replayed events match expected count and type;
- cash plus frozen cash are conserved;
- base asset plus frozen base asset are conserved;
- open orders match account reservations;
- emitted facts can be reconciled against wallet, ledger, provider, or bank
  records.

## Testing Rule

Cross-version tests should assert externally visible semantics, not internal
data structures.

Current green command:

```bash
make test-go
```

Appendix scenario tests live behind `make test-todo-go` and are expected to
fail at TODO boundaries until implemented. Future SQL, memory-core,
replicated-log, and projection implementations should join the exchange-level
contract pattern.

## Documentation Rule

Each active chapter should explain:

1. current system model;
2. semantic guarantees;
3. failure point or pressure;
4. why the next model becomes necessary.

Docs should link to the source of truth instead of repeating long design
arguments. Use `docs/README.md` for document organization and
`docs/07-chapter-roadmap.md` for chapter status.

## 中文

本文档负责项目稳定的工程原则。它合并了原项目指导原则、命令/事件模型，以及
最小重放/恢复检查。

## 项目形状

这个仓库是一个教育型工程项目。它不先构建完整交易所，而是先用小步骤讲清业务
语义，再展示同一语义如何在架构变化中存活：

```text
custody/user liability
  -> balance states
  -> order reservation
  -> match and settlement
  -> ACID SQL exchange
  -> SQL facts and outbox
  -> deterministic memory core
  -> replicated log core
  -> SQL projections and consumer views
```

架构变化本身不是目标。它是在正确性、性能和可靠性之间做取舍。每个阶段都在
选择：排序、持久化、恢复、延迟和运营复杂度分别在哪里付费。

项目使用四个架构分析视角来让这些取舍可见：

- 真相源：什么必须保持正确；
- 排序模型：延迟和竞争在哪里付费；
- 恢复模型：系统如何在故障后存活；
- 状态归属与发布边界：谁拥有状态并发出事实。

## 语义契约优先

业务语义应该在实现底座变化时保持稳定。早期 Go 资金契约位于 `shared/go/funds`；
当前目标契约是 `shared/go/exchange` 中的交易所级表面，并且会从 accounting 和
funds 逐步呈现到 orders、executions、positions、risk 和 projections。

核心形状是：

```text
old_state + command -> new_state + events
```

实现可以是 ACID SQL 事务、SQL facts/outbox 桥接、单机内存核心、复制日志核心
或 SQL projection consumer。它仍应该暴露可解释的命令、事实、余额、订单、仓位、
拒绝原因和重放行为。

## 余额是被解释出来的，不是直接赋值

余额不是第一叙事。journal entries 和 emitted facts 解释余额为什么变化。实现可以
为了速度物化余额，但物化值应该能从 entries 或明确记录的 entries projection 中
重算出来。

因此，早期章节按这个顺序介绍每个业务动作：

```text
用户动作 -> 账户图谱 -> 分录模板 -> 契约检查
```

Modern Treasury 的 wallet ledger 文章是一个有用的外部会计教学参考，但本仓库的真相
源仍然是自己的文档和契约：
[Accounting For Developers, Part II](https://www.moderntreasury.com/journal/accounting-for-developers-part-ii)。

## 命令与事件

命令表达意图。事件表达系统已经产生的事实。

当前和近期命令示例：

- `Deposit`
- `RequestWithdrawal`
- `ConfirmWithdrawal`
- `Transfer`
- `PlaceLimit`
- `Cancel`
- `ApplyExecution`

当前和近期事件示例：

- `Deposited`
- `WithdrawalRequested`
- `WithdrawalConfirmed`
- `Transferred`
- `OrderAccepted`
- `OrderRejected`
- `TradeExecuted`
- `OrderRested`
- `OrderCancelled`
- `CancelRejected`
- `PositionUpdated`
- `MarginRejected`
- `Rejected`

只有事件才能改变下游视图。命令是输入。事件是事实。

## 语义错误

拒绝原因应该描述业务语义，而不是实现细节。

示例：

- `InvalidAmount`
- `DuplicateCallback`
- `DuplicateWithdrawal`
- `DuplicateProviderEvent`
- `UnknownWithdrawal`
- `InsufficientFunds`
- `SequenceGap`
- `InvalidCommand`

SQL 错误、map 查询失败、驱动错误或传输失败，在跨越稳定契约边界前应包装成
显式的语义错误或运营错误。

## 重放与恢复

重放通过再次应用有序输入来重建状态。速度是后面的要求；第一要求是相同命令流
产生相同事实和最终状态。

最小恢复检查：

- 命令序列号连续；
- 重放事件的数量和类型符合预期；
- 现金加冻结现金守恒；
- 基础资产加冻结基础资产守恒；
- 活跃订单与账户预留匹配；
- 发出的事实可以与钱包、账本、provider 或银行记录对账。

## 测试规则

跨版本测试应断言外部可见语义，而不是内部数据结构。

当前绿色命令：

```bash
make test-go
```

附录场景测试位于 `make test-todo-go` 后，在实现前预期会失败在 TODO 边界。未来
SQL、内存核心、复制日志和 projection 实现应接入交易所级契约模式。

## 文档规则

每个活跃章节都应解释：

1. 当前系统模型；
2. 语义保证；
3. 失效点或压力；
4. 为什么下一个模型变得必要。

文档应链接到真相源，而不是到处重复长篇设计论证。文档组织看 `docs/README.md`，
章节状态看 `docs/07-chapter-roadmap.md`。
