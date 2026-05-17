# Position, Matching, Risk, And Margin

[English](#english) · [中文](#中文)

## English

This note defines the next domain surface before implementation. It is not a
complete exchange design. It is a scaffold for the chapters that should come
after ledger, wallet, and replay.

The core idea:

> Matching creates execution facts. Position management interprets fills.
> Pre-trade risk blocks dangerous commands before admission. The risk cluster
> watches continuous exposure after facts are emitted.

## Target Shape

```text
gateway
  -> desk pre-trade risk
  -> sequencer
  -> matching engine
  -> execution events
  -> position manager
  -> risk cluster
  -> cold path
```

Each component has a different responsibility. Keeping those responsibilities
separate is more important than implementing a clever algorithm early.

## Matching Engine

The matching engine owns deterministic order-book state.

First scope:

- accept or reject limit orders;
- maintain bid/ask books;
- expose top-of-book;
- match by price-time priority;
- emit trade execution events;
- cancel by order id.

It should not call databases, pricing services, or remote risk systems on the
hot path. It consumes sequenced commands and emits facts.

For the first implementation, keep it single-threaded. Single-threaded matching
is not a weakness at this stage; it makes correctness, replay, and explanations
simple.

## Position Manager

The position manager owns account and instrument exposure derived from fills.

First scope:

- net position by account and instrument;
- average entry price;
- realized PnL placeholder;
- unrealized PnL placeholder;
- cash and frozen-cash hooks;
- execution-report application.

For the trading domain, position management is often more important than the matching
algorithm. A trading desk can route orders to external exchanges, but it must
always know its position, exposure, available balance, and risk.

## Desk Pre-Trade Risk

Desk pre-trade risk is synchronous and close to order entry. It answers:

- is the account allowed to trade?
- is the instrument enabled?
- is the order notional within limit?
- is the price within band?
- is there enough available balance or margin?
- is the account or strategy killed?

This gate protects command admission. It should be deterministic and fast.

This is not the same as the risk cluster. Pre-trade risk blocks bad orders
before they enter the sequenced core.

## Risk Cluster

The risk cluster is continuous and event-driven. It consumes executions,
positions, mark prices, deposits, withdrawals, funding, and manual adjustments.

First scope:

- consume execution events;
- consume mark-price events;
- aggregate exposure by account and instrument;
- emit risk alerts;
- rebuild state from event history.

It may be warm path rather than the hottest path. It should be replayable and
observable. It can be sharded by account, instrument, or strategy once the
single-node model is clear.

## Hot, Warm, And Cold Paths

Hot path:

```text
sequenced command -> pre-trade check -> match/apply -> emit facts
```

Warm path:

```text
events + marks -> position/risk projections -> alerts/actions
```

Cold path:

```text
events -> reporting/audit/reconciliation/compliance export
```

The cold path is not less important. It is simply not allowed to block order
entry. It consumes facts produced by the hot path and makes them explainable to
humans, regulators, support, finance, and incident reviewers.

## Minimal Margin Model

The first margin implementation should use scaled integers and simple formulas.
Do not start with full portfolio margin or exchange-specific liquidation rules.

Spot:

```text
available_cash = cash_balance - frozen_cash
available_base = base_balance - frozen_base

buy_required_quote = price * quantity + fee_buffer
sell_required_base = quantity
```

Linear contract placeholder:

```text
notional = abs(position_qty) * mark_price

if position_qty > 0:
  unrealized_pnl = position_qty * (mark_price - entry_price)
else:
  unrealized_pnl = abs(position_qty) * (entry_price - mark_price)

initial_margin = notional * initial_margin_rate
maintenance_margin = notional * maintenance_margin_rate
equity = wallet_balance + unrealized_pnl
available = equity - initial_margin - frozen_margin - fee_buffer
```

Liquidation threshold placeholder:

```text
equity <= maintenance_margin + liquidation_fee_buffer
```

These formulas are not meant to be final. They create the first testable model
for explaining:

- why mark price matters;
- why position state differs from order state;
- why risk must keep running after order admission;
- why margin and liquidation should be explicit domain chapters.

## Chapter Plan

1. `10-position-manager`
   - apply execution reports;
   - update position;
   - expose replayable state.

2. `11-simple-matching-engine-java`
   - deterministic price-time matching;
   - top-of-book;
   - execution events.

3. `12-desk-pretrade-risk`
   - simple synchronous order checks;
   - account/instrument/price/notional/balance gates.

4. `13-risk-cluster-projection`
   - consume events and marks;
   - aggregate exposure;
   - emit alerts;
   - rebuild from history.

5. `14-margin-model`
   - spot available balance;
   - linear notional;
   - initial and maintenance margin;
   - liquidation threshold placeholder.

The implementation should come later, chapter by chapter. This document only
defines the boundaries and vocabulary.

---

## 中文

这份笔记在实现之前定义了下一个域表面。它不是完整的交易所设计。它是账本、
钱包和重放之后应该出现的章节的脚手架。

核心思想：

> 撮合生成成交事实。仓位管理解读成交。下单前风控在准入前阻止危险命令。
> 风控集群在事实发出后监视持续敞口。

## 目标形态

```text
gateway
  -> desk pre-trade risk
  -> sequencer
  -> matching engine
  -> execution events
  -> position manager
  -> risk cluster
  -> cold path
```

每个组件有不同的职责。保持这些职责分离比早点实现一个聪明的算法更重要。

## 撮合引擎

撮合引擎拥有确定性订单簿状态。

首个范围：

- 接受或拒绝限价单；
- 维护买/卖盘；
- 暴露最优买卖价；
- 按价格-时间优先级撮合；
- 发出成交执行事件；
- 按订单 ID 撤销。

在热路径上它不应调用数据库、定价服务或远程风控系统。它消费排序命令并发出事实。

对于第一个实现，保持单线程。单线程撮合在这个阶段不是弱点；它使正确性、
重放和解释变得简单。

## 仓位管理器

仓位管理器拥有从成交推导出的账户和合约敞口。

首个范围：

- 按账户和合约的净仓位；
- 平均入场价；
- 已实现盈亏占位符；
- 未实现盈亏占位符；
- 现金和冻结现金钩子；
- 成交报告应用。

对于交易领域，仓位管理往往比撮合算法更重要。交易台可以将订单路由到外部
交易所，但它必须始终知道自己的仓位、敞口、可用余额和风险。

## 柜台下单前风控

柜台下单前风控是同步的，接近订单录入。它回答：

- 账户允许交易吗？
- 合约启用了吗？
- 订单名义值在限额内吗？
- 价格在区间内吗？
- 有足够的可用余额或保证金吗？
- 账户或策略被kill了吗？

这个门保护命令准入。它应该是确定性的和快速的。

这与风控集群不同。下单前风控在订单进入排序核心之前阻止坏订单。

## 风控集群

风控集群是持续的和事件驱动的。它消费成交、仓位、标记价格、入金、出金、
资金费用和人工调整。

首个范围：

- 消费成交事件；
- 消费标记价格事件；
- 按账户和合约聚合敞口；
- 发出风控警报；
- 从事件历史重建状态。

它可能是温路径而不是最热路径。它应该是可重放的和可观测的。一旦单节点
模型清晰，它就可以按账户、合约或策略分片。

## 热、温、冷路径

热路径：

```text
sequenced command -> pre-trade check -> match/apply -> emit facts
```

温路径：

```text
events + marks -> position/risk projections -> alerts/actions
```

冷路径：

```text
events -> reporting/audit/reconciliation/compliance export
```

冷路径不是不重要。它只是不允许阻塞订单录入。它消费热路径生成的事实，
并使其对人类、监管机构、支持、财务和事件审查员可解释。

## 最小保证金模型

第一个保证金实现应使用缩放整数和简单公式。不要从完整的投资组合保证金
或交易所特定清算规则开始。

现货：

```text
available_cash = cash_balance - frozen_cash
available_base = base_balance - frozen_base

buy_required_quote = price * quantity + fee_buffer
sell_required_base = quantity
```

线性合约占位符：

```text
notional = abs(position_qty) * mark_price

if position_qty > 0:
  unrealized_pnl = position_qty * (mark_price - entry_price)
else:
  unrealized_pnl = abs(position_qty) * (entry_price - mark_price)

initial_margin = notional * initial_margin_rate
maintenance_margin = notional * maintenance_margin_rate
equity = wallet_balance + unrealized_pnl
available = equity - initial_margin - frozen_margin - fee_buffer
```

清算阈值占位符：

```text
equity <= maintenance_margin + liquidation_fee_buffer
```

这些公式不是最终的。它们创建了第一个可测试模型，用于解释：

- 为什么标记价格重要；
- 为什么仓位状态与订单状态不同；
- 为什么风控在订单准入后必须持续运行；
- 为什么保证金和清算应该是显式的域章节。

## 章节计划

1. `10-position-manager`
   - 应用成交报告；
   - 更新仓位；
   - 暴露可重放状态。

2. `11-simple-matching-engine-java`
   - 确定性价格-时间撮合；
   - 最优买卖价；
   - 成交事件。

3. `12-desk-pretrade-risk`
   - 简单同步订单检查；
   - 账户/合约/价格/名义值/余额门。

4. `13-risk-cluster-projection`
   - 消费事件和标记价格；
   - 聚合敞口；
   - 发出警报；
   - 从历史重建。

5. `14-margin-model`
   - 现货可用余额；
   - 线性名义值；
   - 初始和维持保证金；
   - 清算阈值占位符。

实现应该稍后逐章进行。这份文档只定义边界和词汇。