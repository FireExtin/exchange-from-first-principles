# Core Principles

[English](#english) · [中文](#中文)

## English

This document owns the stable engineering principles for the project. It
combines the original project principles, the command/event model, and the
minimal replay/recovery checks.

## Project Shape

This repository is an education-oriented engineering project. It does not try
to build a full exchange first. It uses small runnable systems to show how an
exchange evolves:

```text
ACID SQL exchange
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
The first runnable contract was the Go funds contract in `shared/go/funds`.
The active target contract is the exchange-level surface in `shared/go/exchange`.

The core shape is:

```text
old_state + command -> new_state + events
```

An implementation may be an ACID SQL transaction, a SQL facts/outbox bridge, a
single-node memory core, a replicated log core, or a SQL projection consumer.
It should still expose explainable commands, facts, balances, orders,
positions, rejection reasons, and replay behavior.

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

Current command:

```bash
go test ./integration-tests/...
```

The current runnable suite compares appendix prototypes 92 and 93 through the
same shared funds contract. Future SQL, memory-core, replicated-log, and
projection implementations should join the exchange-level contract pattern.

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

这个仓库是一个教育型工程项目。它不先构建完整交易所，而是用小型可运行系统
展示交易所如何演进：

```text
ACID SQL exchange
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

业务语义应该在实现底座变化时保持稳定。第一份可运行契约是 `shared/go/funds`
中的 Go 资金契约；当前目标契约是 `shared/go/exchange` 中的交易所级表面。

核心形状是：

```text
old_state + command -> new_state + events
```

实现可以是 ACID SQL 事务、SQL facts/outbox 桥接、单机内存核心、复制日志核心
或 SQL projection consumer。它仍应该暴露可解释的命令、事实、余额、订单、仓位、
拒绝原因和重放行为。

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

当前命令：

```bash
go test ./integration-tests/...
```

当前可运行套件通过同一份共享资金契约，对比附录原型 92 和 93。未来 SQL、
内存核心、复制日志和 projection 实现应接入交易所级契约模式。

## 文档规则

每个活跃章节都应解释：

1. 当前系统模型；
2. 语义保证；
3. 失效点或压力；
4. 为什么下一个模型变得必要。

文档应链接到真相源，而不是到处重复长篇设计论证。文档组织看 `docs/README.md`，
章节状态看 `docs/07-chapter-roadmap.md`。
