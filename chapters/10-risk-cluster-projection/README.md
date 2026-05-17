# 10 Risk Cluster Projection

[English](#english) · [中文](#中文)

## English

Purpose: consume events and maintain continuous exposure views.

## Status

Status: design scaffold. No runnable implementation exists here yet.

First implementation should cover:

- consume execution events;
- consume mark-price events;
- aggregate account and instrument exposure;
- emit risk alerts;
- rebuild state from event history.

This is warm-path state. It should be replayable and observable.

---

## 中文

目的：消费事件并维护持续敞口视图。

## 状态

状态：设计脚手架。本目录尚无可运行实现。

第一个实现应覆盖：

- 消费成交事件；
- 消费标记价格事件；
- 聚合账户和合约敞口；
- 发出风控警报；
- 从事件历史重建状态。

这是温路径状态。它应该是可重放的和可观测的。
