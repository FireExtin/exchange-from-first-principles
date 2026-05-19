# 05 ACID SQL Exchange

[English](#english) · [中文](#中文)

## English

Purpose: combine chapters 01-04 into the first complete architecture version:
ACID SQL as the source of truth.

The SQL transaction boundary should eventually protect custody/user ledger,
balance states, order reservation, matching, settlement, positions, and risk
admission.

## Status

Status: contract scaffold. No SQL schema, adapter, or business implementation
exists here yet.

The point is to use the model readers already trust after the business
semantics are clear: rows plus transactions.

## Contract Focus

- one committed transaction explains one accepted business mutation;
- ledger postings balance per asset;
- current rows are queryable projections of committed facts;
- constraints protect accounting, order, and risk invariants;
- later versions must preserve the same external semantics.

## 中文

目的：把第 01-04 章组合成第一个完整架构版本：ACID SQL 作为真相源。

SQL 事务边界未来应保护 custody/user ledger、balance states、order reservation、
matching、settlement、positions 和 risk admission。

## 状态

状态：契约脚手架。本章尚无 SQL schema、adapter 或业务实现。

重点是在业务语义清楚之后，使用读者已经信任的模型：数据库行加事务。

## 契约重点

- 一次 committed transaction 解释一次 accepted business mutation；
- ledger postings 按资产分别平衡；
- 当前行是 committed facts 的可查询投影；
- constraints 保护 accounting、order 和 risk invariants；
- 后续版本必须保持同一外部语义。
