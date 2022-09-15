package practicalmemory

import (
	"bytes"
	"testing"
)

var output bytes.Buffer

var in = assembleInputStream()
var find = []byte("elvis")
var repl = []byte("Elvis")

// Capture the time it takes to execute alg on
func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

func BenchmarkAlgorithmTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}
