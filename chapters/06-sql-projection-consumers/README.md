# 06 SQL Projection Consumers

[English](#english) · [中文](#中文)

## English

Purpose: specify v4, where SQL returns as the warm/cold-path projection and
consumer store.

The hot path emits facts. SQL consumers rebuild OMS views, ledger reports,
reconciliation inputs, compliance exports, risk views, cache rebuild state, and
push recovery checkpoints.

## Status

Status: README scaffold. No runnable implementation exists here yet.

## Boundary

- input: event stream plus snapshots;
- output: query tables, reports, reconciliation records, compliance exports,
  recovery cursors;
- required behavior: idempotent consumption, gap detection, rebuild from a known
  cursor;
- not owned here: hot-path command admission or matching.

## 中文

目的：定义 v4，也就是 SQL 回到 warm/cold path，作为 projection 和 consumer store。

热路径发出事实。SQL consumer 重建 OMS view、ledger report、对账输入、合规导出、
risk view、cache rebuild state 和 push recovery checkpoint。

## 状态

状态：README 脚手架。本章尚无可运行实现。

## 边界

- 输入：事件流加快照；
- 输出：查询表、报表、对账记录、合规导出、恢复 cursor；
- 必须行为：幂等消费、gap detection、从已知 cursor 重建；
- 不负责：热路径命令准入或撮合。
