# System Boundary

[English](#english) · [中文](#中文)

## English

This document owns language and system-boundary rules. For the semantic ramp
and later truth migration, see [Truth Source Migration](./04-truth-source-migration.md).

The exchange semantic contract is shared. Ownership changes by version:

```text
Go / SQL edges
  -> Java hot core
  -> Aeron/Raft ordering and replication
  -> SQL warm/cold projections
```

## Go And SQL Own Service Edges

Go and SQL are the natural surfaces for service edges, durable ledgers, and
consumer views:

- ACID SQL truth-source scaffolds;
- double-entry ledger and account models;
- deposits, withdrawals, callbacks, idempotency;
- reconciliation and adjustment workflows;
- outbox producers and consumers;
- SQL projections for OMS, reports, compliance, cache rebuild, and push
  recovery.

Go/SQL code should not pretend to be the low-latency trading core once the
project reaches the memory or replicated-log versions.

## Java Owns Hot-Path State

Java is the primary surface for the deterministic trading core:

- command sequencing;
- private in-memory state;
- reservation state;
- order books;
- matching and execution facts;
- position and risk-admission state;
- snapshots and replay rules.

Java hot-path code should emit facts that Go/SQL consumers can project and
reconcile.

## Aeron/Raft Own Ordering And Replication

Aeron/Raft-style infrastructure owns the replicated-log boundary, not business
rules:

- cluster ingress ordering;
- replicated command log;
- log position;
- snapshot and replay boundary;
- failover;
- flow control and backpressure.

The same exchange command should mean the same thing before and after
replication is introduced.

## Rust Is Exploratory

Rust remains a runtime/hot-path experiment. It can explore replay, parsing,
buffers, FFI, and latency-sensitive components, but it is not the canonical
exchange implementation track in this repository.

---

## 中文

本文档负责语言和系统边界规则。关于语义爬坡和后续真相迁移，见
[Truth Source Migration](./04-truth-source-migration.md)。

交易所语义契约是共享的。不同版本的归属不同：

```text
Go / SQL edges
  -> Java hot core
  -> Aeron/Raft ordering and replication
  -> SQL warm/cold projections
```

## Go 和 SQL 负责服务边界

Go 和 SQL 自然适合服务边界、持久账本和 consumer views：

- ACID SQL truth-source 脚手架；
- double-entry ledger 和 account models；
- 入金、出金、callback、幂等；
- 对账和 adjustment workflows；
- outbox producers 和 consumers；
- OMS、报表、合规、cache rebuild 和 push recovery 的 SQL projections。

项目进入内存或复制日志版本后，Go/SQL 代码不应伪装成低延迟交易核心。

## Java 负责热路径状态

Java 是确定性交易核心的主要实现表面：

- 命令排序；
- 私有内存状态；
- reservation state；
- order books；
- matching 和 execution facts；
- position 和 risk-admission state；
- snapshots 和 replay rules。

Java 热路径代码应该发出事实，供 Go/SQL consumer 投影和对账。

## Aeron/Raft 负责排序和复制

Aeron/Raft-style 基础设施负责复制日志边界，而不是业务规则：

- cluster ingress ordering；
- replicated command log；
- log position；
- snapshot 和 replay boundary；
- failover；
- flow control 和 backpressure。

同一条 exchange command 在引入复制前后应该具有相同业务含义。

## Rust 作为探索实验

Rust 保持 runtime/hot-path experiment 定位。它可以探索 replay、parsing、buffer、
FFI 和 latency-sensitive components，但不是本仓库的规范交易所实现路线。
