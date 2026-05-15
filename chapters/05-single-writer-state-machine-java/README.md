# 05 Single-Writer State Machine Java

This directory is the target home for the first Java trading-core chapter.

Do not build a broad exchange here. Build the smallest Java surface that proves:

- deterministic command application;
- order book and position state;
- append-only command/event logs;
- replay from offset or position;
- low-allocation hot-path choices;
- profiling with JFR and async-profiler;
- clean adapter boundary to Aeron.

## Proposed Modules

```text
exchange-types/       commands, events, ids, price/quantity value rules
exchange-core/        deterministic state machine
exchange-log/         append-only command/event log
exchange-replay/      replay and recovery runner
benchmarks/           allocation and latency micro-benchmarks
```

The replicated-log adapter lives in `../11-replicated-state-machine-aeron-java/`.
It should remain an adapter until the core contract is stable.

## First Java Exercises

1. Top-of-book aggregation from order commands.
2. Position updates from execution reports.
3. Replay start lookup by event position.
4. TTL/LRU dedup cache for client request ids.
5. Allocation comparison between boxed/object-heavy and buffer-oriented paths.

---

## 中文

本目录是第一个 Java 交易核心章节的目标归宿。

不要在这里构建宽泛的交易所。构建最小的 Java 表面来证明：

- 确定性命令应用；
- 订单簿和仓位状态；
- 追加写命令/事件日志；
- 从 offset 或位置重放；
- 低分配热路径选择；
- 用 JFR 和 async-profiler 分析；
- 到 Aeron 的清晰适配器边界。

## 提议的模块

```text
exchange-types/       commands, events, ids, price/quantity value rules
exchange-core/        deterministic state machine
exchange-log/         append-only command/event log
exchange-replay/      replay and recovery runner
benchmarks/           allocation and latency micro-benchmarks
```

复制日志适配器位于 `../11-replicated-state-machine-aeron-java/`。在核心契约
稳定之前它应保持为适配器。

## 第一个 Java 练习

1. 从订单命令聚合最优买卖价。
2. 从成交报告更新仓位。
3. 按事件位置查找重放起点。
4. 用于客户端请求 ID 的 TTL/LRU 去重缓存。
5. boxed/对象重路径与 buffer 导向路径的分配比较。