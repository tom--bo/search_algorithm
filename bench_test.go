package main

import (
	"bytes"
	"strings"
	"testing"
)

func Index(text, pattern string) int {
	ret := strings.Index(text, pattern)
	return ret
}

func prepareText() string {
	var b bytes.Buffer
	for i := 0; i < 100; i++ {
		b.WriteString("abcde")
	}
	b.WriteString("0123")
	return b.String()
}

func bench(b *testing.B, f func(string, string) int) {
	b.ReportAllocs()
	text := prepareText()
	pattern := "0123"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(text, pattern)
	}
}

func BenchmarkIndex(b *testing.B)        { bench(b, Index) }
func BenchmarkSimpleSearch(b *testing.B) { bench(b, simpleSearch) }
func BenchmarkKMP(b *testing.B)          { bench(b, kmp) }
func BenchmarkBM(b *testing.B)           { bench(b, bm) }
