package main

import (
	"coding-problems/test"
	"testing"
)

func TestA(t *testing.T) {
	test.RunTest("A", t)
}

func TestBIG(t *testing.T) {
	g := GenGraph(15)
	t.Log("Big", len(g))
	if hasCycle(len(g), g) {
		t.Fail()
	}
}

var nextId int

func GenGraph(level int) [][]int {
	n := 0
	cur := 1
	for i := 0; i < level; i++ {
		n += cur
		cur *= 2
	}
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	GenGraphDfs(g, 0, level)
	return g
}

func GenGraphDfs(g [][]int, from int, level int) {
	if level == 1 {
		return
	}
	for i := 0; i < 2; i++ {
		nextId++
		g[from] = append(g[from], nextId)
		GenGraphDfs(g, nextId, level-1)
	}
}
