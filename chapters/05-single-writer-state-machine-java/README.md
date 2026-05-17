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

## State Ownership Rule

The single writer is not just a performance trick. It is a state-ownership
boundary.

This chapter should make the owned state explicit:

- order book state;
- position state;
- request-dedup state;
- replay cursor and last applied sequence;
- any derived top-of-book or exposure view maintained on the hot path.

The hot apply function should not call a database, remote service, or shared
cache to decide a mutation. Inputs should arrive as commands or versioned local
state. Outputs should leave as events.

That gives the chapter a clear shape:

```text
owned state + ordered command -> owned state' + events
```

## First Java Exercises

1. Top-of-book aggregation from order commands.
2. Position updates from execution reports.
3. Replay start lookup by event position.
4. TTL/LRU dedup cache for client request ids.
5. Allocation comparison between boxed/object-heavy and buffer-oriented paths.
6. State ownership note: list which state is private, which state is an input
   projection, and which state is emitted as facts.

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

## 状态归属规则

单写者不只是性能技巧，它也是状态归属边界。

本章应显式命名被拥有的状态：

- 订单簿状态；
- 仓位状态；
- 请求去重状态；
- 重放游标和最后应用序列；
- 热路径维护的最优价或敞口派生视图。

热路径 `apply` 函数不应调用数据库、远程服务或共享缓存来决定一次变更。输入应以
命令或带版本的本地状态进入，输出应以事件离开。

本章的形状应保持清楚：

```text
owned state + ordered command -> owned state' + events
```

## 第一个 Java 练习

1. 从订单命令聚合最优买卖价。
2. 从成交报告更新仓位。
3. 按事件位置查找重放起点。
4. 用于客户端请求 ID 的 TTL/LRU 去重缓存。
5. boxed/对象重路径与 buffer 导向路径的分配比较。
6. 状态归属说明：列出哪些状态是私有的，哪些状态是输入投影，哪些状态以事实发出。
