# Position, Matching, Risk, And Margin

[English](#english) · [中文](#中文)

## English

This document owns the hot-path trading semantics that every version-line
chapter must preserve. These semantics are not late add-ons. They are part of
the minimum exchange contract.

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

## Matching

Matching semantics include:

- price-time priority;
- accept/reject limit orders;
- reserve funds or inventory before admission;
- cancel by order id;
- emit partial-fill and full-fill execution facts;
- release unused locked funds when appropriate.

Chapter 07 explains order-book mechanics. The matching semantics themselves are
defined earlier by the exchange contract.

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
execution facts; position management interprets them.

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

本文档负责每个版本线章节都必须保持的热路径交易语义。这些语义不是后期补丁，
而是最小交易所契约的一部分。

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

## 撮合

撮合语义包括：

- 价格-时间优先级；
- 接受/拒绝限价单；
- 准入前冻结资金或库存；
- 按 order id 撤单；
- 发出部分成交和完全成交 execution facts；
- 在适当时释放未使用的 locked funds。

第 07 章解释订单簿内部机制。撮合语义本身由更早的 exchange contract 定义。

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

项目应将仓位视为一等状态。撮合产生成交事实，仓位管理解释它们。

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
