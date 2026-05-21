# 12 Margin And Pre-Trade Risk

[English](#english) · [中文](#中文)

## English

> Before an order enters the book, the system must decide: can this account afford it? Is it too risky to allow? This chapter defines the synchronous gate orders must pass before matching.

Purpose: combine margin checks and synchronous pre-trade risk admission.

This chapter explains how dangerous orders are rejected before they enter the
sequenced core, and how reservation/margin semantics stay consistent across the
architecture migration.
Risk admission can reject a command or require a reservation. It must not
silently post financial truth. If fees, liquidation, funding, or PnL need
ledger impact, that happens through explicit settlement/accrual events.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- account enabled/disabled;
- kill switch;
- available and locked balances;
- initial and maintenance margin placeholders;
- max order notional;
- price band;
- margin accepted/rejected events.

## Boundary

- risk checks read balances, reservations, marks, and limits;
- accepted/rejected decisions are facts;
- risk state can block or alert, but it does not mutate posted balances without
  an explicit ledger event.

## 中文

> 订单进入订单簿之前，系统必须判断：这个账户承受得起吗？这笔单风险太大该拒绝吗？这章定义订单在撮合前必须通过的同步检查门。

目的：合并保证金检查和同步下单前风控准入。

本章解释危险订单如何在进入排序核心前被拒绝，以及 reservation/margin 语义如何
在架构迁移中保持一致。
风控准入可以拒绝命令，也可以要求冻结。它不能偷偷提交财务真相。如果手续费、
强平、funding 或 PnL 需要影响 ledger，必须通过显式 settlement/accrual 事件完成。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- 账户启用/禁用；
- kill switch；
- 可用和冻结余额；
- 初始保证金和维持保证金占位；
- 最大订单名义值；
- 价格区间；
- margin accepted/rejected events。

## 边界

- 风控检查读取 balances、reservations、marks 和 limits；
- accepted/rejected decisions 是事实；
- 风控状态可以阻断或报警，但不能在没有显式 ledger event 的情况下改写已过账余额。
