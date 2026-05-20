# 01 Custody And User Ledger

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines examples, interfaces, fixtures,
and TODO tests only. It does not implement posting logic.

Purpose: teach the first accounting sentence of the project: user balances are
claims against the platform, not platform revenue.

## User Action

A user deposits 100 USD into the exchange. From the user's point of view, their
available balance increases. From the platform's point of view, two things
happen at the same time:

- the platform custody asset increases;
- the platform liability to that user increases.

## Account Map

| Account | Owner | Purpose | Normal side | Meaning |
| --- | --- | --- | --- | --- |
| `Platform_Custody_USD` | platform | custody | debit | External bank, chain, or custody asset controlled by the platform. |
| `User_A_USD_Available` | user A | available | credit | Amount the platform owes user A and can currently be spent or reserved. |

## Journal Template

```text
Deposit 100 USD for user A

Debit   Platform_Custody_USD      100
Credit  User_A_USD_Available      100
```

The entry is balanced in USD. It creates no revenue, profit, or platform equity
movement.

## Contract Checks

- Deposit posts a debit to custody and a credit to user available.
- Debits and credits balance per asset.
- User available balance can be explained from journal entries.
- Fee revenue is not touched by a plain deposit.

## Out Of Scope

- Withdrawal state transitions are introduced in chapter 02.
- Order reservation and locked balances are introduced in chapter 03.
- Matching, fees, and multi-asset settlement are introduced in chapter 04.
- SQL transaction boundaries are introduced in chapter 05.

## 中文

状态：契约脚手架。本章只定义示例、接口、fixture 和 TODO 测试，不实现 posting
逻辑。

目的：讲清本项目第一句会计语义：用户余额是用户对平台的索取权，不是平台收入。

## 用户动作

用户向交易所充值 100 USD。从用户视角看，可用余额增加了。从平台视角看，两件事
同时发生：

- 平台 custody asset 增加；
- 平台欠这个用户的 liability 增加。

## 账户图谱

| 账户 | 归属 | 用途 | 正常方向 | 含义 |
| --- | --- | --- | --- | --- |
| `Platform_Custody_USD` | 平台 | custody | debit | 平台控制的外部银行、链上或托管资产。 |
| `User_A_USD_Available` | 用户 A | available | credit | 平台欠用户 A、且当前可花费或可冻结的金额。 |

## 分录模板

```text
Deposit 100 USD for user A

Debit   Platform_Custody_USD      100
Credit  User_A_USD_Available      100
```

这笔分录在 USD 内平衡。它不产生收入、利润或平台权益变化。

## 契约检查

- 入金对 custody 记 debit，对 user available 记 credit。
- debit 和 credit 按资产分别平衡。
- 用户 available balance 可以由 journal entries 解释出来。
- 普通入金不触碰 fee revenue。

## 本章不做什么

- 出金状态迁移在第 02 章引入。
- 订单冻结和 locked balances 在第 03 章引入。
- 撮合、手续费和多资产结算在第 04 章引入。
- SQL 事务边界在第 05 章引入。
