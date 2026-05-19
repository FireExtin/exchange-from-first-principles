# 08 Replicated Log Core

[English](#english) · [中文](#中文)

## English

> One machine holds all state in memory — if it crashes, the exchange stops. This chapter sends the same commands to multiple machines; each applies them independently and reaches the same state. That is replication.

Purpose: move the same deterministic command stream behind a replicated log /
Raft-style boundary.

Replication changes ordering, durability, recovery, and failover. It should not
change business meaning.

## Status And Run

Status: runnable Java skeleton. Business behavior is intentionally absent.

Run:

```bash
gradle --no-daemon clean test
```

## Boundary

- owns: replicated command order, log position, snapshot/replay contract,
  failover boundary;
- does not own: accounting posting, matching rules, risk rules, SQL projection,
  or a full Raft implementation in this repo.

## 中文

> 一台机器把所有状态放在内存里——崩了，交易所就停了。这章把同一条命令发给多台机器，每台独立执行后得到相同结果。这就是复制。

目的：将同一确定性命令流放到 replicated log / Raft-style 边界之后。

复制改变排序、持久性、恢复和 failover，不应该改变业务含义。

## 状态与运行

状态：可运行 Java 骨架。业务行为刻意留空。

运行：

```bash
gradle --no-daemon clean test
```

## 边界

- 负责：复制命令顺序、日志位置、snapshot/replay 契约、failover 边界；
- 不负责：accounting posting、撮合规则、风控规则、SQL projection 或本仓库内的
  完整 Raft 实现。
