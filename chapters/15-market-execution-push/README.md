# 15 Market And Execution Push

[English](#english) · [中文](#中文)

## English

> Real-time prices and trade confirmations are pushed to clients. What happens when a client misses a message? What about a client that reconnects? This chapter upgrades push from best-effort delivery to a recoverable stream with sequence numbers and replay.

Purpose: specify recoverable public market-data and private execution-report
streams.

Push is not best-effort notification. It is a publication surface over ordered
facts with sequence numbers, snapshots, deltas, replay, and backpressure.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- public book/trade stream;
- private execution-report stream;
- sequence numbers;
- snapshot plus delta recovery;
- gap detection;
- backpressure and replay boundary.

## 中文

> 实时行情和成交回报推给客户端。客户端漏了一条怎么办？断线重连后怎么补？这章把推送从"尽力而为"升级成有顺序号、可重放的可靠流。

目的：定义可恢复的 public market-data 和 private execution-report streams。

Push 不是 best-effort notification。它是基于有序事实的发布表面，需要 sequence
number、snapshot、delta、replay 和 backpressure。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- public book/trade stream；
- private execution-report stream；
- sequence numbers；
- snapshot 加 delta recovery；
- gap detection；
- backpressure 和 replay boundary。
