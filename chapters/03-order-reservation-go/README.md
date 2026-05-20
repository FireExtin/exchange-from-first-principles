# 03 Order Reservation

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines examples, interfaces, fixtures,
and TODO tests only. It does not implement order admission, reservation, or
cancellation logic.

Purpose: separate order intent from execution. A valid order must reserve the
asset it can spend before it enters the book.

## User Action

User A places a buy order for BTC. Before matching, the exchange reserves the
quote asset needed to pay for the order. User B places a sell order for BTC.
Before matching, the exchange reserves the base asset needed to deliver it.

If an order is cancelled before it is fully filled, the remaining locked amount
is released back to available.

## Account Map

| Account | Owner | Purpose | Normal side | Meaning |
| --- | --- | --- | --- | --- |
| `User_A_USD_Available` | buyer | available | credit | Quote asset that can fund a buy order. |
| `User_A_USD_Locked` | buyer | locked | credit | Quote asset reserved for an accepted buy order. |
| `User_B_BTC_Available` | seller | available | credit | Base asset that can fund a sell order. |
| `User_B_BTC_Locked` | seller | locked | credit | Base asset reserved for an accepted sell order. |

## Journal Template

```text
Accept buy order and reserve quote asset

Debit   User_A_USD_Available      60,100
Credit  User_A_USD_Locked         60,100
```

```text
Accept sell order and reserve base asset

Debit   User_B_BTC_Available      1
Credit  User_B_BTC_Locked         1
```

```text
Cancel remaining buy order and release quote asset

Debit   User_A_USD_Locked         60,100
Credit  User_A_USD_Available      60,100
```

## Contract Checks

- Accepted buy orders reserve quote asset before entering the book.
- Accepted sell orders reserve base asset before entering the book.
- Rejected orders emit a reason and reserve nothing.
- Cancellation releases only the remaining locked amount.
- Reservation and release facts are replayable.

## Out Of Scope

- Matching and settlement are introduced in chapter 04.
- Position and PnL effects are introduced later.
- Risk and margin admission are later hot-path deep dives.

## 中文

状态：契约脚手架。本章只定义示例、接口、fixture 和 TODO 测试，不实现订单准入、
冻结或撤单逻辑。

目的：把订单意图和成交执行分开。有效订单入簿前，必须先冻结它将要花费或交付的
资产。

## 用户动作

用户 A 下 BTC 买单。撮合之前，交易所先冻结买单需要支付的 quote asset。用户 B
下 BTC 卖单。撮合之前，交易所先冻结卖单需要交付的 base asset。

如果订单在完全成交前被撤销，剩余 locked amount 会释放回 available。

## 账户图谱

| 账户 | 归属 | 用途 | 正常方向 | 含义 |
| --- | --- | --- | --- | --- |
| `User_A_USD_Available` | 买方 | available | credit | 可为买单付款的 quote asset。 |
| `User_A_USD_Locked` | 买方 | locked | credit | 为已接受买单冻结的 quote asset。 |
| `User_B_BTC_Available` | 卖方 | available | credit | 可为卖单交付的 base asset。 |
| `User_B_BTC_Locked` | 卖方 | locked | credit | 为已接受卖单冻结的 base asset。 |

## 分录模板

```text
Accept buy order and reserve quote asset

Debit   User_A_USD_Available      60,100
Credit  User_A_USD_Locked         60,100
```

```text
Accept sell order and reserve base asset

Debit   User_B_BTC_Available      1
Credit  User_B_BTC_Locked         1
```

```text
Cancel remaining buy order and release quote asset

Debit   User_A_USD_Locked         60,100
Credit  User_A_USD_Available      60,100
```

## 契约检查

- 已接受买单入簿前冻结 quote asset。
- 已接受卖单入簿前冻结 base asset。
- 被拒绝订单发出原因且不冻结资产。
- 撤单只释放剩余 locked amount。
- 冻结和释放 facts 必须可重放。

## 本章不做什么

- 撮合和结算在第 04 章引入。
- 仓位和 PnL 影响在后续章节引入。
- 风控和保证金准入属于后续热路径深挖。
