# Documentation Map

[English](#english) · [中文](#中文)

## English

This page is the canonical map for active documentation. Historical records
live in [history/](./history/) and are not rewritten when the current thesis
changes.

## Recommended Reading Paths

| Goal | Read |
| --- | --- |
| Understand the project | [Goal](./00-goal.md), [Core Principles](./01-core-principles.md), [Design Paper](./02-design-paper.md) |
| Follow the version line | [Chapter Roadmap](./07-chapter-roadmap.md), then chapters 01-06 |
| Understand truth migration | [Truth Source Migration](./04-truth-source-migration.md), [Version Contract And Testing](./06-version-contract-and-testing.md) |
| Understand hot-path semantics | [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md), then chapters 07-12 |
| Run current code | appendix prototypes 90-93, chapter 05, chapter 13 |

## Current Source Of Truth

| Subject | Owner |
| --- | --- |
| Project goal and hierarchy | [00-goal.md](./00-goal.md) |
| Stable engineering principles | [01-core-principles.md](./01-core-principles.md) |
| Full design narrative | [02-design-paper.md](./02-design-paper.md) |
| Language and ownership boundaries | [03-system-boundary.md](./03-system-boundary.md) |
| Truth-source migration | [04-truth-source-migration.md](./04-truth-source-migration.md) |
| Ordering and serial semantics | [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) |
| Cross-version testing | [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) |
| Chapter order and status | [07-chapter-roadmap.md](./07-chapter-roadmap.md) |
| Hot-path trading semantics | [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) |
| Desk extension notes | [09-trading-desk-extension.md](./09-trading-desk-extension.md) |

## Main Docs

| Doc | Owns |
| --- | --- |
| [00-goal.md](./00-goal.md) | What the project is trying to prove. |
| [01-core-principles.md](./01-core-principles.md) | Semantic contract first, command/event rules, replay, recovery, and tests. |
| [02-design-paper.md](./02-design-paper.md) | The full design-paper narrative. |
| [03-system-boundary.md](./03-system-boundary.md) | Java hot core, Go/SQL service edges and projections, Aeron/Raft ordering. |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | Migration from ACID SQL to facts, memory core, replicated log, and projections. |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | Locks, MVCC/CAS, serial history, and consensus. |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | How shared scenarios prove semantics survived architecture changes. |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | Canonical version-line chapter order. |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | Matching, positions, margin, risk, and projection semantics. |
| [09-trading-desk-extension.md](./09-trading-desk-extension.md) | Later desk layer: market data, pricing, routing, hedging, strategies. |
| [90-learning-plan.md](./90-learning-plan.md) | Supporting Java/Go learning guidance, not architecture truth. |

## Chapter Docs

| Chapter | README | Status |
| --- | --- | --- |
| 01 | [Exchange Semantic Contract](../chapters/01-exchange-semantic-contract-go/README.md) | Contract scaffold |
| 02 | [ACID SQL Exchange](../chapters/02-acid-sql-exchange-go/README.md) | Contract scaffold |
| 03 | [SQL Facts And Outbox](../chapters/03-sql-facts-outbox-go/README.md) | Contract scaffold |
| 04 | [Single-Node Memory Core](../chapters/04-single-node-memory-core-java/README.md) | README scaffold |
| 05 | [Replicated Log Core](../chapters/05-replicated-log-core-aeron-java/README.md) | Runnable Java skeleton |
| 06 | [SQL Projection Consumers](../chapters/06-sql-projection-consumers/README.md) | README scaffold |
| 07 | [Order Book Mechanics](../chapters/07-order-book-mechanics/README.md) | README only |
| 08 | [Position And PnL](../chapters/08-position-and-pnl/README.md) | README only |
| 09 | [Margin And Pre-Trade Risk](../chapters/09-margin-and-pretrade-risk/README.md) | README only |
| 10 | [Risk Cluster Projection](../chapters/10-risk-cluster-projection/README.md) | README only |
| 11 | [Cache Coherence And Market State](../chapters/11-cache-coherence-and-market-state/README.md) | README only |
| 12 | [Market And Execution Push](../chapters/12-market-execution-push/README.md) | README only |
| 13 | [Rust Hot Path](../chapters/13-rust-hot-path/README.md) | Runnable Rust experiment |
| 14 | [Low-Latency Runtime And Networking](../chapters/14-low-latency-runtime-networking/README.md) | README only |

## Appendix Prototypes

| Chapter | README | Status |
| --- | --- | --- |
| 90 | [Funds Double-Entry Prototype](../chapters/90-funds-double-entry-prototype-go/README.md) | Runnable Go |
| 91 | [Spot Settlement Transaction Prototype](../chapters/91-spot-settlement-transaction-prototype-go/README.md) | Runnable Go |
| 92 | [Wallet Idempotency Prototype](../chapters/92-wallet-idempotency-prototype-go/README.md) | Runnable Go plus reconciliation lab |
| 93 | [Command Log Replay Prototype](../chapters/93-command-log-replay-prototype-go/README.md) | Runnable Go |

## Supporting Docs

| Area | README |
| --- | --- |
| Shared contracts | [shared/README.md](../shared/README.md) |
| Integration tests | [integration-tests/README.md](../integration-tests/README.md) |
| Go tools | [tools/go/README.md](../tools/go/README.md) |
| Runbooks | [ops/runbook.md](../ops/runbook.md) |

## Historical Records

| Record | Notes |
| --- | --- |
| [history/change_brief_00.md](./history/change_brief_00.md) | Shared funds contract and prototype integration tests. |
| [history/change_brief_01.md](./history/change_brief_01.md) | Reconciliation lab and earlier architecture-doc expansion. |
| [history/change_brief_02.md](./history/change_brief_02.md) | Documentation consolidation and bilingual cleanup. |
| [history/change_brief_03.md](./history/change_brief_03.md) | Exchange semantic version-line reorganization. |

## Maintenance Rules

- Keep active docs bilingual and mirrored.
- Keep `AGENTS.md` English-only.
- Update this map and [07-chapter-roadmap.md](./07-chapter-roadmap.md) when a
  chapter path, role, or status changes.
- Do not rewrite [history/](./history/) to match current terminology.
- Contract tests for unfinished exchange behavior must stay behind an explicit
  build tag.

---

## 中文

本页是 active documentation 的规范地图。历史记录位于 [history/](./history/)，当前
论点变化时不重写历史记录。

## 推荐阅读路径

| 目标 | 阅读 |
| --- | --- |
| 理解项目 | [Goal](./00-goal.md)、[Core Principles](./01-core-principles.md)、[Design Paper](./02-design-paper.md) |
| 跟随版本线 | [Chapter Roadmap](./07-chapter-roadmap.md)，然后读第 01-06 章 |
| 理解真相迁移 | [Truth Source Migration](./04-truth-source-migration.md)、[Version Contract And Testing](./06-version-contract-and-testing.md) |
| 理解热路径语义 | [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md)，然后读第 07-12 章 |
| 运行当前代码 | 附录原型 90-93、第 05 章、第 13 章 |

## 当前真相源

| 主题 | 归属 |
| --- | --- |
| 项目目标和层级 | [00-goal.md](./00-goal.md) |
| 稳定工程原则 | [01-core-principles.md](./01-core-principles.md) |
| 完整设计叙事 | [02-design-paper.md](./02-design-paper.md) |
| 语言和归属边界 | [03-system-boundary.md](./03-system-boundary.md) |
| 真相源迁移 | [04-truth-source-migration.md](./04-truth-source-migration.md) |
| 排序和串行语义 | [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) |
| 跨版本测试 | [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) |
| 章节顺序和状态 | [07-chapter-roadmap.md](./07-chapter-roadmap.md) |
| 热路径交易语义 | [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) |
| 交易台扩展说明 | [09-trading-desk-extension.md](./09-trading-desk-extension.md) |

## 主文档

| 文档 | 负责 |
| --- | --- |
| [00-goal.md](./00-goal.md) | 项目要证明什么。 |
| [01-core-principles.md](./01-core-principles.md) | 语义契约优先、命令/事件、重放、恢复和测试。 |
| [02-design-paper.md](./02-design-paper.md) | 完整设计论文叙事。 |
| [03-system-boundary.md](./03-system-boundary.md) | Java 热核心、Go/SQL 服务边界和 projection、Aeron/Raft 排序。 |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | 从 ACID SQL 到 facts、内存核心、复制日志和 projection 的迁移。 |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | 锁、MVCC/CAS、串行历史和共识。 |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | 共享场景如何证明语义在架构变化后仍然存活。 |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | 规范版本线章节顺序。 |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | 撮合、仓位、保证金、风控和 projection 语义。 |
| [09-trading-desk-extension.md](./09-trading-desk-extension.md) | 后续交易台层：行情、定价、路由、对冲、策略。 |
| [90-learning-plan.md](./90-learning-plan.md) | Java/Go 学习指导，不是架构真相。 |

## 章节文档

| 章节 | README | 状态 |
| --- | --- | --- |
| 01 | [Exchange Semantic Contract](../chapters/01-exchange-semantic-contract-go/README.md) | 契约脚手架 |
| 02 | [ACID SQL Exchange](../chapters/02-acid-sql-exchange-go/README.md) | 契约脚手架 |
| 03 | [SQL Facts And Outbox](../chapters/03-sql-facts-outbox-go/README.md) | 契约脚手架 |
| 04 | [Single-Node Memory Core](../chapters/04-single-node-memory-core-java/README.md) | README 脚手架 |
| 05 | [Replicated Log Core](../chapters/05-replicated-log-core-aeron-java/README.md) | 可运行 Java 骨架 |
| 06 | [SQL Projection Consumers](../chapters/06-sql-projection-consumers/README.md) | README 脚手架 |
| 07 | [Order Book Mechanics](../chapters/07-order-book-mechanics/README.md) | 仅 README |
| 08 | [Position And PnL](../chapters/08-position-and-pnl/README.md) | 仅 README |
| 09 | [Margin And Pre-Trade Risk](../chapters/09-margin-and-pretrade-risk/README.md) | 仅 README |
| 10 | [Risk Cluster Projection](../chapters/10-risk-cluster-projection/README.md) | 仅 README |
| 11 | [Cache Coherence And Market State](../chapters/11-cache-coherence-and-market-state/README.md) | 仅 README |
| 12 | [Market And Execution Push](../chapters/12-market-execution-push/README.md) | 仅 README |
| 13 | [Rust Hot Path](../chapters/13-rust-hot-path/README.md) | 可运行 Rust 实验 |
| 14 | [Low-Latency Runtime And Networking](../chapters/14-low-latency-runtime-networking/README.md) | 仅 README |

## 附录原型

| 章节 | README | 状态 |
| --- | --- | --- |
| 90 | [Funds Double-Entry Prototype](../chapters/90-funds-double-entry-prototype-go/README.md) | 可运行 Go |
| 91 | [Spot Settlement Transaction Prototype](../chapters/91-spot-settlement-transaction-prototype-go/README.md) | 可运行 Go |
| 92 | [Wallet Idempotency Prototype](../chapters/92-wallet-idempotency-prototype-go/README.md) | 可运行 Go，含对账实验 |
| 93 | [Command Log Replay Prototype](../chapters/93-command-log-replay-prototype-go/README.md) | 可运行 Go |

## 支撑文档

| 区域 | README |
| --- | --- |
| 共享契约 | [shared/README.md](../shared/README.md) |
| 集成测试 | [integration-tests/README.md](../integration-tests/README.md) |
| Go 工具 | [tools/go/README.md](../tools/go/README.md) |
| Runbook | [ops/runbook.md](../ops/runbook.md) |

## 历史记录

| 记录 | 说明 |
| --- | --- |
| [history/change_brief_00.md](./history/change_brief_00.md) | 共享资金契约和原型集成测试。 |
| [history/change_brief_01.md](./history/change_brief_01.md) | 对账实验和早期架构文档扩展。 |
| [history/change_brief_02.md](./history/change_brief_02.md) | 文档合并与中英文整理。 |
| [history/change_brief_03.md](./history/change_brief_03.md) | 交易所语义版本线重排。 |

## 维护规则

- active docs 保持中英文镜像。
- `AGENTS.md` 保持 English-only。
- 章节路径、职责或状态变化时，同时更新本地图和 [07-chapter-roadmap.md](./07-chapter-roadmap.md)。
- 不为了适配当前术语而重写 [history/](./history/)。
- 未完成 exchange 行为的契约测试必须放在显式 build tag 后。
