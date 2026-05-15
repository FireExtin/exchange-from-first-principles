# 03 Wallet Deposit Withdrawal

Primitive: outside money enters and leaves the system.

Question: why do callbacks, retries, and duplicate notifications need
idempotency?

This chapter models deposit callbacks and withdrawal requests with explicit
idempotency keys. It does not connect to a chain or payment provider yet.

Run:

```bash
go test ./...
```

---

## 中文

原语：外部资金进入和离开系统。

问题：为什么回调、重试和重复通知需要幂等性？

本章用显式幂等键建模入金回调和出金请求。它尚未连接到链或支付提供商。

运行：

```bash
go test ./...
```