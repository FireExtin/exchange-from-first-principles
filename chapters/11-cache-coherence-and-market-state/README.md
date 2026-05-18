# 13 Cache Coherence And Market State

[English](#english) · [中文](#中文)

## English

Purpose: model the caches that make exchange systems fast while preserving
clear failure semantics.

## Status

Status: design scaffold. No runnable implementation exists here yet.

This chapter is not about adding Redis everywhere. It is about deciding which
state may be cached, how stale it may be, and what the system does when the
cache is missing, old, or contradictory.

First scope:

- user identity and permission cache;
- account status cache: enabled, frozen, reduce-only, killed;
- instrument and trading-rule cache;
- external exchange metadata cache;
- mark-price composition cache;
- risk configuration cache;
- cache versioning, invalidation, TTL, and rebuild.

Key rule:

```text
every cache must name its source of truth, freshness budget, and fail policy
```

Ownership rule:

```text
owner publishes snapshot + versioned deltas; consumers build local projections
```

Some caches must fail closed. Account permission and kill-switch state cannot be
treated as advisory. Some caches may fail stale for a short window, such as
public instrument metadata. Mark-price composition sits in the middle: it can be
derived from multiple feeds, but stale marks directly affect risk and margin.

This is the chapter where data gravity becomes concrete. A hot path should not
ask a remote account service or product service on every order. Instead, each
state owner publishes changes, and consumers keep local projections with
sequence checks, gap detection, and a rebuild path.

This chapter should produce a small table for each cache:

| Cache | Source of truth | Freshness budget | Fail policy | Rebuild path |
| --- | --- | --- | --- | --- |
| Account status | account service / event log | strict | fail closed | replay account events |
| User permissions | auth service / config event | strict | fail closed | refresh snapshot + deltas |
| Mark price | index feeds + formula config | bounded | fail safe | recompute from latest valid feeds |
| Instrument rules | product config | bounded | reject unknown | reload config snapshot |

The implementation should come later. The first deliverable is vocabulary and
test scenarios for stale, missing, and conflicting cache values:

- stale account status must fail closed;
- missing instrument rules must reject unknown orders;
- a mark-price composition gap must stop dependent risk checks or fall back to
  a documented safe policy;
- a restarted consumer must rebuild from snapshot plus deltas.

---

## 中文

目的：建模使交易所系统快速但保持清晰失败语义的缓存。

## 状态

状态：设计脚手架。本目录尚无可运行实现。

本章不是到处加 Redis。它是决定哪些状态可以缓存、可以有多旧，以及当
缓存缺失、过旧或矛盾时系统做什么。

首个范围：

- 用户身份和权限缓存；
- 账户状态缓存：启用、冻结、只减、kill；
- 合约和交易规则缓存；
- 外部交易所元数据缓存；
- 标记价格组成缓存；
- 风控配置缓存；
- 缓存版本控制、失效、TTL 和重建。

关键规则：

```text
每个缓存必须命名其真相源、新鲜度预算和失败策略
```

归属规则：

```text
owner 发布 snapshot + versioned deltas；consumer 构建本地投影
```

某些缓存必须故障关闭。账户权限和 kill 开关状态不能被视为咨询性的。
某些缓存可以在短时间内故障为过期，如公共合约元数据。标记价格组成处于中间：
它可以从多个来源派生，但过期标记直接影响风险和保证金。

data gravity 在这一章变得具体。热路径不应该每个订单都远程查询账户服务或产品
服务。更好的方式是状态 owner 发布变化，consumer 用序列检查、缺口检测和重建
路径维护本地投影。

本章应为每个缓存生成一个小表：

| Cache | Source of truth | Freshness budget | Fail policy | Rebuild path |
| --- | --- | --- | --- | --- |
| Account status | account service / event log | strict | fail closed | replay account events |
| User permissions | auth service / config event | strict | fail closed | refresh snapshot + deltas |
| Mark price | index feeds + formula config | bounded | fail safe | recompute from latest valid feeds |
| Instrument rules | product config | bounded | reject unknown | reload config snapshot |

实现应该稍后进行。第一个交付物是过期、缺失和矛盾缓存值的词汇和测试场景。

- 账户状态过期必须故障关闭；
- 缺失交易规则必须拒绝未知订单；
- 标记价格组成出现缺口时，应停止依赖它的风控检查或进入明确的安全策略；
- 重启后的 consumer 必须能从快照加增量重建。
