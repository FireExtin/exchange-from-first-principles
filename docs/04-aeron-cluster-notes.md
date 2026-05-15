# Aeron Cluster Notes

Aeron Cluster is a later integration path. It owns cluster ordering and
replication, while the Java trading core remains the deterministic business
state machine.

Expected boundary:

```text
Aeron Cluster ingress
  -> replicated ordered log
  -> ClusteredService adapter
  -> Java exchange-core apply(command)
  -> snapshot/load snapshot
```

Do not hand-write Raft for this lab. Use mature infrastructure when high
availability becomes the actual problem.

The current adapter is deliberately thin. It should prove the real Aeron API
boundary, not hide business logic:

```text
onSessionMessage(ClientSession, long, DirectBuffer, int, int, Header)
  -> decode command
  -> call exchange-core
  -> publish/record resulting events
```

Keep business state out of the adapter. The adapter translates committed Aeron
messages into the same command contract used by replay and tests.

---

## 中文

Aeron Cluster 是一个后续的集成路径。它负责集群排序和复制，而 Java 交易核心
保持作为确定性业务状态机。

预期边界：

```text
Aeron Cluster ingress
  -> replicated ordered log
  -> ClusteredService adapter
  -> Java exchange-core apply(command)
  -> snapshot/load snapshot
```

这个实验不要手写 Raft。当高可用成为实际问题时，使用成熟的基础设施。

当前适配器有意保持薄。它应该证明真实的 Aeron API 边界，而不是隐藏业务逻辑：

```text
onSessionMessage(ClientSession, long, DirectBuffer, int, int, Header)
  -> decode command
  -> call exchange-core
  -> publish/record resulting events
```

保持业务状态在适配器之外。适配器将已提交的 Aeron 消息转换为与重放和测试
使用的相同命令契约。