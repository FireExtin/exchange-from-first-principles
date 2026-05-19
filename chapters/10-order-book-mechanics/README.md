# 10 Order Book Mechanics

[English](#english) · [中文](#中文)

## English

> When two orders are at the same price, which one fills first? How does an order book decide, and how fast can it do it? This chapter opens the lid on the matching engine.

Purpose: explain order-book internals after the minimal order and settlement
semantics are already introduced.

This is a domain deep dive, not the first place where matching appears.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- price-time priority;
- bid/ask book structure;
- top-of-book;
- partial and full fills;
- cancel by order id;
- deterministic execution facts.

## 中文

> 同价格的订单哪个先成交？订单簿怎么决定，速度有多快？这章打开撮合引擎的盖子。

目的：在最小订单和结算语义已经引入之后，解释订单簿内部机制。

这是领域 deep dive，不是第一次引入撮合语义的地方。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- 价格-时间优先级；
- bid/ask book 结构；
- top-of-book；
- 部分成交和完全成交；
- 按 order id 撤单；
- 确定性 execution facts。
