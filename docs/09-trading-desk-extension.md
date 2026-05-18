# Trading Desk Extension

[English](#english) · [中文](#中文)

## English

This note defines a later extension of the project: a trading desk, market
maker, or proprietary trading system built on top of the exchange primitives.
It should not be implemented early. It becomes useful only after the exchange
core can produce reliable facts, positions, risk views, market-data streams,
and execution reports.

The distinction matters:

- exchange core systems decide what is true inside the venue;
- trading desk systems decide what to do given market state, inventory, risk,
  and external venues.

## Component Split

| Component | Role | Project Placement |
| --- | --- | --- |
| Pre-Trade Risk | Blocks dangerous orders before admission | Core roadmap |
| Positions | Maintains exposure from execution facts | Core roadmap |
| Post-trade | Feeds settlement, reporting, audit, and compliance | Core roadmap |
| Trade Reconciliator | Compares internal facts with external records | Core roadmap |
| Market Data Push | Publishes venue book, trade, and execution streams | Core roadmap |
| External Market Data | Consumes prices, books, trades, and reference data from venues | Desk extension |
| Pricing Engine | Computes theoretical price, mark price, and quote support | Core support, then desk extension |
| Algo Engine | Converts market state and strategy intent into orders | Desk extension |
| Arbitrage Engine | Detects cross-venue or cross-product opportunities | Desk extension |
| Order Router | Sends child orders to external venues and tracks execution reports | Desk extension |
| Hedger | Reduces inventory or basis exposure after fills | Desk extension |
| Best Execution | Chooses venue, order style, and routing policy | Desk extension |

## Natural Emergence

The desk layer should appear only when the earlier system creates pressure for
it:

1. Matching creates execution facts.
2. Execution facts require positions.
3. Positions create exposure.
4. Exposure requires pre-trade checks and continuous risk.
5. Risk requires market state, marks, and pricing inputs.
6. Public market data and private execution streams leave the core.
7. External venues become relevant only after the system needs to trade outside
   its own book.

At that point, the project can add a desk layer:

```text
external market data
  -> pricing / signals
  -> algo decision
  -> pre-trade risk
  -> order router
  -> external execution reports
  -> positions / hedger / reconciliation
```

## Future Chapters

15. `15-external-market-data-ingestion`
    - Consume venue books, trades, tickers, and reference data.
    - Explain sequence gaps, snapshots, stale data, and mark inputs.

16. `16-pricing-and-signal-engine`
    - Compute theoretical price, fair value, mark price, and simple signals.
    - Explain why pricing is separate from matching and routing.

17. `17-order-router-and-execution-reports`
    - Route child orders to an external venue mock.
    - Consume execution reports and update order state.

18. `18-hedger-and-best-execution`
    - Use positions, risk limits, venue state, and costs to choose a hedge or
      execution venue.
    - Explain the tradeoff between immediacy, fees, liquidity, and risk.

19. `19-arbitrage-strategy-demo`
    - Use multiple venue feeds and router mocks to demonstrate cross-venue
      arbitrage.
    - Keep it educational, not a production trading strategy.

## Boundary Rule

Do not let desk components pollute the exchange core.

The matching engine should not know about arbitrage. The ledger should not know
about best execution. Pre-trade risk should not become a strategy engine. The
desk layer consumes facts and produces orders; the exchange core accepts,
rejects, orders, executes, and emits facts.

---

## 中文

这份笔记定义项目的远期扩展：建立在交易所原语之上的交易台、做市或自营交易
系统。它不应该过早实现。只有当交易所核心已经能够产生可靠事实、仓位、风险
视图、行情流和成交回报之后，这一层才自然出现。

这里要区分两类系统：

- 交易所核心系统决定场内什么是真的；
- 交易台系统根据市场状态、库存、风险和外部场所决定要做什么。

## 组件拆分

| 组件 | 作用 | 项目位置 |
| --- | --- | --- |
| Pre-Trade Risk | 在订单准入前阻止危险订单 | 核心路线 |
| Positions | 从成交事实维护敞口 | 核心路线 |
| Post-trade | 进入结算、报告、审计和合规 | 核心路线 |
| Trade Reconciliator | 对比内部事实与外部记录 | 核心路线 |
| Market Data Push | 发布场内订单簿、成交和执行流 | 核心路线 |
| External Market Data | 从外部场所消费价格、订单簿、成交和参考数据 | 交易台扩展 |
| Pricing Engine | 计算理论价、标记价格和报价支撑 | 核心支撑，然后进入交易台扩展 |
| Algo Engine | 将市场状态和策略意图转换为订单 | 交易台扩展 |
| Arbitrage Engine | 发现跨场所或跨产品机会 | 交易台扩展 |
| Order Router | 向外部场所发送子订单并跟踪成交回报 | 交易台扩展 |
| Hedger | 在成交后降低库存或基差敞口 | 交易台扩展 |
| Best Execution | 选择场所、订单方式和路由策略 | 交易台扩展 |

## 自然出现的时机

交易台层应该在前面系统产生压力之后再出现：

1. 撮合生成成交事实。
2. 成交事实要求系统维护仓位。
3. 仓位产生敞口。
4. 敞口要求下单前检查和持续风控。
5. 风控要求市场状态、标记价格和定价输入。
6. 公共行情和私有成交流离开核心。
7. 只有当系统需要到外部场所交易时，外部交易所才成为核心问题。

这时项目可以增加交易台层：

```text
外部行情
  -> 定价 / 信号
  -> 策略决策
  -> 下单前风控
  -> 订单路由
  -> 外部成交回报
  -> 仓位 / 对冲 / 对账
```

## 远期章节

15. `15-external-market-data-ingestion`
    - 消费外部场所订单簿、成交、ticker 和参考数据。
    - 解释序列缺口、快照、陈旧数据和标记价格输入。

16. `16-pricing-and-signal-engine`
    - 计算理论价、公允价、标记价格和简单信号。
    - 解释为什么定价应与撮合、路由分离。

17. `17-order-router-and-execution-reports`
    - 将子订单路由到外部场所 mock。
    - 消费成交回报并更新订单状态。

18. `18-hedger-and-best-execution`
    - 用仓位、风险限制、场所状态和成本选择对冲或执行场所。
    - 解释即时性、费用、流动性和风险之间的权衡。

19. `19-arbitrage-strategy-demo`
    - 使用多个场所行情和路由 mock 演示跨场所套利。
    - 保持教育用途，不把它写成生产交易策略。

## 边界规则

不要让交易台组件污染交易所核心。

撮合引擎不应该知道套利。账本不应该知道最优执行。下单前风控不应该变成策略
引擎。交易台层消费事实并产生订单；交易所核心负责接受、拒绝、排序、执行并
发出事实。
