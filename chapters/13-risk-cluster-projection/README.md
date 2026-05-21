# 13 Risk Cluster Projection

[English](#english) · [中文](#中文)

## English

> Which accounts are close to liquidation right now? How exposed is the platform if prices move suddenly? This chapter computes continuous risk from the event stream — asynchronously, not at the moment of order placement.

Purpose: model warm-path risk projections that consume facts after hot-path
admission.

This is not the synchronous gate from chapter 12. It is a consumer that rebuilds
continuous exposure from events, marks, deposits, withdrawals, executions, and
manual adjustments.
Its output is derived risk state. It can alert, explain exposure, and feed
operator workflows, but it cannot become posted ledger truth without an
explicit financial event.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- consume execution events;
- consume mark-price events;
- aggregate exposure by account and instrument;
- emit risk alerts;
- rebuild state from event history with a cursor.

## Boundary

- inputs may include posted facts and model inputs;
- outputs must preserve provenance;
- rebuild must recreate the same risk view without changing ledger entries.

## 中文

> 哪些账户正在接近爆仓？价格突然移动，平台的敞口有多大？这章从事实流持续计算风险——是异步的，不是在下单时计算。

目的：建模热路径准入之后消费事实的 warm-path 风控 projection。

这不是第 12 章的同步 gate。它是一个 consumer，从事件、mark、入金、出金、
成交和人工调整中重建连续 exposure。
它的输出是派生风险状态。它可以报警、解释敞口并驱动运营流程，但如果没有显式财务
事件，不能变成已过账 ledger truth。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- 消费 execution events；
- 消费 mark-price events；
- 按账户和 instrument 聚合 exposure；
- 发出 risk alerts；
- 使用 cursor 从事件历史重建状态。

## 边界

- 输入可以包含已过账事实和模型输入；
- 输出必须保留 provenance；
- 重建必须恢复同一风险视图，同时不改变 ledger entries。
