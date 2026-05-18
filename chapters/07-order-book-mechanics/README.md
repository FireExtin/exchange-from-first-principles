# 06 Simple Matching Engine Java

[English](#english) · [中文](#中文)

## English

Purpose: implement the smallest deterministic order book that can explain
price-time priority and trade events.

## Status

Status: design scaffold. No runnable implementation exists here yet.

First implementation should cover:

- limit order accept/reject;
- bid/ask book;
- top-of-book;
- simple match;
- cancel by order id.

Keep the engine single-threaded until the behavior is easy to replay.

---

## 中文

目的：实现最小的确定性订单簿，能解释价格-时间优先级和成交事件。

## 状态

状态：设计脚手架。本目录尚无可运行实现。

第一个实现应覆盖：

- 限价单接受/拒绝；
- 买卖盘；
- 最优买卖价；
- 简单撮合；
- 按订单 ID 撤销。

在行为易于重放之前保持引擎单线程。
