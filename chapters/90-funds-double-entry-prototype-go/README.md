# 90 Funds Double-Entry Prototype

[English](#english) · [中文](#中文)

## English

> When money moves, where does it go? Every dollar that leaves one account must arrive somewhere else — no exceptions. This is double-entry accounting.

Purpose: define the smallest double-entry invariant as a contract scaffold.

This appendix chapter demonstrates the core principle: every transaction must
balance to zero per asset across all accounts. The `external` account absorbs
the other side of incoming funds.

Implementation is intentionally absent. The tests define the expected behavior.
Implement `internal/ledger/ledger.go` to make them pass.

## Status And Run

Status: contract scaffold.

```bash
go test ./...
```

Tests compile and are expected to panic at TODO boundaries until the ledger is
implemented.

## 中文

> 钱移动的时候去哪了？从一个账户离开的每一分钱，必须出现在另一个账户里——没有例外。这就是双分录。

目的：把最小 double-entry 不变量定义为契约脚手架。

这个附录章节展示核心原则：每笔交易在所有账户中对每种资产必须平衡为零。
`external` 账户承接入金的另一侧。

实现刻意留空。测试定义了预期行为。实现 `internal/ledger/ledger.go` 以通过测试。

## 状态与运行

状态：契约脚手架。

```bash
go test ./...
```

测试可以编译，在学习者实现账本之前预期会在 TODO 边界失败。
