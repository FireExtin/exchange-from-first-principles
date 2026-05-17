# 03 Wallet Deposit Withdrawal

Primitive: outside money enters and leaves the system.

Question: why do callbacks, retries, and duplicate notifications need
idempotency?

This chapter models deposit callbacks and withdrawal requests with explicit
idempotency keys. It does not connect to a chain or payment provider yet.

The `adapter` package exposes this workflow through the shared funds contract.
That lets the integration tests compare this direct workflow with the replay
engine in chapter 04.

The reconciliation lab is documented in [RECONCILIATION_LAB.md](./RECONCILIATION_LAB.md).
It adds definitions and opt-in tests for provider callbacks, settlement reports,
bank/custody/on-chain records, partial refunds, timing differences, duplicates,
record matching rules, and adjustment journals. This is reconciliation matching,
not order-book matching. Its core logic is intentionally left as an exercise.

Run:

```bash
go test ./...
```

Run the reconciliation exercise tests:

```bash
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

They are expected to fail at TODO boundaries until you implement the lab.

---

## 中文

原语：外部资金进入和离开系统。

问题：为什么回调、重试和重复通知需要幂等性？

本章用显式幂等键建模入金回调和出金请求。它尚未连接到链或支付提供商。

`adapter` 包将本章工作流暴露为共享资金契约，集成测试可以用它对比本章的
直接处理模型和第 04 章的重放模型。

[RECONCILIATION_LAB.md](./RECONCILIATION_LAB.md) 记录了对账实验。它只加入
定义和可选测试，用来覆盖 provider callback、settlement report、
bank/custody/on-chain record、部分退款、时间差、重复事件和人工 adjustment
journal，以及记录匹配规则。这里说的是对账匹配，不是 order book 撮合。核心逻辑刻意留作练习。

运行：

```bash
go test ./...
```

运行对账练习测试：

```bash
go test -tags reconciliation_lab_todo ./internal/reconciliation
```

在你实现实验前，它们应该失败在 TODO 边界。
