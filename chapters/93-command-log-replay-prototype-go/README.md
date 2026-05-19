# 93 Command Log Replay Prototype

[English](#english) · [中文](#中文)

## English

> The server restarted. How do you rebuild in-memory state? Answer: keep a log of every command in order, replay them from the beginning, and you get the same state back. This chapter proves that replay equals serial execution.

Purpose: define command-log replay semantics as a contract scaffold.

This appendix chapter shows how ordered commands and deterministic replay reach
the same state as a serial transaction model. It implements `funds.Engine` using
the decide/apply pattern:

```text
decide(command) -> event   (pure function, no state mutation)
apply(event)               (mutates state)
```

Separating decision from application means replay is safe: replay events, not
commands. Replaying events rebuilds state without re-running validation logic.

The `adapter/engine.go` wires `replay.FundsEngine` to the `funds.Engine`
interface so the funds conformance suite can prove this chapter produces the
same observable semantics as chapter 92's wallet-based engine.

Implementation is intentionally absent. Implement `internal/replay/` to make
the tests pass.

## Status And Run

Status: contract scaffold.

```bash
go test ./...
```

Tests compile and are expected to panic at TODO boundaries until implemented.

## 中文

> 服务器重启了。内存状态怎么重建？答案：把每条命令按顺序再执行一遍，得到的状态是一样的。这章证明重放等价于串行执行。

目的：把命令日志重放语义定义为契约脚手架。

这个附录章节展示有序命令和确定性重放如何达到与串行事务模型相同的状态。
它使用 decide/apply 模式实现 `funds.Engine`：

```text
decide(command) -> event   （纯函数，不修改状态）
apply(event)               （修改状态）
```

把决策和应用分离意味着重放是安全的：重放 event，不重放 command。重放 event
可以在不重新运行验证逻辑的情况下重建状态。

`adapter/engine.go` 把 `replay.FundsEngine` 接入 `funds.Engine` 接口，
使 funds conformance suite 能证明本章与第 92 章基于 wallet 的引擎产生相同的
可观测语义。

实现刻意留空。实现 `internal/replay/` 以通过测试。

## 状态与运行

状态：契约脚手架。

```bash
go test ./...
```

测试可以编译，在学习者实现之前预期会在 TODO 边界失败。
