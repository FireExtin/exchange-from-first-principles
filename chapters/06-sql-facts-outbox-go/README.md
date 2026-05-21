# 06 SQL Facts And Outbox

[English](#english) · [中文](#中文)

## English

> The SQL transaction committed. How does the rest of the system find out what changed? This chapter writes fact records and an outbox into the same commit, so downstream consumers know what happened without polling.

Purpose: introduce command/event/outbox without changing the business semantics
from chapter 05.

SQL still owns truth, but each accepted mutation also writes facts that can be
consumed, audited, replayed, and projected.

The outbox is the handoff from transaction handling logic to consumers. It says
what happened and where consumers can resume. It should not let consumers infer
ledger postings from hidden row changes.

## Status

Status: contract scaffold. No command table, event table, outbox adapter, or
consumer implementation exists here yet.

## Contract Focus

- SQL mutation, command row, event rows, and outbox intent commit atomically;
- facts explain ledger, orders, executions, positions, and risk decisions;
- facts label whether they are posted financial facts, operational facts, or
  derived/prospective model facts;
- consumers use cursors and idempotency;
- this chapter prepares the move from row-state truth to ordered-log truth.

## 中文

> SQL 事务提交了，下游系统怎么知道发生了什么？这章把事实记录和 outbox 写进同一个提交，让下游消费者不需要轮询就能知道。

目的：在不改变第 05 章业务语义的前提下，引入 command/event/outbox。

SQL 仍然拥有真相，但每次 accepted mutation 也写出可消费、可审计、可重放、可投影
的事实。

outbox 是 transaction handling logic 和消费者之间的交接点。它说明发生了什么，以及
consumer 应该从哪里继续。它不应该让 consumer 从隐藏行变更里猜 ledger postings。

## 状态

状态：契约脚手架。本章尚无 command table、event table、outbox adapter 或 consumer
实现。

## 契约重点

- SQL mutation、command row、event rows 和 outbox intent 原子提交；
- facts 解释 ledger、orders、executions、positions 和 risk decisions；
- facts 标明自己是已过账财务事实、运营事实，还是派生/未来模型事实；
- consumers 使用 cursor 和幂等；
- 本章为从行状态真相迁移到有序日志真相做准备。
