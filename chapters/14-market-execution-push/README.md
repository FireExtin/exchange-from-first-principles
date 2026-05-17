# 14 Market And Execution Push

Purpose: define how public market data and private execution updates leave the
core in a recoverable way.

The trading core emits facts. Push systems turn those facts into client-facing
streams. The push layer must be fast, but it must also make gaps and recovery
explicit.

First scope:

- public market-data push: order-book deltas, trades, tickers, marks;
- private execution push: order accepted, rejected, filled, cancelled;
- sequence numbers and per-stream ordering;
- snapshot plus delta recovery;
- replay window and resubscribe flow;
- batching, compression, and backpressure;
- client disconnect and slow-consumer policy;
- WebSocket/TCP first, UDP or multicast only after semantics are stable.

Design rules:

```text
snapshot(seq=N) + deltas(seq>N) -> reconstructed client view
```

```text
private execution reports are facts, not best-effort notifications
```

Distribution rule:

```text
multicast or fast fan-out distributes bytes; sequence + replay restores truth
```

The chapter should separate public and private streams. A missed ticker update
is annoying. A missed execution report creates user-visible truth ambiguity.

UDP or multicast can be useful for public market data because one publisher can
fan out to many consumers without per-client coordination. It is not sufficient
for reliable private state. Clients still need sequence numbers, gap detection,
snapshots, replay windows, and a policy for slow consumers.

The first lab should simulate:

- a client joining after the snapshot sequence;
- a dropped public delta followed by resync;
- an out-of-order private execution report being rejected until replay;
- a slow consumer being disconnected or forced to resubscribe.

The network discussion belongs here only as far as it affects stream semantics:
ordering, buffering, head-of-line blocking, backpressure, loss detection, and
recovery. Kernel bypass and runtime tuning stay in chapter 16.

---

## 中文

目的：定义公共行情数据和私有成交更新如何以可恢复的方式离开核心。

交易核心发出事实。推送系统将这些事实转化为面向客户端的流。推送层必须快速，
但也必须使间隙和恢复显式。

首个范围：

- 公共行情推送：订单簿增量、成交、行情、标记；
- 私有成交推送：订单接受、拒绝、成交、撤销；
- 序列号和每流排序；
- 快照加增量恢复；
- 重放窗口和重新订阅流程；
- 批处理、压缩和背压；
- 客户端断开和慢消费者策略；
- 先 WebSocket/TCP，在语义稳定后才 UDP 或多播。

设计规则：

```text
snapshot(seq=N) + deltas(seq>N) -> reconstructed client view
```

```text
private execution reports are facts, not best-effort notifications
```

分发规则：

```text
多播或快速 fan-out 分发字节；序列号 + 重放恢复真相
```

本章应分离公共和私有流。错过行情更新是烦人的。错过成交报告造成用户可见的
真相歧义。

UDP 或多播适合某些公共行情，因为一个发布者可以不为每个客户端单独协调地向
很多消费者 fan out。但它不足以承载可靠的私有状态。客户端仍然需要序列号、
缺口检测、快照、重放窗口和慢消费者策略。

第一个 lab 应模拟：

- 客户端在快照序列之后加入；
- 公共增量丢失后重同步；
- 私有成交回报乱序到达，直到重放前被拒绝；
- 慢消费者被断开或要求重新订阅。

网络讨论只在其影响流语义时才属于这里：排序、缓冲、行首阻塞、背压、
丢失检测和恢复。内核绕过和运行时调优留在第 16 章。
