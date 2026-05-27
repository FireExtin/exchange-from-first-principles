# 《How to Scale a Ledger》中文 Markdown 翻译整理版

> 原文：Modern Treasury - *How to Scale A Ledger*  
> 中文题名：如何扩展一个账本系统  
> 说明：本文是面向工程师的中文翻译整理版，保留核心概念、数据模型、架构原则和示例。英文关键术语保留在括号中，便于对照原文。

---

## 0. 这篇文章主要讲什么

这篇 PDF 讲的是：当一个公司开始处理大规模资金流动时，为什么不能继续把“余额”“支付记录”“订单状态”“银行流水”分散塞进普通业务表里，而应该构建一个专门的、可扩展的、支持复式记账的账本数据库（ledger database）。

文章的核心观点是：

1. **资金系统的核心不是 CRUD，而是不可变、可追溯、可核对的状态机。**
2. **所有资金移动都应该被建模为 Account、Entry、Transaction 三类核心对象。**
3. **Entry 不应该随意修改，历史状态必须能重建。**
4. **Transaction 要保证一组 Entry 原子成功或失败。**
5. **复式记账要阻止资金凭空产生或消失。**
6. **在高并发场景中，账本必须处理幂等、锁、余额断言、缓存漂移和分片问题。**
7. **一个好的账本既要能作为“记录系统”（recording），也要能作为“授权系统”（authorizing）。**

这篇文章很适合后端、支付、交易所、钱包、清结算、风控、对账和金融基础设施工程师阅读。

---

## 1. 引言

Modern Treasury 在为大型金融科技公司和 marketplace 提供 API 服务时发现：很多公司一开始只是直接对接银行 API、卡网络、支付处理器或其他金融 SaaS。但业务规模变大后，他们需要跟踪大量资金交易和余额变化，于是开始意识到：资金流动不能只靠散落在业务数据库里的字段来表达。

这些公司真正需要的是一个 **ledger（账本）**：一个可扩展的、复式记账的数据库，用来记录金融交易和余额。

这篇文章不是给会计看的，而是给构建金融产品的工程师看的。它总结了如何实现一个可靠、一致、快速的账本系统。

---

## 2. 第一章：为什么要使用账本数据库

### 2.1 传统模型为什么会在规模化后失效

很多公司早期会把金融状态直接嵌入业务模型：

- 拼车价格存在 ride 对象里。
- 健身房月付金额散落在 booking 记录中。
- 贷款剩余应还金额存在 loan 对象上。

这种方式早期很快，但规模变大后会出现问题：

#### 1. 报表需求膨胀

随着用户增长，公司需要更精细的财务报表和合规报表。每一笔资金移动都必须不可变地记录下来，并且需要追踪资金来源、去向以及每个账户的余额。

#### 2. 数据碎片化

金融事件来自多个地方：银行 API、支付处理器、卡处理器、ACH、加密网络、内部数据库等。没有统一账本时，要重建完整的资金历史会越来越难。

#### 3. 性能下降

用户增长后，付款 cron job 会错过 deadline，授权检查会变慢，其他瓶颈会出现。

#### 4. 失败场景增多

超额扣款、余额不足、欺诈交易、被拒付款等问题都会越来越频繁。如果没有清晰审计线索，问题很难追踪和回滚。

### 2.2 账本数据库的核心特性

一个账本数据库的核心数据模型并不复杂。它包含三个基础对象：

| 对象 | 含义 |
|---|---|
| Account | 一个离散的价值池，用来表示余额 |
| Transaction | 一个原子的资金事件，会影响账户 |
| Entry | 构成交易的单条借记或贷记记录 |

一个可扩展账本数据库应该提供四类保证：

1. **不可变性（immutability）**：所有改变都以历史状态可恢复的方式记录下来。
2. **复式记账约束（double-entry enforcement）**：资金移动必须说明来源和去向。
3. **并发控制（concurrency controls）**：即使并行写入、乱序写入，也不能重复花钱。
4. **高效聚合（efficient aggregations）**：能够快速计算账户余额、历史余额和报表指标。

### 2.3 Fintech Translation Problem：金融科技里的“翻译问题”

产品工程师通常说的是业务语言：订单、乘车、预订、贷款、订阅。金融系统说的是另一套语言：账户、借方、贷方、余额、交易、清算。

当公司没有账本时，工程师必须把业务模型翻译成复式记账模型。这个翻译很难，因为金融事件通常有三个特点：

#### 1. 来源分散且异构

一个 App 可能依赖多个金融 SaaS：发卡、支付处理、银行账户、ACH 等。资金移动可能跨越多个系统。

#### 2. 每个系统说不同语言

一个 fintech 产品可能有自己的 API、自己的文档、自己的历史遗留格式。银行通常用文件格式表达交易，解析起来并不容易。

#### 3. 底层记录可能会变

很多来源系统并不给你一个稳定、不可变的“某一时刻资金值”。如果依赖 mutable records，追踪资金状态会非常困难。

### 2.4 账本错误的真实后果

如果资金归属不清楚，会产生严重后果：

- 商户不知道为什么账上少了一笔钱。
- 钱包用户转账后，另一方没收到钱。
- 贷款客户按时还款，却因为系统记录延迟而被标记为逾期。
- 监管机构要求解释资金流向时，公司无法给出一致答案。

解决方案是：把每个资金事件变成不可变的复式记账记录。

### 2.5 为什么很多公司会拖延建设账本

因为建设一个可靠、高性能、可扩展的复式账本非常难，需要：

- 多年迁移单式记账系统的工程投入。
- 深入理解会计原则并把它映射到系统架构。
- 持续维护对账、报表和修复工具。

但拖延的代价是：资金系统越长越复杂，后期修复成本更高。

---

## 3. 第二章：把金融事件映射到复式记账原语

### 3.1 给金融事件加入复式记账

任何金融事件，无论是用户充值、贷款还款、刷卡授权，还是加密货币交易，都可以抽象为账本中的资金移动。

最基础的核心对象是：

- Account
- Entry
- Transaction

Account 有多个 Entry；Transaction 也有多个 Entry。Entry 是连接 Account 和 Transaction 的核心记录。

### 3.2 Account：所有余额的总和

Account 表示一个离散的价值池，并以某种货币计价。例子包括：

- 数字钱包账户
- 公司运营现金账户
- 贷款本金账户
- 用户卡账户

Account 常见字段：

| 字段 | 类型 | 含义 |
|---|---|---|
| id | uuid | 账户唯一 ID |
| normal_balance | string | 账户的正常余额方向：debit 或 credit |
| posted_balance | Balance | 已结算余额 |
| pending_balance | Balance | 已结算 + 待结算的预期余额 |
| available_balance | Balance | 可用余额，可用于转出或消费 |

### 3.3 三种余额：posted、pending、available

资金在现实世界中并不总是立即结算。例如 ACH 入账可能要一天，信用卡授权和清算之间也有时间差。因此账户通常需要维护三种余额：

| 余额 | 含义 |
|---|---|
| Posted Balance | 已经完全结算的资金 |
| Pending Balance | 已结算资金 + 预计即将结算的资金 |
| Available Balance | 可以立刻使用的资金，不一定等于 pending |

不同产品会选择不同余额作为业务依据。例如：

- 钱包提现通常用 available balance，避免超额转出。
- 贷款账户展示过去状态时，可能用 pending balance。
- 财务报表通常更关注 posted balance。

### 3.4 Entry：不可变的资金移动记录

账户余额不应该被直接修改。余额变化应该通过写入 Entry 来表达。

Entry 是账户余额变化的完整记录，常见字段包括：

| 字段 | 含义 |
|---|---|
| id | Entry ID |
| account_id | 所属账户 |
| status | pending、posted 或 archived |
| direction | debit 或 credit |
| amount | 金额 |
| discarded_at | 如果该 Entry 被废弃，记录废弃时间 |

原文用信用卡生命周期作为例子：

1. 信用卡账户开始有 10,000 美元额度。
2. 用户刷卡买 1,000 美元机票，交易先进入 pending。
3. 当晚交易清算，pending Entry 被废弃，替换为 posted Entry。
4. 用户发起 1,000 美元还款，先进入 pending。
5. 还款完成后，pending Entry 被替换为 posted Entry。
6. 酒店押金 250 美元作为 hold 进入 pending。
7. 离店后 hold 被释放，pending Entry 被 archived。

### 3.5 为什么要 discard Entry，而不是直接修改

账本要求不可变。posted 和 archived Entry 是永久的；pending Entry 可以被替换，但不能被原地修改。

如果一个 pending Entry 的金额或状态变化，系统应该：

1. 给旧 Entry 设置 discarded_at。
2. 创建新的 Entry 表示新状态。

这样做的好处是：

- 账户历史更干净。
- 审计更容易。
- 当前余额可以通过排除 discarded Entry 来计算。
- 不会丢失状态变更历史。

### 3.6 计算账户余额：normal balance 很关键

账户有两类正常余额方向：

| 类型 | 代表 | debit 的影响 | credit 的影响 |
|---|---|---|---|
| Debit-normal | 资产、费用 | 增加余额 | 减少余额 |
| Credit-normal | 负债、权益、收入 | 减少余额 | 增加余额 |

以数字钱包充值为例：

- 公司现金账户增加，因为公司收到现金。现金账户通常是 debit-normal，因此 debit 增加余额。
- 用户钱包账户增加，因为平台欠用户的钱变多。用户钱包从平台角度看是负债，通常是 credit-normal，因此 credit 增加余额。

这解释了为什么“钱从 A 到 B”不能简单理解成“credit 是来源、debit 是用途”。在复式记账中，direction 必须结合账户的 normal_balance 来理解。

### 3.7 余额计算字段

为了快速计算余额，账本可以维护以下聚合字段：

- posted_debits：所有 posted 借记 Entry 之和。
- posted_credits：所有 posted 贷记 Entry 之和。
- pending_debits：posted_debits 加上未废弃 pending 借记 Entry 之和。
- pending_credits：posted_credits 加上未废弃 pending 贷记 Entry 之和。
- normal_balance：credit 或 debit。

余额公式大致如下：

```text
if normal_balance == "credit":
    balance = credits - debits
else:
    balance = debits - credits
```

账本应该成为金融事实的唯一来源（source of financial truth），所有影响金融状态的行为都应映射到 Account、Entry 和 Transaction。

---

## 4. 第三章：用于原子资金移动的 Transaction 模型

### 4.1 为什么 Transaction 很重要

Account 和 Entry 可以提供不可变的余额变化审计线索，但它们本身不能保证一组 Entry 同时成功或失败。

例如，Bob 给 Alice 转 10 美元，需要两条 Entry：

- Bobby 的账户 debit 10 美元。
- Alice 的账户 credit 10 美元。

如果 Bobby 的 Entry 写成功，而 Alice 的 Entry 写失败，系统就进入不一致状态：Bob 少了钱，Alice 没收到钱。

Transaction 的作用是：把一组 Entry 绑定在一起，使它们作为一个整体原子成功或失败。

### 4.2 Transaction 字段

| 字段 | 类型 | 含义 |
|---|---|---|
| id | uuid | Transaction 唯一 ID |
| status | string | pending、posted 或 archived |
| entries | Entry[] | 该 Transaction 下的所有 Entry |

### 4.3 Transaction 生命周期

原文定义了三种状态：

1. **pending**：初始状态，Entry 已经写入，但交易未最终完成。
2. **posted**：最终完成状态。
3. **archived**：在 posted 前取消。

### 4.4 创建 pending Transaction

创建 pending Transaction 时，系统会持久化 debit 和 credit Entry，但交易尚未 finalize。

### 4.5 posted：完成 Transaction

由于 Entry 是不可变的，把 Transaction 从 pending 变成 posted 时不能直接修改原 Entry。正确流程是：

1. 废弃原 pending Entry。
2. 创建新的 posted Entry。
3. 保留历史，保证 auditability。

### 4.6 archived：取消 Transaction

posted Transaction 不能 archived，因为它已经是最终状态。pending Transaction 可以进入 archived，流程类似 posted：

1. 废弃 pending Entry。
2. 创建 archived Entry。

### 4.7 Transaction 模型的开发收益

| 收益 | 含义 |
|---|---|
| Reliability | 避免不平衡 Entry 和意外产生/销毁资金 |
| Auditability | Transaction 级别的可追溯性便于调试和合规 |
| Modularity | 可以封装复杂资金流，比如转账、退款、结算 |

---

## 5. 第四章：向账本写入数据：记录模式与授权模式

### 5.1 两种账本模式

大部分账本可以按使用方式分为两类：

| 模式 | 含义 |
|---|---|
| Recording | 记录已经发生在外部系统中的资金移动 |
| Authorizing | 根据账户状态主动批准或拒绝交易 |

### 5.2 Recording：记录模式

Recording 指捕获外部系统中已经发生的金融事件，例如银行、支付处理器、卡网络等，并把它们翻译成账本中的核心数据模型。

Recording 的特点：

- 高写入吞吐：可能每秒上千甚至更多写入。
- 异步处理：读到的账本状态可能短暂过期，允许最终一致性。
- 支持复杂查询：需要筛选、聚合、生成报表。

### 5.3 Recording 架构

原文架构如下：

```text
Event Source  ->  Application Layer  ->  Ledger Database
                                \->  Domain Object Database
```

各层含义：

| 组件 | 含义 |
|---|---|
| Event Source | 资金移动来源，比如银行、支付处理器、Modern Treasury API |
| Application Layer | 应用层，负责把事件源数据翻译成账本模型和业务对象 |
| Ledger DB | 不可变的资金事件日志 |
| Domain Object DB | 应用自己的业务状态，记录与资金无关的数据 |

### 5.4 Recording 下如何维护一致性

Ledger DB 和 Domain Object DB 都由 Application Layer 写入，因此二者可能短暂不一致。Recording 允许 eventual consistency（最终一致性），但需要支持：

- 客户端提供幂等键。
- Account balance version。
- effective_at 时间戳。

### 5.5 effective_at：事件真实发生时间

很多资金事件在被写入账本之前已经发生。例如银行交易、区块链交易、支付处理器通知等。

因此 Transaction 需要一个 `effective_at` 字段，表示这笔交易在外部系统真实发生的时间，而不是被账本记录的时间。

`effective_at` 的用途：

- 回填历史交易。
- 按真实发生时间重建余额。
- 支持按正确顺序生成历史报表。

### 5.6 Account Balance Versions：账户余额版本

如果允许客户端修改历史余额，会产生一个问题：查询某个时间点的 Entry 和查询该时间点的账户余额可能不一致。

为了解决这个问题，账本会给 Account 和 Entry 都加版本号：

| 对象 | 字段 | 含义 |
|---|---|---|
| Account | version | 每次该账户有 Entry 创建或变更时递增 |
| Entry | account_version | 该 Entry 对应的账户版本 |

这样就能精确回答：某一时刻某个 Account 的 posted_balance 对应哪些 Entry。

### 5.7 Recording 适用场景

#### 1. 展示账户详情

钱包、银行卡、券商等产品通常需要在 UI 展示账户余额。使用 Entry account_version 可以保证展示的余额和 Entry 一致。

#### 2. Payouts

Marketplace 向用户付款时，真实支付状态来自外部处理器。`effective_at` 可以保证哪怕事件乱序到达，也能正确记录。

#### 3. Loan Servicing

贷款服务通常依赖外部 payment processor。利息计算、还款入账和历史状态都需要支持回填与历史余额计算。

#### 4. Crypto

区块链交易发生在链上，而不是应用自己的账本里。应用应当用 `effective_at` 记录链上交易真实发生时间。

### 5.8 Authorizing：授权模式

Authorizing 指账本主动根据账户状态批准或拒绝交易，比如：

- 用户钱包是否有足够余额提现。
- 信用卡授权是否允许消费。
- 账户余额是否会被扣成负数。

Authorizing 的特点：

| 特点 | 含义 |
|---|---|
| Read-after-write consistency | Entry 更新后应立即反映到账户 |
| Lower throughput | 由于需要同步检查，吞吐低于 recording |
| Balance assertions | 写入时检查账户余额不变量 |
| Concurrency control | 用版本或余额锁保证原子性 |

### 5.9 Version Locking：版本锁

版本锁类似 optimistic locking（乐观锁）。流程是：

1. 启动数据库事务。
2. 写入 ledger entry。
3. 更新 Account version，并带上当前 version 条件。
4. 如果更新成功，说明客户端读到的 version 仍然有效，提交事务；否则回滚。

这种方式能防止乱序更新，但对热点账户不友好。

### 5.10 Balance Locking：余额锁

版本锁在高并发热点账户上容易失败。例如卡账户授权频繁，每次交易都会改 version。

更实用的方式是余额断言：在 Entry 上声明交易提交后账户余额必须满足某些条件。

Entry 可以包含：

| 字段 | 含义 |
|---|---|
| pending_balance_amount | Transaction 提交后 pending balance 必须满足的条件 |
| posted_balance_amount | Transaction 提交后 posted balance 必须满足的条件 |
| available_balance_amount | Transaction 提交后 available balance 必须满足的条件 |

BalanceCondition 常见比较操作：

| 操作 | 含义 |
|---|---|
| lt | 小于 |
| lte | 小于等于 |
| eq | 等于 |
| gte | 大于等于 |
| gt | 大于 |

例如：钱包提现可以要求 `available_balance >= 0`。如果写入后余额会小于 0，交易就失败。

### 5.11 Authorizing 适用场景

#### 1. Digital Wallets

用户从钱包提现时，账本必须先确认可用余额足够。否则会发生 double-spend。

#### 2. Card Authorizations

信用卡授权要求账本能够在某一时间点读取准确余额，并把授权金额保留到后续清算。

### 5.12 Mixed-Mode Ledgers：混合模式账本

现实中，一个产品可能既需要 Recording，也需要 Authorizing。

很多账本实现一开始只针对一种模式优化，后来不得不维护多个账本，导致复杂度上升。更好的账本应该允许在 Entry 级别选择模式，而不是只在 Account 级别选择模式。

---

## 6. 第五章：不可变性与复式记账

### 6.1 不可变性：重建每个历史状态

账本最重要的保证之一是：每个状态变化都必须被记录，并且可以容易地重建。

虽然 Account、Transaction、Entry 都有可变字段，但底层必须有不可变日志记录所有变化。

原文提到几种重建历史状态的方式：

- `effective_at` timestamp
- `account_version` number
- `transaction version` fields

### 6.2 查询过去状态

如果要查询某个账户在某个时间点的 Entry，可以按如下条件过滤：

```sql
WHERE account_id = 'account_a'
  AND effective_at <= 'timestamp_a'
  AND (discarded_at IS NULL OR discarded_at > 'timestamp_a')
```

这能排除在该时间点之后才被废弃的 Entry。

### 6.3 为什么需要版本号

时间戳不能解决所有问题，因为：

- Entry 可以用过去的 effective_at 写入。
- Entry 可能共享相同 effective_at 或 discarded_at。
- pending Transaction 在 posted 前可能多次修改。

因此 Transaction 也需要 version 字段，用来表示当前版本和历史版本。

### 6.4 Bill-splitting 示例

原文用分账 App 解释 Transaction version：

1. Version 0：Chuck 加入账单，暂时一个人承担全部金额。
2. Version 1：Dani 加入，账单拆成两份。
3. Version 2：Elio 加入，账单拆成三份。
4. Version 3：账单最终 posted。

每个版本都会保留历史，这样可以重建分账过程中的每个中间状态。

### 6.5 Audit Logs：审计日志

除了记录数据模型本身的历史状态，系统还需要记录：谁做了变更、什么时候做的、通过什么来源做的。

Audit Log 字段：

| 字段 | 含义 |
|---|---|
| id | 审计日志 ID |
| action | Create、Update、Delete 等操作 |
| entity | 被操作的实体 |
| source | 变更来源，例如 API key 或内部用户 |
| data | 变更前后的结构化描述 |
| occurred_at | 变更发生时间 |

### 6.6 复式记账：防止资金产生或消失

一旦账本具备不可变性，下一步就是确保资金不会凭空产生或消失。

每个 Transaction 必须满足：

- 至少两个 Entry：一个 debit，一个 credit。
- 每种货币内部，debit 总额必须等于 credit 总额。

### 6.7 为什么要按货币维度平衡

如果只检查所有 Entry 的 debit 和 credit 总额相等，会在多币种场景下出错。

比如一个加密货币购买交易同时涉及 USD 和 ETH。不能简单把 USD 金额和 ETH 数量加在一起验证平衡。

正确方法是：

```text
按 currency 分组；
每种 currency 内部分别验证 debit == credit。
```

不能把“汇率换算后总额相等”作为账本平衡依据，因为：

1. 汇率会随时间变化，历史验证困难。
2. 不存在所有系统都认同的统一汇率。
3. 现实换汇一定涉及至少四个账户：用户 USD、用户 ETH、平台 USD、平台 ETH。

### 6.8 Card Authorization 示例

信用卡授权和清算看起来简单，但其实很复杂。

一个授权请求通常包含两类 Entry：

- 用户卡账户上的授权 Entry：必须检查余额、不能导致账户低于允许下限，需要同步一致性。
- 处理器或发卡行相关账户上的记录 Entry：可能只需要最终一致性、高吞吐、批量处理。

因此同一个 Transaction 内，部分 Entry 可能需要 Authorizing，部分 Entry 只需要 Recording。

这说明最好的账本不是 Recording 或 Authorizing 二选一，而是二者都支持。

---

## 7. 第六章：并发控制、性能与更多扩展问题

### 7.1 并发控制：防止 double-spend

当应用规模增长后，并发会导致重复花钱。

典型场景：

1. 客户端发送创建 Transaction 请求。
2. 服务端处理较慢，客户端超时。
3. 客户端重试同一个请求。
4. 第一次请求其实已经成功。
5. 第二次请求又成功，导致资金移动两次。

解决方案之一是 idempotency key（幂等键）。

### 7.2 Idempotency Keys：幂等键

客户端每次想表达“同一个业务请求”时，应传入相同的 idempotency key。

伪代码：

```text
idempotency_key = generate_uuid()
response = None
while response is None or not response.is_successful():
    response = create_transaction(idempotency_key)
```

关键点：

- 幂等键应该在重试循环外生成。
- 服务端需要保存一定时间，例如 24 小时。
- 如果同一个 key 已处理过，应返回上一次响应，而不是再次移动资金。

### 7.3 高效聚合：余额读取性能

账本通常不可变、平衡且支持并发写入。但如果每次读取账户余额都从所有 Entry 重新计算，会越来越慢。

因此需要缓存账户余额聚合字段，例如：

- pending_debits
- pending_credits
- posted_debits
- posted_credits

### 7.4 Current Balances：当前余额缓存

当前余额缓存用于：

- 展示账户余额。
- Authorizing 中执行余额锁。

因为余额锁依赖最新余额，所以当前余额缓存必须同步更新。写入 Entry 时，数据库事务应同时更新账户余额缓存，避免 stale read。

例如写入一条 posted debit Entry，系统应自动增加对应的 posted_debits。如果 pending Entry archived，则减少 pending_debits。

### 7.5 Effective Time Balances：按 effective_at 查询历史余额

按 effective_at 查询历史余额更难，因为交易可能乱序写入，也可能回填历史。

原文讨论两种缓存策略。

#### 方案一：Anchoring

每天缓存一次日终余额。查询某个时间点余额时：

1. 找到不晚于查询时间的最近 cache entry。
2. 从该 cache entry 的时间开始，叠加日内 Entry。
3. 得到目标时间余额。

缺点是：如果新写入一条更早 effective_at 的 Entry，之后很多缓存点可能都要更新。

#### 方案二：Resulting Balances

每条缓存记录保存“应用某个 Entry 后的结果余额”。这样可以更精细地定位历史余额。

但缺点是：当插入更早 effective_at 的 Entry 时，后续缓存记录都可能需要修正，更新成本更高。

### 7.6 Monitoring Cache Drift：监控缓存漂移

缓存能提高性能，但缓存值可能和 Entry 真值发生偏离。

系统需要：

1. 定期验证每个账户缓存余额是否等于 Entry 求和结果。
2. 如果发现某个账户 drift，自动关闭该账户的缓存读取。
3. 提供 backfill 工具和 on-call runbook，让工程师能修复缓存漂移。

### 7.7 高级扩展主题

即使账本具备不可变性、复式记账、并发安全和快速聚合，产品复杂度和交易量继续上升后，仍然会遇到更多问题。

#### 1. 账户聚合与报表 UX

有些场景需要把多个账户聚合展示，例如：

- 业务报表：组合多个 Transaction 和账户余额。
- 产品体验：子账户或账户组共享一个用户视角下的余额。

解决方案：Account Categories，使用图结构而不是僵硬层级来聚合账户。

能力包括：

- 父子账户层级。
- 基于账户的权限控制。
- 灵活的余额和交易报表 rollup。

#### 2. 灵活 Transaction 搜索

Marketplace 或券商场景会频繁查询最近交易、生成对账单、按日期范围筛选、模糊匹配、排序等。

扩展策略包括：

- 根据 query shape 优化索引。
- 对高吞吐表做分区。
- 实现基于游标的分页，而不是 offset。
- 监控搜索延迟。

#### 3. 高吞吐写入队列

一些 fintech 账本会超过 5,000 QPS，同步写入可能成为瓶颈。

解决方案：Write-Ahead Queues。

特点：

- 用异步 Entry 处理队列缓冲和批处理账本写入。
- 队列必须 durable、idempotent，并能承接高 QPS。
- 好处包括吸收负载尖峰、把应用逻辑和数据库性能解耦、支持独立重试和监控。

#### 4. 跨数据库分片

当单库无法支撑账本吞吐和存储需求时，需要分片。

分片会带来挑战：

- 如何跨 shard 保持复式记账原子性。
- 如何跨分布式账户高效聚合 Entry 和余额。
- 如何协调后台任务，例如 balance drift checks。

解决方案包括：

- 在 Account 层面提供强一致性保证。
- 管理跨 shard Transaction。
- 每个 shard 独立采集指标和告警。

---

## 8. 工程师视角的核心启发

这篇文章真正想传达的是：账本系统不是普通业务表，也不是简单流水表。

一个可靠的账本至少要同时处理：

1. **Correctness**：余额永远不能错。
2. **Auditability**：任何历史状态都能解释。
3. **Atomicity**：一组资金变化必须一起成功或失败。
4. **Immutability**：历史不能被悄悄改掉。
5. **Idempotency**：重试不能导致重复扣款。
6. **Concurrency**：并发写入不能造成 double-spend。
7. **Performance**：余额读取和报表查询不能每次全量扫描 Entry。
8. **Scalability**：高 QPS、分片、异步队列、缓存漂移都要提前设计。

对交易所、支付平台、钱包、券商和清结算系统来说，这些原则本质上相同。

---

## 9. 关键术语表

| English | IPA | 中文 | 解释 |
|---|---|---|---|
| Ledger | /ˈledʒər/ | 账本 | 记录资金事件和余额的系统 |
| Ledger Database | /ˈledʒər ˈdeɪtəbeɪs/ | 账本数据库 | 支持复式记账和历史重建的数据库 |
| Double-entry | /ˌdʌbəl ˈentri/ | 复式记账 | 每笔资金移动同时记录来源和去向 |
| Account | /əˈkaʊnt/ | 账户 | 一个离散的价值池 |
| Entry | /ˈentri/ | 分录 | 单条借记或贷记记录 |
| Transaction | /trænˈzækʃən/ | 交易 | 一组原子完成的 Entry |
| Immutability | /ɪˌmjuːtəˈbɪləti/ | 不可变性 | 历史状态不能被直接覆盖 |
| Idempotency | /ˌaɪdəmˈpoʊtənsi/ | 幂等性 | 同一请求重试多次只生效一次 |
| Concurrency Control | /kənˈkʌrənsi kənˈtroʊl/ | 并发控制 | 防止并发写入导致状态错误 |
| Double-spend | /ˈdʌbəl spend/ | 双花 | 同一笔钱被重复花掉 |
| Balance Locking | /ˈbæləns ˈlɑːkɪŋ/ | 余额锁 | 用余额断言控制交易是否能提交 |
| Version Locking | /ˈvɜːrʒən ˈlɑːkɪŋ/ | 版本锁 | 用账户版本号防止乱序更新 |
| Effective At | /ɪˈfektɪv æt/ | 生效时间 | 交易在外部系统真实发生的时间 |
| Audit Log | /ˈɔːdɪt lɔːɡ/ | 审计日志 | 记录谁在何时做了什么变更 |
| Cache Drift | /kæʃ drɪft/ | 缓存漂移 | 缓存余额与真实 Entry 求和结果不一致 |
| Sharding | /ˈʃɑːrdɪŋ/ | 分片 | 把数据拆到多个数据库或节点 |

---

## 10. 一句话总结

这篇文章讲的不是“怎么建一个流水表”，而是“怎么把资金系统建成一个可审计、不可变、支持并发和高吞吐的分布式金融状态机”。
