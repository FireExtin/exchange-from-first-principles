# Reconciliation Lab

This lab is intentionally unfinished.

The goal is to practice the shape of a real funds reconciliation service
without building a full payment company. The package defines the contract,
mocks, and opt-in tests. The matching, normalization, and adjustment logic are
left as implementation exercises.

## Core Idea

Reconciliation is not balance arithmetic. It is the process of turning multiple
asynchronous, delayed, duplicated, and sometimes corrected data sources into an
auditable funds fact.

The lab models three external inputs:

- provider callbacks: real-time, duplicated, out of order, and sometimes
  corrective;
- provider settlement reports: batch based, with gross, fee, net, and payout
  identifiers;
- bank/custody/on-chain records: closer to final cash movement, but delayed.

The output should be a report, not an automatic ledger mutation.

## Reconciliation Matching Is Not Trading Matching

In this lab, matching does not mean an order-book matching engine.

It means record matching: deciding whether internal ledger entries, business
orders, provider callbacks, provider settlement rows, bank statements, custody
reports, or on-chain records describe the same money movement.

For example, an internal deposit may credit a user for `100 USD`, a provider
settlement row may show `gross=100, fee=1.5, net=98.5`, and a bank statement may
show one delayed payout of `98.5`. A plain SQL join may report an amount
mismatch. A reconciliation matcher should classify it as a fee-adjusted match.

Use names such as `ReconciliationMatcher`, `record matcher`, or `matching
rules`. Avoid the standalone name `MatchingEngine` here because it is too easy
to confuse with trading/order-book matching.

## Complexity Boundary

This lab keeps v1 deliberately small. It should cover:

- exact id or business id matching;
- gross / fee / net explanation;
- batch-level provider rows matched to one bank or custody payout;
- timing differences before settlement finality;
- duplicates and manual-review exceptions.

It should not cover yet:

- FX, tax, or rounding tolerance systems;
- fuzzy matching over weak descriptions;
- multiple partial refunds over the same charge;
- full chargeback dispute workflows;
- a generic rule DSL or production database schema.

## Exercise Contract

The unfinished package is:

```text
internal/reconciliation
```

It exposes:

- `RawRecord`: immutable external or internal source data;
- `NormalizedRecord`: one common shape for matching;
- `MatchKind`: how records were linked, such as exact or fee-adjusted batch;
- `Reconcile`: deterministic matching and discrepancy classification;
- `AdjustmentProposal`: a suggested correction;
- `AdjustmentJournal`: audit evidence for human-reviewed adjustments.

The TODO methods currently panic. That is deliberate.

## Run The Exercise Tests

Default chapter tests stay green:

```bash
go test ./...
```

To start the reconciliation lab, run:

```bash
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

These tests are expected to fail until the lab is implemented.

## Suggested Implementation Order

1. Implement `ProviderCallbackSimulator`.
   - Produce succeeded, failed, partial refund, chargeback, duplicate, and
     out-of-order callback records.
   - Keep records deterministic so tests remain stable.

2. Implement provider report and bank/chain mocks.
   - Provider report should expose gross, fee, net, settlement date, and batch
     id.
   - Bank/chain record should represent final cash/custody movement.

3. Implement `Normalize`.
   - Preserve source, identifiers, business id, amount, fee, net amount, status,
     settlement date, and raw payload hash.
   - Do not silently repair bad raw data.

4. Implement `Reconcile`.
   - Prefer exact `BusinessID` matching.
   - Detect duplicate external records by `(Source, Provider, ExternalID)`.
   - Return `Match.Kind`, `RuleID`, `Confidence`, and `Explanation` so the
     report says why records were linked.
   - Support the simple batch case where provider gross minus fees equals a
     bank/custody payout.
   - Treat future settlement dates as timing differences, not bugs.
   - Classify amount mismatch and partial refund separately.
   - Generate adjustment proposals, but do not mutate wallet balances.

5. Implement `AdjustmentJournal`.
   - Append operator, note, evidence, and timestamp.
   - Return immutable-looking copies from `Entries`.

## What This Should Teach

- External events can be repeated; internal funds effects must apply once.
- Provider success, provider settlement, and bank/custody finality are different
  facts.
- Gross, fee, and net must close.
- Timing differences are not automatically bugs.
- Unknown differences should enter an exception workflow.
- Manual adjustment is an auditable journal, not an invisible balance edit.

---

## 中文

这个实验刻意没有完成。

目标不是做一个完整支付公司，而是练习真实资金对账服务的形状：多个异步、
延迟、重复、乱序、会修正的数据源，如何归并成一个可审计的资金事实。

本实验建模三类外部输入：

- provider callback：实时，但可能重复、乱序、修正；
- provider settlement report：批量，包含 gross、fee、net 和 payout/batch 标识；
- bank/custody/on-chain record：更接近最终资金事实，但通常延迟。

输出应该是报告，而不是自动修改 ledger 或 wallet balance。

## 对账匹配不是交易撮合

本实验里的 matching 不是 order book 的撮合引擎。

它指的是记录匹配：判断内部 ledger entry、业务订单、provider callback、
provider settlement row、银行流水、托管报告或链上记录，是否描述同一笔资金流。

例如，内部入金给用户记了 `100 USD`，provider settlement row 显示
`gross=100, fee=1.5, net=98.5`，银行流水延迟到账 `98.5`。普通 SQL join
可能只会报 amount mismatch；对账匹配模块应该把它归类为 fee-adjusted match。

这里建议使用 `ReconciliationMatcher`、`record matcher`、`matching rules`
这样的名字。不要单独叫 `MatchingEngine`，因为它太容易和交易撮合混淆。

## 复杂度边界

v1 只保留最小现实复杂度：

- exact id 或 business id 匹配；
- gross / fee / net 的解释；
- 多条 provider row 对一笔 bank/custody payout 的 batch-level match；
- settlement finality 之前的 timing difference；
- duplicate 和 manual-review exception。

v1 暂不处理：

- FX、税费、rounding tolerance 系统；
- 基于弱描述的 fuzzy matching；
- 同一 charge 上多次 partial refund；
- 完整 chargeback dispute 流程；
- 通用规则 DSL 或生产级数据库 schema。

## 练习契约

未完成的包位于：

```text
internal/reconciliation
```

它暴露：

- `RawRecord`：不可变的外部或内部原始数据；
- `NormalizedRecord`：用于匹配的统一形状；
- `MatchKind`：记录是如何被归并的，例如 exact 或 fee-adjusted batch；
- `Reconcile`：确定性的匹配和差异归因；
- `AdjustmentProposal`：调整建议；
- `AdjustmentJournal`：人工处理后的审计证据。

当前 TODO 方法会 panic。这是故意的。

## 运行练习测试

默认章节测试保持通过：

```bash
go test ./...
```

开始练习对账实验时运行：

```bash
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

这些测试在你实现实验前应该失败。

## 建议实现顺序

1. 实现 `ProviderCallbackSimulator`。
   - 生成成功、失败、部分退款、拒付、重复、乱序 callback。
   - 保持 deterministic，避免测试抖动。

2. 实现 provider report 和 bank/chain mock。
   - provider report 表达 gross、fee、net、settlement date 和 batch id。
   - bank/chain record 表达最终现金、托管或链上资金变化。

3. 实现 `Normalize`。
   - 保留 source、标识、business id、金额、手续费、净额、状态、结算日和
     raw payload hash。
   - 不要默默修复坏数据。

4. 实现 `Reconcile`。
   - 优先使用 `BusinessID` 精确匹配。
   - 用 `(Source, Provider, ExternalID)` 识别重复外部记录。
   - 返回 `Match.Kind`、`RuleID`、`Confidence` 和 `Explanation`，让报告说明
     为什么这些记录被归并到一起。
   - 支持最简单的 batch 场景：provider gross 减去 fees 等于
     bank/custody payout。
   - settlement date 未到时归类为 timing difference，而不是 bug。
   - 区分 amount mismatch 和 partial refund。
   - 生成 adjustment proposal，但不能修改 wallet balance。

5. 实现 `AdjustmentJournal`。
   - 记录 operator、note、evidence 和 timestamp。
   - `Entries` 返回拷贝，避免外部改写审计记录。

## 这个实验要练什么

- 外部事件可以重复，内部资金效果只能 apply once。
- provider success、provider settlement、bank/custody finality 是三种不同事实。
- gross、fee、net 必须能闭合。
- 时间差不一定是 bug。
- 未知差异应该进入 exception workflow。
- 人工调整是可审计 journal，不是偷偷改余额。
