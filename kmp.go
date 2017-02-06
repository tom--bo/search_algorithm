package main

import ()

func kmp(text, query string) int {
	tLen := len(text)
	qLen := len(query)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return -1
	}
	next := makeNextIndex(query)
	j := -1

	for i := -1; i < tLen; i++ {
		for j > -1 && j < qLen && query[j] != text[i] {
			j = next[j]
		}
		if j == qLen-1 {
            return i-qLen+1
		}
        j++
	}

	return -1
}

func kmpAll(text, query string) []int {
	ret := []int{}
	tLen := len(text)
	qLen := len(query)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return []int{}
	}
	next := makeNextIndex(query)
	j := -1

	for i := -1; i < tLen; i++ {
		for j > -1 && j < qLen && query[j] != text[i] {
			j = next[j]
		}
		if j == qLen-1 {
			ret = append(ret, i-qLen+1)
			j = -1
			i -= qLen - 1
		}
		j++
	}

	return ret
}

func makeNextIndex(query string) []int {
	qLen := len(query)
	ret := make([]int, qLen)
	t := -1
	ret[0] = -1

	for j := 0; j < qLen-1; j++ {
		for t != -1 && query[j] != query[t] && t < qLen {
			t = ret[t]
		}
		t++
		if query[j+1] == query[t] {
			ret[j+1] = ret[t]
		} else {
			ret[j+1] = t
		}
	}
	return ret
}
