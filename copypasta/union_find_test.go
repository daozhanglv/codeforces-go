package copypasta

import (
	"math/rand"
	"testing"
)

func TestUnionFindMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for n := 1; n <= 40; n++ {
		u := newUnionFind(n)
		label := make([]int, n)
		for i := range label {
			label[i] = i
		}
		groups := n
		for range 200 {
			x, y := rng.Intn(n), rng.Intn(n)
			wantSame := label[x] == label[y]
			if got := u.same(x, y); got != wantSame {
				t.Fatalf("n=%d same(%d,%d)=%v, want %v", n, x, y, got, wantSame)
			}
			if got := u.find(x) == u.findR(x); !got {
				t.Fatalf("n=%d find and findR disagree on %d", n, x)
			}

			root := u.merge(x, y)
			if wantSame {
				if root != -1 {
					t.Fatalf("n=%d merge(%d,%d)=%d, want -1", n, x, y, root)
				}
			} else {
				if root != u.find(x) || root != u.find(y) {
					t.Fatalf("n=%d merge(%d,%d) returned non-root %d", n, x, y, root)
				}
				old := label[x]
				for i, l := range label {
					if l == old {
						label[i] = label[y]
					}
				}
				groups--
			}
			if u.groups != groups {
				t.Fatalf("n=%d groups=%d, want %d", n, u.groups, groups)
			}
		}
	}
}
