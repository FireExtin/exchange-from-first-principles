# Shared Contracts

This directory contains examples and schemas shared across chapters.

The files here are not a framework. They are a small contract layer so each
chapter can talk about the same facts:

- command identity;
- event sequence;
- account and asset identity;
- amount movement;
- replay and reconciliation inputs.

The first concrete Go contract lives in `shared/go`. It defines a minimal funds
language used by chapters 03 and 04:

- deposit;
- withdrawal request and confirmation;
- transfer;
- externally visible success and rejection events.

Rules:

- Shared contracts may be consumed by any chapter.
- Shared contracts must not import code from a chapter.
- Chapter implementations may diverge internally as long as their public events
  can be explained through these contracts.

---

## 中文

### 共享契约

本目录包含章节间共享的示例和 schema。

这里的文件不是框架。它们是一个小的契约层，以便每个章节谈论相同的事实：

- 命令标识；
- 事件序列；
- 账户和资产标识；
- 金额移动；
- 重放和对账输入。

第一份可运行 Go 契约位于 `shared/go`。它定义了一个最小资金语义语言，
供第 03 和第 04 章共用：

- 入金；
- 出金请求与确认；
- 转账；
- 外部可见的成功和拒绝事件。

规则：

- 共享契约可被任何章节消费。
- 共享契约不得从章节导入代码。
- 章节实现可以内部发散，只要其公开事件可以通过这些契约解释。
