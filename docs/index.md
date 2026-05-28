---
layout: home
title: codeforces-go · Algorithm Templates
---

# codeforces-go Algorithm Template Library

A production-quality collection of competitive-programming algorithm templates written in Go.
Over **14 000 commits**, **8 500+ GitHub stars**, covering every major algorithm category.

## Why this library?

- Every template includes **complexity analysis**, **reference links**, and **curated problem lists**
- Templates are **tested** — each file ships with `_test.go` verifying correctness
- Go's powerful standard library (slices, maps, generics) keeps implementations concise
- Covers algorithms from **LeetCode 1000** difficulty up to **Codeforces 3500**

## Algorithm Categories

| Category | Key Templates |
|----------|--------------|
| [Data Structures](algorithms/data-structures) | Segment tree, BIT, Skip list, VEB tree, LCT, Treap |
| [Graph Theory](algorithms/graph-theory) | Dijkstra, Max Flow, LCA, HLD, SCC, 2-SAT |
| [Dynamic Programming](algorithms/dynamic-programming) | Knapsack, Slope Trick, Digit DP, Bitmask DP |
| [String Algorithms](algorithms/string-algorithms) | KMP, SAM, Aho-Corasick, Suffix Array |
| Mathematics | FFT/NTT, Number Theory, Matrix Exponentiation |

## Quick Start

```bash
git clone https://github.com/daozhanglv/codeforces-go
cd codeforces-go
go test ./copypasta/...    # run template unit tests
```

## Fork Highlights

This fork adds:

- 🆕 **Skip List** (`copypasta/skip_list.go`) — O(log n) ordered map with floor/ceiling/range queries
- 🆕 **Van Emde Boas Tree** (`copypasta/van_emde_boas.go`) — O(log log U) successor/predecessor
- ✅ **GitHub Actions CI** — automated build + test on every push
- 📖 **English documentation** — this site and `README_EN.md`

## License

[MIT](https://github.com/daozhanglv/codeforces-go/blob/master/LICENSE) © 2019 EndlessCheng
