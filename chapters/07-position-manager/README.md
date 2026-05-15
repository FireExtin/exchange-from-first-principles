# 07 Position Manager

Purpose: derive account and instrument exposure from execution events.

First implementation should cover:

- apply execution reports;
- maintain net position;
- track average entry price;
- reserve room for realized/unrealized PnL;
- expose replay-friendly state.

Do not implement pricing, liquidation, or cross-margin here yet.

---

## 中文

目的：从成交事件推导账户和合约敞口。

第一个实现应覆盖：

- 应用成交报告；
- 维护净仓位；
- 追踪平均入场价；
- 为已实现/未实现盈亏预留空间；
- 暴露对重放友好的状态。

暂不实现定价、清算或交叉保证金。