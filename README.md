# Exchange From First Principles

> Deriving exchange architecture from first principles through correctness,
> performance, and reliability trade-offs.

[English](#english) · [中文](#中文)

---

## English

### The Question

Why does a modern exchange evolve from a simple database transaction model into
something built around in-memory state machines, event logs, replicated logs,
OMS, risk engines, hot/warm/cold paths, compliance flows, and low-latency
runtime engineering?

This project answers that question by building one stage by stage, with small
runnable programs and design notes.

It is not a production exchange. It is a systems project about correctness,
performance, reliability, risk boundaries, and architecture evolution.

### The Thesis

A modern exchange does not evolve by randomly adding infrastructure.
Architecture changes are not goals by themselves. They are trade-offs among
Correctness, Performance, and Reliability. Each stage chooses where to pay for
ordering, durability, recovery, latency, and operational complexity.

The project still uses a few concrete Architecture Lenses to make those
trade-offs visible: source of truth, ordering model, recovery model, and the
boundary between hot-path facts and downstream views.

Across those versions, the core business shape should remain stable:

```text
old_state + command -> new_state + events
```

This is the precise technical lens: serializable database transactions,
in-memory state machines, and replicated state machines can all present
successful business mutations as an explainable serial history. A database
hides ordering inside storage concurrency control: locks, MVCC, WAL, indexes,
and commit rules. A state-machine model makes ordering explicit through an
append-only command log.

Matching is the cleanest example. Who placed first, who cancelled first, and
who sits ahead at the same price are not implementation details; they define the
correct result. That makes an exchange a useful lens for studying how systems
move from database-centered mutation toward explicit logs, deterministic state
machines, and replicated execution.

### Documentation Map

The canonical documentation index is [docs/README.md](./docs/README.md). Use
that page when you want the full reading order, document catalog, and current
source-of-truth map.

### Roadmap

| Chapter | Topic | Status |
| --- | --- | --- |
| 01 | Double-entry ledger in Go | Runnable |
| 02 | Spot trade with a DB-shaped transaction boundary | Runnable |
| 03 | Wallet deposits, withdrawals, idempotency, and reconciliation | Runnable |
| 04 | Command log, replay, and state reconstruction | Runnable |
| 05 | Java single-writer state-machine boundary | Scaffold |
| 06 | Simple matching engine | Design scaffold |
| 07 | Position manager | Design scaffold |
| 08 | Margin model | Design scaffold |
| 09 | Desk pre-trade risk | Design scaffold |
| 10 | Continuous risk cluster projection | Design scaffold |
| 11 | Aeron/Raft replicated state-machine boundary | Runnable Java skeleton |
| 12 | OMS, ledger, compliance, and hot/warm/cold paths | Scaffold |
| 13 | Caches for users, accounts, permissions, marks, and exchange state | Design scaffold |
| 14 | Market-data and execution push | Design scaffold |
| 15 | Rust hot-path experiment | Runnable experiment |
| 16 | Low-latency runtime and networking | Design scaffold |
| 17-21 | Trading desk extension ideas | Planned notes only |

Each chapter keeps the trade-off concrete by answering four questions:

1. What model is the system in now?
2. What semantic guarantees does it provide?
3. Where does it break?
4. Why is the next stage necessary?

### Repository Layout

```text
chapters/             independent chapter projects
docs/                 design notes, roadmap, and project documentation
shared/               shared event examples, schemas, and Go contracts
integration-tests/    cross-chapter semantic tests
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
- Use [docs/07-chapter-roadmap.md](./docs/07-chapter-roadmap.md) to choose the
  next chapter.
- Use the chapter README before changing code in that chapter.

### Toolchain

The runnable parts of the repo currently expect:

- Go 1.22 or newer for the Go modules and integration tests;
- Java 21 for the Aeron/Gradle chapter;
- Gradle for `make test-java`;
- Rust stable for the chapter 15 workspace.

### First Commands

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-rust

cd chapters/01-double-entry-ledger-go
go run ./cmd/demo
```

Java chapters require Gradle:

```bash
make test-java
```

### Non-Goals

This project is not:

- a production exchange;
- a custody or wallet security system;
- financial software for real funds;
- a complete matching venue;
- a full compliance platform;
- an attempt to implement consensus from scratch.

The goal is to make the architecture explainable from first principles, then
turn each step into small code exercises.

### Author

Built by an engineer with experience in crypto exchange infrastructure, as a
long-term attempt to make trading-system architecture explainable from first
principles.

### License

Apache-2.0. See `LICENSE`.

---

## 中文

### 这是什么

一个从第一性原理出发，逐步构建现代交易所系统的长期项目。

它不是写一个撮合 demo，也不是堆技术栈，而是回答一个问题：

> 一个交易系统，为什么会从最简单的数据库事务模型，逐步演化成现代交易所里的内存状态机、事件日志、复制日志、OMS、风控、冷热路径、合规流程和低延迟运行时工程？

每个阶段都配小型可运行代码和设计说明，解释当前阶段的语义模型、保证、瓶颈，以及为什么需要进入下一阶段。

这不是生产级交易所。它是一个系统设计和工程训练项目，重点是正确性、性能、
可靠性、风控边界和架构演化。

### 核心论点

现代交易所的演化，不是随机添加各种技术。架构变化本身不是目标。它是在
正确性、性能和可靠性之间做取舍。每个阶段都在选择：排序、持久化、恢复、
延迟和运营复杂度分别在哪里付费。

本项目仍会使用几个具体的架构分析视角来让这些取舍可见：真相源、排序模型、
恢复模型，以及热路径事实与下游视图之间的边界。

在这些版本之间，核心业务语义应该保持稳定：

```text
old_state + command -> new_state + events
```

更精确的技术切片是：在足够强的隔离或复制语义下，数据库事务、内存状态机和
复制状态机都可以把成功的业务变更呈现为一条可解释的串行历史。数据库通常把
顺序隐藏在存储层的并发控制里：锁、MVCC、WAL、索引和提交规则。状态机模型则
把顺序显式写进追加命令日志。

撮合是最干净的例子。谁先下单、谁先撤单、同价位谁排在前面，这些不是实现细节，
而是正确结果本身。所以交易所是研究“从数据库中心的状态变更，演化到显式日志、
确定性状态机和复制执行”的一个清晰切片。

当前第一条可运行主线先收在资金域：第 02 章把现货结算暴露为共享资金事件，
第 03 章用直接钱包工作流处理入金、出金和转账，第 04 章用命令日志重放同一组
资金命令。跨章节一致性测试证明：实现底座变了，外部可见的资金语义不变。

### 文档地图

完整文档索引在 [docs/README.md](./docs/README.md)。如果想知道先读什么、
每份文档负责什么、当前真相源在哪里，从那里开始。

### 阶段路线

| 章节 | 主题 | 状态 |
| --- | --- | --- |
| 01 | Go 双分录账本 | 可运行 |
| 02 | 带数据库事务边界的现货成交 | 可运行 |
| 03 | 钱包入金、出金、幂等与对账 | 可运行 |
| 04 | 命令日志、重放与状态重建 | 可运行 |
| 05 | Java 单写者状态机边界 | 脚手架 |
| 06 | 简单撮合引擎 | 设计脚手架 |
| 07 | 仓位管理 | 设计脚手架 |
| 08 | 保证金模型 | 设计脚手架 |
| 09 | 柜台下单前风控 | 设计脚手架 |
| 10 | 持续风控集群投影 | 设计脚手架 |
| 11 | Aeron/Raft 复制状态机边界 | 可运行 Java 骨架 |
| 12 | OMS、账本、合规与冷热路径 | 脚手架 |
| 13 | 用户、账户、权限、标记价格和外部交易所状态缓存 | 设计脚手架 |
| 14 | 行情推送与成交推送 | 设计脚手架 |
| 15 | Rust 热路径实验 | 可运行实验 |
| 16 | 低延迟运行时与网络 | 设计脚手架 |
| 17-21 | 交易台扩展想法 | 仅规划笔记 |

每个章节都通过四个问题把取舍落到具体工程形状上：

1. 当前系统模型是什么？
2. 它提供了什么语义保证？
3. 它在哪里失效？
4. 为什么需要进入下一阶段？

### 仓库结构

```text
chapters/             独立章节项目
docs/                 设计说明、路线图和项目文档
shared/               共享事件样例、schema 和 Go 资金契约
integration-tests/    跨章节一致性测试
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
- 用 [docs/07-chapter-roadmap.md](./docs/07-chapter-roadmap.md) 决定下一章。
- 修改某章代码前，先读该章自己的 README。

### 工具链

当前可运行部分需要：

- Go 1.22 或更新版本，用于 Go modules 和集成测试；
- Java 21，用于 Aeron/Gradle 章节；
- Gradle，用于 `make test-java`；
- Rust stable，用于第 15 章 workspace。

### 第一组命令

```bash
git clone git@github.com:FireExtin/exchange-from-first-principles.git
cd exchange-from-first-principles
make test-go
make test-rust
go test ./integration-tests/...

cd chapters/01-double-entry-ledger-go
go run ./cmd/demo
```

Java 章节需要 Gradle：

```bash
make test-java
```

### 非目标

这个项目不是：

- 生产级交易所；
- 托管或钱包安全系统；
- 处理真实资产的金融软件；
- 完整撮合交易场所；
- 完整合规平台；
- 从零实现共识算法。

目标是从第一性原理把架构讲清楚，再把每一步变成可练习的小代码项目。

### 作者

作者有加密交易所基础设施经验。这个项目是一次长期尝试：从第一性原理解释交易系统架构为什么会长成今天这样。

### License

Apache-2.0。详见 `LICENSE`。
