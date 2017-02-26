package main

import ()

func bm(text, query string) int {
	qLen := len(query)
	tLen := len(text)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return -1
	}
	pos := qLen - 1

	m := make(map[uint8]int)

	for i := 0; i < qLen-1; i++ {
		m[query[i]] = qLen - i - 1
	}
	m[query[qLen-1]] = qLen

	for pos < tLen {
		k := pos
		j := qLen - 1
		for j >= 0 && text[k] == query[j] {
			k -= 1
			j -= 1
		}
		if j == -1 {
			return pos - qLen + 1
		}
		p, ok := m[text[pos]]
		if ok {
			pos = pos + p
		} else {
			pos += qLen
		}
	}

	return -1
}

func bmAll(text, query string) []int {
	ret := []int{}
	qLen := len(query)
	tLen := len(text)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return []int{}
	}
	pos := qLen - 1

	m := make(map[uint8]int)

	for i := 0; i < qLen-1; i++ {
		m[query[i]] = qLen - i - 1
	}
	m[query[qLen-1]] = qLen

	for pos < tLen {
		k := pos
		j := qLen - 1
		for j >= 0 && text[k] == query[j] {
			k -= 1
			j -= 1
		}
		if j == -1 {
			ret = append(ret, pos-qLen+1)
			pos++
			continue
		}
		p, ok := m[text[pos]]
		if ok {
			pos = pos + p
		} else {
			pos += qLen
		}
	}

	return ret
}
