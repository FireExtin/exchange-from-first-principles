# Go Tools

[English](#english) · [中文](#中文)

## English

Go should own service and operational edges, not the trading hot path:

- load generation;
- gateway/API experiments;
- ledger simulation;
- payment callback idempotency;
- reconciliation reports;
- pprof-based diagnosis.

## First Go Exercises

1. Add a small ledger package with debit/credit entries.
2. Add an idempotent callback handler keyed by provider request id.
3. Compare account balance table vs ledger-derived balance.
4. Add a `pprof` endpoint and inspect CPU/heap profiles under load.

Keep goroutines close to IO boundaries. Do not use channels to split a business
ordering path unless the ordering and failure behavior are explicit.

---

## 中文

### Go 工具

Go 应拥有服务和运营边界，而不是交易热路径：

- 压测生成；
- 网关/API 实验；
- 账本模拟；
- 支付回调幂等性；
- 对账报告；
- 基于 pprof 的诊断。

## 第一个 Go 练习

1. 添加一个带借方/贷方条目的小账本包。
2. 添加一个按提供商请求 ID 键控的幂等回调处理器。
3. 比较账户余额表与账本推导余额。
4. 添加一个 `pprof` 端点并在负载下检查 CPU/堆 profile。

将 goroutine 保持在 IO 边界附近。除非排序和失败行为是显式的，
否则不要用 channel 分割业务排序路径。