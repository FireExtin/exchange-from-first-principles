# Replay And Recovery

Replay rebuilds state from the command log by reapplying commands in sequence.

The important property is not speed first. The important property is that the
same command stream produces the same event stream and final state.

## Recovery Checks

- all command sequence numbers are contiguous;
- replayed events match expected event count and type;
- cash + frozen cash are conserved;
- base asset + frozen base asset are conserved;
- open orders match account reservations.

---

## 中文

重放通过按顺序重新应用命令，从命令日志重建状态。

最重要的特性不是速度优先。重要的是相同的命令流产生相同的事件流和最终状态。

## 恢复检查

- 所有命令序列号是连续的；
- 重放的事件与预期的事件数量和类型匹配；
- 现金 + 冻结现金守恒；
- 基础资产 + 冻结基础资产守恒；
- 活跃订单与账户预留匹配。