# 12 Margin And Pre-Trade Risk

[English](#english) · [中文](#中文)

## English

Purpose: combine margin checks and synchronous pre-trade risk admission.

This chapter explains how dangerous orders are rejected before they enter the
sequenced core, and how reservation/margin semantics stay consistent across the
architecture migration.

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

## 中文

目的：合并保证金检查和同步下单前风控准入。

本章解释危险订单如何在进入排序核心前被拒绝，以及 reservation/margin 语义如何
在架构迁移中保持一致。

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
