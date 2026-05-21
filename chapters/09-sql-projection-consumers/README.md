# 09 SQL Projection Consumers

[English](#english) · [中文](#中文)

## English

> The memory core produces a stream of facts. Reports, reconciliation, compliance, and risk all need those facts — but they can't read from hot-path memory directly. This chapter brings SQL back, now as a consumer of the fact stream rather than the source of truth.

Purpose: rebuild warm and cold views from the same emitted facts.

SQL returns as the consumer store for OMS views, ledger reports,
reconciliation, compliance, risk views, cache rebuild, and push recovery.
Projection tables may combine posted ledger facts with marks, model inputs, and
derived risk state, but they must label provenance. A projection can rebuild a
view; it must not directly mutate ledger truth.

## Status

Status: README scaffold. No runnable implementation exists here yet.

## Boundary

- input: event stream plus snapshots;
- output: query tables, reports, reconciliation records, compliance exports,
  recovery cursors;
- required behavior: idempotent consumption, gap detection, rebuild from a known
  cursor, and clear posted-vs-derived provenance;
- out of scope: hot-path command admission or matching.

## 中文

> 内存核心产生事实流。报表、对账、合规、风控都需要这些事实——但不能直接读内存核心。这章让 SQL 回来，这次它是事实流的消费者，不再是真相源。

目的：从同一组 emitted facts 重建温路径和冷路径视图。

SQL 回到 consumer store，用于 OMS views、ledger reports、对账、合规、risk views、
cache rebuild 和 push recovery。
projection tables 可以组合已过账 ledger facts、marks、模型输入和派生风险状态，但
必须标明 provenance。projection 可以重建视图；它不能直接改写 ledger truth。

## 状态

状态：README 脚手架。本章尚无可运行实现。

## 边界

- 输入：事件流加快照；
- 输出：查询表、报表、对账记录、合规导出、恢复 cursor；
- 必须行为：幂等消费、gap detection、从已知 cursor 重建，并清楚标注 posted 与
  derived provenance；
- 范围外：热路径命令准入或撮合。
