package copypasta

import (
	"math/rand"
	"testing"
)

func TestFenwickTreeMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	for n := 1; n <= 40; n++ {
		ft := newFenwickTree(n)
		a := make([]int, n+1) // 下标从 1 开始
		for range 200 {
			i, val := rng.Intn(n)+1, rng.Intn(201)-100
			ft.update(i, val)
			a[i] += val

			l, r := rng.Intn(n)+1, rng.Intn(n)+1
			want := 0
			for j := l; j <= r; j++ {
				want += a[j]
			}
			if got := ft.query(l, r); got != want { // l > r 时两者都是 0
				t.Fatalf("n=%d query(%d,%d)=%d, want %d", n, l, r, got, want)
			}

			pre := 0
			for j := 1; j <= r; j++ {
				pre += a[j]
			}
			if got := ft.pre(r); got != pre {
				t.Fatalf("n=%d pre(%d)=%d, want %d", n, r, got, pre)
			}
		}
		total := 0
		for _, v := range a {
			total += v
		}
		if got := ft.pre(n + 5); got != total { // 超出范围的下标按 n 处理
			t.Fatalf("n=%d pre(%d)=%d, want %d", n, n+5, got, total)
		}
	}
}

func TestFenwickTreeDiffMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(12))
	for n := 1; n <= 40; n++ {
		ft := newFenwickTreeDiff(n)
		a := make([]int, n+1)
		for range 200 {
			l, r := randRange(rng, n)
			l, r = l+1, r+1
			val := rng.Intn(201) - 100
			ft.add(l, r, val)
			for j := l; j <= r; j++ {
				a[j] += val
			}

			l, r = randRange(rng, n)
			l, r = l+1, r+1
			want := 0
			for j := l; j <= r; j++ {
				want += a[j]
			}
			if got := ft.query(l, r); got != want {
				t.Fatalf("n=%d query(%d,%d)=%d, want %d", n, l, r, got, want)
			}
		}
	}
}

func TestFenwickTree2DMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(13))
	for range 60 {
		n, m := rng.Intn(9)+1, rng.Intn(9)+1
		ft := newFenwickTree2D(n, m)
		a := make([][]int, n+1)
		for i := range a {
			a[i] = make([]int, m+1)
		}
		for range 100 {
			x, y, val := rng.Intn(n)+1, rng.Intn(m)+1, rng.Intn(201)-100
			ft.update(x, y, val)
			a[x][y] += val

			x1, x2 := randRange(rng, n)
			y1, y2 := randRange(rng, m)
			x1, x2, y1, y2 = x1+1, x2+1, y1+1, y2+1
			want := 0
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					want += a[i][j]
				}
			}
			if got := ft.query(x1, y1, x2, y2); got != want {
				t.Fatalf("n=%d m=%d query(%d,%d,%d,%d)=%d, want %d", n, m, x1, y1, x2, y2, got, want)
			}
			if got := ft.query2(x1, x2, y1, y2); got != want {
				t.Fatalf("n=%d m=%d query2(%d,%d,%d,%d)=%d, want %d", n, m, x1, x2, y1, y2, got, want)
			}
		}
	}
}
