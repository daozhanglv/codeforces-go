# codeforces-go · Algorithm Competition Template Library

> **English edition** of [算法竞赛模板库](./README.md) — maintained by [@EndlessCheng](https://github.com/EndlessCheng) (Codeforces Master, rating 2208) with additional contributions by this fork.

A production-quality Go library of competitive-programming algorithm templates. Every template ships with:

- Core idea & complexity analysis  
- Curated reference links (papers / editorial posts)  
- Ready-to-paste implementation  
- Representative problems (Codeforces · LeetCode · AtCoder)

---

## Quick Start

```bash
# requires Go 1.23+
git clone https://github.com/daozhanglv/codeforces-go
cd codeforces-go
go test ./copypasta/...        # run all template unit tests
go test ./leetcode/...         # run LeetCode solution tests
```

---

## Algorithm Index

### Data Structures

| File | Contents |
|------|----------|
| [fenwick_tree.go](copypasta/fenwick_tree.go) | Binary Indexed Tree (BIT / Fenwick Tree), 2-D BIT, offline 2-D point queries |
| [segment_tree.go](copypasta/segment_tree.go) | Segment tree with lazy propagation, persistent (Merge-Sort Tree), dynamic nodes, seg-tree merge/split, offline CDQ |
| [sparse_table.go](copypasta/sparse_table.go) | Sparse Table (RMQ), Disjoint Sparse Table, 2-D Sparse Table |
| [union_find.go](copypasta/union_find.go) | Union-Find with weights, rollback, offline dynamic connectivity |
| [heap.go](copypasta/heap.go) | Augmented heap (modify/delete by key), lazy-deletion heap, regret-greedy heap |
| [treap.go](copypasta/treap.go) | Implicit treap (split/merge), rank tree |
| [skip_list.go](copypasta/skip_list.go) | **[New]** Skip list — O(log n) ordered map, lower/upper bound, floor/ceiling |
| [van_emde_boas.go](copypasta/van_emde_boas.go) | **[New]** Van Emde Boas tree — O(log log U) successor/predecessor |
| [link_cut_tree.go](copypasta/link_cut_tree.go) | Link-Cut Tree (LCT), dynamic tree connectivity |
| [top_tree.go](copypasta/top_tree.go) | Top Tree |
| [segment_tree_rect.go](copypasta/segment_tree_rect.go) | Sweep-line area of rectangle union |
| [seg_in_bit.go](copypasta/seg_in_bit.go) | Segment tree inside BIT (2-D point update range query) |
| [kd_tree.go](copypasta/kd_tree.go) | k-d Tree, nearest-neighbor search |
| [sqrt_decomposition.go](copypasta/sqrt_decomposition.go) | Sqrt decomposition, block decomposition, Mo's algorithm |
| [mo.go](copypasta/mo.go) | Mo's algorithm (offline range queries), Mo on trees |

### String Algorithms

| File | Contents |
|------|----------|
| [strings.go](copypasta/strings.go) | KMP, Z-algorithm, string hashing |
| [acam.go](copypasta/acam.go) | Aho-Corasick automaton |
| [sam.go](copypasta/sam.go) | Suffix Automaton (SAM) |
| [pam.go](copypasta/pam.go) | Palindromic Automaton (PAM / Eertree) |
| [trie.go](copypasta/trie.go) | Trie (prefix tree) |
| [trie01.go](copypasta/trie01.go) | Binary trie (XOR trie) |

### Mathematics

| File | Contents |
|------|----------|
| [math.go](copypasta/math.go) | Number theory: primes, factorization, Euler's totient, Möbius function, CRT |
| [math_comb.go](copypasta/math_comb.go) | Combinatorics: binomial coefficients, Catalan, Stirling, partition |
| [math_matrix.go](copypasta/math_matrix.go) | Matrix exponentiation, Gaussian elimination |
| [math_fft.go](copypasta/math_fft.go) | FFT, polynomial operations |
| [math_ntt.go](copypasta/math_ntt.go) | NTT (Number Theoretic Transform) |
| [math_gf.go](copypasta/math_gf.go) | Generating functions, formal power series |
| [math_fwt.go](copypasta/math_fwt.go) | Fast Walsh-Hadamard Transform |
| [math_numerical_analysis.go](copypasta/math_numerical_analysis.go) | Numerical methods: Newton's method, ternary search |
| [math_continued_fraction.go](copypasta/math_continued_fraction.go) | Continued fractions, Stern-Brocot tree |
| [big.go](copypasta/big.go) | Big integer arithmetic |

### Graph Theory

| File | Contents |
|------|----------|
| [graph.go](copypasta/graph.go) | Shortest paths (Dijkstra, Bellman-Ford, SPFA, Floyd), topological sort, SCC, Bridges, 2-SAT |
| [graph_tree.go](copypasta/graph_tree.go) | Tree: LCA, heavy-light decomposition, centroid decomposition, virtual tree, DSU on tree |
| [graph_flow.go](copypasta/graph_flow.go) | Network flow: max flow (Dinic), min cost max flow, bipartite matching (Hopcroft-Karp), Hungarian algorithm |
| [graph_grid.go](copypasta/graph_grid.go) | Grid graph algorithms |

### Dynamic Programming

| File | Contents |
|------|----------|
| [dp.go](copypasta/dp.go) | Knapsack (0-1, unbounded, grouped), interval DP, digit DP, bitmask DP, tree DP, DP optimizations (monotone queue, divide & conquer, Knuth, slope trick) |

### Geometry

| File | Contents |
|------|----------|
| [geometry.go](copypasta/geometry.go) | 2-D / 3-D computational geometry: convex hull, half-plane intersection, rotating calipers, Voronoi |

### Miscellaneous

| File | Contents |
|------|----------|
| [bits.go](copypasta/bits.go) | Bit manipulation tricks |
| [sort.go](copypasta/sort.go) | Sorting algorithms, order statistics |
| [search.go](copypasta/search.go) | Binary search, ternary search |
| [games.go](copypasta/games.go) | Game theory: Sprague-Grundy, Nim |
| [rand.go](copypasta/rand.go) | Randomized algorithms, hashing |

---

## Fork Contributions

This fork extends the original with:

| Contribution | Description |
|-------------|-------------|
| `copypasta/skip_list.go` | Full skip-list implementation with floor/ceiling/lower-bound queries and benchmark |
| `copypasta/van_emde_boas.go` | Complete Van Emde Boas tree with O(log log U) successor/predecessor |
| `.github/workflows/go.yml` | GitHub Actions CI — build + test on every push / PR |
| `docs/` | English algorithm documentation site (GitHub Pages) |

---

## Contributing

We welcome contributions that:

- Add algorithm templates not yet covered  
- Add English explanations or comments to existing templates  
- Provide additional problem references with editorial links  
- Fix bugs or improve performance of existing implementations  

Please open an issue or pull request. For large changes, discuss in an issue first.

### Code Style

Templates follow the conventions of the original project:

```go
// copypasta/your_algorithm.go
package copypasta

/*
Algorithm Name
https://en.wikipedia.org/wiki/...  (reference)

Core idea: ...
Time complexity:  O(...)
Space complexity: O(...)

Related problems:
LC1234 https://leetcode.cn/problems/...
https://codeforces.com/problemset/problem/XXX/Y  RRRR
*/

func yourAlgorithm(n int) int {
    // implementation
}
```

---

## License

[MIT](LICENSE) © 2019 EndlessCheng
