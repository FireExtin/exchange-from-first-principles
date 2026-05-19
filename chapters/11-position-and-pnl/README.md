# 11 Position And PnL

[English](#english) · [中文](#中文)

## English

> You've made five BTC trades this week. What is your net position? What is your average entry price? How much have you gained or lost? This chapter turns execution facts into position numbers.

Purpose: explain how execution facts become positions and PnL placeholders.

Position semantics are part of the cross-version proof: SQL, memory,
replicated log, and projections should interpret the same executions
consistently.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- net position by account and instrument;
- average-entry placeholder;
- realized PnL placeholder;
- unrealized PnL placeholder;
- cash and locked-fund hooks;
- deterministic application of execution reports.

## 中文

> 你这周交易了五次 BTC。净仓位是多少？均价多少？盈亏多少？这章把成交事实转化成仓位数字。

目的：解释 execution facts 如何变成仓位和 PnL 占位。

仓位语义是跨版本证明的一部分：SQL、内存、复制日志和 projection 应该一致解释
同一组成交。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- 按账户和 instrument 计算净仓位；
- 平均开仓价占位；
- 已实现 PnL 占位；
- 未实现 PnL 占位；
- cash 和 locked-fund hook；
- 确定性应用 execution report。
