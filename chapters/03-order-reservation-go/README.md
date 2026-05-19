# 03 Order Reservation

[English](#english) · [中文](#中文)

## English

> You place a buy order. Before matching, the exchange locks the corresponding funds so you can't spend them elsewhere. This chapter explains that lock: where it comes from, and where it goes when the order is cancelled.

Purpose: explain order intent before matching.

A placed order is not a trade yet. It is an intent that must reserve enough
funds or inventory before it can enter the book. Canceling the order releases
the unused locked amount.

Example buy-order reservation:

```text
Debit   User_A_USD_Available      60,100
Credit  User_A_USD_Locked         60,100
```

## Status

Status: contract scaffold. No business implementation exists here yet.

Agents may refine examples, interfaces, fixtures, and TODO tests. They must not
implement order admission, reservation, or cancellation logic.

## Contract Focus

- order placement is intent, not execution;
- accepted orders reserve assets before entering the book;
- rejected orders emit reasons and reserve nothing;
- cancellation releases remaining locked funds;
- reservation facts must be replayable.

## 中文

> 你下了一个买单，在撮合之前，系统把对应的资金"锁起来"，不让你再花。这章解释这把锁：从哪来，撤单后去哪。

目的：在撮合之前，解释订单意图。

下单还不是成交。它是一个意图，必须先冻结足够资金或库存，才可以进入订单簿。
撤单会释放未使用的 locked amount。

买单冻结示例：

```text
Debit   User_A_USD_Available      60,100
Credit  User_A_USD_Locked         60,100
```

## 状态

状态：契约脚手架。本章尚无业务实现。

Agents 可以细化示例、接口、fixture 和 TODO 测试，但不能实现订单准入、冻结或撤单
逻辑。

## 契约重点

- 下单是 intent，不是 execution；
- accepted orders 入簿前先冻结资产；
- rejected orders 发出原因且不冻结；
- 撤单释放剩余 locked funds；
- reservation facts 必须可重放。
