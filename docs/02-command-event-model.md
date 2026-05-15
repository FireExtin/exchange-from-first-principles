# Command And Event Model

Commands are user or system intents. Events are facts produced by the state
machine.

## Commands

- `Deposit`
- `PlaceLimit`
- `Cancel`

## Events

- `Deposited`
- `OrderAccepted`
- `OrderRejected`
- `TradeExecuted`
- `OrderRested`
- `OrderCancelled`
- `CancelRejected`

## Rule

Only events mutate downstream views. Commands are input. Events are facts.

---

## 中文

命令是用户或系统的意图。事件是状态机产生的事实。

## 命令

- `Deposit`
- `PlaceLimit`
- `Cancel`

## 事件

- `Deposited`
- `OrderAccepted`
- `OrderRejected`
- `TradeExecuted`
- `OrderRested`
- `OrderCancelled`
- `CancelRejected`

## 规则

只有事件才能改变下游视图。命令是输入。事件是事实。