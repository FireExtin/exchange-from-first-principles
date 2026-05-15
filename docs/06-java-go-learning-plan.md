# Java And Go Learning Plan

This is the study spine for the project.

## Priority

Java gets priority for the trading hot path. Go is important for service edges
and asset-system concerns.

## Java Track

Learn the Java that strengthens trading-system implementation:

- core collections: `ArrayList`, `HashMap`, `PriorityQueue`, `TreeMap`,
  `ArrayDeque`;
- `Comparator`, generics, `Map.computeIfAbsent`, `getOrDefault`, and common
  idioms;
- represent price and quantity with scaled `long`, not hot-path `BigDecimal`;
- object allocation, escape analysis, young/old generation pressure, and GC
  pause reasoning;
- `ByteBuffer`, Agrona `UnsafeBuffer`, and SBE-style fixed binary messages;
- single-threaded event loop and deterministic state-machine design;
- append-only command log, event log, snapshot, and replay;
- JFR and async-profiler for CPU/allocation profiling;
- Aeron publication/subscription and Aeron Cluster service boundaries.

## Go Track

Learn the Go needed for service edges and asset-system work:

- `context.Context` propagation, cancellation, and deadlines;
- `sync.Mutex`, `sync.RWMutex`, `sync.Map`, `atomic`, and when not to use
  channels;
- goroutine lifecycle discipline;
- gRPC unary calls and basic streaming boundaries;
- Kafka consumer offset handling, retries, and idempotent effects;
- pprof CPU, heap, block, and mutex profiles;
- ledger entries, frozen funds, in-flight funds, reconciliation, and callbacks.

Minimum exercises:

- implement a ledger API that writes debit/credit entries atomically;
- handle duplicate payment callbacks with an idempotency key;
- reconcile ledger-derived balances against a materialized account table;
- add a load generator and inspect the service with pprof.

## Time Rule

Focus on correctness first, then hot-path optimization. The database remains a
projection and audit surface, not the mutation center.

Rust is not deleted. It is paused.

---

## 中文

这是项目的学习主轴。

## 优先级

Java 优先用于交易热路径。Go 对服务边界和资产系统相关工作很重要。

## Java 学习路径

学习加强交易系统实现的 Java 知识：

- 核心集合：`ArrayList`、`HashMap`、`PriorityQueue`、`TreeMap`、`ArrayDeque`；
- `Comparator`、泛型、`Map.computeIfAbsent`、`getOrDefault` 和常用惯用法；
- 用缩放的 `long` 而非热路径上的 `BigDecimal` 表示价格和数量；
- 对象分配、逃逸分析、年轻代/老年代压力和 GC 停顿推理；
- `ByteBuffer`、Agrona `UnsafeBuffer` 和 SBE 风格的固定二进制消息；
- 单线程事件循环和确定性状态机设计；
- 追加写命令日志、事件日志、快照和重放；
- JFR 和 async-profiler 用于 CPU/分配分析；
- Aeron 发布/订阅和 Aeron Cluster 服务边界。

## Go 学习路径

学习服务边界和资产系统工作所需的 Go 知识：

- `context.Context` 传播、取消和截止时间；
- `sync.Mutex`、`sync.RWMutex`、`sync.Map`、`atomic`，以及何时不使用 channel；
- goroutine 生命周期规范；
- gRPC 一元调用和基本流边界；
- Kafka 消费者 offset 处理、重试和幂等效果；
- pprof CPU、堆、阻塞和互斥体分析；
- 账本条目、冻结资金、在途资金、对账和回调。

最低练习：

- 实现一个原子写入借方/贷方条目的账本 API；
- 用幂等键处理重复的支付回调；
- 将账本推导的余额与物化账户表对账；
- 添加压测生成器并用 pprof 检查服务。

## 时间规则

先关注正确性，再优化热路径。数据库保持作为投影和审计面，而不是变更中心。

Rust 并未删除，只是暂停。