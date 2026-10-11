package testutil

import (
	"io"
	"reflect"
	"testing"
)

func TestParseSampleCases(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want [][2]string
	}{
		{"copy markers", "inputCopy\n1\noutputCopy\n2\ninputCopy\n3\noutputCopy\n4", [][2]string{{"1", "2"}, {"3", "4"}}},
		{"plain markers", "input\n1\noutput\n2", [][2]string{{"1", "2"}}},
		{"implicit first input", "1\noutputCopy\n2\ninputCopy\n3\noutputCopy\n4", [][2]string{{"1", "2"}, {"3", "4"}}},
		{"marker words in data", "input\ninputCopycat output values\noutput\nmyinput outputCopycat", [][2]string{{"inputCopycat output values", "myinput outputCopycat"}}},
		{"CRLF and indentation", "  inputCopy \r\n1\r\n  outputCopy \r\n2\r\n", [][2]string{{"1\r", "2"}}},
		{"empty output", "inputCopy\n1\noutputCopy", [][2]string{{"1", ""}}},
		{"multiline data", "input\n1\n2\noutput\n3\n4", [][2]string{{"1\n2", "3\n4"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSampleCases(tt.raw)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestParseSampleCasesRejectsMalformedData(t *testing.T) {
	for _, raw := range []string{"", " \n ", "1", "input\n1", "output\n1", "input\n1\ninput\n2\noutput\n3", "input\n1\noutput\n2\noutput\n3", "input\n1\noutput\n2\ninput\n3"} {
		t.Run(raw, func(t *testing.T) {
			if _, err := parseSampleCases(raw); err == nil {
				t.Fatal("expected a validation error")
			}
		})
	}
}

func TestAssertEqualCasePreservesMarkerWords(t *testing.T) {
	AssertEqualCase(t, "input\ninputCopycat output values\noutput\ninputCopycat output values", 0, func(r io.Reader, w io.Writer) {
		if _, err := io.Copy(w, r); err != nil {
			t.Error(err)
		}
	})
}
