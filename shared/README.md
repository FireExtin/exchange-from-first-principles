# Shared Contracts

[English](#english) · [中文](#中文)

## English

This directory contains semantic contracts shared across chapters.

The contracts are not a framework. They are a language for proving that the same
exchange behavior survives architecture changes. The exchange contract is
presented progressively: accounting and funds first, then orders and
executions, with position, risk, and projection surfaces kept visible for later
chapters. The first four teaching chapters build the contract from smaller
actions: deposit, balance-state movement, order reservation, and settlement.
The contract distinguishes posted financial facts from derived model and
projection state.

## Current Contract Layers

- `shared/go/types`: common identifiers and scaled amounts.
- `shared/go/funds`: the earlier funds contract used by appendix contract
  scaffolds 92 and 93.
- `shared/go/exchange`: the new exchange-level semantic contract covering
  accounting, orders, executions, positions, risk, and projection cursors.
- `shared/go/credit`: optional credit/margin/funding extension contract for
  collateral, borrow liability, funding accrual, repayment, and liquidation
  boundaries.

## Rules

- Shared contracts may be consumed by any chapter.
- Shared contracts must not import code from a chapter.
- Chapter implementations may diverge internally as long as public facts can be
  explained through these contracts.
- Posted financial facts should be expressed as explicit journal entries or
  settlement facts.
- Derived state such as marks, unrealized PnL, risk views, and projections must
  keep provenance and must not silently mutate ledger truth.
- Incomplete exchange behavior belongs behind explicit TODO tests or build
  tags, not in silent partial implementations.

---

## 中文

本目录包含章节间共享的语义契约。

这些契约不是框架，而是一套语言，用来证明同一交易所行为可以在架构变化中存活。
exchange contract 会渐进呈现：先是会计和资金，再到订单与成交；仓位、风控和
projection 表面保留给后续章节。前四个教学章节从更小的动作构建契约：入金、余额
状态移动、订单冻结和结算。契约会区分已过账财务事实，以及派生模型/projection
状态。

## 当前契约层

- `shared/go/types`：通用标识符和缩放金额。
- `shared/go/funds`：早期资金契约，供附录契约脚手架 92 和 93 使用。
- `shared/go/exchange`：新的交易所级语义契约，覆盖会计、订单、成交、仓位、
  风控和 projection cursors。
- `shared/go/credit`：可选的信用/保证金/资金费率扩展契约，用于表达抵押品、
  借款负债、资金费率计提、还款和强平结算边界。

## 规则

- 共享契约可被任何章节消费。
- 共享契约不得从章节导入代码。
- 章节实现可以内部发散，只要公开事实可以通过这些契约解释。
- 已过账财务事实应表达为显式 journal entries 或 settlement facts。
- 标记价格、未实现 PnL、风控视图和 projections 等派生状态必须保留来源信息，
  且不能偷偷改写账本真相。
- 未完成的 exchange 行为应放在显式 TODO 测试或 build tag 后，而不是沉默的半成品实现。
