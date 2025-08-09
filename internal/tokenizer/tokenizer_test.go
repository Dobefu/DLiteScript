package tokenizer

import (
	"testing"
)

func BenchmarkTokenizer(b *testing.B) {
	for b.Loop() {
		NewTokenizer("1 + -2 * 3 / 4")
	}
}
