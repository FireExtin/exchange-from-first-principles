# 15 Rust Hot Path

[English](#english) · [中文](#中文)

## English

This chapter is a runnable experiment, not the active implementation track.

## Status And Run

Status: runnable Rust experiment.

The goal is to isolate what Rust might contribute to an exchange hot path:

- typed command and event envelopes;
- an append-only command log;
- a small state-machine trait;
- replay and benchmark entrypoints;
- a future FFI boundary.

The experiment should answer one question before it grows:

```text
Does Rust make the hot-path contract clearer or faster enough to justify the
extra language boundary?
```

For now, Java remains the primary trading hot-path implementation surface and
Go remains the service-edge surface. Rust stays useful as a measurement and
contract-design lab.

Run the workspace:

```bash
cargo test
```

From the repo root:

```bash
make test-rust
```

## Crates

| Crate | Role |
| --- | --- |
| `exchange-types` | Serializable command and event contracts. |
| `exchange-core` | State-machine trait and core error shape. |
| `exchange-log` | JSONL append-only command log. |
| `exchange-replay` | Replay entrypoint placeholder. |
| `exchange-bench` | Benchmark entrypoint placeholder. |
| `exchange-ffi` | Future foreign-function boundary placeholder. |

---

## 中文

本章是一个可运行实验，不是当前主实现路线。

## 状态与运行

状态：可运行 Rust 实验。

目标是隔离 Rust 可能给交易热路径带来的价值：

- 类型化命令和事件 envelope；
- 追加写命令日志；
- 小型状态机 trait；
- 重放和 benchmark 入口；
- 未来 FFI 边界。

实验在扩张前应先回答一个问题：

```text
Rust 是否能让热路径契约更清晰或更快，且收益足以抵消额外语言边界？
```

目前 Java 仍是交易热路径的主要实现层面，Go 仍负责服务边界。Rust 适合作为
测量和契约设计实验室。

运行 workspace：

```bash
cargo test
```

从仓库根目录运行：

```bash
make test-rust
```
