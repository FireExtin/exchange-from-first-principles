# 16 Hot-Path Runtime Paths

[English](#english) · [中文](#中文)

## English

### Role

This chapter sits after the architecture migration line (chapters 05–09) is
complete. Its purpose is not to build another exchange implementation. Its
purpose is to discuss what the hot-path runtime could look like once the
semantic contract is stable and the dominant cost is measurable latency rather
than missing correctness.

No implementation exists here intentionally. The right runtime choice depends
on measured bottlenecks, which cannot be known until chapters 07–09 are
runnable and profiled.

### The Precondition

Before choosing a runtime, the following must already be true:

- Chapter 07 (single-node memory core) is runnable and the state machine
  boundary is clear.
- Chapter 08 (replicated log core) is runnable and the command-log and
  snapshot format are stable.
- Chapter 09 (SQL projection consumers) is runnable and the projection gap
  and rebuild semantics are defined.
- A latency benchmark exists showing where time is spent.

Only after those conditions hold does it make sense to ask: can we remove the
remaining variance?

### Candidate Paths

The following are not recommendations. They are design spaces worth
understanding before choosing.

#### Path A: Rust + io_uring + monoio

- **Kernel interface**: io_uring submits I/O operations as a ring buffer.
  Completions are polled without syscall overhead on the hot path.
- **Runtime**: monoio is a single-threaded async runtime built directly on
  io_uring. One thread owns all I/O and state mutation.
- **Fit**: matches the single-writer model of the in-memory state machine.
  No cross-thread coordination on the hot path.
- **Cost**: Rust learning curve, Linux kernel version dependency (5.1+ for
  io_uring, 5.11+ for some sqe flags), limited ecosystem.
- **When it makes sense**: the bottleneck is syscall overhead and I/O
  submission, not compute.

#### Path B: Rust + tokio

- **Runtime**: tokio is a multi-threaded async runtime using work-stealing
  schedulers. Mature ecosystem.
- **Fit**: easier to integrate with existing Rust tooling (tonic, sqlx, etc.).
  Better for mixed workloads.
- **Cost**: harder to reason about per-core locality. Task migration can hurt
  cache behavior on latency-sensitive paths.
- **When it makes sense**: the team knows Rust, needs async ecosystem
  integrations, and tail latency (not median latency) is the target.

#### Path C: C++ + DPDK + Seastar

- **Kernel interface**: DPDK bypasses the kernel network stack entirely. Packets
  are handled in user space by polling a NIC ring buffer.
- **Runtime**: Seastar is an event-driven framework with a shard-per-core model.
  Each core owns its state and communicates via message passing.
- **Fit**: extreme tail latency requirements, hardware-aware allocation,
  zero-copy paths, kernel bypass.
- **Cost**: C++ complexity, DPDK hardware requirements (DPDK-supported NIC,
  hugepages, CPU isolation), significantly higher operational overhead.
- **When it makes sense**: the system is co-located with matching hardware,
  the team has C++ and systems expertise, and microsecond tail latency is a
  hard requirement.

### What To Measure First

Before committing to any path, measure the following in the chapter 07–09
implementations:

- Command processing latency (p50, p99, p999) under realistic load.
- Snapshot and replay time for realistic state sizes.
- Network round-trip time between clients and the replicated cluster.
- Allocation and GC behavior in the Java in-memory core.

Low-latency work should follow from measurement, not from assumptions.

### What This Chapter Is Not

- This is not a benchmark contest.
- This is not a claim that Rust or C++ is always faster.
- This is not a plan to rewrite the exchange in a different language.

The exchange semantic contract lives in `shared/go/exchange`. Any runtime
implementation must satisfy that contract and pass the same scenario tests that
all other versions pass.

### Status

Status: README only.

No runnable implementation exists. This chapter will only make sense after
chapters 07–09 are runnable and profiled.

---

## 中文

### 角色

本章在架构迁移线（第 05–09 章）完成后出现，目的不是再造一个交易所实现，
而是讨论：在语义契约已稳定、主要成本来自可测量的延迟而非缺失的正确性时，
热路径运行时可以是什么样的。

本章刻意没有实现。正确的运行时选择依赖可测量的瓶颈，而这些瓶颈在第 07–09
章可运行和可剖析之前是未知的。

### 前提条件

选择运行时之前，以下条件必须已满足：

- 第 07 章（单机内存核心）可运行，状态机边界清晰；
- 第 08 章（复制日志核心）可运行，命令日志和快照格式稳定；
- 第 09 章（SQL 投影消费者）可运行，投影 gap 和重建语义已定义；
- 有延迟基准测试说明时间花在哪里。

只有这些条件成立，才有意义问：能不能消除剩余的延迟抖动？

### 候选路径

以下不是推荐，而是在选择之前值得理解的设计空间。

#### 路径 A：Rust + io_uring + monoio

- **内核接口**：io_uring 通过 ring buffer 提交 I/O 操作，热路径轮询完成，
  没有 syscall 开销。
- **运行时**：monoio 是直接建在 io_uring 上的单线程异步运行时，一个线程拥有
  所有 I/O 和状态变更。
- **契合度**：与内存状态机的单写者模型契合，热路径无跨线程协调。
- **代价**：Rust 学习曲线，Linux 内核版本依赖（io_uring 需要 5.1+，部分特性
  需要 5.11+），生态相对有限。
- **适用场景**：瓶颈在 syscall 开销和 I/O 提交，而非计算。

#### 路径 B：Rust + tokio

- **运行时**：tokio 是使用 work-stealing 调度的多线程异步运行时，生态成熟。
- **契合度**：更容易与 Rust 工具链（tonic、sqlx 等）集成，适合混合负载。
- **代价**：难以推理每核局部性，任务迁移可能损害延迟敏感路径的缓存行为。
- **适用场景**：团队熟悉 Rust，需要异步生态集成，目标是尾延迟而非中位延迟。

#### 路径 C：C++ + DPDK + Seastar

- **内核接口**：DPDK 完全绕过内核网络栈，数据包在用户态通过轮询 NIC ring
  buffer 处理。
- **运行时**：Seastar 是 shard-per-core 模型的事件驱动框架，每个核心拥有自己
  的状态，通过消息传递通信。
- **契合度**：极端尾延迟需求，硬件感知分配，零拷贝路径，内核旁路。
- **代价**：C++ 复杂性，DPDK 硬件要求（支持 DPDK 的 NIC、hugepages、CPU
  隔离），运维开销显著更高。
- **适用场景**：系统与匹配硬件共置，团队有 C++ 和系统编程经验，微秒级尾延迟
  是硬性要求。

### 先要测量什么

在提交任何路径之前，先在第 07–09 章的实现中测量：

- 在真实负载下的命令处理延迟（p50、p99、p999）；
- 真实状态量下的快照和重放时间；
- 客户端和复制集群之间的网络往返时间；
- Java 内存核心的内存分配和 GC 行为。

低延迟工作应该来自测量，而不是来自假设。

### 本章不是什么

- 不是 benchmark 比赛；
- 不是说 Rust 或 C++ 一定更快；
- 不是把交易所用另一种语言重写的计划。

交易所语义契约在 `shared/go/exchange` 中。任何运行时实现都必须满足这个契约，
并通过其他版本已通过的同一组场景测试。

### 状态

状态：仅 README。

无可运行实现。本章只有在第 07–09 章可运行并可剖析之后才有意义。
