# System Boundary

[English](#english) · [中文](#中文)

## English

This document owns language and system-boundary rules. For the migration from
database truth to ordered facts, see
[Truth Source Migration](./04-truth-source-migration.md).

The long-term hot-path source of truth starts at command sequencing:

```text
ingress
  -> assign seq
  -> append command log
  -> apply deterministic state machine
  -> emit events
```

Anything before sequencing is an adapter. Anything after event emission is a
projection.

## Java Owns First

Java is the primary implementation surface for the trading hot path:

- command sequence;
- append-only command log;
- order book;
- account reservation;
- matching;
- event generation;
- snapshot and replay rules.

## Go Owns Service Edges

Go handles service design around assets, ledger, idempotency, reconciliation,
callbacks, and operational glue:

- load generation;
- gateway and admin APIs;
- ledger and account-service simulation;
- materialized view readers;
- reconciliation reports;
- idempotency and callback handling.

## Aeron Owns Integration Ordering

Aeron owns replicated-log integration concerns, not business rules:

- cluster ingress ordering;
- replicated log;
- leader election;
- cluster snapshots;
- flow control and backpressure.

## Rust Is Exploratory

The Rust workspace is a runnable experiment for long-term exploration of a
clean hot-path contract and FFI boundary. It is not the active main
implementation track.

## 中文

本文档负责语言和系统边界规则。关于真相源如何从数据库迁移到有序事实，见
[Truth Source Migration](./04-truth-source-migration.md)。

长期热路径的真相源从命令排序开始：

```text
ingress
  -> assign seq
  -> append command log
  -> apply deterministic state machine
  -> emit events
```

排序之前的任何东西都是适配器。事件发出之后的任何东西都是投影。

## Java 优先

Java 是交易热路径的主要实现层面：

- 命令排序；
- 追加写命令日志；
- 订单簿；
- 账户预留；
- 撮合；
- 事件生成；
- 快照和重放规则。

## Go 负责服务边界

Go 处理资产、账本、幂等、对账、回调和运营胶水代码相关的服务设计：

- 压测生成；
- 网关和管理 API；
- 账本和账户服务模拟；
- 物化视图读取器；
- 对账报告；
- 幂等和回调处理。

## Aeron 负责集成排序

Aeron 负责复制日志集成问题，而不是业务规则：

- 集群入口排序；
- 复制日志；
- leader 选举；
- 集群快照；
- 流量控制和背压。

## Rust 作为探索实验

Rust workspace 是一个可运行实验，用于长期探索干净的热路径契约和 FFI 边界。它
不是当前主实现路线。
