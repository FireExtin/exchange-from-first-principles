# 17 Low-Latency Runtime And Networking

[English](#english) · [中文](#中文)

## English

Purpose: measure runtime and networking costs after semantic correctness is
defined.

This chapter is about variance, allocation, warmup, buffers, and networking
paths. It should not introduce new business semantics.

## Status

Status: README only. No runnable implementation exists here yet.

## First Scope

- warmup and JIT behavior;
- allocation and pooling;
- off-heap/buffer paths;
- network serialization and batching;
- p99/p999 variance measurements.

## 中文

目的：在语义正确性已经定义后，测量 runtime 和 networking 成本。

本章关注 variance、allocation、warmup、buffer 和 networking path，不应引入新的
业务语义。

## 状态

状态：仅 README。本章尚无可运行实现。

## 第一范围

- warmup 和 JIT 行为；
- allocation 和 pooling；
- off-heap/buffer paths；
- 网络序列化和 batching；
- p99/p999 variance measurements。
