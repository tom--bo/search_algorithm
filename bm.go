package main

import ()

func bm(text, query string) []int {
	ret := []int{}
	qLen := len(query)
	tLen := len(text)
	if qLen == 0 || tLen == 0 || tLen < qLen {
		return []int{}
	}
	pos := qLen - 1

	m := preProcess(query)

	for pos < tLen {
		if text[pos] == query[qLen-1] {
			k := pos
			j := qLen - 1
			for j > 0 && text[k] == query[j] {
				k -= 1
				j -= 1
			}
			if j == 0 {
				ret = append(ret, pos-qLen+1)
				pos++
				continue
			}
		}
		pos = pos + skip(m, text[pos], qLen)
	}

	return ret
}

func preProcess(query string) map[uint8]int {
	m := make(map[uint8]int)
	qLen := len(query)

	for i := 0; i < qLen-1; i++ {
		m[query[i]] = qLen - i - 1
	}
	m[query[qLen-1]] = qLen

	return m
}

func skip(m map[uint8]int, ch uint8, qLen int) int {
	pos, ok := m[ch]
	if ok {
		return pos
	}
	return qLen
}
