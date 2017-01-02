package main

import (
	_ "fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	if !reflect.DeepEqual(bm("abcde", "ab"), []int{0}) {
		t.Fatal("Failed at case1")
	}

	if !reflect.DeepEqual(bm("aaaaa", "aa"), []int{0, 1, 2, 3}) {
		t.Fatal("Failed at case2")
	}

	if !reflect.DeepEqual(bm("abcdeab", "ab"), []int{0, 5}) {
		t.Fatal("Failed at case3")
	}

	if !reflect.DeepEqual(bm("abcde\nabcd", "ab"), []int{0, 6}) {
		t.Fatal("Failed at case4")
	}

	if !reflect.DeepEqual(bm("abcde", "ac"), []int{}) {
		t.Fatal("Failed at case5")
	}

	if !reflect.DeepEqual(bm("abcde", "a"), []int{0}) {
		t.Fatal("Failed at case6")
	}

	if !reflect.DeepEqual(bm("abcde", ""), []int{}) {
		t.Fatal("Failed at case7")
	}
}
