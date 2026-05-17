# Ordering And Serial Semantics

[English](#english) · [中文](#中文)

## English

The central claim of this project is that many mechanisms that look different
are really different ways to choose a successful serial history.

```text
old_state + command -> new_state + events
```

The mechanism changes. The business question stays the same:

> In what order did successful commands become true?

## Three Ways To Choose Order

### Pessimistic Concurrency Control

Examples: two-phase locking, row locks, explicit mutexes.

Pessimistic control orders work before or during execution. Whoever obtains the
conflicting lock first gets to run first. Other contenders wait or time out.

This is simple for small systems because storage or lock managers hide the
ordering work. The cost is contention, deadlock handling, and behavior that can
be hard to explain when many locks interact.

### Optimistic Concurrency Control

Examples: MVCC validation, compare-and-swap, optimistic version checks.

Optimistic control allows work to run speculatively, then orders successful
commits at validation time. If two operations conflict, one commits and the
other retries or rolls back.

This can improve throughput when conflicts are rare. The cost is that failure
and retry become part of normal control flow, and the business layer must be
clear about what is safe to repeat.

### Distributed Consensus

Examples: Raft, Paxos, replicated logs.

Consensus orders commands through a replicated log. A command is not allowed to
mutate the state machine until enough replicas have accepted the log position.

This pays a high infrastructure cost: quorum, leader election, replication,
snapshots, backpressure, and operational discipline. In return, the business
layer sees a very direct model: command number `N` runs before command number
`N+1`.

## Serializable, Linearizable, Strictly Serializable

These terms are related but not identical.

Serializable means the final result is equivalent to some serial execution of
transactions. It does not, by itself, require that the serial order matches
real-time order observed by clients.

Linearizable means operations appear to take effect atomically at a single
point between call and response. If operation A finishes before operation B
starts, B must observe A or come after A.

Strict serializability combines serializability with real-time order. This is
often the effect people want when they say "the system behaves as if every
successful command ran one at a time."

For this project, the practical target is:

```text
all externally successful financial commands are explainable as one serial
history, and that history should not violate client-visible real-time order
```

## Why Total Order Feels Expensive But Reads Simply

Global ordering is expensive because the infrastructure must decide who goes
first before business mutation becomes true. That means logs, sequencers,
consensus, failover, replay, and careful backpressure.

But once the order exists, business logic becomes simple:

- no race inside the mutation boundary;
- no dirty reads inside the mutation boundary;
- no non-repeatable reads inside the mutation boundary;
- no phantom writes inside the mutation boundary;
- no distributed transaction ambiguity inside the mutation boundary;
- no need to reason about thread scheduling for core state.

The upper layer only asks:

```text
if command A runs before command B, what should happen?
```

That is why exchanges, ledgers, and risk systems often accept a heavy ordering
boundary. The cost is paid once in the substrate. The benefit is that every
business rule above it becomes easier to explain, test, replay, and audit.

## What This Does Not Mean

It does not mean every service should be single-threaded.

It does not mean every read must be strongly consistent.

It does not mean all reporting, compliance, push, cache, and analytics work
must sit on the hot path.

It means the commands that create money, fills, positions, and risk facts need a
clear ordering boundary. Everything else should state whether it is a hot-path
fact, a warm-path projection, or a cold-path view.

---

## 中文

本项目的核心主张是：许多看起来不同的机制实际上都是选择成功串行历史的不同方式。

```text
old_state + command -> new_state + events
```

机制在变。业务问题不变：

> 成功的命令以什么顺序变为真？

### 三种选择顺序的方式

#### 悲观并发控制

示例：两阶段锁、行锁、显式互斥锁。

悲观控制在执行前或执行期间对工作排序。先获得冲突锁的人先运行。
其他竞争者等待或超时。

这对小系统很简单，因为存储或锁管理器隐藏了排序工作。
代价是争用、死锁处理，以及当许多锁交互时行为可能难以解释。

#### 乐观并发控制

示例：MVCC 验证、compare-and-swap、乐观版本检查。

乐观控制允许工作投机运行，然后在验证时对成功提交排序。
如果两个操作冲突，一个提交，另一个重试或回滚。

当冲突很少时这可以提高吞吐量。代价是失败和重试成为正常控制流的一部分，
业务层必须清楚什么可以安全重复。

#### 分布式共识

示例：Raft、Paxos、复制日志。

共识通过复制日志对命令排序。命令在被足够多副本接受日志位置之前不允许
变更状态机。

这付出高基础设施成本：仲裁、领导者选举、复制、快照、背压和运营规范。
作为回报，业务层看到一个非常直接的模型：命令号 `N` 在命令号 `N+1` 之前运行。

### 可串行化、线性化、严格可串行化

这些术语相关但不相同。

可串行化意味着最终结果等价于某个串行执行的事务。它本身不要求
串行顺序与客户观察到的实时顺序匹配。

线性化意味着操作看起来在调用和响应之间的单个点原子生效。
如果操作 A 在操作 B 开始之前完成，B 必须观察 A 或在 A 之后。

严格可串行化结合了可串行化和实时顺序。这通常是当人们说"系统表现得
好像每个成功命令一次运行一个"时想要的效果。

对于本项目，实践目标是：

```text
所有外部成功的金融命令都可以解释为一个串行历史，
且该历史不应违反客户可见的实时顺序
```

### 为什么全排序感觉昂贵但读取简单

全局排序是昂贵的，因为基础设施必须在业务变更变为真之前决定谁先。
这意味着日志、排序器、共识、故障转移、重放和仔细的背压。

但一旦顺序存在，业务逻辑变得简单：

- 变更边界内没有竞态；
- 变更边界内没有脏读；
- 变更边界内没有不可重复读；
- 变更边界内没有幻写；
- 变更边界内没有分布式事务歧义；
- 不需要为核状态推理线程调度。

上层只问：

```text
如果命令 A 在命令 B 之前运行，会发生什么？
```

这就是为什么交易所、账本和风险系统经常接受一个重的排序边界。
成本在底层付一次。好处是其上每个业务规则变得更容易解释、测试、重放和审计。

### 这不意味着什么

它不意味着每个服务应该是单线程的。

它不意味着每次读取都必须强一致。

它不意味着所有报告、合规、推送、缓存和分析工作都必须坐在热路径上。

它意味着创造资金、成交、仓位和风险事实的命令需要一个清晰的排序边界。
其他一切都应该说明它是热路径事实、温路径投影还是冷路径视图。
