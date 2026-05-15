# 16 Low-Latency Runtime And Networking

Purpose: optimize the runtime and network path after the semantic model is
stable.

This chapter intentionally comes late. Low-latency work is valuable only after
the state transitions, logs, replay, risk boundaries, and push semantics are
clear enough to measure.

First scope:

- allocation profiling and zero-GC Java patterns;
- JFR and async-profiler for JVM code;
- `perf`, `strace`, eBPF, and flame graphs for native/runtime paths;
- CPU isolation, thread pinning, NUMA locality, and cache behavior;
- socket options, epoll/io_uring, RSS, NIC queues, and buffer sizing;
- DPDK, XDP, RDMA, and kernel-bypass research notes;
- PTP and timestamping notes for latency measurement.

Rule:

```text
measurement before optimization
```

Every optimization should name the workload, the metric, the baseline, the
variance, and the operational cost. The project should avoid heroic low-level
work until the simpler architecture has produced a real bottleneck.

---

## 中文

目的：在语义模型稳定后优化运行时和网络路径。

本章故意放在后面。低延迟工作只在状态转换、日志、重放、风控边界和推送
语义足够清晰以至于可以测量时才有价值。

首个范围：

- 分配分析和零 GC Java 模式；
- JFR 和 async-profiler 用于 JVM 代码；
- `perf`、`strace`、eBPF 和火焰图用于原生/运行时路径；
- CPU 隔离、线程固定、NUMA 局部性和缓存行为；
- socket 选项、epoll/io_uring、RSS、网卡队列和缓冲大小；
- DPDK、XDP、RDMA 和内核绕过研究笔记；
- PTP 和时间戳笔记用于延迟测量。

规则：

```text
measurement before optimization
```

每个优化应命名工作负载、指标、基线、方差和运营成本。在更简单架构产生
真实瓶颈之前，项目应避免英雄般的底层工作。