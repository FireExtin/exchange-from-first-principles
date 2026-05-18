# 01 Exchange Semantic Contract

[English](#english) · [中文](#中文)

## English

Purpose: define the exchange-level semantic contract before any storage or
runtime implementation.

This chapter owns the vocabulary for accounting, reservation, orders,
executions, positions, risk decisions, projection cursors, and adapter
boundaries. The concrete Go definitions live in `../../shared/go/exchange`.

## Status

Status: contract scaffold. No exchange engine implementation belongs here.

Agents may refine interfaces, fixtures, and contract tests. They must not fill
in posting, matching, risk, replay, projection, SQL, memory-core, or replicated
log implementation.

## Contract Surface

- accounting: account references, owner types, purposes, normal sides,
  debit/credit entries, journal transactions;
- funds: deposit, withdrawal, transfer, available, locked, pending withdrawal,
  fee revenue;
- orders: place, accept/reject, reserve, cancel, release, partial fill, full
  fill;
- executions: trade id, price, quantity, maker/taker, fees, execution reports;
- positions: position updates, average-entry placeholders, realized/unrealized
  PnL placeholders;
- risk: pre-trade decisions, margin decisions, kill switch;
- projection: snapshot cursor, event cursor, gap detection, rebuild result.

## 中文

目的：在任何存储或运行时实现之前，定义交易所级语义契约。

本章负责 accounting、reservation、orders、executions、positions、risk
decisions、projection cursors 和 adapter boundaries 的词汇。具体 Go 定义位于
`../../shared/go/exchange`。

## 状态

状态：契约脚手架。本章不实现交易所引擎。

Agents 可以细化接口、fixture 和契约测试，但不能补 posting、matching、risk、
replay、projection、SQL、内存核心或复制日志实现。

## 契约表面

- accounting：账户引用、owner 类型、purpose、normal side、debit/credit entry、
  journal transaction；
- funds：入金、出金、转账、可用、冻结、提现中、手续费收入；
- orders：下单、接受/拒绝、冻结、撤单、释放、部分成交、完全成交；
- executions：trade id、价格、数量、maker/taker、手续费、execution report；
- positions：仓位更新、平均开仓价占位、已实现/未实现 PnL 占位；
- risk：下单前风控、保证金决策、kill switch；
- projection：snapshot cursor、event cursor、gap detection、rebuild result。
