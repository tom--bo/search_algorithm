package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	text, query string
	want        int
}{
	{"", "", -1},
	{"", "a", -1},
	{"abc", "", -1},
	{"abcde", "ab", 0},
	{"aaaaa", "aa", 0},
	{"abcde\nabcd", "ab", 0},
	{"abcde", "ac", -1},
	{"abcde", "a", 0},
	{"abcabeabcabeababcabe", "abcabe", 0},
	{"abcababcabeababcabe", "abcabe", 5},
	{"テスト", "テ", 0},
	{"漢字な感じ", "感じ", 9},
	{"which finally halts.  at that point", "at that", 22},
	{"grep  searches the named input FILEs for lines containing a match to the given PATTERN.  If no files are specified, or if the", "are", 101},
}

var testCasesForMultiInput = []struct {
	text  string
	query []string
	want  int
}{
	{"", []string{}, -1},
	{"", []string{"a"}, -1},
	{"abc", []string{"", ""}, -1},
	{"abcde", []string{"ab"}, 0},
	{"xbabcdex", []string{"ab", "abcde"}, 2},
	{"xbabcdex", []string{"x", "ab", "abcde"}, 0},
	{"xbabcdex", []string{"ab", "bc", "bab", "d", "abcde"}, 1},
	{"aaaa", []string{"a", "b"}, 0},
	{"which finally halts.  at that point", []string{"at that"}, 22},
}

func test(t *testing.T, f func(string, string) int) {
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
