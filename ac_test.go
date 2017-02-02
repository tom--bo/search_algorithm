package main

import (
	"reflect"
	"testing"
)

func TestACSearch(t *testing.T) {
	var tests = []struct {
		text  string
		query []string
		want  []int
	}{
		{"", []string{}, []int{}},
		{"", []string{"a"}, []int{}},
		//{"abc", []string{""}, []int{}},
		{"abcde", []string{"ab"}, []int{0}},
		{"xbabcdex", []string{"ab", "abcde"}, []int{2, 2}},
		{"xbabcdex", []string{"ab", "bc", "bab", "d", "abcde"}, []int{1, 2, 2, 3, 5}},
	}

	for num, tc := range tests {
		got := ac(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
