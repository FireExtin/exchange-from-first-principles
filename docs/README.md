# Documentation Map

[English](#english) · [中文](#中文)

## English

This directory is the project knowledge base. The root `README.md` is the
front door; this file is the map for everything deeper.

Documentation roles are intentionally narrow:

- project docs explain why the system evolves;
- chapter docs explain one local pressure and one local model;
- shared, integration, tool, and ops docs explain supporting contracts;
- history docs record old changes and are not the current source of truth.

## Recommended Reading Paths

For a first pass:

1. [Goal](./00-goal.md)
2. [Core Principles](./01-core-principles.md)
3. [Minimal Exchange Design Paper](./02-design-paper.md)
4. [Chapter Roadmap](./07-chapter-roadmap.md)
5. Start running chapters from `chapters/01-double-entry-ledger-go`.

For the architecture thread:

1. [System Boundary](./03-system-boundary.md)
2. [Truth Source Migration](./04-truth-source-migration.md)
3. [Ordering And Serial Semantics](./05-ordering-and-serial-semantics.md)
4. [Version Contract And Testing](./06-version-contract-and-testing.md)

For the trading-domain thread:

1. [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md)
2. [Trading Desk Extension](./09-trading-desk-extension.md)
3. Chapter READMEs from `chapters/05-*` through `chapters/16-*`.

For implementation work:

1. [Shared Contracts](../shared/README.md)
2. [Integration Test Contract](../integration-tests/README.md)
3. [Chapter Roadmap](./07-chapter-roadmap.md)
4. The README inside the chapter you are changing.

## Current Source Of Truth

| Area | Current source |
| --- | --- |
| Project goal | [00-goal.md](./00-goal.md) |
| Core engineering principles | [01-core-principles.md](./01-core-principles.md) |
| Full design narrative | [02-design-paper.md](./02-design-paper.md) |
| Language and ownership boundaries | [03-system-boundary.md](./03-system-boundary.md) |
| Chapter sequence and status | [07-chapter-roadmap.md](./07-chapter-roadmap.md) |
| Runtime/toolchain entrypoint | [../README.md](../README.md) |
| Cross-chapter semantic contract | [../shared/README.md](../shared/README.md), [../integration-tests/README.md](../integration-tests/README.md) |
| Historical change notes | [history/change_brief_00.md](./history/change_brief_00.md), [history/change_brief_01.md](./history/change_brief_01.md), [history/change_brief_02.md](./history/change_brief_02.md) |

## Main Docs

| Document | Role |
| --- | --- |
| [00-goal.md](./00-goal.md) | Smallest statement of what the project is trying to prove. |
| [01-core-principles.md](./01-core-principles.md) | Semantic contract, command/event rules, replay, recovery, and testing principles. |
| [02-design-paper.md](./02-design-paper.md) | Full bilingual design paper. |
| [03-system-boundary.md](./03-system-boundary.md) | Ownership boundaries: Java, Go, Aeron, Rust. |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | How truth moves from DB rows to ordered facts. |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | Serial history, locks, optimistic control, and consensus. |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | How tests preserve business semantics across architecture changes. |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | Canonical chapter order and status. |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | Trading-domain boundary across matching, positions, risk, and margin. |
| [09-trading-desk-extension.md](./09-trading-desk-extension.md) | Future desk layer: market data, pricing, routing, hedging, arbitrage. |
| [90-learning-plan.md](./90-learning-plan.md) | Supporting Java/Go learning guidance, not an architecture source of truth. |

## Chapter Docs

| Chapter | README | Status |
| --- | --- | --- |
| 01 | [Double-Entry Ledger](../chapters/01-double-entry-ledger-go/README.md) | Runnable Go |
| 02 | [Spot Trade DB Transaction](../chapters/02-spot-trade-db-go/README.md) | Runnable Go |
| 03 | [Wallet Deposit Withdrawal](../chapters/03-wallet-deposit-withdrawal-go/README.md) | Runnable Go plus reconciliation lab |
| 03 lab | [Reconciliation Lab](../chapters/03-wallet-deposit-withdrawal-go/RECONCILIATION_LAB.md) | TODO exercise behind build tag |
| 04 | [Command Log Replay](../chapters/04-command-log-replay-go/README.md) | Runnable Go |
| 05 | [Single-Writer State Machine Java](../chapters/05-single-writer-state-machine-java/README.md) | Design scaffold |
| 06 | [Simple Matching Engine Java](../chapters/06-simple-matching-engine-java/README.md) | Design scaffold |
| 07 | [Position Manager](../chapters/07-position-manager/README.md) | Design scaffold |
| 08 | [Margin Model](../chapters/08-margin-model/README.md) | Design scaffold |
| 09 | [Desk Pre-Trade Risk](../chapters/09-desk-pretrade-risk/README.md) | Design scaffold |
| 10 | [Risk Cluster Projection](../chapters/10-risk-cluster-projection/README.md) | Design scaffold |
| 11 | [Replicated State Machine Aeron Java](../chapters/11-replicated-state-machine-aeron-java/README.md) | Runnable Java skeleton |
| 12 | [OMS, Ledger, Compliance, And Paths](../chapters/12-oms-ledger-compliance-java-go/README.md) | Design scaffold |
| 13 | [Cache Coherence And Market State](../chapters/13-cache-coherence-and-market-state/README.md) | Design scaffold |
| 14 | [Market And Execution Push](../chapters/14-market-execution-push/README.md) | Design scaffold |
| 15 | [Rust Hot Path](../chapters/15-rust-hot-path/README.md) | Runnable Rust experiment |
| 16 | [Low-Latency Runtime And Networking](../chapters/16-low-latency-runtime-networking/README.md) | Design scaffold |

Chapters 17-21 are future desk-extension ideas described in
[09-trading-desk-extension.md](./09-trading-desk-extension.md). They do not yet
have chapter directories.

## Supporting Docs

| Document | Role |
| --- | --- |
| [../shared/README.md](../shared/README.md) | Shared command/event contracts. |
| [../integration-tests/README.md](../integration-tests/README.md) | Cross-chapter semantic test contract. |
| [../tools/go/README.md](../tools/go/README.md) | Go tooling placeholders. |
| [../tools/readme-matrix/README.md](../tools/readme-matrix/README.md) | README comparison format helper. |
| [../ops/runbook.md](../ops/runbook.md) | Operational runbook notes. |
| [../ops/flamegraph-notes.md](../ops/flamegraph-notes.md) | Profiling and flamegraph notes. |

## Historical Records

| Document | Role |
| --- | --- |
| [history/change_brief_00.md](./history/change_brief_00.md) | Shared funds contract and integration-test change record. |
| [history/change_brief_01.md](./history/change_brief_01.md) | Reconciliation lab and design-paper change record. |
| [history/change_brief_02.md](./history/change_brief_02.md) | Documentation organization refresh record. |

## Maintenance Rules

- Update the root `README.md` when setup commands, runnable status, or the
  first reading path changes.
- Update this file when a Markdown document is added, removed, or reclassified.
- Update [07-chapter-roadmap.md](./07-chapter-roadmap.md) when a chapter changes
  phase or status.
- Keep change briefs append-only. They explain what changed at a point in time;
  they should not be edited to describe the latest state.

---

## 中文

本目录是项目知识库。根目录 `README.md` 是入口；本文件是更完整的文档地图。

文档职责有意保持窄边界：

- 项目文档解释系统为什么演化；
- 章节文档解释一个局部压力和一个局部模型；
- shared、integration、tools、ops 文档解释支撑契约；
- history 文档记录历史变更，不是当前真相源。

## 推荐阅读路径

第一次阅读：

1. [Goal](./00-goal.md)
2. [Core Principles](./01-core-principles.md)
3. [Minimal Exchange Design Paper](./02-design-paper.md)
4. [Chapter Roadmap](./07-chapter-roadmap.md)
5. 从 `chapters/01-double-entry-ledger-go` 开始运行章节代码。

架构主线：

1. [System Boundary](./03-system-boundary.md)
2. [Truth Source Migration](./04-truth-source-migration.md)
3. [Ordering And Serial Semantics](./05-ordering-and-serial-semantics.md)
4. [Version Contract And Testing](./06-version-contract-and-testing.md)

交易领域主线：

1. [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md)
2. [Trading Desk Extension](./09-trading-desk-extension.md)
3. `chapters/05-*` 到 `chapters/16-*` 的章节 README。

实现工作：

1. [Shared Contracts](../shared/README.md)
2. [Integration Test Contract](../integration-tests/README.md)
3. [Chapter Roadmap](./07-chapter-roadmap.md)
4. 正在修改的章节 README。

## 当前真相源

| 范围 | 当前来源 |
| --- | --- |
| 项目目标 | [00-goal.md](./00-goal.md) |
| 核心工程原则 | [01-core-principles.md](./01-core-principles.md) |
| 完整设计叙事 | [02-design-paper.md](./02-design-paper.md) |
| 语言与归属边界 | [03-system-boundary.md](./03-system-boundary.md) |
| 章节顺序和状态 | [07-chapter-roadmap.md](./07-chapter-roadmap.md) |
| 运行时和工具链入口 | [../README.md](../README.md) |
| 跨章节语义契约 | [../shared/README.md](../shared/README.md), [../integration-tests/README.md](../integration-tests/README.md) |
| 历史变更记录 | [history/change_brief_00.md](./history/change_brief_00.md), [history/change_brief_01.md](./history/change_brief_01.md), [history/change_brief_02.md](./history/change_brief_02.md) |

## 主线文档

| 文档 | 职责 |
| --- | --- |
| [00-goal.md](./00-goal.md) | 项目要证明什么的最小陈述。 |
| [01-core-principles.md](./01-core-principles.md) | 语义契约、命令/事件规则、重放、恢复和测试原则。 |
| [02-design-paper.md](./02-design-paper.md) | 完整双语设计论文。 |
| [03-system-boundary.md](./03-system-boundary.md) | Java、Go、Aeron、Rust 的归属边界。 |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | 真相如何从 DB 行迁移到有序事实。 |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | 串行历史、锁、乐观控制和共识。 |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | 测试如何在架构变化中保持业务语义。 |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | 规范章节顺序和状态。 |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | 撮合、仓位、风控和保证金的交易领域边界。 |
| [09-trading-desk-extension.md](./09-trading-desk-extension.md) | 未来交易台层：行情、定价、路由、对冲、套利。 |
| [90-learning-plan.md](./90-learning-plan.md) | Java/Go 学习指导，不是架构真相源。 |

## 章节文档

| 章节 | README | 状态 |
| --- | --- | --- |
| 01 | [Double-Entry Ledger](../chapters/01-double-entry-ledger-go/README.md) | 可运行 Go |
| 02 | [Spot Trade DB Transaction](../chapters/02-spot-trade-db-go/README.md) | 可运行 Go |
| 03 | [Wallet Deposit Withdrawal](../chapters/03-wallet-deposit-withdrawal-go/README.md) | 可运行 Go，包含对账实验 |
| 03 lab | [Reconciliation Lab](../chapters/03-wallet-deposit-withdrawal-go/RECONCILIATION_LAB.md) | build tag 后的 TODO 练习 |
| 04 | [Command Log Replay](../chapters/04-command-log-replay-go/README.md) | 可运行 Go |
| 05 | [Single-Writer State Machine Java](../chapters/05-single-writer-state-machine-java/README.md) | 设计脚手架 |
| 06 | [Simple Matching Engine Java](../chapters/06-simple-matching-engine-java/README.md) | 设计脚手架 |
| 07 | [Position Manager](../chapters/07-position-manager/README.md) | 设计脚手架 |
| 08 | [Margin Model](../chapters/08-margin-model/README.md) | 设计脚手架 |
| 09 | [Desk Pre-Trade Risk](../chapters/09-desk-pretrade-risk/README.md) | 设计脚手架 |
| 10 | [Risk Cluster Projection](../chapters/10-risk-cluster-projection/README.md) | 设计脚手架 |
| 11 | [Replicated State Machine Aeron Java](../chapters/11-replicated-state-machine-aeron-java/README.md) | 可运行 Java 骨架 |
| 12 | [OMS, Ledger, Compliance, And Paths](../chapters/12-oms-ledger-compliance-java-go/README.md) | 设计脚手架 |
| 13 | [Cache Coherence And Market State](../chapters/13-cache-coherence-and-market-state/README.md) | 设计脚手架 |
| 14 | [Market And Execution Push](../chapters/14-market-execution-push/README.md) | 设计脚手架 |
| 15 | [Rust Hot Path](../chapters/15-rust-hot-path/README.md) | 可运行 Rust 实验 |
| 16 | [Low-Latency Runtime And Networking](../chapters/16-low-latency-runtime-networking/README.md) | 设计脚手架 |

第 17-21 章是 [09-trading-desk-extension.md](./09-trading-desk-extension.md)
描述的未来交易台扩展想法，目前尚无章节目录。

## 支撑文档

| 文档 | 职责 |
| --- | --- |
| [../shared/README.md](../shared/README.md) | 共享命令/事件契约。 |
| [../integration-tests/README.md](../integration-tests/README.md) | 跨章节语义测试契约。 |
| [../tools/go/README.md](../tools/go/README.md) | Go 工具占位。 |
| [../tools/readme-matrix/README.md](../tools/readme-matrix/README.md) | README 比较格式辅助说明。 |
| [../ops/runbook.md](../ops/runbook.md) | 运维 runbook 笔记。 |
| [../ops/flamegraph-notes.md](../ops/flamegraph-notes.md) | profiling 和 flamegraph 笔记。 |

## 历史记录

| 文档 | 职责 |
| --- | --- |
| [history/change_brief_00.md](./history/change_brief_00.md) | 共享资金契约和集成测试变更记录。 |
| [history/change_brief_01.md](./history/change_brief_01.md) | 对账实验和设计论文变更记录。 |
| [history/change_brief_02.md](./history/change_brief_02.md) | 文档组织刷新记录。 |

## 维护规则

- 当 setup 命令、可运行状态或第一阅读路径变化时，更新根目录 `README.md`。
- 新增、删除或重分类 Markdown 文档时，更新本文件。
- 章节阶段或状态变化时，更新 [07-chapter-roadmap.md](./07-chapter-roadmap.md)。
- change brief 保持追加式历史记录。它们解释某个时间点发生了什么，不应被改写为
  最新状态说明。
