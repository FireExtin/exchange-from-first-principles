# 02 Spot Trade DB Transaction

Primitive: one buyer and one seller settle a spot trade.

Question: why does a trade need an atomic boundary?

This chapter uses an in-memory transaction boundary that snapshots state before
settlement and commits only when all account checks pass. It is intentionally
small; a future iteration can replace the in-memory store with SQL and the same
test cases should still describe the behavior.

Run:

```bash
go test ./...
```

---

## 中文

原语：买方和卖方结算一笔现货交易。

问题：为什么交易需要一个原子边界？

本章使用内存事务边界，在结算前快照状态，只在所有账户检查通过后才提交。
它故意做得很小；未来的迭代可以用 SQL 替换内存存储，相同的测试用例
仍应描述行为。

运行：

```bash
go test ./...
```