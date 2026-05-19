# 14 Cache Coherence And Market State

[English](#english) · [中文](#中文)

## English

> Permissions, mark prices, and instrument rules are cached everywhere. How stale is too stale? What happens when a cache falls behind? This chapter gives caches explicit rules: declare your freshness guarantee, and declare what to do when you can't meet it.

Purpose: explain how reference data, permissions, marks, and market state are
cached without breaking the semantic contract.

Caches are consumer views. They must declare freshness, fail policy, rebuild
path, and cursor/gap behavior.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- account and permission cache;
- instrument and reference-data cache;
- mark-price and market-state cache;
- snapshot plus versioned deltas;
- gap detection and rebuild;
- fail-open versus fail-closed policy.

## 中文

> 权限、mark price、instrument 规则——到处都在缓存。缓存多旧算太旧？落后了会怎样？这章给缓存立规矩：声明你的新鲜度保证，声明达不到时怎么办。

目的：解释 reference data、permission、mark 和 market state 如何缓存，同时不破坏
语义契约。

缓存是 consumer view。它们必须声明新鲜度、失败策略、重建路径，以及 cursor/gap
行为。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- account 和 permission cache；
- instrument 和 reference-data cache；
- mark-price 和 market-state cache；
- snapshot 加 versioned deltas；
- gap detection 和 rebuild；
- fail-open 与 fail-closed 策略。
