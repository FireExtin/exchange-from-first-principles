# 92 Wallet Idempotency Prototype

[English](#english) · [中文](#中文)

## English

Purpose: preserve the original runnable wallet workflow prototype.

This appendix prototype covers deposits, withdrawal requests, withdrawal
confirmations, transfers, idempotency, and the reconciliation lab. It remains
the current runnable adapter for the funds conformance suite, but it is not the
complete exchange semantic contract.

## Status And Run

Status: runnable Go prototype. The reconciliation lab is an opt-in TODO
exercise behind the `reconciliation_lab_todo` build tag.

Run:

```bash
go test ./...
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

The second command is expected to fail at TODO boundaries until the lab is
implemented by the learner.

## 中文

目的：保留原来的可运行钱包工作流原型。

这个 appendix prototype 覆盖入金、出金请求、出金确认、转账、幂等和对账实验。
它仍然是当前 funds conformance suite 的可运行 adapter，但不是完整 exchange
semantic contract。

## 状态与运行

状态：可运行 Go 原型。对账实验在 `reconciliation_lab_todo` build tag 后，是可选
TODO 练习。

运行：

```bash
go test ./...
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

第二条命令在学习者实现实验前应失败在 TODO 边界。
