# Exchange From First Principles

> Business semantics first, architecture migration second: prove the same
> exchange behavior across ACID SQL, memory state machines, replicated logs,
> and SQL projections.

[English](#english) · [中文](#中文)

---

## English

### The Question

Why does a modern exchange evolve from a simple ACID SQL transaction model into
something built around command/event facts, in-memory trading cores, replicated
logs, OMS, risk engines, hot/warm/cold paths, cache rebuilds, push streams, and
low-latency runtime work?

This project answers that question in two steps: first make the smallest
exchange business semantics legible, then keep those semantics stable while the
architecture changes underneath them.

It is not a production exchange. It is a systems project about correctness,
performance, reliability, risk boundaries, and architecture evolution.

### The Thesis

Architecture changes are not goals by themselves. They are trade-offs among
Correctness, Performance, and Reliability. Each stage chooses where to pay for
ordering, durability, recovery, latency, and operational complexity.

The business semantics should remain stable:

```text
ordered command + current state -> new state + emitted facts
```

The first chapters climb the business semantics slowly: custody and user
liability, balance states, order reservation, cancellation, matching,
settlement, fees, and release. The first complete architecture version then
uses ACID SQL because it is familiar, durable, and good at explaining
double-entry accounting and transaction boundaries. Later versions move the
same exchange semantics through SQL facts/outbox, a single-node in-memory
trading core, a replicated log core, and SQL projection consumers.

The contract is exchange-level, not only funds-level: accounting, reservation,
matching, executions, positions, margin/risk admission, projections, caches,
and push recovery all have to survive the architecture migration.

### Documentation Map

The canonical documentation index is [docs/README.md](./docs/README.md). Use
that page for reading order, source-of-truth ownership, and chapter status.

### Main Chapters

| Chapter | Topic | Status |
| --- | --- | --- |
| 01 | [Custody and user ledger](./chapters/01-custody-and-user-ledger-go/README.md) | Contract scaffold |
| 02 | [Balance states](./chapters/02-balance-states-go/README.md) | Contract scaffold |
| 03 | [Order reservation](./chapters/03-order-reservation-go/README.md) | Contract scaffold |
| 04 | [Match and settlement](./chapters/04-match-and-settlement-go/README.md) | Contract scaffold |
| 05 | [ACID SQL exchange](./chapters/05-acid-sql-exchange-go/README.md) | Contract scaffold |
| 06 | [SQL facts and outbox](./chapters/06-sql-facts-outbox-go/README.md) | Contract scaffold |
| 07 | [Single-node memory core](./chapters/07-single-node-memory-core-java/README.md) | README scaffold |
| 08 | [Replicated log core](./chapters/08-replicated-log-core-aeron-java/README.md) | Runnable Java skeleton |
| 09 | [SQL projection consumers](./chapters/09-sql-projection-consumers/README.md) | README scaffold |
| 10 | [Order book mechanics](./chapters/10-order-book-mechanics/README.md) | README only |
| 11 | [Position and PnL](./chapters/11-position-and-pnl/README.md) | README only |
| 12 | [Margin and pre-trade risk](./chapters/12-margin-and-pretrade-risk/README.md) | README only |
| 13 | [Risk cluster projection](./chapters/13-risk-cluster-projection/README.md) | README only |
| 14 | [Cache coherence and market state](./chapters/14-cache-coherence-and-market-state/README.md) | README only |
| 15 | [Market and execution push](./chapters/15-market-execution-push/README.md) | README only |
| 16 | [Rust hot path](./chapters/16-rust-hot-path/README.md) | README only |
| 17 | [Low-latency runtime and networking](./chapters/17-low-latency-runtime-networking/README.md) | README only |
| 18-20 | Credit, margin, and funding extension ideas | Planned notes only |
| 21-25 | Trading desk extension ideas | Planned notes only |

### Appendix Prototypes

The current Go funds examples are preserved as contract scaffolds. They are
useful exercises, but they are not the new canonical teaching sequence.

| Chapter | Topic | Status |
| --- | --- | --- |
| 90 | [Funds double-entry prototype](./chapters/90-funds-double-entry-prototype-go/README.md) | Contract scaffold |
| 91 | [Spot settlement transaction prototype](./chapters/91-spot-settlement-transaction-prototype-go/README.md) | Contract scaffold |
| 92 | [Wallet idempotency prototype](./chapters/92-wallet-idempotency-prototype-go/README.md) | Contract scaffold plus reconciliation lab |
| 93 | [Command log replay prototype](./chapters/93-command-log-replay-prototype-go/README.md) | Contract scaffold |

### Repository Layout

```text
chapters/             semantic ramp, architecture chapters, and appendix prototypes
docs/                 design notes, roadmap, and project documentation
shared/               shared semantic contracts and schemas
integration-tests/    cross-version semantic tests
protocol/             SBE schema placeholder and fixtures
tools/go/             Go load generator and reconciler placeholders
testdata/scenarios/   end-to-end scenario fixtures
ops/                  runbooks and deployment notes
```

### How To Read

- Start with [docs/README.md](./docs/README.md).
- Read [docs/00-goal.md](./docs/00-goal.md),
  [docs/01-core-principles.md](./docs/01-core-principles.md), and
  [docs/02-design-paper.md](./docs/02-design-paper.md).
- Use [docs/07-chapter-roadmap.md](./docs/07-chapter-roadmap.md) for the
  semantic ramp and architecture migration order.
- Use the chapter README before changing code in that chapter.

### Toolchain

The runnable checks in the repo currently expect:

- Go 1.22 or newer for shared contracts and tools;
- Java 21 and Gradle for the replicated-log skeleton.

### First Commands

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-java
```

Appendix contract scaffolds are intentionally incomplete:

```bash
make test-todo-go
```

That target is expected to fail at TODO boundaries until a learner implements
the exercises.

The incomplete exchange contract scenarios are intentionally behind a build
tag:

```bash
cd shared/go
go test -tags exchange_contract_todo ./exchange
```

They are expected to fail until a learner provides adapters and implementations.

### Non-Goals

This project is not:

- a production exchange;
- a custody or wallet security system;
- financial software for real funds;
- a full compliance platform;
- an attempt to implement consensus from scratch.

The goal is to make the architecture explainable from first principles, then
turn each step into small code and contract exercises.

### License

Apache-2.0. See `LICENSE`.

---

## 中文

### 这是什么

为什么现代交易所会从简单的 ACID SQL 事务模型，逐步演进到 command/event facts、
内存交易核心、复制日志、OMS、风控、冷热路径、缓存重建、推送流和低延迟运行时？

本项目分两步回答这个问题：先把最小交易所业务语义讲清楚，再在底层架构不断
变化时保持这些语义稳定。

这不是生产级交易所。它是一个系统设计和工程训练项目，重点是正确性、性能、
可靠性、风控边界和架构演化。

### 核心论点

架构变化本身不是目标。它是在正确性、性能和可靠性之间做取舍。每个阶段都在
选择：排序、持久化、恢复、延迟和运营复杂度分别在哪里付费。

业务语义应该保持稳定：

```text
ordered command + current state -> new state + emitted facts
```

前几章先缓慢爬坡业务语义：custody 与 user liability、余额状态、订单冻结、
撤单释放、撮合、结算、手续费和差额释放。第一版完整架构再使用 ACID SQL，
因为它熟悉、持久，也最适合讲清 double-entry accounting 和事务边界。后续版本
把同一套交易所语义迁移到 SQL facts/outbox、单机内存交易核心、复制日志核心，
以及 SQL projection consumers。

契约是交易所级别的，不只是资金级别：accounting、reservation、matching、
executions、positions、margin/risk admission、projections、caches 和 push
recovery 都必须在架构迁移中保持可解释。

### 文档地图

规范文档索引在 [docs/README.md](./docs/README.md)。阅读顺序、真相源归属和章节
状态都从那里开始。

### 主章节

| 章节 | 主题 | 状态 |
| --- | --- | --- |
| 01 | [托管与用户账本](./chapters/01-custody-and-user-ledger-go/README.md) | 契约脚手架 |
| 02 | [余额状态](./chapters/02-balance-states-go/README.md) | 契约脚手架 |
| 03 | [订单冻结](./chapters/03-order-reservation-go/README.md) | 契约脚手架 |
| 04 | [撮合与结算](./chapters/04-match-and-settlement-go/README.md) | 契约脚手架 |
| 05 | [ACID SQL exchange](./chapters/05-acid-sql-exchange-go/README.md) | 契约脚手架 |
| 06 | [SQL facts and outbox](./chapters/06-sql-facts-outbox-go/README.md) | 契约脚手架 |
| 07 | [单机内存核心](./chapters/07-single-node-memory-core-java/README.md) | README 脚手架 |
| 08 | [复制日志核心](./chapters/08-replicated-log-core-aeron-java/README.md) | 可运行 Java 骨架 |
| 09 | [SQL projection consumers](./chapters/09-sql-projection-consumers/README.md) | README 脚手架 |
| 10 | [订单簿机制](./chapters/10-order-book-mechanics/README.md) | 仅 README |
| 11 | [仓位和 PnL](./chapters/11-position-and-pnl/README.md) | 仅 README |
| 12 | [保证金与下单前风控](./chapters/12-margin-and-pretrade-risk/README.md) | 仅 README |
| 13 | [风控集群 projection](./chapters/13-risk-cluster-projection/README.md) | 仅 README |
| 14 | [缓存一致性与市场状态](./chapters/14-cache-coherence-and-market-state/README.md) | 仅 README |
| 15 | [行情与成交推送](./chapters/15-market-execution-push/README.md) | 仅 README |
| 16 | [Rust 热路径](./chapters/16-rust-hot-path/README.md) | 仅 README |
| 17 | [低延迟运行时与网络](./chapters/17-low-latency-runtime-networking/README.md) | 仅 README |
| 18-20 | Credit、margin 和 funding 扩展想法 | 仅规划笔记 |
| 21-25 | 交易台扩展想法 | 仅规划笔记 |

### 附录原型

当前 Go 资金示例保留为 contract scaffold。它们仍然适合练习，但不是新的规范教学顺序。

| 章节 | 主题 | 状态 |
| --- | --- | --- |
| 90 | [资金 double-entry 原型](./chapters/90-funds-double-entry-prototype-go/README.md) | 契约脚手架 |
| 91 | [现货结算事务原型](./chapters/91-spot-settlement-transaction-prototype-go/README.md) | 契约脚手架 |
| 92 | [钱包幂等原型](./chapters/92-wallet-idempotency-prototype-go/README.md) | 契约脚手架，含对账实验 |
| 93 | [命令日志重放原型](./chapters/93-command-log-replay-prototype-go/README.md) | 契约脚手架 |

### 仓库结构

```text
chapters/             语义爬坡、架构章节和附录原型
docs/                 设计说明、路线图和项目文档
shared/               共享语义契约和 schema
integration-tests/    跨版本语义测试
protocol/             SBE schema 占位和测试数据
tools/go/             Go 压测器与对账工具占位
testdata/scenarios/   端到端场景数据
ops/                  runbook 和部署说明
```

### 如何阅读

- 先从 [docs/README.md](./docs/README.md) 开始。
- 再读 [docs/00-goal.md](./docs/00-goal.md)、
  [docs/01-core-principles.md](./docs/01-core-principles.md) 和
  [docs/02-design-paper.md](./docs/02-design-paper.md)。
- 用 [docs/07-chapter-roadmap.md](./docs/07-chapter-roadmap.md) 查看语义爬坡和架构迁移顺序。
- 修改某章代码前，先读该章自己的 README。

### 工具链

当前绿色检查需要：

- Go 1.22 或更新版本，用于共享契约和工具；
- Java 21 和 Gradle，用于复制日志骨架。

### 第一组命令

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-java
```

附录契约脚手架刻意不完整：

```bash
make test-todo-go
```

这个 target 在学习者实现练习前，预期会失败在 TODO 边界。

未完成的 exchange contract 场景刻意放在 build tag 后：

```bash
cd shared/go
go test -tags exchange_contract_todo ./exchange
```

在学习者提供 adapter 和实现前，这些测试应该失败。

### 非目标

这个项目不是：

- 生产级交易所；
- 托管或钱包安全系统；
- 处理真实资产的金融软件；
- 完整合规平台；
- 从零实现共识算法。

目标是从第一性原理把架构讲清楚，再把每一步变成可练习的小代码和契约项目。

### License

Apache-2.0。详见 `LICENSE`。
