# 91 Spot Settlement Transaction Prototype

[English](#english) · [中文](#中文)

## English

Purpose: preserve the original runnable spot-settlement transaction prototype.

This appendix prototype uses an in-memory transaction boundary to show why spot
settlement needs atomicity. It is DB-shaped, but it is not the new ACID SQL
exchange chapter.

## Status And Run

Status: runnable Go prototype.

Run:

```bash
go test ./...
```

## 中文

目的：保留原来的可运行现货结算事务原型。

这个 appendix prototype 使用内存事务边界展示为什么现货结算需要原子性。它具有
DB-shaped 结构，但不是新的 ACID SQL exchange 章节。

## 状态与运行

状态：可运行 Go 原型。

运行：

```bash
go test ./...
```
