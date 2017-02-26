package main

import (
	"reflect"
	"testing"
)

var testsMulti = []struct {
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

func TestAC(t *testing.T) {
	for num, tc := range testsMulti {
		got := ac(tc.text, tc.query)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got -> %v, want -> %v\n", got, want)
			t.Fatal("Failed at case:", num)
		}
	}

}
