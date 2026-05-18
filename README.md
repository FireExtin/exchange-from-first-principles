# Exchange From First Principles

> Proving the same exchange semantics across ACID SQL, in-memory state
> machines, replicated logs, and SQL projections.

[English](#english) · [中文](#中文)

---

## English

### The Question

Why does a modern exchange evolve from a simple ACID SQL transaction model into
something built around command/event facts, in-memory trading cores, replicated
logs, OMS, risk engines, hot/warm/cold paths, cache rebuilds, push streams, and
low-latency runtime work?

This project answers that question by keeping one exchange semantic contract
stable while the architecture changes underneath it.

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

The first full version starts from ACID SQL because it is familiar, durable, and
good at explaining double-entry accounting and transaction boundaries. Later
versions move the same exchange semantics through SQL facts/outbox, a
single-node in-memory trading core, a replicated log core, and SQL projection
consumers.

The contract is exchange-level, not only funds-level: accounting, reservation,
matching, executions, positions, margin/risk admission, projections, caches,
and push recovery all have to survive the architecture migration.

### Documentation Map

The canonical documentation index is [docs/README.md](./docs/README.md). Use
that page for reading order, source-of-truth ownership, and chapter status.

### Version Line

| Chapter | Topic | Status |
| --- | --- | --- |
| 01 | [Exchange semantic contract](./chapters/01-exchange-semantic-contract-go/README.md) | Contract scaffold |
| 02 | [ACID SQL exchange](./chapters/02-acid-sql-exchange-go/README.md) | Contract scaffold |
| 03 | [SQL facts and outbox](./chapters/03-sql-facts-outbox-go/README.md) | Contract scaffold |
| 04 | [Single-node memory core](./chapters/04-single-node-memory-core-java/README.md) | README scaffold |
| 05 | [Replicated log core](./chapters/05-replicated-log-core-aeron-java/README.md) | Runnable Java skeleton |
| 06 | [SQL projection consumers](./chapters/06-sql-projection-consumers/README.md) | README scaffold |
| 07 | [Order book mechanics](./chapters/07-order-book-mechanics/README.md) | README only |
| 08 | [Position and PnL](./chapters/08-position-and-pnl/README.md) | README only |
| 09 | [Margin and pre-trade risk](./chapters/09-margin-and-pretrade-risk/README.md) | README only |
| 10 | [Risk cluster projection](./chapters/10-risk-cluster-projection/README.md) | README only |
| 11 | [Cache coherence and market state](./chapters/11-cache-coherence-and-market-state/README.md) | README only |
| 12 | [Market and execution push](./chapters/12-market-execution-push/README.md) | README only |
| 13 | [Rust hot path](./chapters/13-rust-hot-path/README.md) | Runnable Rust experiment |
| 14 | [Low-latency runtime and networking](./chapters/14-low-latency-runtime-networking/README.md) | README only |
| 15-19 | Trading desk extension ideas | Planned notes only |

### Appendix Prototypes

The current runnable Go funds examples are preserved as prototypes. They are
useful exercises, but they are not the new canonical version line.

| Chapter | Topic | Status |
| --- | --- | --- |
| 90 | [Funds double-entry prototype](./chapters/90-funds-double-entry-prototype-go/README.md) | Runnable Go |
| 91 | [Spot settlement transaction prototype](./chapters/91-spot-settlement-transaction-prototype-go/README.md) | Runnable Go |
| 92 | [Wallet idempotency prototype](./chapters/92-wallet-idempotency-prototype-go/README.md) | Runnable Go plus reconciliation lab |
| 93 | [Command log replay prototype](./chapters/93-command-log-replay-prototype-go/README.md) | Runnable Go |

### Repository Layout

```text
chapters/             version-line chapters and appendix prototypes
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
  version-line order.
- Use the chapter README before changing code in that chapter.

### Toolchain

The runnable parts of the repo currently expect:

- Go 1.22 or newer for Go modules and integration tests;
- Java 21 and Gradle for the replicated-log skeleton;
- Rust stable for the Rust experiment.

### First Commands

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-rust

cd chapters/90-funds-double-entry-prototype-go
go run ./cmd/demo
```

Java chapters require Gradle:

```bash
make test-java
```

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

本项目通过一件事回答这个问题：保持同一套交易所语义契约稳定，同时不断替换它
下面的架构底座。

这不是生产级交易所。它是一个系统设计和工程训练项目，重点是正确性、性能、
可靠性、风控边界和架构演化。

### 核心论点

架构变化本身不是目标。它是在正确性、性能和可靠性之间做取舍。每个阶段都在
选择：排序、持久化、恢复、延迟和运营复杂度分别在哪里付费。

业务语义应该保持稳定：

```text
ordered command + current state -> new state + emitted facts
```

第一版完整语义从 ACID SQL 开始，因为它熟悉、持久，也最适合讲清 double-entry
accounting 和事务边界。后续版本把同一套交易所语义迁移到 SQL facts/outbox、
单机内存交易核心、复制日志核心，以及 SQL projection consumers。

契约是交易所级别的，不只是资金级别：accounting、reservation、matching、
executions、positions、margin/risk admission、projections、caches 和 push
recovery 都必须在架构迁移中保持可解释。

### 文档地图

规范文档索引在 [docs/README.md](./docs/README.md)。阅读顺序、真相源归属和章节
状态都从那里开始。

### 版本线

| 章节 | 主题 | 状态 |
| --- | --- | --- |
| 01 | [交易所语义契约](./chapters/01-exchange-semantic-contract-go/README.md) | 契约脚手架 |
| 02 | [ACID SQL exchange](./chapters/02-acid-sql-exchange-go/README.md) | 契约脚手架 |
| 03 | [SQL facts and outbox](./chapters/03-sql-facts-outbox-go/README.md) | 契约脚手架 |
| 04 | [单机内存核心](./chapters/04-single-node-memory-core-java/README.md) | README 脚手架 |
| 05 | [复制日志核心](./chapters/05-replicated-log-core-aeron-java/README.md) | 可运行 Java 骨架 |
| 06 | [SQL projection consumers](./chapters/06-sql-projection-consumers/README.md) | README 脚手架 |
| 07 | [订单簿机制](./chapters/07-order-book-mechanics/README.md) | 仅 README |
| 08 | [仓位和 PnL](./chapters/08-position-and-pnl/README.md) | 仅 README |
| 09 | [保证金与下单前风控](./chapters/09-margin-and-pretrade-risk/README.md) | 仅 README |
| 10 | [风控集群 projection](./chapters/10-risk-cluster-projection/README.md) | 仅 README |
| 11 | [缓存一致性与市场状态](./chapters/11-cache-coherence-and-market-state/README.md) | 仅 README |
| 12 | [行情与成交推送](./chapters/12-market-execution-push/README.md) | 仅 README |
| 13 | [Rust 热路径](./chapters/13-rust-hot-path/README.md) | 可运行 Rust 实验 |
| 14 | [低延迟运行时与网络](./chapters/14-low-latency-runtime-networking/README.md) | 仅 README |
| 15-19 | 交易台扩展想法 | 仅规划笔记 |

### 附录原型

当前可运行 Go 资金示例保留为 prototype。它们仍然适合练习，但不是新的规范版本线。

| 章节 | 主题 | 状态 |
| --- | --- | --- |
| 90 | [资金 double-entry 原型](./chapters/90-funds-double-entry-prototype-go/README.md) | 可运行 Go |
| 91 | [现货结算事务原型](./chapters/91-spot-settlement-transaction-prototype-go/README.md) | 可运行 Go |
| 92 | [钱包幂等原型](./chapters/92-wallet-idempotency-prototype-go/README.md) | 可运行 Go，含对账实验 |
| 93 | [命令日志重放原型](./chapters/93-command-log-replay-prototype-go/README.md) | 可运行 Go |

### 仓库结构

```text
chapters/             版本线章节和附录原型
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
- 用 [docs/07-chapter-roadmap.md](./docs/07-chapter-roadmap.md) 查看版本线顺序。
- 修改某章代码前，先读该章自己的 README。

### 工具链

当前可运行部分需要：

- Go 1.22 或更新版本，用于 Go modules 和集成测试；
- Java 21 和 Gradle，用于复制日志骨架；
- Rust stable，用于 Rust 实验。

### 第一组命令

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-rust

cd chapters/90-funds-double-entry-prototype-go
go run ./cmd/demo
```

Java 章节需要 Gradle：

```bash
make test-java
```

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
