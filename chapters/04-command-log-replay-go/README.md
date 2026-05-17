# 04 Command Log Replay

[English](#english) · [中文](#中文)

## English

Primitive: state changes are written as commands and events.

Question: how do we recover after a crash and explain what happened?

This chapter introduces an append-only command log, strict sequence checks, and
deterministic replay into account state.

It also contains a small equivalence test: a serializable DB transaction model
and a sequenced state machine apply the same ordered commands and reach the
same balances. This is the bridge from "DB rows are truth" to "ordered facts are
truth".

The chapter also exposes a funds replay engine through `adapter`, using the same
shared command and event contract as chapter 03.

## Status And Run

Status: runnable Go chapter.

Run:

```bash
go test ./...
```

---

## 中文

原语：状态变更作为命令和事件写入。

问题：崩溃后如何恢复并解释发生了什么？

本章在账户状态中引入追加写命令日志、严格序列检查和确定性重放。

它还包含一个小等价测试：可串行化 DB 事务模型和排序状态机应用相同的排序命令
并达到相同的余额。这是"DB 行是真相"到"有序事实是真相"的桥梁。

本章也通过 `adapter` 暴露了一个资金重放引擎，使用与第 03 章相同的共享
命令和事件契约。

## 状态与运行

状态：可运行 Go 章节。

运行：

```bash
go test ./...
```
