# 11 Replicated State Machine Aeron Java

This module is intentionally a skeleton.

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

On Java 21, tests using Agrona `UnsafeBuffer` need:

```text
--add-opens=java.base/jdk.internal.misc=ALL-UNNAMED
```

The Gradle test task already sets this flag.

---

## 中文

本模块故意是一个骨架。

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

在 Java 21 上，使用 Agrona `UnsafeBuffer` 的测试需要：

```text
--add-opens=java.base/jdk.internal.misc=ALL-UNNAMED
```

Gradle test 任务已设置此标志。