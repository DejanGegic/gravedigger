package checker_test

import (
	"testing"

	"github.com/dejangegic/gravedigger/checker"
)

func BenchmarkTestRunAll(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		checker.RunAll("/home/dejan/dev/go")
	}
}
