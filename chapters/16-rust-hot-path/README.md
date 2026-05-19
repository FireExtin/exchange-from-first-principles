# 16 Rust Hot Path

[English](#english) · [中文](#中文)

## English

Purpose: keep the existing Rust experiment as a runtime/hot-path experiment,
not as the canonical exchange implementation.

The semantic contract still comes from `shared/go/exchange` and the main
chapters. Rust may later be used to test command-log parsing, replay, or
latency-sensitive components.

## Status And Run

Status: runnable Rust experiment. No exchange business implementation should be
added in this pass.

Run:

```bash
cargo test
```

## 中文

目的：保留现有 Rust 实验作为 runtime/hot-path experiment，而不是规范交易所实现。

语义契约仍来自 `shared/go/exchange` 和主章节。Rust 后续可以用于测试
command-log parsing、replay 或 latency-sensitive components。

## 状态与运行

状态：可运行 Rust 实验。本轮不应加入交易所业务实现。

运行：

```bash
cargo test
```
