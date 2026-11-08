package copypasta

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestPowMatchesBigInt(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	bigMod := big.NewInt(mod)
	for range 1000 {
		x, n := rng.Intn(2*mod), rng.Intn(1e9)
		want := new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(n)), bigMod).Int64()
		if got := pow(x, n); int64(got) != want {
			t.Fatalf("pow(%d,%d)=%d, want %d", x, n, got, want)
		}

		p := rng.Intn(1e9) + 1
		want = new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(n)), big.NewInt(int64(p))).Int64()
		if got := powM(x, n, p); int64(got) != want {
			t.Fatalf("powM(%d,%d,%d)=%d, want %d", x, n, p, got, want)
		}
	}
	if got := powM(5, 0, 1); got != 0 {
		t.Fatalf("powM(5,0,1)=%d, want 0", got)
	}
}

func TestGeometricSum(t *testing.T) {
	rng := rand.New(rand.NewSource(8))
	for range 300 {
		x := rng.Intn(mod)
		sum, p := 0, 1
		for n := range 200 {
			sum = (sum + p) % mod
			p = p * x % mod
			if got := gp(x, n); got != sum {
				t.Fatalf("gp(%d,%d)=%d, want %d", x, n, got, sum)
			}
		}
	}
}

func TestMulMatchesBigInt(t *testing.T) {
	rng := rand.New(rand.NewSource(9))
	bigMod := big.NewInt(mod)
	for range 1000 {
		a, b := rng.Intn(mod), rng.Intn(mod)
		want := new(big.Int).Mul(big.NewInt(int64(a)), big.NewInt(int64(b)))
		want.Mod(want, bigMod)
		if got := mul(a, b); int64(got) != want.Int64() {
			t.Fatalf("mul(%d,%d)=%d, want %d", a, b, got, want.Int64())
		}
	}
}

func TestLogFunctions(t *testing.T) {
	for k := 2; k <= 12; k++ {
		for n := 1; n <= 5000; n++ {
			floor, p := 0, 1
			for p*k <= n {
				p *= k
				floor++
			}
			if got := floorLog(n, k); got != floor {
				t.Fatalf("floorLog(%d,%d)=%d, want %d", n, k, got, floor)
			}

			ceil, q := 0, 1
			for q < n {
				q *= k
				ceil++
			}
			if got := ceilLog(n, k); got != ceil {
				t.Fatalf("ceilLog(%d,%d)=%d, want %d", n, k, got, ceil)
			}
			if got := ceilLogSafe(n, k); got != ceil {
				t.Fatalf("ceilLogSafe(%d,%d)=%d, want %d", n, k, got, ceil)
			}
		}
	}
}
