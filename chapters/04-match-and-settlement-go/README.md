# 04 Match And Settlement

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines examples, interfaces, fixtures,
and TODO tests only. It does not implement matching, settlement, fee posting, or
release logic.

Purpose: translate a minimal matched trade into execution facts and
double-entry postings, one asset at a time.

## User Action

User A buys 1 BTC from user B at 60,000 USD. The buyer pays a 10 USD fee. The
buyer had reserved 60,100 USD, so 90 USD is no longer needed after the fill and
must be released.

The match creates execution facts. Settlement explains those facts as journal
entries.

## Account Map

| Account | Owner | Purpose | Normal side | Meaning |
| --- | --- | --- | --- | --- |
| `User_A_USD_Locked` | buyer | locked | credit | Reserved quote asset consumed by settlement. |
| `User_A_USD_Available` | buyer | available | credit | Surplus quote asset released after partial or cheaper fill. |
| `User_B_USD_Available` | seller | available | credit | Quote asset received by the seller. |
| `Platform_FeeRevenue_USD` | platform | fee revenue | credit | Fee captured by the platform. |
| `User_B_BTC_Locked` | seller | locked | credit | Reserved base asset consumed by settlement. |
| `User_A_BTC_Available` | buyer | available | credit | Base asset received by the buyer. |

## Journal Template

```text
USD settlement for 1 BTC at 60,000 USD, 10 USD buyer fee, 90 USD surplus release

Debit   User_A_USD_Locked         60,100
Credit  User_B_USD_Available      60,000
Credit  Platform_FeeRevenue_USD       10
Credit  User_A_USD_Available          90
```

```text
BTC settlement for 1 BTC

Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

USD balances against USD. BTC balances against BTC. The two assets never
balance against each other.

## Contract Checks

- A match emits execution facts before downstream projections consume them.
- Settlement posts balanced entries per asset.
- Buyer fee posts to platform fee revenue.
- Surplus locked funds are released back to available.
- The same execution facts can later update positions and projections.

## Out Of Scope

- Order-book data structures are introduced in chapter 10.
- Position, PnL, margin, and risk are introduced in later deep dives.
- SQL transaction boundaries are introduced in chapter 05.

## 中文

状态：契约脚手架。本章只定义示例、接口、fixture 和 TODO 测试，不实现撮合、结算、
手续费 posting 或释放逻辑。

目的：把一笔最小成交翻译成 execution facts 和 double-entry postings，并且每种资产
各自解释。

## 用户动作

用户 A 以 60,000 USD 从用户 B 买入 1 BTC。买方支付 10 USD 手续费。买方原先冻结
了 60,100 USD，所以成交后有 90 USD 不再需要，必须释放。

撮合产生成交事实。结算用 journal entries 解释这些事实。

## 账户图谱

| 账户 | 归属 | 用途 | 正常方向 | 含义 |
| --- | --- | --- | --- | --- |
| `User_A_USD_Locked` | 买方 | locked | credit | 结算消耗的已冻结 quote asset。 |
| `User_A_USD_Available` | 买方 | available | credit | 部分成交或低价成交后释放的多余 quote asset。 |
| `User_B_USD_Available` | 卖方 | available | credit | 卖方收到的 quote asset。 |
| `Platform_FeeRevenue_USD` | 平台 | fee revenue | credit | 平台收取的手续费。 |
| `User_B_BTC_Locked` | 卖方 | locked | credit | 结算消耗的已冻结 base asset。 |
| `User_A_BTC_Available` | 买方 | available | credit | 买方收到的 base asset。 |

## 分录模板

```text
USD settlement for 1 BTC at 60,000 USD, 10 USD buyer fee, 90 USD surplus release

Debit   User_A_USD_Locked         60,100
Credit  User_B_USD_Available      60,000
Credit  Platform_FeeRevenue_USD       10
Credit  User_A_USD_Available          90
```

```text
BTC settlement for 1 BTC

Debit   User_B_BTC_Locked              1
Credit  User_A_BTC_Available           1
```

USD 只和 USD 平衡。BTC 只和 BTC 平衡。两种资产不能互相借贷平衡。

## 契约检查

- 撮合先发出 execution facts，再由下游 projection 消费。
- 结算按资产分别提交平衡 entries。
- 买方手续费进入平台 fee revenue。
- 多余 locked funds 释放回 available。
- 同一组 execution facts 后续可以更新 positions 和 projections。

## 本章不做什么

- 订单簿数据结构在第 10 章引入。
- 仓位、PnL、保证金和风控在后续深挖章节引入。
- SQL 事务边界在第 05 章引入。
