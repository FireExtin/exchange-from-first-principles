# 04 Single-Node Memory Core

[English](#english) · [中文](#中文)

## English

Purpose: specify v2, where the hot trading path becomes a deterministic
single-node in-memory state machine.

The memory core owns hot-path state: reservations, order books, executions,
positions, and risk-admission state. It emits facts that later projection and
ledger consumers can explain.

## Status

Status: README scaffold. No runnable implementation exists here yet.

Implementation must wait until the exchange semantic contract is stable.

## Boundary

- input: sequenced commands;
- state: private in-memory hot state;
- output: ordered facts and snapshots;
- not owned here: SQL ledger storage, reconciliation, reporting, compliance,
  cache rebuild, or push-client recovery.

## 中文

目的：定义 v2，也就是交易热路径进入确定性的单机内存状态机。

内存核心拥有热路径状态：reservation、order book、execution、position 和
risk-admission state。它发出事实，供后续 projection 和 ledger consumer 解释。

## 状态

状态：README 脚手架。本章尚无可运行实现。

实现必须等 exchange semantic contract 稳定后再写。

## 边界

- 输入：已排序命令；
- 状态：私有内存热状态；
- 输出：有序事实和快照；
- 不归本章负责：SQL ledger 存储、对账、报表、合规、缓存重建或 push client 恢复。
