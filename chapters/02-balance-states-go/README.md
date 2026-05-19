# 02 Balance States

[English](#english) · [中文](#中文)

## English

Purpose: explain user balance states before orders and matching.

An exchange does not only track one balance number. It tracks user claims by
state: available, locked, and pending withdrawal. It also names platform
purposes such as fee revenue. State changes are ledger movements, not invisible
balance edits.

Example withdrawal request:

```text
Debit   User_A_USD_Available          20
Credit  User_A_USD_PendingWithdrawal  20
```

## Status

Status: contract scaffold. No business implementation exists here yet.

Agents may refine examples, interfaces, fixtures, and TODO tests. They must not
implement balance-state transitions.

## Contract Focus

- available funds can be spent or reserved;
- locked funds back open orders;
- pending withdrawal funds cannot be reused;
- fee revenue is a platform account purpose, not a user claim;
- every state movement is explainable as entries.

## 中文

目的：在订单和撮合之前，解释用户余额状态。

交易所不只跟踪一个 balance 数字。它按状态跟踪用户索取权：available、locked 和
pending withdrawal。它也会命名平台用途，例如 fee revenue。状态变化是账本移动，
不是隐藏改余额。

出金请求示例：

```text
Debit   User_A_USD_Available          20
Credit  User_A_USD_PendingWithdrawal  20
```

## 状态

状态：契约脚手架。本章尚无业务实现。

Agents 可以细化示例、接口、fixture 和 TODO 测试，但不能实现余额状态迁移。

## 契约重点

- available funds 可以花费或冻结；
- locked funds 支撑 open orders；
- pending withdrawal funds 不能再次使用；
- fee revenue 是平台账户用途，不是用户索取权；
- 每个状态移动都必须能解释成 entries。
