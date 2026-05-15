# System Boundary

The long-term hot-path source of truth starts at command sequencing.

Earlier chapters deliberately start with database transactions as the source of
truth, then migrate toward ordered facts and a deterministic state machine. See
`docs/08-truth-source-migration.md` for the full model.

```text
ingress
  -> assign seq
  -> append command log
  -> apply deterministic state machine
  -> emit events
```

Anything before sequencing is an adapter. Anything after event emission is a
projection.

## Java Owns First

- command sequence;
- append-only command log;
- order book;
- account reservation;
- matching;
- event generation;
- snapshot and replay rules.

Java is the primary implementation surface for the trading hot path.

## Go Owns Service Edges

- load generation;
- gateway and admin APIs;
- ledger and account-service simulation;
- materialized view readers;
- reconciliation reports;
- idempotency and callback handling;
- operational glue.

Go handles service design around assets, ledger, idempotency, reconciliation,
callbacks, and engineering tradeoffs.

## Aeron Owns Integration Ordering

- cluster ingress ordering;
- replicated log;
- leader election;
- cluster snapshots;
- flow control and backpressure.

## Rust Is Parked

The Rust workspace remains useful for long-term exploration of a clean hot path,
but it is not the active development track. Do not let Rust absorb time needed
for Java or Go development.

---

## 中文

长期热路径的真相源从命令排序开始。

早期章节有意从数据库事务作为真相源开始，然后迁移到有序事实和确定性状态机。详见
`docs/08-truth-source-migration.md`。

```text
ingress
  -> assign seq
  -> append command log
  -> apply deterministic state machine
  -> emit events
```

排序之前的任何东西都是适配器。事件发出之后的任何东西都是投影。

## Java 优先

- 命令排序；
- 追加写命令日志；
- 订单簿；
- 账户预留；
- 撮合；
- 事件生成；
- 快照和重放规则。

Java 是交易热路径的主要实现层面。

## Go 负责服务边界

- 压测生成；
- 网关和管理 API；
- 账本和账户服务模拟；
- 物化视图读取器；
- 对账报告；
- 幂等和回调处理；
- 运营胶水代码。

Go 处理资产、账本、幂等、对账、回调和工程权衡相关的服务设计。

## Aeron 负责集成排序

- 集群入口排序；
- 复制日志；
- 领导者选举；
- 集群快照；
- 流量控制和背压。

## Rust 暂停

Rust 工作区对长期探索干净的热路径仍有价值，但它不是活跃的开发主线。
不要让 Rust 占用 Java 或 Go 开发所需的时间。