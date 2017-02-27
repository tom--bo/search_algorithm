package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	text, query string
	want        []int
}{
	{"", "", []int{-1}},
	{"", "a", []int{-1}},
	{"abc", "", []int{-1}},
	{"abcde", "ab", []int{0}},
	{"aaaaa", "aa", []int{0, 1, 2, 3}},
	{"abcde\nabcd", "ab", []int{0, 6}},
	{"abcde", "ac", []int{-1}},
	{"abcde", "a", []int{0}},
	{"abcabeabcabeababcabe", "abcabe", []int{0, 6, 14}},
	{"abcababcabeababcabe", "abcabe", []int{5, 13}},
	{"テスト", "テ", []int{0}},
	{"漢字な感じ", "感じ", []int{9}},
	{"which finally halts.  at that point", "at that", []int{22}},
	{"grep  searches the named input FILEs for lines containing a match to the given PATTERN.  If no files are specified, or if the", "are", []int{101}},
}

var testCasesForMultiInput = []struct {
	text  string
	query []string
	want  []int
}{
	{"", []string{}, []int{-1}},
	{"", []string{"a"}, []int{-1}},
	{"abc", []string{"", ""}, []int{-1}},
	{"abcde", []string{"ab"}, []int{0}},
	{"xbabcdex", []string{"ab", "abcde"}, []int{2, 2}},
	{"xbabcdex", []string{"x", "ab", "abcde"}, []int{0, 2, 2, 7}},
	{"xbabcdex", []string{"ab", "bc", "bab", "d", "abcde"}, []int{1, 2, 2, 3, 5}},
	{"aaaa", []string{"a", "b"}, []int{0, 1, 2, 3}},
	{"which finally halts.  at that point", []string{"at that"}, []int{22}},
}

func test(t *testing.T, f func(string, string) int) {
	for num, tc := range testCases {
		got := f(tc.text, tc.query)
		want := tc.want[0]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func testAllMatchedPatterns(t *testing.T, f func(string, string) []int) {
	for num, tc := range testCases {
		got := f(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func testForMultiImput(t *testing.T, f func(string, []string) int) {
	for num, tc := range testCasesForMultiInput {
		got := f(tc.text, tc.query)
		want := tc.want[0]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func testAllMatchedPatternsForMultiImput(t *testing.T, f func(string, []string) []int) {
	for num, tc := range testCasesForMultiInput {
		got := f(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func TestSimpleSearch(t *testing.T) {
	test(t, simpleSearch)
}

func TestKMPSearch(t *testing.T) {
	test(t, kmp)
}

func TestBMSearch(t *testing.T) {
	test(t, bm)
}

func TestAC(t *testing.T) {
	testForMultiImput(t, ac)
}

func TestSimpleSearchAll(t *testing.T) {
	testAllMatchedPatterns(t, simpleSearchAll)
}

func TestKMPAll(t *testing.T) {
	testAllMatchedPatterns(t, kmpAll)
}

func TestBMAll(t *testing.T) {
	testAllMatchedPatterns(t, bmAll)
}

func TestACAll(t *testing.T) {
	testAllMatchedPatternsForMultiImput(t, acAll)
}
