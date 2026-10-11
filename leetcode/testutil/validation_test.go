package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInvalidTestDataReturnsError(t *testing.T) {
	var nilFunc func(int) int
	var nilConstructor func() foo
	file := filepath.Join(t.TempDir(), "cases.txt")
	if err := os.WriteFile(file, []byte("1\n"), 0600); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		run  func(*testing.T) error
	}{
		{"typed nil function", func(t *testing.T) error {
			return RunLeetCodeFuncWithExamples(t, nilFunc, [][]string{{"1", "1"}}, 0)
		}},
		{"typed nil constructor", func(t *testing.T) error {
			return RunLeetCodeClassWithExamples(t, nilConstructor, [][3]string{{"[]", "[]", "[]"}}, 0)
		}},
		{"non function", func(t *testing.T) error {
			return RunLeetCodeFuncWithExamples(t, 1, [][]string{{"1"}}, 0)
		}},
		{"non constructor", func(t *testing.T) error {
			return RunLeetCodeClassWithExamples(t, 1, [][3]string{{"[]", "[]", "[]"}}, 0)
		}},
		{"nil function", func(t *testing.T) error {
			return RunLeetCodeFuncWithExamples(t, nil, [][]string{{"1"}}, 0)
		}},
		{"nil function with file", func(t *testing.T) error {
			return RunLeetCodeFuncWithFile(t, nil, file, 0)
		}},
		{"zero arity function with file", func(t *testing.T) error {
			return RunLeetCodeFuncWithFile(t, func() {}, file, 0)
		}},
		{"nil constructor", func(t *testing.T) error {
			return RunLeetCodeClassWithExamples(t, nil, [][3]string{{"[]", "[]", "[]"}}, 0)
		}},
		{"missing function output", func(t *testing.T) error {
			return RunLeetCodeFuncWithCase(t, func(x int) int { return x }, [][]string{{"1"}}, nil, 0)
		}},
		{"extra function output", func(t *testing.T) error {
			return RunLeetCodeFuncWithCase(t, func(x int) int { return x }, nil, [][]string{{"1"}}, 0)
		}},
		{"missing class output", func(t *testing.T) error {
			return RunLeetCodeClassWithCase(t, constructor, []string{"[\"foo\"]\n[[]]"}, nil, 0)
		}},
		{"extra class output", func(t *testing.T) error {
			return RunLeetCodeClassWithCase(t, constructor, nil, []string{"[]"}, 0)
		}},
		{"missing class arguments", func(t *testing.T) error {
			return RunLeetCodeClassWithCase(t, constructor, []string{"[\"foo\"]"}, []string{"[null]"}, 0)
		}},
		{"extra class input line", func(t *testing.T) error {
			return RunLeetCodeClassWithCase(t, constructor, []string{"[\"foo\"]\n[[]]\n[]"}, []string{"[null]"}, 0)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if p := recover(); p != nil {
					t.Errorf("invalid test data caused panic: %v", p)
				}
			}()
			if err := tt.run(t); err == nil {
				t.Error("expected a validation error")
			}
		})
	}
}
