# Position, Matching, Risk, And Margin

[English](#english) · [中文](#中文)

## English

This document owns the later hot-path trading surfaces that every architecture
version must eventually preserve. It is not the first place where orders or
matching are introduced; chapters 03 and 04 introduce the small business
semantics first.

## Target Shape

```text
command
  -> pre-trade risk and margin admission
  -> reservation
  -> sequenced trading core
  -> matching / execution facts
  -> position updates
  -> warm and cold projections
```

Different versions may place these responsibilities in SQL transactions,
in-memory state, a replicated log, or SQL consumers. The external meaning should
stay stable.

## Posted Versus Derived State

Execution settlement can create posted ledger facts. Positions are first-class
state derived from execution facts. Mark-based valuation, unrealized PnL,
margin requirements, and continuous risk views are derived or prospective state
until an explicit accrual, settlement, liquidation, or fee event posts journal
entries.

Pre-trade risk can reject a command before it reaches matching. Warm-path risk
can alert or trigger downstream workflows. Neither boundary should silently
post ledger truth.

## Matching

Matching semantics include:

- price-time priority;
- accept/reject limit orders;
- reserve funds or inventory before admission;
- cancel by order id;
- emit partial-fill and full-fill execution facts;
- release unused locked funds when appropriate.

Chapter 04 introduces minimal matching and settlement. Chapter 10 then explains
order-book mechanics in more detail.

## Double-Entry Execution Settlement

Execution facts must be explainable as ledger postings. For example, if user A
buys 1 BTC from user B at 60,000 USD and pays a 10 USD buyer fee:

```text
USD ledger:
Debit   User_A_USD_Locked         60,010
Credit  User_B_USD_Available      60,000
Credit  Fee_Revenue_USD               10

BTC ledger:
Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

Rules:

- USD and BTC balance separately;
- fees are facts and postings, not hidden balance edits;
- reservation, execution, release, and cancellation all need ledger
  explanations;
- memory and replicated-log versions may emit facts first, but those facts must
  still explain the same postings.

## Positions

Position semantics include:

- net quantity by account and instrument;
- average-entry placeholder;
- realized and unrealized PnL placeholders;
- deterministic application of execution reports;
- rebuild from execution facts.

The project should treat positions as first-class state. Matching creates
execution facts; position management interprets them. Realized PnL only becomes
posted ledger truth at an explicit settlement/accrual boundary. Unrealized PnL
is a valuation view until then.

## Margin And Pre-Trade Risk

Admission semantics include:

- account enabled/disabled;
- kill switch;
- available and locked balances;
- notional limit;
- price band;
- initial and maintenance margin placeholders;
- accept/reject facts with reasons.

Pre-trade risk is synchronous and protects command admission. It is separate
from the continuous risk cluster.

## Risk Projection

Continuous risk is a warm-path consumer:

- consume executions, marks, deposits, withdrawals, funding, and adjustments;
- aggregate exposure by account and instrument;
- emit alerts;
- rebuild from event history with cursors and gap detection.

## Hot, Warm, And Cold Paths

Hot path:

```text
sequenced command -> admission -> reservation/match/apply -> facts
```

Warm path:

```text
facts + marks -> position/risk/cache projections
```

Cold path:

```text
facts -> ledger reports, reconciliation, compliance, audit
```

The cold path is not less important. It is simply not allowed to block order
entry.

---

## 中文

本文档负责每个架构版本最终都必须保持的后续热路径交易表面。它不是第一次介绍
订单或撮合的地方；第 03 和第 04 章会先引入小的业务语义。

## 目标形状

```text
command
  -> pre-trade risk and margin admission
  -> reservation
  -> sequenced trading core
  -> matching / execution facts
  -> position updates
  -> warm and cold projections
```

不同版本可以把这些职责放在 SQL 事务、内存状态、复制日志或 SQL consumer 中。
外部含义应该保持稳定。

## 已过账状态与派生状态

成交结算可以产生已过账 ledger facts。仓位是从 execution facts 派生出来的一等状态。
基于 mark 的估值、未实现 PnL、保证金要求和持续风控视图，在显式 accrual、
settlement、liquidation 或 fee 事件提交 journal entries 之前，都属于派生或未来状态。

下单前风控可以在命令进入撮合前拒绝它。温路径风控可以报警或触发下游流程。两者都
不应该偷偷提交 ledger truth。

## 撮合

撮合语义包括：

- 价格-时间优先级；
- 接受/拒绝限价单；
- 准入前冻结资金或库存；
- 按 order id 撤单；
- 发出部分成交和完全成交 execution facts；
- 在适当时释放未使用的 locked funds。

第 04 章先介绍最小撮合和结算。第 10 章再更细地解释订单簿内部机制。

## Double-Entry 成交结算

成交事实必须能解释成账本分录。例如用户 A 以 60,000 USD 从用户 B 买 1 BTC，
并支付 10 USD 买方手续费：

```text
USD ledger:
Debit   User_A_USD_Locked         60,010
Credit  User_B_USD_Available      60,000
Credit  Fee_Revenue_USD               10

BTC ledger:
Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

规则：

- USD 和 BTC 分别平衡；
- 手续费是事实和 posting，不是隐藏的余额改动；
- reservation、execution、release 和 cancellation 都需要账本解释；
- 内存和复制日志版本可以先发事实，但这些事实仍必须解释同一组 postings。

## 仓位

仓位语义包括：

- 按账户和 instrument 计算净数量；
- 平均开仓价占位；
- 已实现和未实现 PnL 占位；
- 确定性应用 execution report；
- 从 execution facts 重建。

项目应将仓位视为一等状态。撮合产生成交事实，仓位管理解释它们。已实现 PnL 只有在
显式 settlement/accrual 边界才会成为已过账 ledger truth。未实现 PnL 在此之前是
估值视图。

## 保证金和下单前风控

准入语义包括：

- 账户启用/禁用；
- kill switch；
- 可用和冻结余额；
- 名义值限制；
- 价格区间；
- 初始和维持保证金占位；
- 带原因的 accept/reject facts。

下单前风控是同步的，保护 command admission。它不同于持续风控集群。

## 风控 Projection

持续风控是 warm-path consumer：

- 消费 executions、marks、deposits、withdrawals、funding 和 adjustments；
- 按账户和 instrument 聚合 exposure；
- 发出 alert；
- 使用 cursor 和 gap detection 从事件历史重建。

## 热、温、冷路径

热路径：

```text
sequenced command -> admission -> reservation/match/apply -> facts
```

温路径：

```text
facts + marks -> position/risk/cache projections
```

冷路径：

```text
facts -> ledger reports, reconciliation, compliance, audit
```

冷路径并不次要。它只是不允许阻塞订单入口。
