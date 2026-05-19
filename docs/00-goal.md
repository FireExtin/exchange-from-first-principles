# Goal

[English](#english) · [中文](#中文)

## English

Build the smallest exchange semantic contract that is still honest about
trading-system failure modes, and teach it in business-sized steps before
changing the architecture underneath it.

The target is not a full exchange. The target is a clear kernel:

- custody assets and user liabilities are explicit;
- balance states, reservations, cancellation, matching, settlement, fees, and
  release are ledger-explainable;
- every accepted command has a sequence number;
- every command is durable before state is changed;
- every state transition is deterministic;
- every resulting event can be replayed and checked;
- every balance movement can be reconciled;
- every order, execution, position, and risk decision can be explained by
  emitted facts.

## Documentation Shape

The project docs are organized around these layers:

- `README.md` explains the project quickly.
- `docs/README.md` is the canonical documentation map.
- `docs/01-core-principles.md` owns stable engineering principles.
- `docs/02-design-paper.md` is the full design narrative.
- chapter READMEs explain local pressure, local guarantees, and how to run the
  local exercises.

When these files disagree, prefer the more specific source: chapter README for
chapter behavior, `docs/07-chapter-roadmap.md` for chapter status, and
`docs/README.md` for document organization.

## Non-goals

- No hand-written consensus in version one.
- No full KYC, wallet, custody, compliance, or web admin system.
- No broad exporter or framework-first architecture.
- No mixed-language hot path during active development.

---

## 中文

构建一个最小化的交易所语义契约，它仍然诚实地面对交易系统的故障模式，并且先用
业务小步骤讲清楚，再替换底层架构。

目标不是完整的交易所，而是一个清晰的核：

- custody assets 和 user liabilities 是显式的；
- balance states、reservation、撤单、撮合、结算、手续费和释放都能用账本解释；
- 每条接受的命令都有一个序列号；
- 每条命令在状态变更前都是持久化的；
- 每个状态转换都是确定性的；
- 每个生成的事件都可以重放和检验；
- 每笔余额变动都可以对账；
- 每个订单、成交、仓位和风控决策都可以由发出的事实解释。

## 文档形状

项目文档分成这些层：

- `README.md` 快速解释项目；
- `docs/README.md` 是规范文档地图；
- `docs/01-core-principles.md` 负责稳定工程原则；
- `docs/02-design-paper.md` 是完整设计叙事；
- 各章节 README 解释本章的局部压力、局部保证和运行方式。

如果这些文件出现不一致，优先相信更具体的来源：章节行为看章节 README，
章节状态看 `docs/07-chapter-roadmap.md`，文档组织看 `docs/README.md`。

## 非目标

- 版本一不手写共识算法。
- 不做完整的 KYC、钱包、托管、合规或后台管理系统。
- 不做宽泛的导出器或框架优先的架构。
- 开发期间不做混语言热路径。
