package main

import ()

func simpleSearch(text, query string) []int {
	ret := []int{}
	tLen := len(text)
	qLen := len(query)
	for i := 0; i < tLen-qLen+1; i++ {
		for j := 0; j < qLen; j++ {
			if text[i+j] != query[j] {
				break
			}
			if j == qLen-1 {
				ret = append(ret, i)
			}
		}
	}

	return ret
}
