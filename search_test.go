package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	if !reflect.DeepEqual(simpleSearch("abcde", "ab"), []int{0}) {
		t.Fatal("Failed at case1")
	}

	if !reflect.DeepEqual(simpleSearch("aaaaa", "aa"), []int{0, 1, 2, 3}) {
		t.Fatal("Failed at case2")
	}

	if !reflect.DeepEqual(simpleSearch("abcdeab", "ab"), []int{0, 5}) {
		t.Fatal("Failed at case3")
	}

	if !reflect.DeepEqual(simpleSearch("abcde\nabcd", "ab"), []int{0, 6}) {
		t.Fatal("Failed at case4")
	}

	if !reflect.DeepEqual(simpleSearch("abcde", "ac"), []int{}) {
		t.Fatal("Failed at case5")
	}

	if !reflect.DeepEqual(simpleSearch("abcde", "a"), []int{0}) {
		t.Fatal("Failed at case6")
	}

	if !reflect.DeepEqual(simpleSearch("abcde", ""), []int{}) {
		t.Fatal("Failed at case7")
	}
}
