package main

import (
	"sort"
)

type Set map[string]bool
type Node struct {
	BeforeNode int
	FromLink   rune
	Link       map[rune]int
	HasOutput  bool
	Output     Set
	Failure    int
}

func ac(text string, query []string) []int {
	ret := []int{}
	qLen := len(query)
	tLen := len(text)
	if qLen == 0 || tLen == 0 {
		return []int{}
	}

	Goto := buildMachine(query)

	now := 0
	for i, ch := range text {
		next, ok := Goto[now].Link[ch]
		for !ok && now != 0 {
			now = Goto[now].Failure
			next, ok = Goto[now].Link[ch]
		}
		now = next

		if Goto[now].HasOutput {
			for _, v := range makeRet(i, Goto[now].Output) {
				ret = append(ret, v)
			}
		}
	}

	sort.Ints(ret)
	return ret
}

func buildMachine(query []string) []Node {
	Goto := buildGoto(query)
	buildFailure(query, Goto)
	return Goto
}

// build Failure & update Output
func buildFailure(query []string, Goto []Node) {
	queue := []int{}
	for _, n := range Goto[0].Link {
		Goto[n].Failure = 0
		for _, m := range Goto[n].Link {
			queue = append(queue, m)
		}
	}

	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, l := range Goto[pos].Link {
			queue = append(queue, l)
		}

		fromLink := Goto[pos].FromLink
		before := Goto[pos].BeforeNode
		fbefore := Goto[before].Failure

		if tgt, ok := Goto[fbefore].Link[fromLink]; ok {
			Goto[pos].Failure = tgt
			if Goto[tgt].HasOutput {
				for o, _ := range Goto[tgt].Output {
					Goto[pos].Output[o] = true
					Goto[pos].HasOutput = true
				}
			}
		}
		// fix later
		// 2回目以降は Goto[pos].Failure = tgt これはだめ
		// fbeforeが0になってから1度だけ実行したい...
		for fbefore != 0 {
			if tgt, ok := Goto[fbefore].Link[fromLink]; ok {
				Goto[pos].Failure = tgt
				if Goto[tgt].HasOutput {
					for o, _ := range Goto[tgt].Output {
						Goto[pos].Output[o] = true
						Goto[pos].HasOutput = true
					}
				}
			}
			fbefore = Goto[fbefore].Failure
		}
		if tgt, ok := Goto[fbefore].Link[fromLink]; ok {
			if Goto[tgt].HasOutput {
				Goto[pos].Failure = tgt
				for o, _ := range Goto[tgt].Output {
					Goto[pos].Output[o] = true
					Goto[pos].HasOutput = true
				}
			}
		}
	}
}

// build simple slice, Goto
func buildGoto(query []string) []Node {
	var Goto []Node
	Goto = append(Goto, makeNode(0, -1))
	cnt := 0
	now := 0

	for _, q := range query {
		if q == "" {
			continue
		}
		l := len(q)
		now = 0
		for i, ch := range q {
			if val, ok := Goto[now].Link[ch]; ok {
				now = val
				continue
			} else {
				Goto = append(Goto, makeNode(ch, now))
				cnt++
				Goto[now].Link[ch] = cnt
				now = cnt
			}
			if i == l-1 {
				Goto[cnt].Output[q] = true
				Goto[cnt].HasOutput = true
			}
		}
	}
	return Goto
}

// to change what values the "ret" return
func makeRet(now int, Output Set) []int {
	var ret []int
	for k, _ := range Output {
		ret = append(ret, now-len(k)+1)
	}
	return ret
}

func makeNode(fl rune, before int) Node {
	return Node{
		BeforeNode: before,
		FromLink:   fl,
		Link:       map[rune]int{},
		HasOutput:  false,
		Output:     map[string]bool{},
		Failure:    0,
	}
}
