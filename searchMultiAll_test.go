package main

import (
	"reflect"
	"testing"
)

var testsMultiAll = []struct {
    text  string
    query []string
    want  []int
}{
    {"", []string{}, []int{}},
    {"", []string{"a"}, []int{}},
    {"abc", []string{"", ""}, []int{}},
    {"abcde", []string{"ab"}, []int{0}},
    {"xbabcdex", []string{"ab", "abcde"}, []int{2, 2}},
    {"xbabcdex", []string{"x", "ab", "abcde"}, []int{0, 2, 2, 7}},
    {"xbabcdex", []string{"ab", "bc", "bab", "d", "abcde"}, []int{1, 2, 2, 3, 5}},
    {"aaaa", []string{"a", "b"}, []int{0, 1, 2, 3}},
    {"which finally halts.  at that point", []string{"at that"}, []int{22}},
}

func TestSearchMultiAll(t *testing.T) {
	for num, tc := range testsMultiAll {
		got := acAll(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
