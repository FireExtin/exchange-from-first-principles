# Version Contract And Testing

This project should become stronger as its architecture changes. The way to
prove that is to keep the business contract stable.

## Stable Business Contract

The concrete names may differ by language, but every runnable version should
preserve this shape:

```go
type StateMachine interface {
    Apply(event Event) error
    Snapshot() ([]byte, error)
    Restore(snapshot []byte) error
}
```

Or, for command-oriented chapters:

```go
func applyOrder(state *State, order Order) (*State, error) {
    // Same business rule, different ordering substrate.
}
```

The important invariant is not the exact method name. The invariant is:

```text
ordered input + deterministic transition -> state change + emitted facts
```

If the storage engine changes from SQL rows to an in-memory state machine, the
business rule should not silently change. If replication is added, the state
transition should still mean the same thing.

The first concrete contract in this repository is the Go funds contract in
`shared/go`: command in, facts out, then the same observable balances and
withdrawal states across implementations.

## README Comparison Format

Every implemented version or major chapter should eventually include:

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

The most powerful line in `Semantic Change` should often be:

```text
None. The same external scenario suite passes unchanged.
```

That sentence is only credible if the tests make it true.

## Same Tests Across Versions

The current Go suite already has the first thin-adapter shape:

```bash
go test ./integration-tests/...
```

It currently compares the chapter 03 wallet workflow with the chapter 04
command-log replay engine through the same shared funds contract. Later
versions can add DB, single-writer, and replicated-state-machine adapters
without changing the scenario assertions.

The user-facing examples can also be presented as separate directories:

```bash
cd exchange-v0-db && go test ../integration-tests/...
cd exchange-v1-single-writer && go test ../integration-tests/...
cd exchange-v2-raft && go test ../integration-tests/...
```

This repo currently uses `chapters/` rather than separate version directories,
but the contract is the same: a scenario should assert behavior rather than
internal structure.

## Scenarios To Preserve

Already covered by the current Go conformance suite:

- duplicate deposits are idempotent;
- withdrawals cannot spend unavailable funds;
- duplicate withdrawal confirmations are idempotent;
- transfers emit replayable facts and move balances.

Next scenarios to add as the project grows:

- deposits conserve value across ledger entries;
- spot trade settlement is atomic and emits standard funding facts;
- replay produces the same balance state as the original run;
- price-time order determines fills;
- execution reports update positions deterministically;
- margin rejects orders that exceed available equity;
- risk projection rebuilds from the event stream;
- push clients can recover from snapshot plus deltas.

## DB ACID Versus Raft

The DB version is best when:

- concurrency is modest;
- storage latency is acceptable;
- the database can be the mutation boundary;
- the team needs operational simplicity.

The replicated state-machine version is best when:

- hot-path ordering must be explicit;
- replay and failover are core requirements;
- database locks and retries dominate behavior;
- the team can afford the operational cost of a replicated log.

They are not equivalent as implementations. They can be equivalent as externally
observable business semantics when the same successful commands appear in the
same valid order and run the same deterministic transition logic.

---

## 中文

### 版本契约与测试

本项目应该在架构变化中变得更强。证明这一点的方法是保持业务契约稳定。

### 稳定的业务契约

具体名称可能因语言而异，但每个可运行版本应保留这个形状：

```go
type StateMachine interface {
    Apply(event Event) error
    Snapshot() ([]byte, error)
    Restore(snapshot []byte) error
}
```

或者对于面向命令的章节：

```go
func applyOrder(state *State, order Order) (*State, error) {
    // Same business rule, different ordering substrate.
}
```

重要的不变量不是确切的方法名。不变量是：

```text
ordered input + deterministic transition -> state change + emitted facts
```

如果存储引擎从 SQL 行变为内存状态机，业务规则不应静默改变。
如果添加了复制，状态转换仍应意味着相同的事情。

本仓库的第一份具体契约是 `shared/go` 中的 Go 资金契约：输入命令、输出事实，
并在不同实现之间保持相同的可观察余额和出金状态。

### README 比较格式

每个已实现的版本或主要章节最终应包含：

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

`Semantic Change` 中最有力的行通常应该是：

```text
None. The same external scenario suite passes unchanged.
```

只有测试使其为真，这句话才可信。

### 跨版本相同测试

当前 Go 套件已经具备第一版薄适配器形态：

```bash
go test ./integration-tests/...
```

它当前通过同一份共享资金契约，对比第 03 章的钱包工作流和第 04 章的
命令日志重放引擎。后续可以继续加入 DB、单写者、复制状态机适配器，
而不改变场景断言。

用户面向的示例也可以呈现为单独的目录：

```bash
cd exchange-v0-db && go test ../integration-tests/...
cd exchange-v1-single-writer && go test ../integration-tests/...
cd exchange-v2-raft && go test ../integration-tests/...
```

本仓库当前使用 `chapters/` 而不是单独的版本目录，但契约是相同的：
场景应断言行为而不是内部结构。

### 要保留的场景

当前 Go 一致性测试已经覆盖：

- 重复入金是幂等的；
- 出金不能花费不可用资金；
- 重复出金确认是幂等的；
- 转账会产生可重放事实并移动余额。

项目继续演进时需要补充：

- 入金在账务分录上守恒；
- 现货交易结算是原子的；
- 重放产生与原始运行相同的余额状态；
- 价格-时间顺序决定成交；
- 成交报告确定性地更新仓位；
- 保证金拒绝超过可用权益的订单；
- 风控投影从事件流重建；
- 推送客户端可以从快照加增量恢复。

### DB ACID 与 Raft

DB 版本适用于：

- 并发适中；
- 存储延迟可接受；
- 数据库可以是变更边界；
- 团队需要运营简单性。

复制状态机版本适用于：

- 热路径排序必须是显式的；
- 重放和故障转移是核心需求；
- 数据库锁和重试主导行为；
- 团队能承受复制日志的运营成本。

作为实现它们不等价。当相同的成功命令以相同有效顺序出现并运行相同的
确定性转换逻辑时，作为外部可观察业务语义它们可以等价。
