# 04 Match And Settlement

[English](#english) · [中文](#中文)

## English

> Two orders match. Now what? Something has to actually move money and assets between accounts. This chapter translates a matched trade into double-entry postings, one set per asset.

Purpose: explain the smallest trade before a full exchange architecture.

Matching creates execution facts. Settlement explains those facts as
double-entry postings per asset. USD and BTC never balance against each other.

Example: user A buys 1 BTC from user B at 60,000 USD and pays a 10 USD fee.

```text
USD ledger:
Debit   User_A_USD_Locked         60,010
Credit  User_B_USD_Available      60,000
Credit  Fee_Revenue_USD               10

BTC ledger:
Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

## Status

Status: contract scaffold. No business implementation exists here yet.

Agents may refine examples, interfaces, fixtures, and TODO tests. They must not
implement matching, settlement, fee posting, or release logic.

## Contract Focus

- matching emits execution facts;
- settlement balances per asset;
- buyer/seller transfers and fee postings are explicit;
- partial fills release unused locked funds;
- execution facts later update positions.

## 中文

> 两个订单撮合成功了。然后呢？钱和资产需要在账户之间实际移动。这章把一笔成交翻译成账本条目，每种资产各自平衡。

目的：在完整交易所架构之前，解释最小成交。

撮合产生成交事实。结算把这些事实解释为按资产分别平衡的 double-entry postings。
USD 和 BTC 不能互相平衡。

示例：用户 A 以 60,000 USD 从用户 B 买 1 BTC，并支付 10 USD 手续费。

```text
USD ledger:
Debit   User_A_USD_Locked         60,010
Credit  User_B_USD_Available      60,000
Credit  Fee_Revenue_USD               10

BTC ledger:
Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

## 状态

状态：契约脚手架。本章尚无业务实现。

Agents 可以细化示例、接口、fixture 和 TODO 测试，但不能实现撮合、结算、手续费
posting 或释放逻辑。

## 契约重点

- 撮合发出 execution facts；
- 结算按资产分别平衡；
- 买卖双方转移和手续费 posting 都显式记录；
- 部分成交释放未使用的 locked funds；
- execution facts 后续更新仓位。
