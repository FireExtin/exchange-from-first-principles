# 13 Cache Coherence And Market State

Purpose: model the caches that make exchange systems fast while preserving
clear failure semantics.

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

Some caches must fail closed. Account permission and kill-switch state cannot be
treated as advisory. Some caches may fail stale for a short window, such as
public instrument metadata. Mark-price composition sits in the middle: it can be
derived from multiple feeds, but stale marks directly affect risk and margin.

This chapter should produce a small table for each cache:

| Cache | Source of truth | Freshness budget | Fail policy | Rebuild path |
| --- | --- | --- | --- | --- |
| Account status | account service / event log | strict | fail closed | replay account events |
| User permissions | auth service / config event | strict | fail closed | refresh snapshot + deltas |
| Mark price | index feeds + formula config | bounded | fail safe | recompute from latest valid feeds |
| Instrument rules | product config | bounded | reject unknown | reload config snapshot |

The implementation should come later. The first deliverable is vocabulary and
test scenarios for stale, missing, and conflicting cache values.

---

## 中文

目的：建模使交易所系统快速但保持清晰失败语义的缓存。

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

某些缓存必须故障关闭。账户权限和 kill 开关状态不能被视为咨询性的。
某些缓存可以在短时间内故障为过期，如公共合约元数据。标记价格组成处于中间：
它可以从多个来源派生，但过期标记直接影响风险和保证金。

本章应为每个缓存生成一个小表：

| Cache | Source of truth | Freshness budget | Fail policy | Rebuild path |
| --- | --- | --- | --- | --- |
| Account status | account service / event log | strict | fail closed | replay account events |
| User permissions | auth service / config event | strict | fail closed | refresh snapshot + deltas |
| Mark price | index feeds + formula config | bounded | fail safe | recompute from latest valid feeds |
| Instrument rules | product config | bounded | reject unknown | reload config snapshot |

实现应该稍后进行。第一个交付物是过期、缺失和矛盾缓存值的词汇和测试场景。