# Credit, Margin, And Funding Extension

[English](#english) · [中文](#中文)

## English

This note defines an optional exchange-core extension: credit, margin, funding,
collateral, and liquidation settlement. It is not a Modern Lending product
track. It exists only where lending-style concepts strengthen exchange
semantics.

The extension should appear after the spot exchange can already explain
custody, user liabilities, reservations, matching, positions, risk views, and
posted-versus-derived state.

## Why It Exists

Spot exchange semantics answer:

```text
what does the user own, reserve, trade, and settle?
```

Credit and margin semantics add:

```text
what can the user borrow, collateralize, accrue, repay, or liquidate?
```

The useful borrowing from lending systems is not a full loan marketplace. The
useful borrowing is the boundary between posted ledger facts, future schedules,
accruals, collateral, risk views, and settlement events.

## Semantic Surface

| Surface | Meaning | Ledger Boundary |
| --- | --- | --- |
| Collateral | Assets pledged to support credit or margin exposure | Ledger state when pledged, released, or seized. |
| Borrow liability | Amount owed by the user to the platform or lending pool | Ledger state when borrow, repay, or write-off posts. |
| Funding / interest | Cost of carrying a position or borrow | Derived until an explicit accrual posts entries. |
| Margin requirement | Risk model output used for admission and liquidation checks | Derived/prospective; not ledger truth by itself. |
| Liquidation | Forced settlement that reduces risk after margin failure | Ledger event when execution, fee, or seizure posts. |

## Journal Boundaries

Examples of events that may post entries:

- collateral pledge or release;
- borrow drawdown;
- repayment;
- interest or funding accrual;
- liquidation execution;
- liquidation fee;
- bad-debt or insurance-fund transfer.

Examples of values that should not post entries by themselves:

- mark-price changes;
- margin requirement recalculation;
- unrealized PnL movement;
- liquidation warning;
- risk projection alert.

## Chapter 18 Contract Lab

Chapter 18 is the contract lab for this extension:

18. [`18-credit-and-collateral-accounts-go`](../chapters/18-credit-and-collateral-accounts-go/README.md)
    - Defines the shared `credit` contract surface.
    - Adds TODO contract tests behind `credit_contract_todo`.
    - Keeps all credit, funding, liquidation, and posting implementation blank.

## Planned Notes

These remain planned notes, not active chapter directories:

19. `19-funding-and-interest-accrual`
    - Explain accrual boundaries for funding, interest, and carry cost.
    - Show why model schedules are not posted ledger truth until accrual.

20. `20-liquidation-and-repayment-settlement`
    - Explain repayment, collateral release, liquidation execution, liquidation
      fee, and bad-debt handling as explicit settlement events.
    - Keep liquidation risk checks separate from posted settlement.

## Boundary Rule

Do not turn the exchange into a lending marketplace. Borrow only the concepts
needed to explain margin, funding, collateral, and settlement boundaries.

This extension consumes exchange facts and risk views, then emits explicit
credit or settlement facts. It should not silently mutate the spot ledger,
order book, or risk model.

---

## 中文

本文定义一个可选的交易所核心扩展：credit、margin、funding、collateral 和
liquidation settlement。它不是 Modern Lending 产品线。它只在 lending-style 概念能
强化交易所语义的地方出现。

这个扩展应该在现货交易所已经能解释 custody、user liabilities、reservations、
matching、positions、risk views，以及 posted-versus-derived state 之后再出现。

## 为什么存在

现货交易所语义回答：

```text
用户拥有什么、冻结什么、交易什么、结算什么？
```

credit 和 margin 语义继续追问：

```text
用户能借什么、用什么抵押、如何计提、如何偿还、如何强平？
```

从 lending 系统里真正有用的，不是完整 loan marketplace，而是 posted ledger facts、
future schedules、accruals、collateral、risk views 和 settlement events 之间的边界。

## 语义表面

| 表面 | 含义 | Ledger 边界 |
| --- | --- | --- |
| Collateral | 为 credit 或 margin exposure 质押的资产 | 质押、释放或被扣划时成为 ledger state。 |
| Borrow liability | 用户欠平台或 lending pool 的金额 | borrow、repay 或 write-off 过账时成为 ledger state。 |
| Funding / interest | 持仓或借款的持有成本 | 在显式 accrual 提交 entries 前是派生值。 |
| Margin requirement | 用于准入和强平检查的风险模型输出 | 派生/未来状态，本身不是 ledger truth。 |
| Liquidation | margin failure 后降低风险的强制结算 | execution、fee 或 seizure 过账时成为 ledger event。 |

## 分录边界

可能提交 entries 的事件：

- collateral pledge 或 release；
- borrow drawdown；
- repayment；
- interest 或 funding accrual；
- liquidation execution；
- liquidation fee；
- bad-debt 或 insurance-fund transfer。

不应自己提交 entries 的值：

- mark-price changes；
- margin requirement recalculation；
- unrealized PnL movement；
- liquidation warning；
- risk projection alert。

## 第 18 章契约实验

第 18 章是这个扩展的契约实验：

18. [`18-credit-and-collateral-accounts-go`](../chapters/18-credit-and-collateral-accounts-go/README.md)
    - 定义共享 `credit` contract surface。
    - 在 `credit_contract_todo` 后增加 TODO 契约测试。
    - 保持所有信用、资金费率、强平和过账实现空白。

## 规划笔记

这些仍然是规划笔记，不是 active chapter directories：

19. `19-funding-and-interest-accrual`
    - 解释 funding、interest 和 carry cost 的 accrual 边界。
    - 说明为什么 model schedules 在 accrual 之前不是 posted ledger truth。

20. `20-liquidation-and-repayment-settlement`
    - 把 repayment、collateral release、liquidation execution、liquidation fee 和
      bad-debt handling 解释成显式 settlement events。
    - 保持 liquidation risk checks 和 posted settlement 分离。

## 边界规则

不要把交易所变成 lending marketplace。只借用能解释 margin、funding、collateral 和
settlement boundaries 的概念。

这个扩展消费 exchange facts 和 risk views，然后发出显式 credit 或 settlement facts。
它不应该偷偷改写 spot ledger、order book 或 risk model。
