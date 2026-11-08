package copypasta

import (
	"math"
	"math/rand"
	"slices"
	"testing"
)

// 无向图的 Bellman-Ford，作为最短路模板的对照；不可达时为 inf
func bellmanFordUndirected(n, st int, edges [][]int, inf int) []int {
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	for range n {
		for _, e := range edges {
			v, w, wt := e[0], e[1], e[2]
			if dis[v] < inf {
				dis[w] = min(dis[w], dis[v]+wt)
			}
			if dis[w] < inf {
				dis[v] = min(dis[v], dis[w]+wt)
			}
		}
	}
	return dis
}

func randEdges(rng *rand.Rand, n int) [][]int {
	edges := make([][]int, 0, n*2)
	for range rng.Intn(n*2 + 1) {
		v, w := rng.Intn(n), rng.Intn(n)
		if v != w {
			edges = append(edges, []int{v, w, rng.Intn(20) + 1})
		}
	}
	return edges
}

func TestDijkstraShortestPathMatchesBruteForce(t *testing.T) {
	const inf int = 1e18
	rng := rand.New(rand.NewSource(4))
	for range 500 {
		n := rng.Intn(12) + 1
		edges := randEdges(rng, n)
		st := rng.Intn(n)
		got := (&graph{}).dijkstraShortestPath(n, st, edges)
		if want := bellmanFordUndirected(n, st, edges, inf); !slices.Equal(got, want) {
			t.Fatalf("n=%d st=%d edges=%v\n got %v\nwant %v", n, st, edges, got, want)
		}
	}
}

func TestDijkstraShortestPath2MatchesBruteForce(t *testing.T) {
	const inf int = 1e18
	rng := rand.New(rand.NewSource(5))
	for range 500 {
		n := rng.Intn(12) + 1
		edges := randEdges(rng, n)
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				g[i][j] = inf
			}
			g[i][i] = 0
		}
		for _, e := range edges {
			v, w, wt := e[0], e[1], e[2]
			g[v][w] = min(g[v][w], wt)
			g[w][v] = min(g[w][v], wt)
		}
		st := rng.Intn(n)
		got := (&graph{}).dijkstraShortestPath2(g, st)[:n]
		if want := bellmanFordUndirected(n, st, edges, inf); !slices.Equal(got, want) {
			t.Fatalf("n=%d st=%d edges=%v\n got %v\nwant %v", n, st, edges, got, want)
		}
	}
}

func TestFloydShortestPathMatchesBruteForce(t *testing.T) {
	const inf = math.MaxInt / 2
	rng := rand.New(rand.NewSource(6))
	for range 300 {
		n := rng.Intn(12) + 1
		edges := randEdges(rng, n)
		got := (&graph{}).floydShortestPath(n, edges)
		for st := range n {
			if want := bellmanFordUndirected(n, st, edges, inf); !slices.Equal(got[st], want) {
				t.Fatalf("n=%d st=%d edges=%v\n got %v\nwant %v", n, st, edges, got[st], want)
			}
		}
	}
}
