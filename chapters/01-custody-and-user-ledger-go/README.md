# 01 Custody And User Ledger

[English](#english) · [中文](#中文)

## English

> When a user deposits $100, is that the exchange's money? No — it's a debt the exchange owes them. This chapter writes that debt into a ledger.

Purpose: explain the first accounting idea before any exchange architecture:
user balances are platform liabilities.

When a user deposits USD, the platform receives an asset and owes the user the
same amount. That is not revenue.

```text
Debit   Platform_Custody_USD      100
Credit  User_A_USD_Available      100
```

## Status

Status: contract scaffold. No business implementation exists here yet.

Agents may refine examples, interfaces, fixtures, and TODO tests. They must not
implement posting logic.

## Contract Focus

- custody account is debit-normal platform asset;
- user available account is credit-normal platform liability;
- deposit creates both sides;
- every journal transaction balances per asset.

## 中文

> 用户存了 $100 进来，这笔钱是平台的收入吗？不是——这是平台欠用户的债。这章把这个"欠"字写进账本。

目的：在任何交易所架构之前，先解释第一条会计语义：用户余额是平台负债。

用户充值 USD 时，平台收到一项资产，同时欠用户同等金额。这不是收入。

```text
Debit   Platform_Custody_USD      100
Credit  User_A_USD_Available      100
```

## 状态

状态：契约脚手架。本章尚无业务实现。

Agents 可以细化示例、接口、fixture 和 TODO 测试，但不能实现 posting 逻辑。

## 契约重点

- custody account 是 debit-normal 平台资产；
- user available account 是 credit-normal 平台负债；
- 入金同时产生两边；
- 每笔 journal transaction 在每种资产内分别平衡。
