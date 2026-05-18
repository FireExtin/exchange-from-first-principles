# 01 Double-Entry Ledger

[English](#english) · [中文](#中文)

## English

Primitive: two accounts exchange value.

Question: how do we prove money was not created or destroyed?

This chapter keeps only the smallest useful invariant: every transaction must
balance to zero per asset. It is deliberately in-memory so the accounting rule
is visible before storage, API, wallet, and trading concerns enter the system.

## Status And Run

Status: runnable Go chapter.

Run:

```bash
go test ./...
go run ./cmd/demo
```

---

## 中文

原语：两个账户交换价值。

问题：如何证明资金没有被创造或销毁？

本章只保留最小有用的不变量：每笔交易必须按资产平衡为零。它故意放在内存中，
以便在存储、API、钱包和交易问题进入系统之前，账务规则是可见的。

## 状态与运行

状态：可运行 Go 章节。

运行：

```bash
go test ./...
go run ./cmd/demo
```
