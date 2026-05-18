# 02 ACID SQL Exchange

[English](#english) · [中文](#中文)

## English

Purpose: specify v0, where ACID SQL is the first source of truth for the full
exchange semantic contract.

This chapter should eventually show one transaction boundary protecting
double-entry ledger postings, account reservation, orders, trades, positions,
and risk-admission state.

## Status

Status: contract scaffold. No SQL schema, adapter, or business implementation
exists here yet.

The point is to start from the model readers already trust: database rows plus
transactions. The implementation slot is intentionally blank.

## Required Semantics

- every posted transaction balances per asset;
- user balances are platform liabilities;
- custody accounts are platform assets;
- order placement reserves available funds into locked accounts;
- matching emits execution facts and ledger postings;
- positions and risk state are updated inside the same semantic boundary.

## 中文

目的：定义 v0，也就是用 ACID SQL 作为完整交易所语义的第一真相源。

本章未来应展示一个事务边界如何同时保护 double-entry ledger postings、账户冻结、
订单、成交、仓位和风控准入状态。

## 状态

状态：契约脚手架。本章尚无 SQL schema、adapter 或业务实现。

重点是从读者已经熟悉并信任的模型开始：数据库行加事务。实现位置刻意留白。

## 必须保持的语义

- 每笔 posted transaction 在每种资产内分别平衡；
- 用户余额是平台负债；
- custody 账户是平台资产；
- 下单将可用资金冻结到 locked 账户；
- 撮合产生成交事实和账本分录；
- 仓位和风控状态在同一语义边界内更新。
