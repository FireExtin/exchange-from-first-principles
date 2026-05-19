# 07 Single-Node Memory Core

[English](#english) · [中文](#中文)

## English

Purpose: move the same business semantics into a deterministic single-node
memory core.

The memory core owns hot-path state, but not external custody or cold-path
reporting. It consumes sequenced commands and emits facts.

## Status

Status: README scaffold. No runnable implementation exists here yet.

Implementation must wait until the semantic ramp and SQL contract are stable.

## Boundary

- input: sequenced commands;
- state: private in-memory reservations, orders, positions, and risk admission;
- output: ordered facts and snapshots;
- out of scope: SQL ledger storage, reconciliation, compliance, cache rebuild,
  push recovery.

## 中文

目的：将同一业务语义迁入确定性的单机内存核心。

内存核心拥有热路径状态，但不拥有外部托管或冷路径报表。它消费已排序命令并发出
事实。

## 状态

状态：README 脚手架。本章尚无可运行实现。

实现必须等语义爬坡和 SQL 契约稳定后再写。

## 边界

- 输入：已排序命令；
- 状态：私有内存 reservations、orders、positions 和 risk admission；
- 输出：有序事实和快照；
- 范围外：SQL ledger 存储、对账、合规、缓存重建、push recovery。
