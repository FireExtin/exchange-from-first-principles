# Chapter Roadmap

[English](#english) · [中文](#中文)

## English

This roadmap owns chapter order and status. The main sequence now has three
parts:

```text
business semantic ramp
  -> architecture migration line
  -> domain and runtime deep dives
```

The ramp exists so readers do not meet the whole exchange contract at once.
Custody, balance states, reservation, cancellation, matching, settlement, fees,
and release are introduced before the first complete ACID SQL exchange.
Early ramp chapters use the same teaching shape: user action, account map,
journal template, contract checks, and out-of-scope boundaries.

For the full documentation catalog, see [Documentation Map](./README.md).

## Status Vocabulary

- Runnable: normal tests or demos run without special TODO work.
- Runnable skeleton: the module compiles and tests a boundary, but most
  business behavior is intentionally absent.
- Contract scaffold: interfaces, docs, and TODO tests define behavior; business
  implementation is intentionally absent.
- README scaffold: the chapter explains its role, but has no code yet.
- README only: domain note or future chapter without runnable code.
- Planned note: the idea is described in docs, but no chapter directory exists.

## Business Semantic Ramp

1. `01-custody-and-user-ledger-go`
   - Role: explain platform custody assets and user liability accounts.
   - Covers: deposits as Debit custody / Credit user available, and why user
     balances are not revenue.
   - Status: contract scaffold.

2. `02-balance-states-go`
   - Role: split user balances into operational states.
   - Covers: available, locked, pending withdrawal, fee revenue as a platform
     account purpose, and state-changing journal entries.
   - Status: contract scaffold.

3. `03-order-reservation-go`
   - Role: explain order intent before matching.
   - Covers: order acceptance/rejection, reservation, cancellation, and release.
   - Status: contract scaffold.

4. `04-match-and-settlement-go`
   - Role: explain the smallest match and settlement semantics.
   - Covers: execution facts, USD/BTC separate balancing, buyer fee, partial
     fill release, and ledger-explainable settlement.
   - Status: contract scaffold.

## Architecture Migration Line

5. `05-acid-sql-exchange-go`
   - Version role: v0 ACID SQL truth source.
   - Covers: composing chapters 01-04 inside SQL transaction boundaries.
   - Status: contract scaffold.

6. `06-sql-facts-outbox-go`
   - Version role: v1 SQL facts and outbox bridge.
   - Covers: SQL mutation plus command/event/outbox records in the same commit.
   - Status: contract scaffold.

7. `07-single-node-memory-core-java`
   - Version role: v2 deterministic single-node memory trading core.
   - Covers: sequenced commands, private hot state, facts, snapshots.
   - Status: README scaffold.

8. `08-replicated-log-core-aeron-java`
   - Version role: v3 replicated log core.
   - Covers: command order, log position, snapshot/replay, failover boundary.
   - Status: runnable Java skeleton.

9. `09-sql-projection-consumers`
   - Version role: v4 SQL projection and consumer views.
   - Covers: OMS, ledger reports, reconciliation, compliance, risk views,
     cache rebuild, push recovery.
   - Status: README scaffold.

## Domain And Runtime Deep Dives

10. `10-order-book-mechanics`
    - Purpose: explain order-book internals after basic matching semantics are
      already introduced in chapter 04.
    - Status: README only.

11. `11-position-and-pnl`
    - Purpose: explain how execution facts become positions and PnL placeholders.
    - Status: README only.

12. `12-margin-and-pretrade-risk`
    - Purpose: combine margin checks, reservation, risk admission, and kill
      switch semantics.
    - Status: README only.

13. `13-risk-cluster-projection`
    - Purpose: model warm-path risk projections that consume facts.
    - Status: README only.

14. `14-cache-coherence-and-market-state`
    - Purpose: define cache freshness, fail policy, gap detection, and rebuild.
    - Status: README only.

15. `15-market-execution-push`
    - Purpose: define recoverable public market-data and private
      execution-report streams.
    - Status: README only.

16. `16-rust-hot-path`
    - Purpose: discuss hot-path runtime options (Rust+io_uring, C++/DPDK, etc.)
      after chapter 07-09 semantics are stable and measured.
    - Status: README only.

17. `17-low-latency-runtime-networking`
    - Purpose: measure runtime and network variance after semantics are stable.
    - Status: README only.

## Appendix Prototypes

The Go funds appendix chapters are preserved as contract scaffolds. They are
useful exercises, but they are not the canonical teaching sequence and their
tests are expected to fail at TODO boundaries until implemented.

90. `90-funds-double-entry-prototype-go`
    - Status: contract scaffold.

91. `91-spot-settlement-transaction-prototype-go`
    - Status: contract scaffold.

92. `92-wallet-idempotency-prototype-go`
    - Status: contract scaffold. Reconciliation lab is an opt-in TODO exercise
      behind the `reconciliation_lab_todo` build tag.

93. `93-command-log-replay-prototype-go`
    - Status: contract scaffold.

## Credit, Margin, And Funding Extension

These are optional exchange-core extensions. They borrow only the lending
concepts needed for margin and funding semantics; they are not a Modern Lending
product track.

18. `18-credit-and-collateral-accounts-go`
    - Purpose: define collateral, borrow liability, funding accrual, repayment,
      and liquidation settlement contract boundaries.
    - Status: contract scaffold.

19. `19-funding-and-interest-accrual`
    - Status: planned note.

20. `20-liquidation-and-repayment-settlement`
    - Status: planned note.

## Trading Desk Extension

Trading desk chapters are planned notes only and are not part of the exchange
core. They remain consumers of exchange-core facts and come after the
credit/margin/funding extension.

21. `21-external-market-data-ingestion`
22. `22-pricing-and-signal-engine`
23. `23-order-router-and-execution-reports`
24. `24-hedger-and-best-execution`
25. `25-arbitrage-strategy-demo`

## Rule

Each chapter should say:

- what business semantic, architecture role, or domain role it owns;
- what semantic contract it must preserve;
- what implementation is intentionally absent;
- what would count as a future runnable proof.

Chapters 01-05 should additionally keep the teaching flow concrete: start from
a user action, name the accounts involved, show the journal template, then name
the TODO contract checks.

Unfinished exchange behavior belongs behind explicit TODO tests or build tags,
not in silent partial implementations.

## Core Theory Notes

- `shared/go/exchange` is the exchange-level semantic contract.
- `integration-tests` owns cross-version scenario tests.
- `docs/04-truth-source-migration.md` explains how truth moves after the
  business semantic ramp is defined.
- `docs/06-version-contract-and-testing.md` defines how tests prove semantic
  equivalence.
- `docs/08-position-matching-risk-margin.md` defines later hot-path trading
  surfaces preserved by each version.
- `docs/09-credit-margin-funding-extension.md` defines the optional
  credit/margin/funding extension before desk systems.

---

## 中文

本路线图负责章节顺序和状态。主线现在分成三段：

```text
业务语义爬坡
  -> 架构迁移线
  -> 领域与运行时深挖
```

语义爬坡的目的，是避免读者一上来就面对完整 exchange contract。custody、余额
状态、reservation、撤单、撮合、结算、手续费和释放，会先于第一版完整 ACID SQL
exchange 出现。
早期爬坡章节使用同一教学形状：用户动作、账户图谱、分录模板、契约检查和本章
不做什么。

完整文档目录见 [Documentation Map](./README.md)。

## 状态术语

- 可运行：常规测试或 demo 可以运行，不依赖 TODO 实现。
- 可运行骨架：模块能编译并测试边界，但多数业务行为故意留空。
- 契约脚手架：接口、文档和 TODO 测试定义行为；业务实现刻意留空。
- README 脚手架：章节说明自己的角色，但尚无代码。
- 仅 README：领域说明或未来章节，无可运行代码。
- 规划笔记：想法已写进文档，但尚无章节目录。

## 业务语义爬坡

1. `01-custody-and-user-ledger-go`
   - 角色：解释平台 custody asset 和用户 liability account。
   - 覆盖：入金是 Debit custody / Credit user available，以及为什么用户余额
     不是收入。
   - 状态：契约脚手架。

2. `02-balance-states-go`
   - 角色：把用户余额拆成操作状态。
   - 覆盖：available、locked、pending withdrawal、fee revenue 作为平台账户用途，
     以及改变状态的 journal entries。
   - 状态：契约脚手架。

3. `03-order-reservation-go`
   - 角色：在撮合前解释订单意图。
   - 覆盖：订单接受/拒绝、冻结、撤单和释放。
   - 状态：契约脚手架。

4. `04-match-and-settlement-go`
   - 角色：解释最小撮合和结算语义。
   - 覆盖：execution facts、USD/BTC 分别平衡、买方手续费、部分成交释放，以及
     可由账本解释的结算。
   - 状态：契约脚手架。

## 架构迁移线

5. `05-acid-sql-exchange-go`
   - 版本角色：v0 ACID SQL 真相源。
   - 覆盖：把第 01-04 章组合进 SQL 事务边界。
   - 状态：契约脚手架。

6. `06-sql-facts-outbox-go`
   - 版本角色：v1 SQL facts 和 outbox 桥接。
   - 覆盖：SQL mutation 加 command/event/outbox records 在同一次 commit 内完成。
   - 状态：契约脚手架。

7. `07-single-node-memory-core-java`
   - 版本角色：v2 确定性单机内存交易核心。
   - 覆盖：已排序命令、私有热状态、事实、快照。
   - 状态：README 脚手架。

8. `08-replicated-log-core-aeron-java`
   - 版本角色：v3 复制日志核心。
   - 覆盖：命令顺序、日志位置、snapshot/replay、failover 边界。
   - 状态：可运行 Java 骨架。

9. `09-sql-projection-consumers`
   - 版本角色：v4 SQL projection 和 consumer views。
   - 覆盖：OMS、ledger reports、对账、合规、risk views、cache rebuild、push recovery。
   - 状态：README 脚手架。

## 领域与运行时深挖

10. `10-order-book-mechanics`
    - 目的：在第 04 章已经引入基础撮合语义之后，解释订单簿内部机制。
    - 状态：仅 README。

11. `11-position-and-pnl`
    - 目的：解释 execution facts 如何变成仓位和 PnL 占位。
    - 状态：仅 README。

12. `12-margin-and-pretrade-risk`
    - 目的：合并保证金检查、reservation、risk admission 和 kill switch 语义。
    - 状态：仅 README。

13. `13-risk-cluster-projection`
    - 目的：建模消费事实的 warm-path risk projection。
    - 状态：仅 README。

14. `14-cache-coherence-and-market-state`
    - 目的：定义缓存新鲜度、失败策略、gap detection 和 rebuild。
    - 状态：仅 README。

15. `15-market-execution-push`
    - 目的：定义可恢复的 public market-data 和 private execution-report streams。
    - 状态：仅 README。

16. `16-rust-hot-path`
    - 目的：在第 07-09 章语义稳定和可测量之后，讨论热路径运行时选项
（Rust+io_uring、C++/DPDK 等）。
    - 状态：仅 README。

17. `17-low-latency-runtime-networking`
    - 目的：在语义稳定后测量 runtime 和 network variance。
    - 状态：仅 README。

## 附录原型

Go 资金附录章节保留为契约脚手架。它们是有用练习，但不是规范教学顺序；在实现
之前，它们的测试预期会失败在 TODO 边界。

90. `90-funds-double-entry-prototype-go`
    - 状态：契约脚手架。

91. `91-spot-settlement-transaction-prototype-go`
    - 状态：契约脚手架。

92. `92-wallet-idempotency-prototype-go`
    - 状态：契约脚手架。对账实验在 `reconciliation_lab_todo` build tag 后，
      是可选 TODO 练习。

93. `93-command-log-replay-prototype-go`
    - 状态：契约脚手架。

## Credit, Margin, And Funding 扩展

这些是可选 exchange-core 扩展。它们只借用解释 margin 和 funding 语义所需的
lending 概念；不是 Modern Lending 产品线。

18. `18-credit-and-collateral-accounts-go`
    - 目的：定义 collateral、borrow liability、funding accrual、repayment 和
      liquidation settlement 的契约边界。
    - 状态：契约脚手架。

19. `19-funding-and-interest-accrual`
    - 状态：规划笔记。

20. `20-liquidation-and-repayment-settlement`
    - 状态：规划笔记。

## 交易台扩展

交易台章节目前只是规划笔记，不属于交易所核心。它们仍然是 exchange-core facts 的
消费者，并排在 credit/margin/funding 扩展之后。

21. `21-external-market-data-ingestion`
22. `22-pricing-and-signal-engine`
23. `23-order-router-and-execution-reports`
24. `24-hedger-and-best-execution`
25. `25-arbitrage-strategy-demo`

## 规则

每章都应该说明：

- 自己拥有的业务语义、架构角色或领域角色；
- 必须保持的语义契约；
- 哪些实现刻意缺席；
- 未来怎样才算可运行证明。

第 01-05 章还应保持具体教学流程：先从用户动作出发，命名涉及账户，展示分录模板，
然后说明 TODO 契约检查。

未完成的 exchange 行为应放在显式 TODO 测试或 build tag 后，而不是沉默的半成品实现。

## 核心理论说明

- `shared/go/exchange` 是交易所级语义契约。
- `integration-tests` 负责跨版本场景测试。
- `docs/04-truth-source-migration.md` 解释业务语义爬坡之后，真相如何迁移。
- `docs/06-version-contract-and-testing.md` 定义测试如何证明语义等价。
- `docs/08-position-matching-risk-margin.md` 定义每个版本必须保持的后续热路径交易表面。
- `docs/09-credit-margin-funding-extension.md` 定义交易台系统之前的可选
  credit/margin/funding 扩展。
