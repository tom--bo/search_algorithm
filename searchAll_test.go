package main

import (
	"reflect"
	"testing"
)

var testsAll = []struct {
    text, query string
    want        []int
}{
    {"", "", []int{}},
    {"", "a", []int{}},
    {"abc", "", []int{}},
    {"abcde", "ab", []int{0}},
    {"aaaaa", "aa", []int{0, 1, 2, 3}},
    {"abcde\nabcd", "ab", []int{0, 6}},
    {"abcde", "ac", []int{}},
    {"abcde", "a", []int{0}},
    {"abcabeabcabeababcabe", "abcabe", []int{0, 6, 14}},
    {"abcababcabeababcabe", "abcabe", []int{5, 13}},
    {"テスト", "テ", []int{0}},
    {"漢字な感じ", "感じ", []int{9}},
    {"which finally halts.  at that point", "at that", []int{22}},
}

func TestSimpleSearchAll(t *testing.T) {
	for num, tc := range testsAll {
		got := simpleSearchAll(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}

func TestKMPSearchAll(t *testing.T) {
	for num, tc := range testsAll {
		got := kmpAll(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
func TestBMSearchAll(t *testing.T) {
	for num, tc := range testsAll {
		got := bmAll(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
