# 05 Replicated Log Core

[English](#english) · [中文](#中文)

## English

Purpose: specify v3, where a Raft/Aeron-style replicated log provides command
ordering and failover for the same exchange semantics.

The replicated log does not change business meaning. It changes who owns order,
durability, replay, and failover.

## Status And Run

Status: runnable Java skeleton. Business behavior is intentionally absent.

Run:

```bash
gradle --no-daemon clean test
```

## Boundary

- owns: replicated command order, log position, snapshot/replay contract,
  failover boundary;
- does not own: matching, accounting posting, risk rules, SQL projections, or
  full Raft implementation in this repo.

## 中文

目的：定义 v3，也就是用 Raft/Aeron-style replicated log 为同一套交易所语义提供
命令排序和故障切换。

复制日志不改变业务含义。它改变的是排序、持久性、重放和 failover 由谁承担。

## 状态与运行

状态：可运行 Java 骨架。业务行为刻意留空。

运行：

```bash
gradle --no-daemon clean test
```

## 边界

- 负责：复制命令顺序、日志位置、snapshot/replay 契约、failover 边界；
- 不负责：撮合、会计 posting、风控规则、SQL projection 或本仓库内的完整 Raft 实现。
