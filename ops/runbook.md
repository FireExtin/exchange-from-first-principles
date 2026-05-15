# Runbook

## Local Rust Core

```bash
cd chapters/15-rust-hot-path
cargo test
cargo run -p exchange-replay
cargo run -p exchange-bench
```

## Later Aeron Cluster

Start with one local node. Only move to three local nodes after replay and
snapshot semantics are stable.

---

## 中文

### 运行手册

## 本地 Rust 核心

```bash
cd chapters/15-rust-hot-path
cargo test
cargo run -p exchange-replay
cargo run -p exchange-bench
```

## 后续 Aeron 集群

从一个本地节点开始。只有在重放和快照语义稳定后才移到三个本地节点。