# Chapter Roadmap

[English](#english) · [中文](#中文)

## English

This roadmap owns chapter order and status. The main sequence now follows a
version line: the same exchange semantics should survive as the source of truth
moves from ACID SQL to SQL facts/outbox, a single-node memory core, a replicated
log core, and SQL projection consumers.

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

## Main Version Line

1. `01-exchange-semantic-contract-go`
   - Version role: shared exchange semantic contract.
   - Covers: accounting, reservation, orders, executions, positions, risk,
     projection cursors, adapter boundaries.
   - Status: contract scaffold.

2. `02-acid-sql-exchange-go`
   - Version role: v0 ACID SQL truth source.
   - Covers: double-entry ledger, account reservation, orders, trades,
     positions, risk-admission state inside one transaction boundary.
   - Status: contract scaffold.

3. `03-sql-facts-outbox-go`
   - Version role: v1 SQL facts and outbox bridge.
   - Covers: SQL mutation plus command/event/outbox records in the same commit.
   - Status: contract scaffold.

4. `04-single-node-memory-core-java`
   - Version role: v2 deterministic single-node memory trading core.
   - Covers: sequenced commands, private hot state, facts, snapshots.
   - Status: README scaffold.

5. `05-replicated-log-core-aeron-java`
   - Version role: v3 replicated log core.
   - Covers: command order, log position, snapshot/replay, failover boundary.
   - Status: runnable Java skeleton.

6. `06-sql-projection-consumers`
   - Version role: v4 SQL projection and consumer views.
   - Covers: OMS, ledger reports, reconciliation, compliance, risk views,
     cache rebuild, push recovery.
   - Status: README scaffold.

## Domain Deep Dives

7. `07-order-book-mechanics`
   - Purpose: explain order-book internals after matching semantics are already
     part of the shared contract.
   - Status: README only.

8. `08-position-and-pnl`
   - Purpose: explain how execution facts become positions and PnL placeholders.
   - Status: README only.

9. `09-margin-and-pretrade-risk`
   - Purpose: combine margin checks, reservation, risk admission, and kill
     switch semantics.
   - Status: README only.

10. `10-risk-cluster-projection`
    - Purpose: model warm-path risk projections that consume facts.
    - Status: README only.

11. `11-cache-coherence-and-market-state`
    - Purpose: define cache freshness, fail policy, gap detection, and rebuild.
    - Status: README only.

12. `12-market-execution-push`
    - Purpose: define recoverable public market-data and private
      execution-report streams.
    - Status: README only.

13. `13-rust-hot-path`
    - Purpose: keep the Rust code as a runtime/hot-path experiment.
    - Status: runnable Rust experiment.

14. `14-low-latency-runtime-networking`
    - Purpose: measure runtime and network variance after semantics are stable.
    - Status: README only.

## Appendix Prototypes

The current runnable Go funds chapters are preserved as prototypes. They are
useful exercises and current test fixtures, but they are not the canonical
version-line implementation.

90. `90-funds-double-entry-prototype-go`
    - Status: runnable Go prototype.

91. `91-spot-settlement-transaction-prototype-go`
    - Status: runnable Go prototype.

92. `92-wallet-idempotency-prototype-go`
    - Status: runnable Go prototype with reconciliation lab.

93. `93-command-log-replay-prototype-go`
    - Status: runnable Go prototype.

## Trading Desk Extension

Trading desk chapters are planned notes only and are not part of the exchange
core. They remain consumers of exchange-core facts.

15. `15-external-market-data-ingestion`
16. `16-pricing-and-signal-engine`
17. `17-order-router-and-execution-reports`
18. `18-hedger-and-best-execution`
19. `19-arbitrage-strategy-demo`

## Rule

Each chapter should say:

- what version-line role or domain role it owns;
- what semantic contract it must preserve;
- what implementation is intentionally absent;
- what would count as a future runnable proof.

Unfinished exchange behavior belongs behind explicit TODO tests or build tags,
not in silent partial implementations.

## Core Theory Notes

- `shared/go/exchange` is the exchange-level semantic contract.
- `integration-tests` owns cross-version scenario tests.
- `docs/04-truth-source-migration.md` explains how truth moves across v0-v4.
- `docs/06-version-contract-and-testing.md` defines how tests prove semantic
  equivalence.
- `docs/08-position-matching-risk-margin.md` defines the hot-path trading
  semantics preserved by each version.

---

## 中文

本路线图负责章节顺序和状态。主线现在遵循版本线：同一套交易所语义应该在真相源
从 ACID SQL 迁移到 SQL facts/outbox、单机内存核心、复制日志核心和 SQL projection
consumer 时保持不变。

完整文档目录见 [Documentation Map](./README.md)。

## 状态术语

- 可运行：常规测试或 demo 可以运行，不依赖 TODO 实现。
- 可运行骨架：模块能编译并测试边界，但多数业务行为故意留空。
- 契约脚手架：接口、文档和 TODO 测试定义行为；业务实现刻意留空。
- README 脚手架：章节说明自己的角色，但尚无代码。
- 仅 README：领域说明或未来章节，无可运行代码。
- 规划笔记：想法已写进文档，但尚无章节目录。

## 主版本线

1. `01-exchange-semantic-contract-go`
   - 版本角色：共享交易所语义契约。
   - 覆盖：accounting、reservation、orders、executions、positions、risk、
     projection cursors、adapter boundaries。
   - 状态：契约脚手架。

2. `02-acid-sql-exchange-go`
   - 版本角色：v0 ACID SQL 真相源。
   - 覆盖：double-entry ledger、account reservation、orders、trades、positions、
     risk-admission state 位于同一个事务边界内。
   - 状态：契约脚手架。

3. `03-sql-facts-outbox-go`
   - 版本角色：v1 SQL facts 和 outbox 桥接。
   - 覆盖：SQL mutation 加 command/event/outbox records 在同一次 commit 内完成。
   - 状态：契约脚手架。

4. `04-single-node-memory-core-java`
   - 版本角色：v2 确定性单机内存交易核心。
   - 覆盖：已排序命令、私有热状态、事实、快照。
   - 状态：README 脚手架。

5. `05-replicated-log-core-aeron-java`
   - 版本角色：v3 复制日志核心。
   - 覆盖：命令顺序、日志位置、snapshot/replay、failover 边界。
   - 状态：可运行 Java 骨架。

6. `06-sql-projection-consumers`
   - 版本角色：v4 SQL projection 和 consumer views。
   - 覆盖：OMS、ledger reports、对账、合规、risk views、cache rebuild、push recovery。
   - 状态：README 脚手架。

## 领域深挖

7. `07-order-book-mechanics`
   - 目的：在撮合语义已经属于共享契约之后，解释订单簿内部机制。
   - 状态：仅 README。

8. `08-position-and-pnl`
   - 目的：解释 execution facts 如何变成仓位和 PnL 占位。
   - 状态：仅 README。

9. `09-margin-and-pretrade-risk`
   - 目的：合并保证金检查、reservation、risk admission 和 kill switch 语义。
   - 状态：仅 README。

10. `10-risk-cluster-projection`
    - 目的：建模消费事实的 warm-path risk projection。
    - 状态：仅 README。

11. `11-cache-coherence-and-market-state`
    - 目的：定义缓存新鲜度、失败策略、gap detection 和 rebuild。
    - 状态：仅 README。

12. `12-market-execution-push`
    - 目的：定义可恢复的 public market-data 和 private execution-report streams。
    - 状态：仅 README。

13. `13-rust-hot-path`
    - 目的：保留 Rust 代码作为 runtime/hot-path experiment。
    - 状态：可运行 Rust 实验。

14. `14-low-latency-runtime-networking`
    - 目的：在语义稳定后测量 runtime 和 network variance。
    - 状态：仅 README。

## 附录原型

当前可运行 Go 资金章节保留为 prototype。它们是有用练习和当前测试 fixture，但
不是规范版本线实现。

90. `90-funds-double-entry-prototype-go`
    - 状态：可运行 Go 原型。

91. `91-spot-settlement-transaction-prototype-go`
    - 状态：可运行 Go 原型。

92. `92-wallet-idempotency-prototype-go`
    - 状态：可运行 Go 原型，含对账实验。

93. `93-command-log-replay-prototype-go`
    - 状态：可运行 Go 原型。

## 交易台扩展

交易台章节目前只是规划笔记，不属于交易所核心。它们仍然是 exchange-core facts 的
消费者。

15. `15-external-market-data-ingestion`
16. `16-pricing-and-signal-engine`
17. `17-order-router-and-execution-reports`
18. `18-hedger-and-best-execution`
19. `19-arbitrage-strategy-demo`

## 规则

每章都应该说明：

- 自己拥有的 version-line role 或 domain role；
- 必须保持的语义契约；
- 哪些实现刻意缺席；
- 未来怎样才算可运行证明。

未完成的 exchange 行为应放在显式 TODO 测试或 build tag 后，而不是沉默的半成品实现。

## 核心理论说明

- `shared/go/exchange` 是交易所级语义契约。
- `integration-tests` 负责跨版本场景测试。
- `docs/04-truth-source-migration.md` 解释真相如何在 v0-v4 之间迁移。
- `docs/06-version-contract-and-testing.md` 定义测试如何证明语义等价。
- `docs/08-position-matching-risk-margin.md` 定义每个版本必须保持的热路径交易语义。
