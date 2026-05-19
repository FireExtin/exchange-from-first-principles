# 06 SQL Facts And Outbox

[English](#english) · [中文](#中文)

## English

Purpose: introduce command/event/outbox without changing the business semantics
from chapter 05.

SQL still owns truth, but each accepted mutation also writes facts that can be
consumed, audited, replayed, and projected.

## Status

Status: contract scaffold. No command table, event table, outbox adapter, or
consumer implementation exists here yet.

## Contract Focus

- SQL mutation, command row, event rows, and outbox intent commit atomically;
- facts explain ledger, orders, executions, positions, and risk decisions;
- consumers use cursors and idempotency;
- this chapter prepares the move from row-state truth to ordered-log truth.

## 中文

目的：在不改变第 05 章业务语义的前提下，引入 command/event/outbox。

SQL 仍然拥有真相，但每次 accepted mutation 也写出可消费、可审计、可重放、可投影
的事实。

## 状态

状态：契约脚手架。本章尚无 command table、event table、outbox adapter 或 consumer
实现。

## 契约重点

- SQL mutation、command row、event rows 和 outbox intent 原子提交；
- facts 解释 ledger、orders、executions、positions 和 risk decisions；
- consumers 使用 cursor 和幂等；
- 本章为从行状态真相迁移到有序日志真相做准备。
