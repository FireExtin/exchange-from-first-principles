# Integration Test Contract

This directory is reserved for cross-version scenarios.

The goal is to prove that architecture can change while business semantics stay
stable. A test scenario should be written once, then run against every runnable
version through a small adapter.

Current runnable shape:

```bash
go test ./integration-tests/...
```

The first suite compares the chapter 03 wallet workflow and the chapter 04
command-log replay engine through the shared funds contract.

Over time, adapters should expose the same behavioral operations:

- create account;
- deposit;
- withdraw;
- place order;
- cancel order;
- apply execution;
- query balance;
- query position;
- query emitted facts.

The tests should assert externally visible semantics, not internal data
structures. If the DB version, single-writer version, and replicated-log version
all pass the same scenario suite, the repo has hard evidence that the semantic
contract survived the architecture migration.

---

## 中文

### 集成测试契约

本目录保留用于跨版本场景。

目标是证明架构可以变化而业务语义保持稳定。测试场景应写一次，然后
通过一个薄适配器针对每个可运行版本运行。

当前可运行形态：

```bash
go test ./integration-tests/...
```

第一套测试通过共享资金契约，对比第 03 章的钱包工作流和第 04 章的
命令日志重放引擎。

长期来看，适配器应暴露相同的行为操作：

- 创建账户；
- 入金；
- 出金；
- 下单；
- 撤单；
- 应用成交；
- 查询余额；
- 查询仓位；
- 查询发出的事实。

测试应断言外部可见语义，而不是内部数据结构。如果 DB 版本、单写者版本
和复制日志版本都通过相同的场景套件，仓库就有了语义契约在架构迁移中
存活的有力证据。
