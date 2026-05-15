# Goal

Build the smallest exchange core that is still honest about trading-system
failure modes.

The target is not a full exchange. The target is a clear kernel:

- every accepted command has a sequence number;
- every command is durable before state is changed;
- every state transition is deterministic;
- every resulting event can be replayed and checked;
- every balance movement can be reconciled.

## Non-goals

- No hand-written consensus in version one.
- No full KYC, wallet, custody, compliance, or web admin system.
- No broad exporter or framework-first architecture.
- No mixed-language hot path during active development.

---

## 中文

构建一个最小化的交易所核心，它仍然诚实地面对交易系统的故障模式。

目标不是完整的交易所，而是一个清晰的核：

- 每条接受的命令都有一个序列号；
- 每条命令在状态变更前都是持久化的；
- 每个状态转换都是确定性的；
- 每个生成的事件都可以重放和检验；
- 每笔余额变动都可以对账。

## 非目标

- 版本一不手写共识算法。
- 不做完整的 KYC、钱包、托管、合规或后台管理系统。
- 不做宽泛的导出器或框架优先的架构。
- 开发期间不做混语言热路径。