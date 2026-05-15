# README Matrix Tool

Future home for a small tool that extracts comparison sections from chapter or
version READMEs.

Each README should eventually expose these headings:

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

The tool should produce a table like:

| Version | Problem solved | Architecture change | Semantic change | Cost |
| --- | --- | --- | --- | --- |
| DB | atomic settlement | storage-level ordering | baseline | lock contention |
| Single writer | hot-path determinism | explicit command order | none | in-memory recovery |
| Replicated state machine | failover | replicated total order | none | quorum and ops cost |

This is intentionally not implemented yet. The first step is to keep README
section names stable enough that a script can extract them later.

---

## 中文

### README 矩阵工具

未来存放一个小型工具的地方，该工具从章节或版本 README 中提取比较部分。

每个 README 最终应暴露这些标题：

```text
## Problem Solved
## Architecture Change
## Performance Result
## Semantic Change
## Cost
```

该工具应生成如下表格：

| Version | Problem solved | Architecture change | Semantic change | Cost |
| --- | --- | --- | --- | --- |
| DB | atomic settlement | storage-level ordering | baseline | lock contention |
| Single writer | hot-path determinism | explicit command order | none | in-memory recovery |
| Replicated state machine | failover | replicated total order | none | quorum and ops cost |

这故意尚未实现。第一步是保持 README 章节名称足够稳定，以便脚本稍后提取。