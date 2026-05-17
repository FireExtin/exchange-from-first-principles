# Chapter Roadmap

This repo is a long-running explanation project: derive exchange and payment
systems from small primitives, then show why each new mechanism becomes
necessary.

The chapter order intentionally starts with money correctness, then builds the
trading hot path, then adds replication, warm/cold paths, caches, push, and
runtime optimization. The point is not to chase advanced infrastructure first;
it is to make the pressure for that infrastructure visible.

For the full documentation catalog, see [Documentation Map](./README.md).

## Status Vocabulary

- Runnable: normal tests or demos run without special TODO work.
- Runnable skeleton: the module compiles and tests a boundary, but most
  business behavior is intentionally absent.
- Design scaffold: the chapter has a problem statement and target model, but no
  runnable implementation yet.
- Planned note: the idea is described in docs, but no chapter directory exists.

## Phase 1: Funds Correctness

Phase 1 now has a first shared semantic layer: `shared/go` defines the funds
commands, events, and typed identifiers; `integration-tests` runs the same
external scenarios against multiple chapter implementations. Chapter 02 exposes
spot settlement as funding events, while chapters 03 and 04 prove that a direct
wallet workflow and a replay-oriented engine can preserve the same observable
funds semantics.

1. `01-double-entry-ledger-go`
   - Primitive: two accounts exchange value.
   - Question: how do we prove money was not created or destroyed?

2. `02-spot-trade-db-go`
   - Primitive: a buyer and seller settle a spot trade.
   - Question: why does a trade need an atomic boundary?

3. `03-wallet-deposit-withdrawal-go`
   - Primitive: outside money enters or leaves the system.
   - Question: why do callbacks, retries, and duplicate notifications need
     idempotency?

4. `04-command-log-replay-go`
   - Primitive: state changes are written as commands/events.
   - Question: how do we recover after a crash and explain what happened?

## Phase 2: Trading Hot Path

5. `05-single-writer-state-machine-java`
   - Primitive: state transitions move into one in-memory writer.
   - Question: why remove database locks from the hot mutation path?

6. `06-simple-matching-engine-java`
   - Primitive: deterministic price-time matching.
   - Question: why does order determine correctness?

7. `07-position-manager`
   - Primitive: execution facts become account/instrument exposure.
   - Question: why is position often more important than the matching
     algorithm for desks, market makers, and external exchange routing?

8. `08-margin-model`
   - Primitive: equity, marks, positions, and frozen funds become admissibility
     constraints.
   - Question: why must margin be explicit instead of hidden inside order
     handlers?

9. `09-desk-pretrade-risk`
   - Primitive: a synchronous gate decides whether an order may enter the
     sequenced core.
   - Question: what should be rejected before it can affect trading state?

10. `10-risk-cluster-projection`
    - Primitive: warm-path consumers continuously rebuild exposure from facts
      and marks.
    - Question: what risks must keep running after order admission?

## Phase 3: Replication And System Boundaries

11. `11-replicated-state-machine-aeron-java`
    - Primitive: ordered commands cross an Aeron/Raft-style replicated log
      boundary.
    - Question: how do replay, backpressure, and failover preserve the same
      business semantics?
    - Status: runnable Java skeleton.

12. `12-oms-ledger-compliance-java-go`
    - Primitive: hot facts feed OMS, ledger, reporting, reconciliation, and
      compliance views.
    - Question: how do we keep cold-path truth explainable without blocking
      the hot path?

13. `13-cache-coherence-and-market-state`
    - Primitive: user, account, permission, instrument, mark-price, and
      external-exchange state must be fast but cannot be casually stale.
    - Question: which caches fail closed, which can lag, and how is every cache
      rebuilt or invalidated?

14. `14-market-execution-push`
    - Primitive: public market data and private execution reports leave the
      core as ordered streams.
    - Question: how do sequence numbers, snapshots, deltas, backpressure, and
      replay make push clients recoverable?

## Phase 4: Runtime Experiments

15. `15-rust-hot-path`
    - Primitive: isolate a possible Rust hot path.
    - Question: what does Rust improve, and what complexity does it add?
    - Status: runnable Rust experiment, not the active development track.

16. `16-low-latency-runtime-networking`
    - Primitive: measure and tune the runtime and network path after semantics
      are stable.
    - Question: what should be optimized only after correctness is nailed down?

## Lab Contract: Ownership, Publication, And Recovery

The in-memory and low-latency principles should enter through existing chapters,
not as a detached theory chapter.

- Chapter 05 should name the private state owned by the single writer and show
  why hot mutation avoids shared storage reads.
- Chapter 11 should separate fast distribution from reliable recovery: committed
  order, snapshot, replay, and failover are the contract.
- Chapter 13 should model owner-published reference and market state:
  snapshot plus versioned deltas, gap detection, rebuild, and fail policy.
- Chapter 14 should show public and private push streams as recoverable
  publications, not best-effort notifications.
- Chapter 16 should measure warmup, allocation, pooling, off-heap/buffer paths,
  and variance before claiming an optimization.

New chapters are useful only when a pressure cannot be expressed inside these
existing boundaries.

## Phase 5: Trading Desk Extension

Phase 5 is not part of the exchange core. It models a trading desk, market
maker, or proprietary trading system built on top of the earlier primitives.
It appears only after the project has reliable market-data streams, execution
reports, positions, risk views, and reconciliation boundaries.

These are planned notes only. Chapter directories for 17-21 do not exist yet.

17. `17-external-market-data-ingestion`
    - Primitive: consume external venue books, trades, tickers, and reference
      data.
    - Question: how do stale data, sequence gaps, and snapshots affect
      pricing and risk?

18. `18-pricing-and-signal-engine`
    - Primitive: convert market state into fair values, marks, and simple
      signals.
    - Question: why is pricing separate from matching, routing, and risk?

19. `19-order-router-and-execution-reports`
    - Primitive: send child orders to external venue mocks and consume
      execution reports.
    - Question: what changes when execution happens outside the local book?

20. `20-hedger-and-best-execution`
    - Primitive: reduce exposure and choose venues under cost, latency, and
      liquidity constraints.
    - Question: how do positions become execution decisions?

21. `21-arbitrage-strategy-demo`
    - Primitive: use multiple venue feeds and router mocks to show a simple
      arbitrage loop.
    - Question: what does a strategy need from market data, pricing, risk,
      routing, and reconciliation?

## Rule

Each chapter should include:

- one README explaining the problem pressure;
- one minimal runnable demo or test when the chapter reaches implementation;
- one post draft in Chinese and one in English when the chapter stabilizes;
- one explicit comparison against the previous model: what changed, what stayed
  semantically identical, and what new operational cost was introduced.

## Core Theory Notes

- `docs/00-project-principles.zh.md` records the project-level engineering
  principles: semantic contract first, shared tests, explicit docs, and simple
  executable examples before advanced infrastructure.
- `shared/go` is the first concrete semantic contract: typed funds commands,
  events, rejection reasons, and a minimal engine interface.
- `integration-tests` contains the first cross-chapter proof that different
  execution styles can preserve the same external business semantics.
- `docs/10-design-paper.md` is the full design-paper version of the project:
  DB truth source, command log, deterministic state machine, replicated log,
  matching, position, risk, margin, caches, push, implementation chapters,
  state ownership, publication, data gravity, and recovery semantics.
- `docs/11-ordering-and-serial-semantics.md` explains why locks, MVCC/CAS, and
  Raft/Paxos are all ways to choose a successful serial history.
- `docs/12-version-contract-and-testing.md` defines the cross-version contract:
  business semantics should survive architecture changes, and the same
  integration scenarios should run against every version.
- `docs/08-truth-source-migration.md` explains how the source of truth moves
  from DB ACID transactions to ordered facts and a replicated state machine.
- `docs/09-position-matching-risk-margin.md` defines the trading-domain
  surface: matching, position management, margin, pre-trade risk, continuous
  risk, and hot/warm/cold paths.
- `docs/13-trading-desk-extension.md` defines the later desk layer: external
  market data, pricing, algos, order routing, hedging, best execution, and
  arbitrage as consumers of exchange-core facts.

---

## 中文

这个仓库是一个长期的解释性项目：从小型原语推导出交易所和支付系统，然后展示
每种新机制为何变得必要。

章节顺序有意从资金正确性开始，然后构建交易热路径，再添加复制、冷热路径、
缓存、推送和运行时优化。重点不是先追逐高级基础设施，而是让那种基础设施的
压力变得可见。

完整文档目录见 [Documentation Map](./README.md)。

## 状态术语

- 可运行：常规测试或 demo 可以运行，不依赖 TODO 实现。
- 可运行骨架：模块能编译并测试边界，但多数业务行为故意留空。
- 设计脚手架：章节已有问题陈述和目标模型，但还没有可运行实现。
- 规划笔记：想法已经写进文档，但尚无章节目录。

## 阶段一：资金正确性

阶段一现在已有第一层共享语义：`shared/go` 定义资金命令、事件和类型化标识；
`integration-tests` 用同一套外部场景测试多个章节实现。第 02 章把现货结算暴露为
资金事件，第 03 和第 04 章证明直接钱包工作流与重放型引擎可以保持相同的外部
资金语义。

1. `01-double-entry-ledger-go`
   - 原语：两个账户交换价值。
   - 问题：如何证明资金没有被创造或销毁？

2. `02-spot-trade-db-go`
   - 原语：买方和卖方结算一笔现货交易。
   - 问题：为什么交易需要一个原子边界？

3. `03-wallet-deposit-withdrawal-go`
   - 原语：外部资金进入或离开系统。
   - 问题：为什么回调、重试和重复通知需要幂等性？

4. `04-command-log-replay-go`
   - 原语：状态变更作为命令/事件写入。
   - 问题：崩溃后如何恢复并解释发生了什么？

## 阶段二：交易热路径

5. `05-single-writer-state-machine-java`
   - 原语：状态转换移入一个内存写入者。
   - 问题：为什么从热变更路径移除数据库锁？

6. `06-simple-matching-engine-java`
   - 原语：确定性价格-时间撮合。
   - 问题：为什么顺序决定正确性？

7. `07-position-manager`
   - 原语：成交事实成为账户/合约敞口。
   - 问题：对于交易台、做市商和外部交易所路由，为什么仓位往往比撮合算法更重要？

8. `08-margin-model`
   - 原语：权益、标记价格、仓位和冻结资金成为准入约束。
   - 问题：为什么保证金必须是显式的，而不是隐藏在订单处理器内部？

9. `09-desk-pretrade-risk`
   - 原语：同步门决定订单是否可以进入排序核心。
   - 问题：什么应该在影响交易状态之前被拒绝？

10. `10-risk-cluster-projection`
    - 原语：温路径消费者持续从事实和标记价格重建敞口。
    - 问题：订单准入后，哪些风险必须持续运行？

## 阶段三：复制和系统边界

11. `11-replicated-state-machine-aeron-java`
    - 原语：有序命令跨过 Aeron/Raft 风格的复制日志边界。
    - 问题：重放、背压和故障转移如何保持相同的业务语义？
    - 状态：可运行 Java 骨架。

12. `12-oms-ledger-compliance-java-go`
    - 原语：热事实流向 OMS、账本、报告、对账和合规视图。
    - 问题：如何在不阻塞热路径的前提下让冷路径真相可解释？

13. `13-cache-coherence-and-market-state`
    - 原语：用户、账户、权限、合约、标记价格和外部交易所状态必须快速但不能随意过期。
    - 问题：哪些缓存故障关闭，哪些可以滞后，每个缓存如何重建或失效？

14. `14-market-execution-push`
    - 原语：公共行情数据和私有成交报告作为有序流离开核心。
    - 问题：序列号、快照、增量、背压和重放如何使推送客户端可恢复？

## 阶段四：运行时实验

15. `15-rust-hot-path`
    - 原语：隔离一个可能的 Rust 热路径。
    - 问题：Rust 改进了什么，又增加了什么复杂度？
    - 状态：可运行 Rust 实验，但不是当前主开发线。

16. `16-low-latency-runtime-networking`
    - 原语：在语义稳定后测量和调优运行时和网络路径。
    - 问题：什么应该只在正确性确定后才优化？

## Lab 契约：归属、发布与恢复

内存化和低延迟原则应进入已有章节，而不是单独变成一章抽象理论。

- 第 05 章应命名单写者拥有的私有状态，并展示热变更为什么避免共享存储读取。
- 第 11 章应区分快速分发和可靠恢复：提交顺序、快照、重放和故障转移才是契约。
- 第 13 章应建模 owner 发布的参考数据和市场状态：快照加版本化增量、缺口检测、
  重建和失败策略。
- 第 14 章应把公共行情和私有成交推送建模为可恢复发布，而不是 best-effort 通知。
- 第 16 章应在声称优化前测量 warmup、分配、对象池、堆外或 buffer 路径以及方差。

只有当某个压力无法被这些已有边界表达时，才新增章节。

## 阶段五：交易台扩展

阶段五不是交易所核心。它建模建立在前面原语之上的交易台、做市或自营交易
系统。只有当项目已经拥有可靠的行情流、成交回报、仓位、风险视图和对账边界
之后，这一层才自然出现。

这些目前只是规划笔记。第 17-21 章的目录尚不存在。

17. `17-external-market-data-ingestion`
    - 原语：消费外部场所订单簿、成交、ticker 和参考数据。
    - 问题：陈旧数据、序列缺口和快照如何影响定价与风控？

18. `18-pricing-and-signal-engine`
    - 原语：把市场状态转换成公允价、标记价格和简单信号。
    - 问题：为什么定价应与撮合、路由和风控分离？

19. `19-order-router-and-execution-reports`
    - 原语：向外部场所 mock 发送子订单并消费成交回报。
    - 问题：当执行发生在本地订单簿之外时，系统发生了什么变化？

20. `20-hedger-and-best-execution`
    - 原语：在成本、延迟和流动性约束下减少敞口并选择场所。
    - 问题：仓位如何变成执行决策？

21. `21-arbitrage-strategy-demo`
    - 原语：用多个场所行情和路由 mock 展示一个简单套利闭环。
    - 问题：策略需要从行情、定价、风控、路由和对账获得什么？

## 规则

每个章节应包含：

- 一个 README 解释问题压力；
- 章节达到实现时，一个最小的可运行演示或测试；
- 章节稳定时，一份中文和一份英文的文章草稿；
- 与上一模型的显式比较：什么改变了，什么语义上保持不变，引入什么新的运营成本。

## 核心理论笔记

- `docs/00-project-principles.zh.md` 记录项目级工程原则：语义契约优先、
  共享测试、显式文档，以及在高级基础设施之前先提供简单可执行示例。
- `shared/go` 是第一份具体语义契约：类型化资金命令、事件、拒绝原因和最小引擎接口。
- `integration-tests` 包含第一组跨章节证明：不同执行方式可以保持相同的外部业务语义。
- `docs/10-design-paper.md` 是项目的完整设计论文版本：DB 真相源、命令日志、
  确定性状态机、复制日志、撮合、仓位、风控、保证金、缓存、推送、实现章节、
  状态归属、发布、data gravity 和恢复语义。
- `docs/11-ordering-and-serial-semantics.md` 解释为什么锁、MVCC/CAS 和 Raft/Paxos
  都是选择成功串行历史的不同方式。
- `docs/12-version-contract-and-testing.md` 定义跨版本契约：业务语义应能在
  架构变化中存活，相同的集成场景应能针对每个版本运行。
- `docs/08-truth-source-migration.md` 解释真相源如何从 DB ACID 事务迁移到
  有序事实和复制状态机。
- `docs/09-position-matching-risk-margin.md` 定义交易域表面：撮合、仓位管理、
  保证金、下单前风控、持续风控和冷热路径。
- `docs/13-trading-desk-extension.md` 定义后期交易台层：外部行情、定价、
  策略、订单路由、对冲、最优执行和套利，作为交易所核心事实的消费者。
