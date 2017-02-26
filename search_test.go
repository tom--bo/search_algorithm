package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
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

func TestSimpleSearch(t *testing.T) {
	for num, tc := range tests {
		got := simpleSearch(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func TestKMPSearch(t *testing.T) {
	for num, tc := range tests {
		got := kmp(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}

func TestBMSearch(t *testing.T) {
	for num, tc := range tests {
		got := bm(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}
}
