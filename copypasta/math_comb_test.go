package copypasta

import "testing"

func TestCombMatchesPascalTriangle(t *testing.T) {
	const mx = 300
	pascal := make([][]int, mx+1)
	for n := range pascal {
		pascal[n] = make([]int, n+1)
		pascal[n][0], pascal[n][n] = 1, 1
		for k := 1; k < n; k++ {
			pascal[n][k] = (pascal[n-1][k-1] + pascal[n-1][k]) % mod
		}
	}

	c := newComb()
	// 先查大的再查小的、以及从小到大逐步扩容，两种顺序都要正确
	for _, order := range [][2]int{{mx, -1}, {0, 1}} {
		for n := order[0]; 0 <= n && n <= mx; n += order[1] {
			for k := -1; k <= n+1; k++ {
				want := 0
				if 0 <= k && k <= n {
					want = pascal[n][k]
				}
				if got := c.c(n, k); got != want {
					t.Fatalf("c(%d,%d)=%d, want %d", n, k, got, want)
				}
			}
		}
		c = newComb()
	}
}

func TestCombFactorialAndPermutation(t *testing.T) {
	c := newComb()
	fact := 1
	for n := 0; n <= 300; n++ {
		if n > 0 {
			fact = fact * n % mod
		}
		if got := c.f(n); got != fact {
			t.Fatalf("f(%d)=%d, want %d", n, got, fact)
		}
		if got := c.f(n) * c.invF(n) % mod; got != 1 {
			t.Fatalf("f(%d)*invF(%d)=%d, want 1", n, n, got)
		}

		perm := 1 // n*(n-1)*...*(n-k+1)
		for k := 0; k <= n; k++ {
			if got := c.p(n, k); got != perm {
				t.Fatalf("p(%d,%d)=%d, want %d", n, k, got, perm)
			}
			perm = perm * (n - k) % mod
		}
		if got := c.p(n, n+1); got != 0 {
			t.Fatalf("p(%d,%d)=%d, want 0", n, n+1, got)
		}
		if got := c.p(n, -1); got != 0 {
			t.Fatalf("p(%d,-1)=%d, want 0", n, got)
		}
	}
}
