# 05 ACID SQL Exchange

[English](#english) · [中文](#中文)

## English

Status: contract scaffold. This chapter defines the first complete
architecture target, but it does not provide a SQL schema, adapter, or business
implementation.

Purpose: compose chapters 01-04 inside a single ACID transaction boundary.
Rows and transactions are the first model most readers already trust.

## User Action

A complete accepted business mutation may touch several surfaces at once:
ledger entries, balance projections, order state, reservation state, execution
facts, and position facts. The first architecture proof is simple:

```text
if the SQL transaction commits, all related facts commit together;
if it rolls back, none of the mutation becomes visible.
```

## Account Map

The SQL version should eventually persist the same conceptual surfaces that
chapters 01-04 introduced:

| Surface | Meaning |
| --- | --- |
| `accounts` | Account owner, asset, purpose, and normal side. |
| `journal_transactions` | One accepted business mutation or explicit reversal. |
| `journal_entries` | Debit/credit postings that balance per asset. |
| `orders` | Accepted, rejected, open, filled, and cancelled order states. |
| `reservations` | Locked assets backing open orders. |
| `executions` | Trade facts emitted by matching and consumed by settlement/projections. |
| `positions` | Position facts derived from executions. |

This is a conceptual map, not a schema requirement for this scaffold.

## Journal Template

```text
ACID transaction for an accepted match

1. Validate command and current rows.
2. Write execution facts.
3. Write balanced journal entries per asset.
4. Update order and reservation rows.
5. Update position/projection rows that belong in the same boundary.
6. Commit once, or roll everything back.
```

The SQL transaction does not change the business meaning of the entries from
chapters 01-04. It only gives them an atomic durability boundary.

## Contract Checks

- One committed transaction explains one accepted business mutation.
- Ledger entries balance per asset inside the committed transaction.
- Order/reservation state and ledger state cannot commit separately.
- Rows are queryable projections of committed facts.
- Later outbox, memory-core, replicated-log, and projection versions must keep
  the same external semantics.

## Out Of Scope

- No SQL schema is implemented in this scaffold.
- No adapter or posting logic is implemented here.
- SQL facts/outbox are introduced in chapter 06.
- Hot-path memory and replicated-log execution are introduced in chapters 07
  and 08.

## 中文

状态：契约脚手架。本章定义第一版完整架构目标，但不提供 SQL schema、adapter 或
业务实现。

目的：把第 01-04 章组合进一个 ACID 事务边界。数据库行和事务，是大多数读者已经
信任的第一个模型。

## 用户动作

一次完整的 accepted business mutation 可能同时触碰多个表面：ledger entries、
balance projections、order state、reservation state、execution facts 和 position
facts。第一版架构证明很简单：

```text
if the SQL transaction commits, all related facts commit together;
if it rolls back, none of the mutation becomes visible.
```

## 账户图谱

SQL 版本未来应该持久化第 01-04 章已经引入的同一组概念表面：

| 表面 | 含义 |
| --- | --- |
| `accounts` | 账户归属、资产、用途和正常方向。 |
| `journal_transactions` | 一次 accepted business mutation 或显式 reversal。 |
| `journal_entries` | 按资产分别平衡的 debit/credit postings。 |
| `orders` | accepted、rejected、open、filled、cancelled 等订单状态。 |
| `reservations` | 支撑 open orders 的 locked assets。 |
| `executions` | 撮合发出的成交事实，供结算和 projection 消费。 |
| `positions` | 从 executions 推导的仓位事实。 |

这是概念图谱，不是本脚手架要求实现的 schema。

## 分录模板

```text
ACID transaction for an accepted match

1. Validate command and current rows.
2. Write execution facts.
3. Write balanced journal entries per asset.
4. Update order and reservation rows.
5. Update position/projection rows that belong in the same boundary.
6. Commit once, or roll everything back.
```

SQL 事务不改变第 01-04 章 entries 的业务含义。它只给这些事实一个原子持久化边界。

## 契约检查

- 一次 committed transaction 解释一次 accepted business mutation。
- ledger entries 在 committed transaction 内按资产分别平衡。
- order/reservation state 和 ledger state 不能分开提交。
- rows 是 committed facts 的可查询 projection。
- 后续 outbox、memory-core、replicated-log 和 projection 版本必须保持同一外部语义。

## 本章不做什么

- 本脚手架不实现 SQL schema。
- 本章不实现 adapter 或 posting logic。
- SQL facts/outbox 在第 06 章引入。
- 热路径内存执行和复制日志执行在第 07、08 章引入。
