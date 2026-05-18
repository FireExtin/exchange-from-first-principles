# Minimal Exchange Design Paper

[English](#english) · [中文](#中文)

## English

### 1. Thesis

This project derives an exchange-like system from small, inspectable
primitives. It does not start with Aeron, Raft, DPDK, Rust, or a complete
matching engine. It starts with the smallest financial question:

```text
can two accounts exchange value without creating or losing money?
```

From there, each stage introduces one new pressure:

- more explicit money facts;
- more explicit ordering;
- more explicit recovery;
- clearer system boundaries;
- lower variance on the hot path;
- better auditability on the cold path.

The central thesis is:

```text
Modern exchange architecture is not a pile of components.
Architecture changes are not goals by themselves.
They are trade-offs among Correctness, Performance, and Reliability.
Each stage chooses where to pay for ordering, durability, recovery,
latency, and operational complexity.
```

The core transition can be written as:

```text
old_state + command -> new_state + events
```

A database transaction, an in-memory state machine, and a replicated state
machine can all present successful business changes as an explainable serial
history. They differ in where they pay for ordering, durability, recovery, and
operational complexity.

This paper is the project-level design map. Individual chapters should stay
smaller: one pressure, one mechanism, one runnable experiment or test.

### 2. Architecture Lenses For Trade-Offs

Correctness, performance, and reliability are the objective frame. The project
uses four architecture lenses to make the trade-offs concrete.

| Dimension | Early Form | Later Form | Why It Changes |
| --- | --- | --- | --- |
| Truth source | Database rows | Ordered commands, events, snapshots, projections | Correctness and audit need explainable facts, not only current rows |
| Ordering model | SQL locks, MVCC, commit order | Sequencer, single writer, replicated log | Performance and fairness depend on where contention is paid |
| Recovery model | Restore DB backup and inspect rows | Snapshot plus replay from a known position | Reliability depends on rebuilding exact state after failure |
| Ownership boundary | Shared store queried by many services | Named state owner publishes facts | Performance and reliability improve when state ownership is explicit |

The migration is not a rejection of databases. Databases remain excellent for
ledgers, reports, settlement state, reconciliation, compliance archives, and
many payment workflows. The migration happens when the hot path needs clearer
ordering and more predictable latency than general-purpose storage can provide.

### 3. System Evolution

#### 3.1 Stage One: Database ACID As Source Of Truth

The first correct system can be built with SQL transactions:

```text
request
  -> begin transaction
  -> validate current rows
  -> write ledger/order/wallet rows
  -> commit
```

This is the right starting point for:

- double-entry ledger;
- basic spot settlement;
- deposit callbacks;
- withdrawal state transitions;
- provider records;
- reconciliation reports;
- admin adjustments.

The main invariant is conservation of value:

```text
sum(debits) == sum(credits)
```

The database gives atomicity, durability, and inspectability. It is the most
understandable first truth source.

Its limits appear when the system needs to answer questions like:

- which command produced this exact state?
- what was the system state after sequence `N`?
- why did this order win the race?
- can another process replay the same facts and rebuild the same view?
- can the hot path avoid storage locks and storage I/O?

Those are not database defects. They are signals that the architecture is ready
to make ordering and facts explicit.

#### 3.2 Stage Two: Command Log And Outbox Inside The Database

The bridge stage keeps SQL but records facts:

```text
request
  -> begin transaction
  -> insert command_log
  -> update authoritative rows
  -> insert event_log / outbox
  -> commit
```

Now the source of explanation begins to shift:

```text
facts = command_log + event_log
views = balances + orders + positions + reports
```

The outbox pattern teaches several core lessons:

- state change and publication must be tied together;
- duplicate messages are normal;
- consumers must be idempotent;
- projections can lag behind facts;
- reconciliation is a system boundary, not a cleanup script.

This is still not the final hot path. It is the bridge from row-centered
thinking to fact-centered thinking.

#### 3.3 Stage Three: Deterministic In-Memory State Machine

The trading hot path eventually wants a stricter mutation boundary:

```text
command
  -> sequence number
  -> deterministic apply()
  -> events
  -> snapshot / replay boundary
```

Inside this boundary, one ordered command mutates core state at a time. This
does not mean the whole system is single-threaded. It means the most sensitive
business mutation has a named ordering point.

The state machine owns private hot state:

- order book;
- order status;
- reservations;
- execution facts;
- minimal position hooks;
- risk-admission inputs that must be synchronous.

It should not call databases, remote services, clocks, random generators, or
slow dependencies from the mutation function. Inputs such as timestamps, ids,
marks, and configuration versions must be part of the command stream or a
versioned snapshot.

The business equivalence with a serializable database can be expressed as:

```text
DB:             State_0 --tx_1-->  State_1 --tx_2-->  State_2
State machine:  State_0 --cmd_1--> State_1 --cmd_2--> State_2
```

If the same operations are applied in the same order by a deterministic
transition function, they produce the same externally observable business
history. The difference is where ordering is enforced:

- the database hides ordering inside locks, MVCC, indexes, WAL, and commit
  rules;
- the state machine exposes ordering before business mutation.

This is why the same scenario tests should run against multiple
implementations. The substrate may change, but the contract should not.

#### 3.4 Stage Four: Replicated State Machine

Once one state machine is understandable, replication becomes the next pressure:

```text
client command
  -> replicated log / Aeron Cluster / Raft-style ordering
  -> node A apply(command)
  -> node B apply(command)
  -> node C apply(command)
  -> same state if same snapshot and same log
```

The project should not implement consensus from scratch. It should define what
the exchange core asks from the replicated log:

- command encoding;
- session identity;
- accepted log position;
- snapshot format;
- replay start point;
- backpressure semantics;
- failover behavior;
- event emission boundary.

The learning goal is not to become a consensus implementer. The learning goal
is to understand why a trading core asks for total order, replay, and failover
at this boundary.

#### 3.5 Stage Five: Projections, Caches, Push, And Runtime

After facts are ordered and recoverable, the system has to serve many readers:

- OMS;
- ledger projection;
- reconciliation;
- compliance;
- public market data;
- private execution reports;
- user/account permission caches;
- mark-price and risk projections;
- dashboards and reports.

These components should not pull hot state directly from the core. They should
consume versioned publications:

```text
snapshot(version=N) + deltas(version>N) -> local projection
```

Later runtime work should optimize only after the semantics are stable. Warmup,
allocation control, pooling, off-heap/buffer usage, CPU isolation, and
networking experiments belong here, not at the beginning.

### 4. Design Principles

#### 4.1 Truth Should Have A Shape

Every stage must say what the current truth source is.

Early truth:

```text
balances, ledger_entries, orders, trades, deposits, withdrawals
```

Later truth:

```text
command_log, event_log, snapshots, replicated log position
```

Views such as `available_balance`, `position_qty`, `top_of_book`, and
`risk_exposure` are useful, but they should be explainable from facts.

#### 4.2 Facts Before Views

The system should prefer immutable facts before mutable summaries:

```text
deposit accepted
withdrawal requested
cash reserved
order accepted
trade executed
position updated
margin checked
provider record ingested
bank record matched
```

When facts and views disagree, reconciliation needs a path back to the raw
facts and the rule version that interpreted them.

#### 4.3 Prefer Explicit Order Over Hidden Contention

Locks are acceptable in many payment and ledger workflows. They become harder
to explain in the trading hot path when tail latency, retries, and fairness are
part of the business contract.

The state-machine model pays more infrastructure cost up front:

- sequencing;
- logs;
- snapshots;
- replay;
- backpressure;
- failover;
- operational tooling.

In return, the business question becomes simpler:

```text
if command A ran before command B, what should happen?
```

#### 4.4 Ordering Mechanisms Are Different Places To Pay

| Mechanism | Where Ordering Is Paid | Strength | Cost |
| --- | --- | --- | --- |
| Pessimistic locking / 2PL | Before or during execution | Simple correctness for shared rows | Contention and deadlocks |
| MVCC / CAS / optimistic validation | At commit time | High read concurrency | Retries and conflict handling |
| Single writer / sequencer | Before mutation | Simple deterministic core | Partitioning and recovery design |
| Raft / Paxos / replicated log | Before replicated execution | Failover and consistency | Operational and latency cost |

Serializable isolation gives equivalence to some serial history.
Linearizability adds real-time constraints. Strict serializability combines
both effects. The project uses these ideas as comparison language, not as
academic decoration.

#### 4.5 Separate Hot, Warm, And Cold Paths

Hot path:

```text
admitted command -> ordered mutation -> emitted facts
```

Warm path:

```text
facts + market data -> positions -> exposure -> risk alerts/actions
```

Cold path:

```text
facts -> reports -> reconciliation -> audit -> compliance export
```

Cold path work is not less important. It is simply not allowed to block the hot
path. The design should be explainable without making every explanation
synchronous.

#### 4.6 State Ownership And Data Gravity

Hot state becomes expensive when many components repeatedly fetch, join, and
reinterpret it. The design should ask:

```text
should the function move to the state instead of moving the state to every function?
```

Every hot state should have an owner:

- who may mutate it;
- how updates are published;
- how consumers detect gaps;
- how consumers rebuild after restart;
- what stale behavior is allowed.

This applies to order books, positions, account status, risk limits,
instrument rules, and mark-price inputs.

#### 4.7 Publication Before Pull

Reference data, account permissions, risk configuration, execution facts, and
market state should not be repeatedly pulled on the hot path.

The owner should publish versioned changes. Consumers should maintain local
projections and know their failure policy:

- fail closed;
- continue for a bounded stale window;
- rebuild from snapshot plus deltas;
- reject until a missing gap is replayed.

This makes responsibility visible: the owner publishes, the consumer detects
staleness and rebuilds.

#### 4.8 Distribution Is Not Durability

Fast distribution is not the same as reliable truth.

UDP multicast, for example, lets a sender transmit once to a multicast group
that many receivers subscribe to. It is efficient for fan-out, but the
transport does not remember late joiners, lost packets, slow consumers, or
restart recovery.

A reliable stream needs more than fast delivery:

- sequence numbers;
- gap detection;
- snapshots;
- replay;
- backpressure;
- a durable or replicated source of facts;
- explicit recovery ownership.

Public market data can sometimes tolerate bounded loss if clients can resync.
Private execution reports, ledger facts, and risk decisions cannot be treated
as best-effort notifications.

#### 4.9 Performance Is A Consequence Of Shape

Low latency should not be a bag of tricks. It should follow from the system
shape:

- hot mutable state has an owner;
- commands are ordered before mutation;
- consumers use local projections;
- durable facts are written at clear boundaries;
- replay and snapshot semantics are explicit;
- allocation and warmup behavior are measured.

Only after this shape exists does it make sense to study pooling, off-heap
buffers, zero-allocation paths, CPU pinning, NUMA locality, kernel bypass, or
runtime-specific tuning.

### 5. Domain Boundaries

| Boundary | Owns | Emits | Must Not Do |
| --- | --- | --- | --- |
| Ledger | Value movement and double-entry facts | Journal entries, balance facts | Hide unmatched external money |
| Wallet | Deposit/withdrawal lifecycle | Provider callbacks, withdrawal state, raw records | Apply duplicate external events twice |
| Reconciliation | Raw external records and matching reports | Exceptions, match results, adjustment proposals | Silently mutate ledger |
| Matching | Price-time ordering and executions | Trades, fills, order status | Depend on remote calls during mutation |
| Position | Exposure by account/instrument | Position updates, realized/unrealized PnL inputs | Treat open orders as executions |
| Margin | Equity and requirement calculation | Margin snapshots, admissibility inputs | Hide formulas inside handlers |
| Pre-trade risk | Admission decision | Accept/reject facts | Become a strategy engine |
| Risk cluster | Continuous exposure projection | Alerts, risk actions, projection state | Pretend eventual projections are synchronous truth |
| OMS / compliance | Workflow, reporting, audit | Reports, case state, regulatory exports | Block core mutation |
| Cache / market state | Local projections of owned facts | Snapshot/delta versions, gap status | Serve unbounded stale data silently |
| Push gateway | Public/private stream delivery | Sequenced snapshots and deltas | Claim best-effort delivery is truth |

### 6. Trading Desk Extension

The exchange core decides what is true inside the venue. A trading desk decides
what to do given market state, inventory, risk, and external venues.

Desk components should appear later, after the core can emit reliable market
data, execution reports, positions, risk views, and reconciliation facts.

Natural emergence:

```text
matching
  -> executions
  -> positions
  -> exposure
  -> pre-trade and continuous risk
  -> market data and mark inputs
  -> external venue routing
  -> hedging / best execution / strategy
```

Possible later desk layer:

```text
external market data
  -> pricing / signals
  -> algo decision
  -> pre-trade risk
  -> order router
  -> external execution reports
  -> positions / hedger / reconciliation
```

Boundary rule: do not let desk concerns pollute the exchange core. The
matching engine should not know about arbitrage. The ledger should not know
about best execution. Pre-trade risk should not become a strategy engine.

### 7. Failure And Recovery Model

| Stage | Typical Failure | Required Recovery Story |
| --- | --- | --- |
| DB truth source | Transaction abort, duplicate callback, partial workflow | Atomic rollback, idempotency keys, raw external records |
| Outbox / command log | Published state and message diverge | Transactional outbox, idempotent consumers, resend |
| In-memory state machine | Process crash after sequence `N` | Snapshot plus replay from durable command/event log |
| Replicated state machine | Leader crash, follower lag, client retry | Log position, session identity, replay, failover |
| Cache / projection | Gap, stale snapshot, consumer restart | Versioned snapshot, delta replay, fail-open/closed policy |
| Push stream | Packet loss, late joiner, slow client | Sequence numbers, snapshots, replay, backpressure |
| Reconciliation | Provider, bank, custody, and internal records disagree | Immutable raw records, normalized records, match reports, manual adjustment journal |
| Runtime | Warmup variance, GC/allocation spikes, noisy network path | Benchmarks, profiles, warmup discipline, allocation control, runbooks |

### 8. Chapter Map

| Phase | Chapters | Purpose |
| --- | --- | --- |
| Version line | 01-06 | Exchange contract, ACID SQL, SQL facts/outbox, memory core, replicated log, SQL projections |
| Domain deep dives | 07-12 | Order book mechanics, position/PnL, margin/risk, risk projection, cache coherence, push |
| Runtime experiments | 13-14 | Rust hot path experiments and low-latency runtime/networking measurement |
| Desk extension | 15-19 | External market data, pricing, routing, hedging, best execution, simple strategy |
| Appendix prototypes | 90-93 | Earlier runnable Go funds prototypes |

The chapter rule is:

```text
one pressure -> one mechanism -> one explicit semantic comparison
```

Each chapter should answer:

1. What correctness property must stay stable?
2. What performance pressure is being introduced?
3. What reliability or recovery failure does this model expose?
4. Which architecture lens changed: truth, order, recovery, or ownership?
5. What new operational cost did the next model introduce?

### 9. Non-Goals

This project is not trying to be:

- a production exchange;
- a complete matching engine library;
- a consensus implementation;
- a full wallet or custody system;
- a real trading strategy;
- a benchmark contest;
- a pile of impressive infrastructure without a semantic reason.

The goal is to build just enough code and tests to make each architectural
pressure visible.

### 10. Success Criteria

The project succeeds if a reader can explain:

- why DB ACID is the right first model;
- why ordered facts become more useful than current rows;
- why the hot path eventually wants a deterministic mutation boundary;
- why replication needs a log, snapshots, replay, and backpressure;
- why reconciliation is record matching and evidence, not just arithmetic;
- why caches need owners, versions, and rebuild policies;
- why private execution reports cannot be best-effort notifications;
- why low latency work should start from ownership and measurement, not tricks.

The strongest proof is cross-version testing:

```text
same business scenario
same observable semantics
different execution substrate
```

That is the heart of the project.

---

## 中文

### 1. 核心论点

这个项目不是从 Aeron、Raft、DPDK、Rust 或完整撮合引擎开始，而是从最小的
金融问题开始：

```text
两个账户交换价值，系统能不能证明没有凭空造钱，也没有丢钱？
```

然后每一章只引入一个新的压力：

- 钱的事实要更清楚；
- 并发顺序要更清楚；
- 故障恢复要更清楚；
- 系统边界要更清楚；
- 热路径延迟要更可解释；
- 冷路径审计要更可靠。

核心论点是：

```text
现代交易所架构不是一堆组件的堆叠，
架构变化是在正确性、性能和可靠性之间做取舍。
每个阶段都在选择：排序、持久化、恢复、延迟和运营复杂度分别在哪里付费。
```

最核心的业务转移可以写成：

```text
old_state + command -> new_state + events
```

数据库事务、内存状态机、复制状态机都可以把成功的业务变更呈现为一条可解释
的串行历史。它们的区别在于：排序、持久化、恢复和运维复杂度分别由谁承担。

### 2. 用于取舍分析的架构视角

正确性、性能和可靠性是目标框架。项目使用四个架构分析视角把取舍具体化。

| 维度 | 早期形态 | 后期形态 | 为什么会变化 |
| --- | --- | --- | --- |
| 真相源 | 数据库行 | 有序命令、事件、快照、投影 | 正确性和审计需要可解释事实，而不只是当前行 |
| 排序模型 | SQL 锁、MVCC、提交顺序 | sequencer、单写者、复制日志 | 性能和公平性取决于竞争在哪里付费 |
| 恢复模型 | 恢复数据库后查当前行 | 从已知位置快照加重放 | 可靠性依赖故障后能重建精确状态 |
| 归属边界 | 多个服务查询共享存储 | 明确的状态 owner 发布事实 | 状态归属显式后，性能和可靠性更可控 |

这不是说数据库不好。数据库仍然很适合账本、报表、结算状态、对账、合规归档
和很多支付工作流。迁移发生在热路径需要更清晰的顺序和更稳定的延迟时。

### 3. 演进路径

#### 3.1 数据库 ACID 作为真相源

第一版正确系统可以用 SQL 事务完成：

```text
request
  -> begin transaction
  -> validate current rows
  -> write ledger/order/wallet rows
  -> commit
```

它适合双分录账本、现货结算、充值回调、提现状态流转、provider records、对
账报表和人工调整。

核心不变量是：

```text
sum(debits) == sum(credits)
```

数据库给了原子性、持久性和可检查性，所以它是最容易理解的第一个真相源。

#### 3.2 数据库里的命令日志和 outbox

桥接阶段仍然使用 SQL，但开始记录事实：

```text
request
  -> begin transaction
  -> insert command_log
  -> update authoritative rows
  -> insert event_log / outbox
  -> commit
```

系统解释开始从当前行转向事实：

```text
facts = command_log + event_log
views = balances + orders + positions + reports
```

这个阶段教会几个重要原则：状态变更和消息发布要绑定；重复消息是常态；消费
者必须幂等；投影可以落后于事实；对账是系统边界，不是事后脚本。

#### 3.3 确定性内存状态机

交易热路径最终会需要更严格的变更边界：

```text
command
  -> sequence number
  -> deterministic apply()
  -> events
  -> snapshot / replay boundary
```

这个边界里，一次只有一个有序命令修改核心状态。这不是说整个系统都单线程，
而是最敏感的业务变更必须有明确排序点。

状态机拥有私有热状态：订单簿、订单状态、资金冻结、成交事实、最小仓位钩子
和同步风控输入。它不应该在状态转移中调用数据库、远程服务、时钟、随机数或
慢依赖。

数据库和状态机的等价关系可以这样看：

```text
DB:             State_0 --tx_1-->  State_1 --tx_2-->  State_2
State machine:  State_0 --cmd_1--> State_1 --cmd_2--> State_2
```

如果同样的操作按同样的顺序执行，并且转移函数是确定性的，对外可观察的业务
历史就应该一致。区别只在于：数据库把顺序藏在锁、MVCC、索引、WAL 和提交
规则里；状态机把顺序放在业务变更之前。

#### 3.4 复制状态机

当单个状态机可理解后，下一个压力就是复制：

```text
client command
  -> replicated log / Aeron Cluster / Raft-style ordering
  -> node A apply(command)
  -> node B apply(command)
  -> node C apply(command)
```

项目不应该自己从零实现共识，而是定义交易核心向复制日志要求什么能力：命令
编码、session identity、日志位置、快照格式、重放起点、背压语义、故障切换
和事件发出边界。

#### 3.5 投影、缓存、推送和运行时

当事实已经有序且可恢复，系统还要服务大量读者：OMS、账本投影、对账、合规、
公共行情、私有成交回报、权限缓存、mark-price 输入、风险投影和报表。

这些组件不应该直接拉取核心热状态，而应该消费有版本的发布：

```text
snapshot(version=N) + deltas(version>N) -> local projection
```

运行时优化应该在语义稳定之后才进入。warmup、allocation、pooling、off-heap
buffer、CPU isolation、网络路径和 kernel bypass 都属于后期实验，不是第一
步。

### 4. 设计原则

#### 4.1 真相必须有形状

每个阶段都要说清楚当前真相源是什么。早期是数据库行，后期是命令日志、事件
日志、快照和复制日志位置。余额、仓位、top-of-book、风险敞口都是视图，必须
能从事实解释回来。

#### 4.2 事实先于视图

系统应优先保存不可变事实，再生成可变汇总。充值 accepted、提现 requested、
订单 accepted、成交 executed、仓位 updated、保证金 checked、provider record
ingested、bank record matched，这些事实比当前余额更适合审计和恢复。

#### 4.3 显式顺序优先于隐式竞争

锁在支付和账本工作流里完全合理。但在交易热路径里，锁竞争、重试和尾延迟会
让公平性和可解释性变差。

状态机需要更高基础设施成本：sequencing、log、snapshot、replay、backpressure
和 failover。换来的好处是业务问题变简单：

```text
如果 A 先于 B 执行，应该发生什么？
```

#### 4.4 不同排序机制是在不同地方付费

| 机制 | 在哪里为排序付费 | 优势 | 代价 |
| --- | --- | --- | --- |
| 悲观锁 / 2PL | 执行前或执行中 | 共享行正确性简单 | 锁竞争、死锁 |
| MVCC / CAS | 提交时 | 读并发好 | 冲突重试 |
| 单写者 / sequencer | 业务变更前 | 核心语义简单确定 | 分片和恢复设计 |
| Raft / Paxos / 复制日志 | 复制执行前 | 故障切换和一致性 | 运维和延迟成本 |

Serializable 给出等价于某条串行历史的效果；linearizability 加上真实时间约束；
strict serializability 同时包含两者。项目使用这些概念来比较工程模型，而不是
为了堆术语。

#### 4.5 冷热路径分离

热路径是命令准入、有序变更和事实发出。温路径是事实加市场数据生成仓位、敞
口和风险动作。冷路径是报表、对账、审计和合规导出。

冷路径不是不重要，只是不能阻塞热路径。

#### 4.6 状态归属与数据重力

热状态被很多组件反复拉取、join、解释时会变贵。设计应该问：

```text
是不是应该让函数靠近状态，而不是让状态到处移动？
```

每个热状态都要有 owner：谁能修改它，如何发布更新，消费者如何发现 gap，重
启后如何 rebuild，允许多旧的数据。

#### 4.7 发布优先于拉取

参考数据、账户权限、风险配置、成交事实和市场状态不应该在热路径里反复 pull。
owner 应该发布带版本的 snapshot 和 delta。消费者维护本地投影，并明确知道：
fail closed、允许有限 stale、从快照加 delta 重建，还是在 gap 回放前拒绝服务。

#### 4.8 分发不等于持久可靠

快速分发不等于可靠真相。

比如 UDP multicast 是发送者向一个组播地址发一次，多个订阅者接收。它很适合
低延迟 fan-out，但传输层不会替你记住晚加入者、丢包、慢消费者和重启恢复。

可靠流还需要 sequence number、gap detection、snapshot、replay、backpressure、
以及一个 durable 或 replicated 的事实源。

公共行情有时可以容忍有限丢失，因为客户端能 resync。私有成交回报、账本事实
和风控决策不能当成 best-effort notification。

#### 4.9 性能是系统形状的结果

低延迟不应该是一堆技巧，而应该来自系统形状：热状态有 owner；命令先排序再
变更；消费者维护本地投影；持久事实写在明确边界；恢复语义清楚；allocation
和 warmup 都有测量。

只有这些成立后，pooling、off-heap buffer、zero allocation、CPU pinning、
NUMA、kernel bypass 和 runtime tuning 才有意义。

### 5. 领域边界

| 边界 | 拥有什么 | 发出什么 | 不应该做什么 |
| --- | --- | --- | --- |
| Ledger | 价值移动和双分录事实 | journal entries、balance facts | 隐藏未匹配外部资金 |
| Wallet | 充值提现生命周期 | provider callbacks、withdrawal state、raw records | 重复 apply 外部事件 |
| Reconciliation | 外部原始记录和匹配报告 | exceptions、match results、adjustment proposals | 静默修改 ledger |
| Matching | 价格时间优先和成交 | trades、fills、order status | 状态转移中依赖远程调用 |
| Position | 账户和品种敞口 | position updates、PnL 输入 | 把 open order 当 execution |
| Margin | 权益和保证金需求 | margin snapshots、准入输入 | 把公式藏在 handler 里 |
| Pre-trade risk | 下单准入 | accept/reject facts | 变成策略引擎 |
| Risk cluster | 持续风险投影 | alerts、risk actions、projection state | 把 eventual projection 当同步真相 |
| OMS / compliance | 工作流、报告、审计 | reports、case state、监管导出 | 阻塞核心变更 |
| Cache / market state | owned facts 的本地投影 | snapshot/delta versions、gap status | 静默提供无限 stale 数据 |
| Push gateway | 公共和私有流发布 | sequenced snapshots and deltas | 把 best-effort 分发说成真相 |

### 6. 交易台扩展

交易所核心决定场内什么是真的；交易台根据市场状态、库存、风险和外部场所决
定要做什么。

交易台应该在核心已经能产生可靠行情、成交回报、仓位、风险视图和对账事实之
后再出现。

自然演进是：

```text
matching
  -> executions
  -> positions
  -> exposure
  -> pre-trade and continuous risk
  -> market data and mark inputs
  -> external venue routing
  -> hedging / best execution / strategy
```

边界规则：不要让交易台污染交易所核心。撮合引擎不应该知道套利；账本不应该
知道 best execution；pre-trade risk 不应该变成策略引擎。

### 7. 故障与恢复模型

| 阶段 | 典型故障 | 需要的恢复叙事 |
| --- | --- | --- |
| DB 真相源 | 事务 abort、重复 callback、流程半失败 | 原子回滚、幂等 key、外部原始记录 |
| Outbox / command log | 状态提交和消息发布不一致 | transactional outbox、幂等消费者、重发 |
| 内存状态机 | sequence N 后进程崩溃 | snapshot + durable log replay |
| 复制状态机 | leader 崩溃、follower 落后、client 重试 | log position、session identity、replay、failover |
| Cache / projection | gap、stale snapshot、consumer 重启 | versioned snapshot、delta replay、fail-open/closed policy |
| Push stream | 丢包、晚加入、慢 client | sequence、snapshot、replay、backpressure |
| Reconciliation | provider、bank、custody、internal records 不一致 | immutable raw records、normalized records、match reports、manual adjustment journal |
| Runtime | warmup 抖动、GC/allocation spike、网络噪声 | benchmark、profile、warmup discipline、allocation control、runbook |

### 8. 章节地图

| 阶段 | 章节 | 目的 |
| --- | --- | --- |
| 版本线 | 01-06 | 交易所契约、ACID SQL、SQL facts/outbox、内存核心、复制日志、SQL projections |
| 领域深挖 | 07-12 | 订单簿机制、仓位/PnL、保证金/风控、风险投影、缓存一致性、推送 |
| 运行时实验 | 13-14 | Rust 热路径实验和低延迟运行时/网络测量 |
| 交易台扩展 | 15-19 | 外部行情、定价、路由、对冲、最优执行、简单策略 |
| 附录原型 | 90-93 | 早期可运行 Go 资金原型 |

每章都应该回答：

1. 哪个正确性属性必须保持稳定？
2. 引入了什么性能压力？
3. 这个模型暴露了什么可靠性或恢复故障？
4. 哪个架构视角发生变化：真相、排序、恢复或归属？
5. 下一个模型引入了什么运维成本？

### 9. 非目标

这个项目不是生产交易所、完整撮合库、共识实现、真实钱包、真实策略或 benchmark
比赛。它的目标是用刚好足够的代码和测试，让每个架构压力都变得可见。

### 10. 成功标准

如果读者能讲清楚下面这些事，这个项目就成功了：

- 为什么 DB ACID 是合理起点；
- 为什么有序事实比当前行更适合解释和恢复；
- 为什么热路径最终需要确定性变更边界；
- 为什么复制需要 log、snapshot、replay 和 backpressure；
- 为什么对账是记录匹配和证据链，不只是算术；
- 为什么缓存需要 owner、version 和 rebuild policy；
- 为什么私有成交回报不能是 best-effort notification；
- 为什么低延迟工作应该从 ownership 和 measurement 开始，而不是从技巧开始。

最强证明是跨版本测试：

```text
same business scenario
same observable semantics
different execution substrate
```

这就是整个项目的核心。
