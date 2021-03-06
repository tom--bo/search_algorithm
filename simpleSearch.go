package main

import ()

func simpleSearch(text, query string) int {
	tLen := len(text)
	qLen := len(query)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return -1
	}
	for i := 0; i < tLen-qLen+1; i++ {
		for j := 0; j < qLen; j++ {
			if text[i+j] != query[j] {
				break
			}
			if j == qLen-1 {
				return i
			}
		}
	}

	return -1
}

func simpleSearchAll(text, query string) []int {
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

	if len(ret) == 0 {
		ret = append(ret, -1)
	}
	return ret
}
