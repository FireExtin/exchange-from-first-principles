# 项目指导范式：从第一性原理构建交易所系统

## 1. 项目目标

这个仓库是一个长期的教育型工程项目。

它的目标不是写一个玩具撮合引擎，也不是简单堆叠现代基础设施组件。
它想通过可运行的代码和文档，解释一个交易系统如何从第一性原理出发
一步步演化：

```text
数据库事务
  -> 内存状态机
  -> 事件日志 / 重放
  -> 复制状态机
  -> OMS / 风控 / 钱包 / 合规 / 冷热路径
```

核心论点：

> 现代交易所的演化，不是随机添加各种技术，而是逐步迁移系统的真相源、
> 并发模型、恢复模型和系统边界。

最开始，数据库事务是系统的真相源。后来，随着性能、恢复、确定性、复制
和审计要求提高，系统逐步演化为显式事件日志和确定性状态机模型。

这个项目需要通过代码、测试、文档和可执行脚本，把这个演化过程显式展示
出来。

## 2. 当前阶段

当前阶段只关注四种显式化：

```text
1. 接口显式化
2. 测试显式化
3. 文档显式化
4. 可执行显式化
```

现阶段不要过度构建高级模块。

不要在第一阶段实现 Raft、DPDK、OMS、风控集群、真实钱包集成或者复杂的
行情基础设施。

当前目标是构建一个最小但扎实的基础，用来证明：

```text
同样的业务语义，可以有不同的技术实现。
```

第一批目标实现是：

```text
dbtx      - 基于数据库事务的引擎
memstate  - 基于内存确定性状态机的引擎
```

这两个实现都必须遵守同一个核心语义接口，并通过同一套一致性测试。

## 3. 工程原则

### 3.1 语义契约优先

最重要的架构思想是：

> 业务语义应该保持稳定，技术实现可以持续演进。

核心契约应该描述系统能接受什么命令、产生什么事实、暴露什么错误，而不是
绑定到某个数据库、消息队列或运行时。

建议的最小接口：

```go
type Engine interface {
    Apply(ctx context.Context, cmd Command) ([]Event, error)
}
```

它背后的抽象是：

```text
old_state + command -> new_state + events
```

数据库事务引擎可以实现这个接口。内存状态机也可以实现这个接口。未来的复制
状态机可以先把 command 提交到复制日志，再 apply 到状态机。

### 3.2 命令表达意图

Command 表示用户或系统的意图。

初始 command 类型只需要覆盖：

```text
CreateAccount
Deposit
SubmitLimitOrder
CancelOrder
```

Command 应该保持明确、小而具体。不要过早添加市价单、止损单、杠杆、合约、
强平等复杂能力。

### 3.3 事件表达事实

Event 表示引擎已经产生的事实。

初始 event 类型只需要覆盖：

```text
AccountCreated
BalanceDeposited
OrderAccepted
OrderRejected
OrderMatched
TradeExecuted
OrderCancelled
LedgerEntryCreated
```

Event 很重要，因为后续阶段会基于事件日志、重放、projection 和审计路径继续
演进。

### 3.4 错误表达业务语义

Error 应该表达业务语义失败，而不是实现细节。

例如：

```text
ErrAccountNotFound
ErrInsufficientBalance
ErrOrderNotFound
ErrOrderAlreadyClosed
ErrInvalidCommand
```

不要让 SQL 错误、数据库驱动错误、内部 map 查询错误直接泄漏到语义接口。
实现细节错误应该被包装或转换为核心层的语义错误。

## 4. 测试显式化

最重要的测试套件是一致性测试套件。

一致性测试只验证核心语义，不依赖具体实现内部细节。每个具体实现把自己传入
共享测试套件。

示例结构：

```go
func RunEngineConformanceSuite(t *testing.T, newEngine func(t *testing.T) core.Engine) {
    t.Run("simple full match", func(t *testing.T) {
        engine := newEngine(t)
        // run shared scenario
    })

    t.Run("partial fill", func(t *testing.T) {
        engine := newEngine(t)
        // run shared scenario
    })

    t.Run("cancel resting order", func(t *testing.T) {
        engine := newEngine(t)
        // run shared scenario
    })

    t.Run("reject insufficient balance", func(t *testing.T) {
        engine := newEngine(t)
        // run shared scenario
    })
}
```

测试应该证明：

```text
给定同样的入金和订单，
当这些 command 被应用到不同 engine 上时，
最终得到的成交、余额、订单状态和账务流水应该等价。
```

初始测试场景：

```text
1. 创建两个账户
2. 给卖方账户充值 BTC
3. 给买方账户充值 USDT
4. 卖方提交限价卖单
5. 买方提交可成交限价买单
6. 引擎产生成交事件
7. 余额正确更新
8. 账务流水正确生成
```

之后再添加：

```text
1. 部分成交
2. 取消挂单
3. 因余额不足拒单
4. 拒绝重复或非法 command
5. 保持价格时间优先
```

测试应该优先关注正确性，而不是性能。

## 5. 文档显式化

每个阶段都必须有 README，解释四件事：

```text
1. 当前系统模型是什么？
2. 它提供了什么语义保证？
3. 它无法解决什么问题？
4. 为什么下一个阶段变得必要？
```

数据库事务阶段应该解释：

```text
数据库事务引擎把关系型数据库视为系统真相源。

它使用 ACID 事务原子地完成：
- 校验余额
- 冻结资金
- 接收订单
- 撮合订单
- 结算成交
- 写入账务流水

这是交易所第一版最简单且正确的模型。

但是，当撮合吞吐要求提高后，这个模型会出现问题，因为热路径上的每一步都
依赖数据库事务、锁和 I/O。
```

内存状态机阶段应该解释：

```text
内存状态机引擎把撮合热路径从数据库中移出。

它在内存中按顺序处理 command，并产生 events。

数据库不再是撮合热路径的真相源。确定性的 command 顺序和产生的 events，
开始成为后续重放与恢复能力的基础。
```

重要设计决策应该写成 ADR，放在：

```text
docs/adr/
```

初始 ADR：

```text
0001-why-semantic-contract.md
0002-why-database-transaction-first.md
0003-why-conformance-tests.md
0004-why-in-memory-state-machine.md
```

ADR 格式：

```markdown
# ADR-0001: 标题

## Status

Accepted

## Context

我们正在解决什么问题？

## Decision

我们做出了什么决定？

## Consequences

什么事情变容易了？
什么事情变难了？
我们接受了什么权衡？
```

## 6. 可执行显式化

这个项目应该很容易在本地运行。

读者不应该在理解整个代码库之前，才能跑起示例。

初始命令：

```bash
make test
make test-dbtx
make test-memstate
make run-dbtx
make run-memstate
```

后续可以添加：

```bash
make bench
make report
```

现阶段保持执行方式简单。

除非必要，不要构建复杂的 Docker Compose 基础设施。如果 `dbtx` 需要数据库，
第一版优先使用 SQLite 降低本地启动成本。如果需要更真实的事务行为，再使用
PostgreSQL 和本地容器编排。

第一阶段应该优先保证可理解和可复现，而不是生产级真实感。

## 7. 暂不做事项

第一阶段不要实现：

```text
Raft
Kafka
Aeron
DPDK
Seastar
Rust FFI
真实链上钱包集成
真实 KYC / AML
风控集群
OMS 微服务
行情分发
合约 / 杠杆 / 强平
分布式部署
高性能 benchmark 套件
```

这些都是未来阶段的内容。

第一阶段只专注于：

```text
核心语义契约
数据库事务实现
内存状态机实现
共享一致性测试
清晰文档
简单可执行脚本
```

## 8. 工程风格

优先选择清晰，而不是炫技。

这是一个教育型系统项目，不是代码炫技项目。

代码应该：

```text
简单
明确
命名清楚
范围小
容易阅读
容易测试
```

不要把重要业务逻辑藏在过度泛化的抽象后面。不要过早做性能优化。除非框架
直接服务于项目目标，否则不要引入框架。

## 9. 第一里程碑

第一里程碑是：

```text
M1: Database Transaction as Source of Truth
```

完成标准：

```text
1. 存在核心 Engine 接口
2. 存在基础 Command 和 Event 类型
3. dbtx engine 实现 Engine
4. 存在一致性测试套件
5. dbtx 通过一致性测试
6. README 解释为什么数据库事务是第一版真相源
7. make test 可以运行
8. make run-dbtx 可以运行一个简单 demo
```

Demo 场景应该展示：

```text
1. 创建卖方和买方账户
2. 给卖方充值 BTC
3. 给买方充值 USDT
4. 卖方提交卖单
5. 买方提交买单
6. 成交执行
7. 打印余额和账务流水
```

## 10. 第二里程碑

第二里程碑是：

```text
M2: In-Memory State Machine with Same Semantics
```

完成标准：

```text
1. memstate engine 实现 Engine
2. memstate 通过完全相同的一致性测试
3. README 解释为什么撮合热路径会移出数据库
4. make run-memstate 可以运行同样的 demo 场景
5. 文档对比 dbtx 和 memstate
```

这个里程碑应该证明核心论点：

```text
实现变了。
业务语义没变。
同一套测试证明了这一点。
```

## 11. 长期方向

后续阶段会逐步演化到：

```text
事件日志
快照与重放
确定性恢复
复制状态机
OMS
风控引擎
钱包与账务
合规
热路径 / 温路径 / 冷路径
低延迟运行时
```

但只有在前两个里程碑足够干净、足够容易理解之后，才应该继续添加这些内容。

## 12. 核心信息

这个仓库想展示的是：交易系统不是一堆技术组件的集合。

它是一系列语义迁移过程：

```text
数据库作为真相源
  -> command log 作为真相源
  -> 确定性状态机作为真相源
  -> 复制日志作为真相源
```

每个阶段都应该回答：

```text
当前的真相源是什么？
它保证了什么语义？
为什么这个模型开始不够用了？
下一个模型把什么东西显式化了？
```
