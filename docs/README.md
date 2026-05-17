# Documentation Map

This directory is the project knowledge base. The root `README.md` is the
front door; this file is the map for everything deeper.

The organizing rule is simple:

- project docs explain why the system evolves;
- chapter docs explain one local pressure and one local model;
- shared, integration, tool, and ops docs explain supporting contracts;
- change briefs are historical records, not the current source of truth.

## Recommended Reading Paths

For a first pass:

1. [Goal](./00-goal.md)
2. [Minimal Exchange Design Paper](./10-design-paper.md)
3. [Chapter Roadmap](./07-chapter-roadmap.md)
4. [Version Contract And Testing](./12-version-contract-and-testing.md)
5. Start running chapters from `chapters/01-double-entry-ledger-go`.

For the architecture thread:

1. [System Boundary](./01-system-boundary.md)
2. [Command And Event Model](./02-command-event-model.md)
3. [Replay And Recovery](./03-replay-and-recovery.md)
4. [Truth Source Migration](./08-truth-source-migration.md)
5. [Ordering And Serial Semantics](./11-ordering-and-serial-semantics.md)

For the trading-domain thread:

1. [Position, Matching, Risk, And Margin](./09-position-matching-risk-margin.md)
2. [Trading Desk Extension](./13-trading-desk-extension.md)
3. Chapter READMEs from `chapters/05-*` through `chapters/16-*`.

For implementation work:

1. [Shared Contracts](../shared/README.md)
2. [Integration Test Contract](../integration-tests/README.md)
3. [Chapter Roadmap](./07-chapter-roadmap.md)
4. The README inside the chapter you are changing.

## Current Status

| Area | Current source of truth |
| --- | --- |
| Project goal | [00-goal.md](./00-goal.md) |
| Full design narrative | [10-design-paper.md](./10-design-paper.md) |
| Chapter sequence and status | [07-chapter-roadmap.md](./07-chapter-roadmap.md) |
| Runtime/toolchain entrypoint | [../README.md](../README.md) |
| Cross-chapter semantic contract | [../shared/README.md](../shared/README.md), [../integration-tests/README.md](../integration-tests/README.md) |
| Historical change notes | [change_brief_00.md](./change_brief_00.md), [change_brief_01.md](./change_brief_01.md), [change_brief_02.md](./change_brief_02.md) |

## Project Docs

| Document | Role |
| --- | --- |
| [00-goal.md](./00-goal.md) | Smallest statement of what the project is trying to prove. |
| [00-project-principles.zh.md](./00-project-principles.zh.md) | Chinese project principles and engineering style. |
| [01-system-boundary.md](./01-system-boundary.md) | Language and ownership boundaries: Java, Go, Aeron, Rust. |
| [02-command-event-model.md](./02-command-event-model.md) | The command/event vocabulary used by later stages. |
| [03-replay-and-recovery.md](./03-replay-and-recovery.md) | Minimal recovery checks and replay expectations. |
| [04-aeron-cluster-notes.md](./04-aeron-cluster-notes.md) | Aeron Cluster integration notes. |
| [06-java-go-learning-plan.md](./06-java-go-learning-plan.md) | Learning priorities for Java and Go. |
| [07-chapter-roadmap.md](./07-chapter-roadmap.md) | Canonical chapter order and status. |
| [08-truth-source-migration.md](./08-truth-source-migration.md) | How truth moves from DB rows to ordered facts. |
| [09-position-matching-risk-margin.md](./09-position-matching-risk-margin.md) | Trading-domain boundary across matching, positions, risk, and margin. |
| [10-design-paper.md](./10-design-paper.md) | Full bilingual design paper. |
| [11-ordering-and-serial-semantics.md](./11-ordering-and-serial-semantics.md) | Serial history, locks, optimistic control, and consensus. |
| [12-version-contract-and-testing.md](./12-version-contract-and-testing.md) | How tests preserve business semantics across architecture changes. |
| [13-trading-desk-extension.md](./13-trading-desk-extension.md) | Future desk layer: market data, pricing, routing, hedging, arbitrage. |

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
[13-trading-desk-extension.md](./13-trading-desk-extension.md). They do not yet
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

组织规则：

- 项目文档解释系统为什么演化；
- 章节文档解释一个局部压力和一个局部模型；
- shared、integration、tools、ops 文档解释支撑契约；
- change brief 是历史记录，不是当前真相源。

推荐先读：

1. [Goal](./00-goal.md)
2. [Minimal Exchange Design Paper](./10-design-paper.md)
3. [Chapter Roadmap](./07-chapter-roadmap.md)
4. [Version Contract And Testing](./12-version-contract-and-testing.md)
5. 从 `chapters/01-double-entry-ledger-go` 开始跑代码。

后续新增、删除或重分类文档时，先更新本页。
