# 03 Wallet Deposit Withdrawal

Primitive: outside money enters and leaves the system.

Question: why do callbacks, retries, and duplicate notifications need
idempotency?

This chapter models deposit callbacks and withdrawal requests with explicit
idempotency keys. It does not connect to a chain or payment provider yet.

The `adapter` package exposes this workflow through the shared funds contract.
That lets the integration tests compare this direct workflow with the replay
engine in chapter 04.

Run:

```bash
go test ./...
```

---

## 中文

原语：外部资金进入和离开系统。

问题：为什么回调、重试和重复通知需要幂等性？

本章用显式幂等键建模入金回调和出金请求。它尚未连接到链或支付提供商。

`adapter` 包将本章工作流暴露为共享资金契约，集成测试可以用它对比本章的
直接处理模型和第 04 章的重放模型。

运行：

```bash
go test ./...
```
