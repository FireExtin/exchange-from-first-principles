# 91 Spot Settlement Transaction Prototype

[English](#english) · [中文](#中文)

## English

> Alice buys 1 BTC from Bob. Two things happen at once: USDT from Alice to Bob, BTC from Bob to Alice. Either both legs succeed, or neither does. This chapter proves that atomicity.

Purpose: define spot settlement atomicity as a contract scaffold.

This appendix chapter introduces the two-leg settlement pattern: buyer pays
QuoteAsset, seller delivers BaseAsset. Both legs must apply or neither does.
The `SettleEvents` function must return the two `funds.EventTransferred` facts
that record the movement.

Implementation is intentionally absent. The tests define the expected behavior.
Implement `internal/spot/settlement.go` to make them pass.

## Status And Run

Status: contract scaffold.

```bash
go test ./...
```

Tests compile and are expected to panic at TODO boundaries until the settlement
logic is implemented.

## 中文

> Alice 从 Bob 那里买了 1 BTC。两件事同时发生：USDT 从 Alice 到 Bob，BTC 从 Bob 到 Alice。要么都发生，要么都不发生。这章证明这个原子性。

目的：把现货结算原子性定义为契约脚手架。

这个附录章节引入两腿结算模式：买方支付 QuoteAsset，卖方交付 BaseAsset，两腿
必须同时成功或都不发生。`SettleEvents` 函数必须返回记录资金移动的两个
`funds.EventTransferred` 事实。

实现刻意留空。测试定义了预期行为。实现 `internal/spot/settlement.go` 以通过测试。

## 状态与运行

状态：契约脚手架。

```bash
go test ./...
```

测试可以编译，在学习者实现结算逻辑之前预期会在 TODO 边界失败。
