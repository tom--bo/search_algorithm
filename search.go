package main

import (
	"fmt"
)

func main() {
	text := "abcdeabcde"
	query := "ab"
	ret := simpleSearch(text, query)

	for _, s := range ret {
		fmt.Println(s)
	}
}