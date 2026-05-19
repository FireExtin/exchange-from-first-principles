# 92 Wallet Idempotency Prototype

[English](#english) · [中文](#中文)

## English

> A deposit webhook fires twice. The funds should arrive once, not twice. This chapter builds a wallet that is safe for at-least-once delivery: the same callback can arrive many times without double-crediting the account.

Purpose: define wallet idempotency and the reconciliation lab as a contract
scaffold.

This appendix chapter covers:

- deposits with callbackID deduplication (idempotent, at-least-once safe);
- withdrawal requests and confirmations as a two-phase lifecycle;
- confirmation deduplication via providerEventID;
- transfers with balance checking;
- reconciliation: turning provider callbacks, settlement reports, and bank
  records into an auditable match report with adjustment proposals.

The `adapter/engine.go` file wires up `wallet.Processor` to the `funds.Engine`
interface so the funds conformance suite can test this chapter alongside chapter
93.

Implementation is intentionally absent. Implement `internal/wallet/wallet.go`
and the adapter to make the main tests pass. The reconciliation lab is a
separate opt-in exercise.

## Status And Run

Status: contract scaffold. The reconciliation lab is an opt-in TODO exercise
behind the `reconciliation_lab_todo` build tag.

```bash
# Main wallet tests (panic at TODO boundaries until implemented)
go test ./...

# Reconciliation lab (panic at TODO boundaries until implemented)
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

Both commands are expected to fail at TODO boundaries until implemented.

## 中文

> 充值的 webhook 触发了两次。钱应该只到账一次，不能两次。这章实现一个对 at-least-once 传递安全的钱包：同一条回调来多少次，账户只记一次。

目的：把钱包幂等和对账实验定义为契约脚手架。

这个附录章节覆盖：

- 使用 callbackID 去重的入金（幂等，at-least-once 安全）；
- 出金请求和确认的两阶段生命周期；
- 使用 providerEventID 去重的确认；
- 带余额检查的转账；
- 对账：把 provider callbacks、结算报告和银行记录转化为可审计的匹配报告，
  带调整提案。

`adapter/engine.go` 把 `wallet.Processor` 接入 `funds.Engine` 接口，
使 funds conformance suite 能同时测试本章和第 93 章。

实现刻意留空。实现 `internal/wallet/wallet.go` 和 adapter 以通过主测试。
对账实验是独立的可选练习。

## 状态与运行

状态：契约脚手架。对账实验在 `reconciliation_lab_todo` build tag 后，是可选
TODO 练习。

```bash
# 主要钱包测试（在实现之前预期在 TODO 边界失败）
go test ./...

# 对账实验（在实现之前预期在 TODO 边界失败）
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

两条命令在学习者实现之前均预期失败在 TODO 边界。
