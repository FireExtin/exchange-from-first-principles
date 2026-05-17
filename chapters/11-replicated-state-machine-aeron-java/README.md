# 11 Replicated State Machine Aeron Java

This module is a runnable skeleton.

Later boundary:

```text
Aeron Cluster committed message
  -> ExchangeClusteredService
  -> decode command
  -> call Java exchange-core
  -> encode events / snapshot
```

The skeleton is compiled against `io.aeron:aeron-cluster:1.51.0`.
`ExchangeClusteredService` implements the real
`io.aeron.cluster.service.ClusteredService` interface. In that API,
`onSessionMessage(ClientSession, long, DirectBuffer, int, int, Header)` is the
boundary where the service receives a committed ingress message from the
cluster log.

Do not put business rules here. This adapter should only translate between
Aeron buffers and the Java exchange-core contract.

## Reliability Rule

Distribution is not durability.

Aeron, UDP, multicast, or any fast messaging layer can move bytes quickly. This
chapter should not call that reliable until the recovery contract is explicit:

- every committed command has an ordered position;
- consumers can detect gaps;
- a restarted service can restore from snapshot and replay;
- a lagging consumer has a backpressure or disconnect policy;
- failover preserves the same externally visible command history.

The adapter boundary should therefore test transport shape separately from
business semantics. If a message is private, financial, or state-changing, it
must have a replay story.

## Run

This chapter expects Java 21. The repo has been tested with Azul Zulu 21 and
Gradle 9.5.

From this directory:

```bash
gradle --no-daemon clean test
```

From the repo root:

```bash
make test-java
```

On Java 21, tests using Agrona `UnsafeBuffer` need:

```text
--add-opens=java.base/jdk.internal.misc=ALL-UNNAMED
```

The Gradle test task already sets this flag.

---

## 中文

本模块是一个可运行骨架。

后续边界：

```text
Aeron Cluster committed message
  -> ExchangeClusteredService
  -> decode command
  -> call Java exchange-core
  -> encode events / snapshot
```

骨架编译为 `io.aeron:aeron-cluster:1.51.0`。`ExchangeClusteredService`
实现真实的 `io.aeron.cluster.service.ClusteredService` 接口。在该 API 中，
`onSessionMessage(ClientSession, long, DirectBuffer, int, int, Header)` 是
服务从集群日志接收已提交入口消息的边界。

不要在这里放业务规则。这个适配器只应在 Aeron 缓冲区和 Java exchange-core
契约之间转换。

## 可靠性规则

分发不等于持久，不等于可靠。

Aeron、UDP、多播或任何快速消息层都可以快速移动字节。但本章不应在恢复契约
显式之前称它可靠：

- 每条已提交命令都有有序位置；
- 消费者可以检测缺口；
- 重启后的服务可以从快照恢复并重放；
- 落后的消费者有背压或断开策略；
- 故障转移保持同一条外部可见命令历史。

因此适配器边界应把传输形态和业务语义分开测试。如果消息是私有的、金融性的或
会改变状态，它就必须有重放故事。

## 运行

本章需要 Java 21。仓库已用 Azul Zulu 21 和 Gradle 9.5 验证。

在本目录运行：

```bash
gradle --no-daemon clean test
```

从仓库根目录运行：

```bash
make test-java
```

在 Java 21 上，使用 Agrona `UnsafeBuffer` 的测试需要：

```text
--add-opens=java.base/jdk.internal.misc=ALL-UNNAMED
```

Gradle test 任务已设置此标志。
