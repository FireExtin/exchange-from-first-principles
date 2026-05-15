# 12 OMS, Ledger, Compliance, And Paths

Primitive: the trading engine cannot also be every product view.

Question: why separate OMS, risk, ledger, and user-facing state?

This chapter is intentionally empty for now. It will connect the earlier Go
asset/account chapters with the Java trading-system chapters after the command
and event contracts stabilize.

Planned split:

- Java owns deterministic trading-state application.
- Go owns service workflows, reconciliation, and product-facing APIs.
- Shared events describe the handoff between hot-path facts and warm/cold-path
  views.

---

## 中文

原语：交易引擎不能同时也是每个产品视图。

问题：为什么要分离 OMS、风控、账本和面向用户的状态？

本章目前故意留空。它将在命令和事件契约稳定后，连接早期的 Go 资产/账户
章节和 Java 交易系统章节。

计划分离：

- Java 拥有确定性交易状态应用。
- Go 拥有服务工作流、对账和面向产品的 API。
- 共享事件描述热路径事实与温/冷路径视图之间的交接。