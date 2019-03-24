package concat_test

import (
	"strings"
	"testing"
)

var args = []string{"hi", "there", "buddy", "5", "hi", "there", "buddy", "boy", "5", "6", "7", "8", "9", "hi", "there", "buddy", "boy", "5", "6", "7", "8", "9"}

func concat(args []string) {
	var r, sep string
	for _, a := range args {
		r += sep + a
		sep = " "
	}

}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
