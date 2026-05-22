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
| Follow the business semantic ramp | [Chapter Roadmap](./07-chapter-roadmap.md), then chapters 01-04; each chapter starts from user action, account map, journal template, and contract checks |
| Follow architecture migration | [Truth Source Migration](./04-truth-source-migration.md), then chapters 05-09 |
| Understand cross-version proof | [Version Contract And Testing](./06-version-contract-and-testing.md), [shared contracts](../shared/README.md), [integration tests](../integration-tests/README.md) |
| Understand hot-path deep dives | [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md), then chapters 10-15 |
| Understand credit and funding extension | [Credit, Margin, And Funding Extension](./09-credit-margin-funding-extension.md), after the hot-path risk docs |
| Run current code | shared Go checks and chapter 08; appendix chapters 90-93 are TODO contract scaffolds |

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
| Credit, margin, and funding extension | [09-credit-margin-funding-extension.md](./09-credit-margin-funding-extension.md) |
| Desk extension notes | [10-trading-desk-extension.md](./10-trading-desk-extension.md) |

## Main Docs

| Doc | Owns |
| --- | --- |
| [00-goal.md](./00-goal.md) | What the project is trying to prove. |
| [01-core-principles.md](./01-core-principles.md) | Semantic contract first, command/event rules, replay, recovery, and tests. |
| [02-design-paper.md](./02-design-paper.md) | The full design-paper narrative. |
| [03-system-boundary.md](./03-system-boundary.md) | Java hot core, Go/SQL service edges and projections, Aeron/Raft ordering. |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | Migration after the minimal business semantics are defined. |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | Locks, MVCC/CAS, serial history, and consensus. |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | How small scenarios compose into cross-version proof. |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | Canonical chapter order: semantic ramp, migration line, deep dives, appendix. |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | Later hot-path surfaces: matching internals, positions, margin, risk, projections. |
| [09-credit-margin-funding-extension.md](./09-credit-margin-funding-extension.md) | Optional credit, collateral, funding, accrual, and liquidation settlement extension. |
| [10-trading-desk-extension.md](./10-trading-desk-extension.md) | Later desk layer: market data, pricing, routing, hedging, strategies. |
| [90-learning-plan.md](./90-learning-plan.md) | Supporting Java/Go learning guidance, not architecture truth. |

## Chapter Docs

| Chapter | README | Status |
| --- | --- | --- |
| 01 | [Custody And User Ledger](../chapters/01-custody-and-user-ledger-go/README.md) | Contract scaffold |
| 02 | [Balance States](../chapters/02-balance-states-go/README.md) | Contract scaffold |
| 03 | [Order Reservation](../chapters/03-order-reservation-go/README.md) | Contract scaffold |
| 04 | [Match And Settlement](../chapters/04-match-and-settlement-go/README.md) | Contract scaffold |
| 05 | [ACID SQL Exchange](../chapters/05-acid-sql-exchange-go/README.md) | Contract scaffold |
| 06 | [SQL Facts And Outbox](../chapters/06-sql-facts-outbox-go/README.md) | Contract scaffold |
| 07 | [Single-Node Memory Core](../chapters/07-single-node-memory-core-java/README.md) | README scaffold |
| 08 | [Replicated Log Core](../chapters/08-replicated-log-core-aeron-java/README.md) | Runnable Java skeleton |
| 09 | [SQL Projection Consumers](../chapters/09-sql-projection-consumers/README.md) | README scaffold |
| 10 | [Order Book Mechanics](../chapters/10-order-book-mechanics/README.md) | README only |
| 11 | [Position And PnL](../chapters/11-position-and-pnl/README.md) | README only |
| 12 | [Margin And Pre-Trade Risk](../chapters/12-margin-and-pretrade-risk/README.md) | README only |
| 13 | [Risk Cluster Projection](../chapters/13-risk-cluster-projection/README.md) | README only |
| 14 | [Cache Coherence And Market State](../chapters/14-cache-coherence-and-market-state/README.md) | README only |
| 15 | [Market And Execution Push](../chapters/15-market-execution-push/README.md) | README only |
| 16 | [Rust Hot Path](../chapters/16-rust-hot-path/README.md) | README only |
| 17 | [Low-Latency Runtime And Networking](../chapters/17-low-latency-runtime-networking/README.md) | README only |
| 18 | [Credit And Collateral Accounts](../chapters/18-credit-and-collateral-accounts-go/README.md) | Contract scaffold |

## Appendix Prototypes

| Chapter | README | Status |
| --- | --- | --- |
| 90 | [Funds Double-Entry Prototype](../chapters/90-funds-double-entry-prototype-go/README.md) | Contract scaffold |
| 91 | [Spot Settlement Transaction Prototype](../chapters/91-spot-settlement-transaction-prototype-go/README.md) | Contract scaffold |
| 92 | [Wallet Idempotency Prototype](../chapters/92-wallet-idempotency-prototype-go/README.md) | Contract scaffold plus reconciliation lab |
| 93 | [Command Log Replay Prototype](../chapters/93-command-log-replay-prototype-go/README.md) | Contract scaffold |

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
| [history/change_brief_04.md](./history/change_brief_04.md) | Business semantic ramp before architecture migration. |
| [history/change_brief_05.md](./history/change_brief_05.md) | Posted facts versus derived/prospective state boundary. |
| [history/change_brief_06.md](./history/change_brief_06.md) | Credit, margin, and funding extension before trading desk. |
| [history/change_brief_07.md](./history/change_brief_07.md) | Credit/margin/funding contract lab. |

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
| 跟随业务语义爬坡 | [Chapter Roadmap](./07-chapter-roadmap.md)，然后读第 01-04 章；每章从用户动作、账户图谱、分录模板和契约检查开始 |
| 跟随架构迁移 | [Truth Source Migration](./04-truth-source-migration.md)，然后读第 05-09 章 |
| 理解跨版本证明 | [Version Contract And Testing](./06-version-contract-and-testing.md)、[shared contracts](../shared/README.md)、[integration tests](../integration-tests/README.md) |
| 理解热路径深挖 | [Position, Matching, Risk, And Margin](./08-position-matching-risk-margin.md)，然后读第 10-15 章 |
| 理解 credit 和 funding 扩展 | [Credit, Margin, And Funding Extension](./09-credit-margin-funding-extension.md)，在热路径风控文档之后阅读 |
| 运行当前代码 | shared Go 检查和第 08 章；附录 90-93 是 TODO 契约脚手架 |

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
| Credit、margin 和 funding 扩展 | [09-credit-margin-funding-extension.md](./09-credit-margin-funding-extension.md) |
| 交易台扩展说明 | [10-trading-desk-extension.md](./10-trading-desk-extension.md) |

## 主文档

| 文档 | 负责 |
| --- | --- |
| [00-goal.md](./00-goal.md) | 项目要证明什么。 |
| [01-core-principles.md](./01-core-principles.md) | 语义契约优先、命令/事件、重放、恢复和测试。 |
| [02-design-paper.md](./02-design-paper.md) | 完整设计论文叙事。 |
| [03-system-boundary.md](./03-system-boundary.md) | Java 热核心、Go/SQL 服务边界和 projection、Aeron/Raft 排序。 |
| [04-truth-source-migration.md](./04-truth-source-migration.md) | 最小业务语义定义完成后的真相源迁移。 |
| [05-ordering-and-serial-semantics.md](./05-ordering-and-serial-semantics.md) | 锁、MVCC/CAS、串行历史和共识。 |
| [06-version-contract-and-testing.md](./06-version-contract-and-testing.md) | 小场景如何组合成跨版本证明。 |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | 规范章节顺序：语义爬坡、迁移线、深挖、附录。 |
| [08-position-matching-risk-margin.md](./08-position-matching-risk-margin.md) | 后续热路径表面：撮合内部、仓位、保证金、风控、projection。 |
| [09-credit-margin-funding-extension.md](./09-credit-margin-funding-extension.md) | 可选 credit、collateral、funding、accrual 和强平结算扩展。 |
| [10-trading-desk-extension.md](./10-trading-desk-extension.md) | 后续交易台层：行情、定价、路由、对冲、策略。 |
| [90-learning-plan.md](./90-learning-plan.md) | Java/Go 学习指导，不是架构真相。 |

## 章节文档

| 章节 | README | 状态 |
| --- | --- | --- |
| 01 | [Custody And User Ledger](../chapters/01-custody-and-user-ledger-go/README.md) | 契约脚手架 |
| 02 | [Balance States](../chapters/02-balance-states-go/README.md) | 契约脚手架 |
| 03 | [Order Reservation](../chapters/03-order-reservation-go/README.md) | 契约脚手架 |
| 04 | [Match And Settlement](../chapters/04-match-and-settlement-go/README.md) | 契约脚手架 |
| 05 | [ACID SQL Exchange](../chapters/05-acid-sql-exchange-go/README.md) | 契约脚手架 |
| 06 | [SQL Facts And Outbox](../chapters/06-sql-facts-outbox-go/README.md) | 契约脚手架 |
| 07 | [Single-Node Memory Core](../chapters/07-single-node-memory-core-java/README.md) | README 脚手架 |
| 08 | [Replicated Log Core](../chapters/08-replicated-log-core-aeron-java/README.md) | 可运行 Java 骨架 |
| 09 | [SQL Projection Consumers](../chapters/09-sql-projection-consumers/README.md) | README 脚手架 |
| 10 | [Order Book Mechanics](../chapters/10-order-book-mechanics/README.md) | 仅 README |
| 11 | [Position And PnL](../chapters/11-position-and-pnl/README.md) | 仅 README |
| 12 | [Margin And Pre-Trade Risk](../chapters/12-margin-and-pretrade-risk/README.md) | 仅 README |
| 13 | [Risk Cluster Projection](../chapters/13-risk-cluster-projection/README.md) | 仅 README |
| 14 | [Cache Coherence And Market State](../chapters/14-cache-coherence-and-market-state/README.md) | 仅 README |
| 15 | [Market And Execution Push](../chapters/15-market-execution-push/README.md) | 仅 README |
| 16 | [Rust Hot Path](../chapters/16-rust-hot-path/README.md) | 仅 README |
| 17 | [Low-Latency Runtime And Networking](../chapters/17-low-latency-runtime-networking/README.md) | 仅 README |
| 18 | [信用与抵押账户](../chapters/18-credit-and-collateral-accounts-go/README.md) | 契约脚手架 |

## 附录原型

| 章节 | README | 状态 |
| --- | --- | --- |
| 90 | [Funds Double-Entry Prototype](../chapters/90-funds-double-entry-prototype-go/README.md) | 契约脚手架 |
| 91 | [Spot Settlement Transaction Prototype](../chapters/91-spot-settlement-transaction-prototype-go/README.md) | 契约脚手架 |
| 92 | [Wallet Idempotency Prototype](../chapters/92-wallet-idempotency-prototype-go/README.md) | 契约脚手架，含对账实验 |
| 93 | [Command Log Replay Prototype](../chapters/93-command-log-replay-prototype-go/README.md) | 契约脚手架 |

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
| [history/change_brief_04.md](./history/change_brief_04.md) | 架构迁移前增加业务语义爬坡。 |
| [history/change_brief_05.md](./history/change_brief_05.md) | 已过账事实与派生/未来状态边界。 |
| [history/change_brief_06.md](./history/change_brief_06.md) | 交易台之前增加 credit、margin 和 funding 扩展。 |
| [history/change_brief_07.md](./history/change_brief_07.md) | Credit/margin/funding 契约实验。 |

## 维护规则

- active docs 保持中英文镜像。
- `AGENTS.md` 保持 English-only。
- 章节路径、职责或状态变化时，同时更新本地图和 [07-chapter-roadmap.md](./07-chapter-roadmap.md)。
- 不为了适配当前术语而重写 [history/](./history/)。
- 未完成 exchange 行为的契约测试必须放在显式 build tag 后。
