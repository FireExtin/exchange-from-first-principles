# 09 SQL Projection Consumers

[English](#english) · [中文](#中文)

## English

Purpose: rebuild warm and cold views from the same emitted facts.

SQL returns as the consumer store for OMS views, ledger reports,
reconciliation, compliance, risk views, cache rebuild, and push recovery.

## Status

Status: README scaffold. No runnable implementation exists here yet.

## Boundary

- input: event stream plus snapshots;
- output: query tables, reports, reconciliation records, compliance exports,
  recovery cursors;
- required behavior: idempotent consumption, gap detection, rebuild from a known
  cursor;
- out of scope: hot-path command admission or matching.

## 中文

目的：从同一组 emitted facts 重建温路径和冷路径视图。

SQL 回到 consumer store，用于 OMS views、ledger reports、对账、合规、risk views、
cache rebuild 和 push recovery。

## 状态

状态：README 脚手架。本章尚无可运行实现。

## 边界

- 输入：事件流加快照；
- 输出：查询表、报表、对账记录、合规导出、恢复 cursor；
- 必须行为：幂等消费、gap detection、从已知 cursor 重建；
- 范围外：热路径命令准入或撮合。
