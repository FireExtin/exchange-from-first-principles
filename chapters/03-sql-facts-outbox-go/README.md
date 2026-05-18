# 03 SQL Facts And Outbox

[English](#english) · [中文](#中文)

## English

Purpose: specify v1, where SQL still owns truth but every mutation also writes
ordered commands, events, and outbox records.

Rows remain durable and transactional. Facts begin to explain why those rows
exist.

## Status

Status: contract scaffold. No command table, event table, outbox adapter, or
consumer implementation exists here yet.

## Required Semantics

- SQL mutation, command record, event facts, and outbox publication intent are
  committed atomically;
- balances, orders, positions, and risk views can be explained from emitted
  facts;
- downstream consumers use cursors and idempotency, not best-effort callbacks;
- this bridge prepares the later move from row-state truth to ordered-log truth.

## 中文

目的：定义 v1，也就是 SQL 仍然拥有真相，但每次 mutation 同步写入有序 command、
event 和 outbox 记录。

数据库行仍然是持久且事务性的；事实开始解释这些行为什么存在。

## 状态

状态：契约脚手架。本章尚无 command table、event table、outbox adapter 或 consumer
实现。

## 必须保持的语义

- SQL mutation、command record、event facts 和 outbox 发布意图必须原子提交；
- balances、orders、positions 和 risk views 都能由发出的事实解释；
- 下游 consumer 使用 cursor 和幂等，而不是 best-effort callback；
- 这一桥接阶段为之后从行状态真相迁移到有序日志真相做准备。
