package pit

import "bytes"
import "strings"
import "unicode/utf8"
import "testing"


/*

go test count_test.go -bench ".*"

*/

func f1(s string) int {
	return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
	return strings.Count(s, "") - 1
}

func f3(s string) int {
	return len([]rune(s))
}

func f4(s string) int {
	return utf8.RuneCountInString(s)
}

var s = "Hello, 世界"

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1(s)
	}
}


func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2(s)
	}
}

func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f3(s)
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f4(s)
	}
}