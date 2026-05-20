# Version Contract And Testing

[English](#english) · [中文](#中文)

## English

This project should become stronger as architecture changes. The way to prove
that is to keep the exchange semantic contract stable and run the same scenarios
against each version. The scenario suite should be learned in small pieces
first, then composed into a full exchange run.

## Stable Exchange Contract

The exact method names may differ by language, but each runnable version should
preserve this shape:

```text
ordered command + deterministic transition -> new state + emitted facts
```

The contract is not only balances. Tests should inspect the journal entries
that explain a balance and the externally visible state that users and
downstream systems observe. It includes:

- double-entry ledger postings;
- account reservation and release;
- order acceptance, rejection, cancellation, and matching;
- execution reports and fee facts;
- position updates;
- margin and pre-trade risk decisions;
- projection cursors, gap detection, rebuild, and read models.

`shared/go/exchange` defines the current contract surface. It is intentionally
interface-first. Business implementations belong in version chapters, and only
after the contract tests are clear.

## Same Tests Across Versions

The desired proof is:

```text
same scenario suite
  -> v0 ACID SQL
  -> v1 SQL facts/outbox
  -> v2 single-node memory core
  -> v3 replicated log core
  -> v4 SQL projections
  -> same externally visible semantics
```

The current appendix funds scenarios are contract scaffolds, not normal green
tests:

```bash
make test-todo-go
```

That target is expected to fail at TODO boundaries until the exercises are
implemented. The future exchange contract scenarios are intentionally behind a
build tag:

```bash
cd shared/go
go test -tags exchange_contract_todo ./exchange
```

Those tests are also expected to fail until adapters and implementations exist.

## Scenario Composition

The full exchange scenarios are built from smaller teaching scenarios:

- chapter 01: deposit posts debit custody and credit user available;
- chapter 02: balance states move through journal entries between available,
  locked, pending withdrawal, custody, and fee revenue;
- chapter 03: buy orders reserve quote asset, sell orders reserve base asset,
  and cancellation releases remaining locked funds;
- chapter 04: execution emits facts, balances USD/BTC separately, posts fees,
  and releases surplus.

Those pieces then compose into architecture-version scenarios:

- the same scenario runs against ACID SQL, outbox, memory core, replicated log,
  and projection adapters;
- execution facts update positions;
- margin/risk rejects dangerous orders;
- replay preserves balances, orders, positions, and emitted facts;
- replicated nodes applying the same command stream reach the same state;
- SQL projection rebuilds a consistent read model from snapshot plus events.

Current funds-prototype scenarios:

- duplicate deposits are idempotent;
- withdrawals cannot spend unavailable funds;
- duplicate withdrawal confirmations are idempotent;
- transfers emit replayable facts and move balances.

## README Comparison Format

Every implemented version should eventually include:

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

The strongest possible `Semantic Change` is:

```text
None. The same external scenario suite passes unchanged.
```

That sentence is credible only when the tests make it true.

## DB ACID Versus Replicated Log

The DB version is best when:

- transaction latency is acceptable;
- the database can be the mutation boundary;
- SQL constraints make accounting and audit simple;
- operational simplicity matters most.

The replicated-log version is best when:

- hot-path ordering must be explicit;
- failover and replay are core requirements;
- database locks and retries dominate behavior;
- the team can afford replicated-log operational complexity.

They are not equivalent as implementations. They can be equivalent as external
business semantics when successful commands appear in the same valid order and
run the same deterministic transition logic.

---

## 中文

本项目应该在架构变化中变得更强。证明方法是保持 exchange semantic contract 稳定，
并让同一套场景跑过每个版本。场景套件应先被拆成小块学习，再组合成完整交易所
流程。

## 稳定交易所契约

具体方法名可能因语言而异，但每个可运行版本都应保持这个形状：

```text
ordered command + deterministic transition -> new state + emitted facts
```

契约不只是余额。测试应该检查解释余额的 journal entries，也要检查用户和下游系统
能观察到的外部状态。它包括：

- double-entry ledger postings；
- account reservation 和 release；
- 订单接受、拒绝、撤单和撮合；
- execution reports 和 fee facts；
- 仓位更新；
- margin 和 pre-trade risk decisions；
- projection cursors、gap detection、rebuild 和 read models。

`shared/go/exchange` 定义当前契约表面。它刻意 interface-first。业务实现属于版本
章节，而且要等契约测试清楚后再写。

## 跨版本相同测试

期望证明是：

```text
same scenario suite
  -> v0 ACID SQL
  -> v1 SQL facts/outbox
  -> v2 single-node memory core
  -> v3 replicated log core
  -> v4 SQL projections
  -> same externally visible semantics
```

当前附录资金场景是契约脚手架，不是常规绿色测试：

```bash
make test-todo-go
```

这个 target 在练习实现前，预期会失败在 TODO 边界。未来 exchange contract 场景
刻意放在 build tag 后：

```bash
cd shared/go
go test -tags exchange_contract_todo ./exchange
```

这些测试在 adapter 和实现出现前也应该失败。

## 场景组合

完整交易所场景由更小的教学场景组成：

- 第 01 章：入金对 custody 记 debit，对 user available 记 credit；
- 第 02 章：余额状态通过 journal entries 在 available、locked、pending withdrawal、
  custody 和 fee revenue 之间移动；
- 第 03 章：买单冻结 quote asset，卖单冻结 base asset，撤单释放剩余 locked funds；
- 第 04 章：成交发出 facts，分别平衡 USD/BTC，记录手续费并释放差额。

这些小块再组合成架构版本场景：

- 同一场景运行在 ACID SQL、outbox、memory core、replicated log 和 projection
  adapter 上；
- execution facts 更新仓位；
- margin/risk 拒绝危险订单；
- replay 保持 balances、orders、positions 和 emitted facts 一致；
- 复制节点应用同一命令流后达到相同状态；
- SQL projection 从 snapshot 加 events 重建一致 read model。

当前 funds prototype 场景：

- 重复入金是幂等的；
- 出金不能花费不可用资金；
- 重复出金确认是幂等的；
- 转账会产生可重放事实并移动余额。

## README 对比格式

每个已实现版本最终应包含：

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

`Semantic Change` 中最强的一句话是：

```text
None. The same external scenario suite passes unchanged.
```

只有测试使其为真，这句话才可信。

## DB ACID 与复制日志

DB 版本适合：

- 事务延迟可接受；
- 数据库可以作为 mutation boundary；
- SQL 约束让会计和审计简单；
- 运营简单性最重要。

复制日志版本适合：

- 热路径排序必须显式；
- failover 和 replay 是核心要求；
- 数据库锁和 retry 主导行为；
- 团队能承担复制日志运营复杂度。

它们作为实现并不等价。但当成功命令以同一有效顺序出现，并运行同一套确定性转换
逻辑时，它们可以拥有等价的外部业务语义。
