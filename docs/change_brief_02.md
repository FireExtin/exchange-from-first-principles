# Change Brief 02: Documentation Organization Refresh

## 1. Summary

This change reorganizes the documentation system without changing project code.

The main decision is to separate documentation roles:

- root `README.md` is the project front door;
- `docs/README.md` is the canonical documentation map;
- `docs/00-goal.md` states the goal and source-of-truth rules;
- `docs/07-chapter-roadmap.md` owns chapter order and status vocabulary;
- chapter READMEs own local problem statements and run commands;
- change briefs remain append-only historical records.

## 2. Changes

| Area | Change |
| --- | --- |
| Documentation map | Added `docs/README.md` with reading paths and a full document catalog. |
| Root README | Added toolchain notes and redirected detailed reading guidance to `docs/README.md`. |
| Goal doc | Added documentation-layer rules and conflict-resolution guidance. |
| Roadmap | Added status vocabulary and clarified that chapters 17-21 are planned notes only. |
| Chapter 11 | Marked the Aeron Java module as a runnable skeleton and added run commands. |
| Chapter 15 | Added a README for the runnable Rust hot-path experiment. |
| System boundary | Clarified that Rust is exploratory and runnable, but not the active track. |

## 3. Current Reading Order

For new readers:

1. `README.md`
2. `docs/README.md`
3. `docs/00-goal.md`
4. `docs/10-design-paper.md`
5. `docs/07-chapter-roadmap.md`
6. The README for the chapter being run or changed

For implementation work:

1. `shared/README.md`
2. `integration-tests/README.md`
3. `docs/12-version-contract-and-testing.md`
4. The target chapter README

## 4. Verification

Docs-only change. Recommended checks:

```bash
make test-go
make test-rust
make test-java
```

---

## 中文

本次变更只整理文档系统，不修改项目代码。

核心决定是区分文档职责：

- 根目录 `README.md` 是项目入口；
- `docs/README.md` 是规范文档地图；
- `docs/00-goal.md` 描述目标和真相源规则；
- `docs/07-chapter-roadmap.md` 负责章节顺序和状态术语；
- 各章节 README 负责局部问题、局部模型和运行命令；
- change brief 保持追加式历史记录。
