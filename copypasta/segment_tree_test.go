package copypasta

import (
	"math/rand"
	"slices"
	"testing"
)

func firstIndex(a []int, l, r int, ok func(int) bool) int {
	for i := l; i <= r; i++ {
		if ok(a[i]) {
			return i
		}
	}
	return -1
}

func lastIndex(a []int, l, r int, ok func(int) bool) int {
	for i := r; i >= l; i-- {
		if ok(a[i]) {
			return i
		}
	}
	return -1
}

func randRange(rng *rand.Rand, n int) (l, r int) {
	l, r = rng.Intn(n), rng.Intn(n)
	if l > r {
		l, r = r, l
	}
	return
}

func TestSegmentTreeMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	for n := 1; n <= 40; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = rng.Intn(100)
		}
		st := newSegmentTree(a)
		for range 200 {
			i := rng.Intn(n)
			val := rng.Intn(100)
			st.update(1, i, val)
			a[i] = max(a[i], val) // 模板的 set 把新值与旧值取 max

			l, r := randRange(rng, n)
			if got, want := st.query(1, l, r), slices.Max(a[l:r+1]); got != want {
				t.Fatalf("n=%d query(%d,%d)=%d, want %d", n, l, r, got, want)
			}
			if got, want := st.queryAll(), slices.Max(a); got != want {
				t.Fatalf("n=%d queryAll=%d, want %d", n, got, want)
			}

			target := rng.Intn(120)
			nodeOK := func(v int) bool { return v >= target }
			if got, want := st.findFirst(1, l, r, nodeOK), firstIndex(a, l, r, nodeOK); got != want {
				t.Fatalf("n=%d findFirst(%d,%d,>=%d)=%d, want %d", n, l, r, target, got, want)
			}
			if got, want := st.findLast(1, l, r, nodeOK), lastIndex(a, l, r, nodeOK); got != want {
				t.Fatalf("n=%d findLast(%d,%d,>=%d)=%d, want %d", n, l, r, target, got, want)
			}
		}
	}
}

func TestLazySegmentTreeMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	sum := func(a []int) (s int) {
		for _, v := range a {
			s += v
		}
		return
	}
	for n := 1; n <= 40; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = rng.Intn(2)
		}
		st := newLazySegmentTree(a)
		positive := func(v int) bool { return v > 0 } // 元素非负，区间和大于 0 等价于区间内存在正数
		for range 200 {
			l, r := randRange(rng, n)
			f := rng.Intn(3)
			st.update(1, l, r, f)
			for i := l; i <= r; i++ {
				a[i] += f
			}

			l, r = randRange(rng, n)
			if got, want := st.query(1, l, r), sum(a[l:r+1]); got != want {
				t.Fatalf("n=%d query(%d,%d)=%d, want %d", n, l, r, got, want)
			}
			if got, want := st.queryAll(), sum(a); got != want {
				t.Fatalf("n=%d queryAll=%d, want %d", n, got, want)
			}
			if got, want := st.findFirst(1, l, r, positive), firstIndex(a, l, r, positive); got != want {
				t.Fatalf("n=%d findFirst(%d,%d)=%d, want %d", n, l, r, got, want)
			}
			if got, want := st.findLast(1, l, r, positive), lastIndex(a, l, r, positive); got != want {
				t.Fatalf("n=%d findLast(%d,%d)=%d, want %d", n, l, r, got, want)
			}
		}
	}
}
