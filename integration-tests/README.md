# Integration Test Contract

[English](#english) · [中文](#中文)

## English

This directory is reserved for cross-version scenarios.

The goal is to prove that architecture can change while exchange semantics stay
stable. A scenario should be written once, then run against every runnable
version through a small adapter.
Future exchange scenarios should be composed from the first four teaching
steps: custody/liability, balance states, order reservation, and
match/settlement.

## Current Runnable Shape

```bash
go test ./integration-tests/...
```

The current suite still compares appendix funds prototypes 92 and 93 through
the earlier `shared/go/funds` contract.

## Target Exchange Shape

Future adapters should expose the `shared/go/exchange` behavior surface:

- submit commands;
- query balances, orders, and positions;
- inspect emitted facts;
- snapshot and restore where relevant;
- verify projection cursors and rebuild results.

Tests should assert externally visible semantics, not internal storage shape.
If the SQL version, memory version, replicated-log version, and projection
version pass the same scenario suite, the repo has hard evidence that the
semantic contract survived the architecture migration.

Unimplemented exchange scenarios should remain behind the
`exchange_contract_todo` build tag until adapters exist.

---

## 中文

本目录保留用于跨版本场景。

目标是证明架构可以变化，而交易所语义保持稳定。场景应写一次，然后通过薄 adapter
针对每个可运行版本运行。
未来 exchange 场景应由前四个教学步骤组合而来：custody/liability、余额状态、
订单冻结、撮合/结算。

## 当前可运行形态

```bash
go test ./integration-tests/...
```

当前套件仍然通过早期 `shared/go/funds` 契约，对比附录资金原型 92 和 93。

## 目标交易所形态

未来 adapter 应暴露 `shared/go/exchange` 行为表面：

- 提交 commands；
- 查询 balances、orders 和 positions；
- 检查 emitted facts；
- 在相关版本中 snapshot 和 restore；
- 验证 projection cursors 和 rebuild results。

测试应断言外部可见语义，而不是内部存储形状。如果 SQL 版本、内存版本、
复制日志版本和 projection 版本都通过同一套场景，仓库就有了语义契约在架构迁移中
存活的有力证据。

未实现的 exchange 场景应继续放在 `exchange_contract_todo` build tag 后，直到
adapter 出现。
