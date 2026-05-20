# 02 Balance States

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines examples, interfaces, fixtures,
and TODO tests only. It does not implement balance-state transitions.

Purpose: show that a user balance is not one mutable number. It is a set of
ledger-explainable account states.

## User Action

A user can have funds that are available, reserved, waiting for withdrawal, or
already captured as platform fees. The visible balance screen is a projection
of those journaled states, not the source of truth by itself.

The teaching actions are:

- move available funds into locked funds;
- request a withdrawal by moving available funds into pending withdrawal;
- confirm a withdrawal by reducing pending withdrawal and platform custody;
- record fee revenue as a platform account movement.

## Account Map

| Account | Owner | Purpose | Normal side | Meaning |
| --- | --- | --- | --- | --- |
| `User_A_USD_Available` | user A | available | credit | Spendable or reservable user claim. |
| `User_A_USD_Locked` | user A | locked | credit | User claim reserved for an order or other hold. |
| `User_A_USD_PendingWithdrawal` | user A | pending withdrawal | credit | User claim approved for external payout and no longer spendable. |
| `Platform_FeeRevenue_USD` | platform | fee revenue | credit | Platform revenue account. |
| `Platform_Custody_USD` | platform | custody | debit | External custody asset. |

## Journal Template

```text
Reserve 60 USD from available to locked

Debit   User_A_USD_Available      60
Credit  User_A_USD_Locked         60
```

```text
Request withdrawal of 20 USD

Debit   User_A_USD_Available          20
Credit  User_A_USD_PendingWithdrawal  20
```

```text
Confirm withdrawal of 20 USD

Debit   User_A_USD_PendingWithdrawal  20
Credit  Platform_Custody_USD          20
```

```text
Capture 1 USD fee from user available

Debit   User_A_USD_Available      1
Credit  Platform_FeeRevenue_USD   1
```

## Contract Checks

- Available, locked, and pending withdrawal are separate account purposes.
- State movement is represented by entries, not invisible balance edits.
- Withdrawal confirmation reduces both pending withdrawal liability and
  custody asset.
- Fee revenue is a platform account, not a user claim.
- Balances are derived from entries or projections of entries.

## Out Of Scope

- Order acceptance and cancellation rules are introduced in chapter 03.
- Trade execution and fee calculation are introduced in chapter 04.
- SQL storage and constraints are introduced in chapter 05.

## 中文

状态：契约脚手架。本章只定义示例、接口、fixture 和 TODO 测试，不实现余额状态
迁移。

目的：说明用户余额不是一个可随手修改的数字，而是一组可以由账本解释的账户状态。

## 用户动作

用户资金可能处在可用、已冻结、提现中，或已经被平台确认为手续费收入。页面上的
余额是这些 journaled states 的 projection，本身不是唯一真相源。

本章教学动作是：

- 把 available funds 移到 locked funds；
- 通过把 available funds 移到 pending withdrawal 来发起提现；
- 通过减少 pending withdrawal 和 platform custody 来确认提现；
- 把手续费收入记录为平台账户移动。

## 账户图谱

| 账户 | 归属 | 用途 | 正常方向 | 含义 |
| --- | --- | --- | --- | --- |
| `User_A_USD_Available` | 用户 A | available | credit | 可花费或可冻结的用户索取权。 |
| `User_A_USD_Locked` | 用户 A | locked | credit | 为订单或其他 hold 保留的用户索取权。 |
| `User_A_USD_PendingWithdrawal` | 用户 A | pending withdrawal | credit | 已批准外部支付、不可再花费的用户索取权。 |
| `Platform_FeeRevenue_USD` | 平台 | fee revenue | credit | 平台收入账户。 |
| `Platform_Custody_USD` | 平台 | custody | debit | 外部托管资产。 |

## 分录模板

```text
Reserve 60 USD from available to locked

Debit   User_A_USD_Available      60
Credit  User_A_USD_Locked         60
```

```text
Request withdrawal of 20 USD

Debit   User_A_USD_Available          20
Credit  User_A_USD_PendingWithdrawal  20
```

```text
Confirm withdrawal of 20 USD

Debit   User_A_USD_PendingWithdrawal  20
Credit  Platform_Custody_USD          20
```

```text
Capture 1 USD fee from user available

Debit   User_A_USD_Available      1
Credit  Platform_FeeRevenue_USD   1
```

## 契约检查

- available、locked、pending withdrawal 是不同账户用途。
- 状态移动由 entries 表达，而不是隐藏改余额。
- 提现确认同时减少 pending withdrawal liability 和 custody asset。
- fee revenue 是平台账户，不是用户索取权。
- 余额由 entries 或 entries 的 projection 推导。

## 本章不做什么

- 订单接受和撤单规则在第 03 章引入。
- 成交执行和手续费计算在第 04 章引入。
- SQL 存储和约束在第 05 章引入。
