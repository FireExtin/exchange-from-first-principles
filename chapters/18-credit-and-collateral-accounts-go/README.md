# 18 Credit And Collateral Accounts

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines examples, interfaces, fixtures,
and TODO tests only. It does not implement credit, margin, funding, liquidation,
or ledger posting logic.

Purpose: make the optional credit/margin/funding extension concrete without
turning the project into a Modern Lending product.

## User Action

A user can pledge collateral, borrow against it, accrue funding or interest,
repay the borrow, or be liquidated after a margin failure.

The important boundary is:

```text
collateral / borrow / funding schedule / liquidation risk
  != posted ledger truth
```

Only explicit accrual, repayment, liquidation, or settlement events create
journal entries.

## Account Map

| Account | Owner | Purpose | Meaning |
| --- | --- | --- | --- |
| `User_A_USD_Available` | user A | available | Spendable user claim. |
| `User_A_USD_Collateral` | user A | collateral | User claim pledged to support credit or margin exposure. |
| `User_A_USD_BorrowLiability` | user A | borrow liability | Amount the user owes after a borrow drawdown. |
| `Platform_FundingRevenue_USD` | platform | funding revenue | Funding or interest revenue once accrual posts. |
| `Platform_InsuranceFund_USD` | platform | insurance fund | Buffer used for liquidation loss or bad debt workflows. |

## Journal Boundaries

Examples that may post entries:

```text
Pledge collateral:
Debit   User_A_USD_Available
Credit  User_A_USD_Collateral
```

```text
Accrue funding:
Debit   User_A_USD_Available
Credit  Platform_FundingRevenue_USD
```

```text
Repay borrow:
Debit   User_A_USD_BorrowLiability
Credit  User_A_USD_Available
```

The exact account set can evolve, but the boundary must not: marks, margin
requirements, schedules, warnings, and unrealized PnL do not post by
themselves.

## Contract Checks

- Collateral pledge moves available funds into a collateral account.
- Borrow drawdown creates borrow liability and usable funds.
- Mark-price changes do not create ledger entries.
- Funding schedules do not post until an accrual event.
- Repayment reduces borrow liability through explicit entries.
- Liquidation settles collateral, borrow liability, and fees through explicit
  entries.

## How To Run

There is no normal runnable implementation in this chapter yet. The contract
tests live in the shared credit package:

```bash
cd shared/go
go test -tags credit_contract_todo ./credit
```

Those tests are expected to fail at explicit TODO boundaries until the project
owner implements adapters and exercise logic.

## Out Of Scope

- No loan marketplace.
- No amortization engine.
- No SQL schema or adapter.
- No risk model implementation.
- No liquidation engine implementation.

## 中文

状态：契约脚手架。本章只定义示例、接口、fixture 和 TODO 测试，不实现信用、
保证金、资金费率、强平或账本过账逻辑。

目的：把可选信用/保证金/资金费率扩展具体化，但不把项目变成 Modern Lending
产品。

## 用户动作

用户可以质押抵押品、基于抵押借出资产、计提资金费率或利息、偿还借款，或者在
保证金不足后被强平。

关键边界是：

```text
抵押品 / 借款 / 资金费率计划 / 强平风险
  != 已过账账本真相
```

只有显式计提、还款、强平或结算事件会创建 journal entries。

## 账户图谱

| 账户 | 归属 | 用途 | 含义 |
| --- | --- | --- | --- |
| `User_A_USD_Available` | 用户 A | available | 可花费的用户索取权。 |
| `User_A_USD_Collateral` | 用户 A | collateral | 为信用或保证金敞口质押的用户索取权。 |
| `User_A_USD_BorrowLiability` | 用户 A | borrow liability | 借款提款后用户欠下的金额。 |
| `Platform_FundingRevenue_USD` | 平台 | funding revenue | 计提过账后的资金费率或利息收入。 |
| `Platform_InsuranceFund_USD` | 平台 | insurance fund | 用于强平损失或坏账流程的缓冲。 |

## 分录边界

可能提交 entries 的例子：

```text
Pledge collateral:
Debit   User_A_USD_Available
Credit  User_A_USD_Collateral
```

```text
Accrue funding:
Debit   User_A_USD_Available
Credit  Platform_FundingRevenue_USD
```

```text
Repay borrow:
Debit   User_A_USD_BorrowLiability
Credit  User_A_USD_Available
```

具体账户集合可以演化，但边界不能变：标记价格、保证金要求、计划、预警和未实现
PnL 本身都不会过账。

## 契约检查

- 质押抵押品会把 available funds 移入 collateral account。
- 借款提款会创建 borrow liability 和 usable funds。
- 标记价格变化不创建 ledger entries。
- 资金费率计划在 accrual event 前不过账。
- 还款通过显式 entries 减少 borrow liability。
- 强平通过显式 entries 结算 collateral、borrow liability 和 fees。

## 如何运行

本章还没有普通可运行实现。契约测试位于共享 credit 包：

```bash
cd shared/go
go test -tags credit_contract_todo ./credit
```

这些测试在项目所有者实现 adapter 和练习逻辑前，预期会失败在显式 TODO 边界。

## 本章不做什么

- 不做 loan marketplace。
- 不做 amortization engine。
- 不做 SQL schema 或 adapter。
- 不实现 risk model。
- 不实现 liquidation engine。
