package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		text, query string
		want        []int
	}{
		{"abcde", "ab", []int{0}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"abcde\nabcd", "ab", []int{0, 6}},
		{"abcde", "ac", []int{}},
		{"abcde", "a", []int{0}},
		{"abcde", "", []int{}},
		{"which finally halts.  at that point", "at that", []int{22}},
	}

	for num, tc := range tests {
		got := bm(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
