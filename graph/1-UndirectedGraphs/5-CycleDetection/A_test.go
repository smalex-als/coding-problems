package main

import (
	"coding-problems/test"
	"math/rand"
	"testing"
)

func TestA(t *testing.T) {
	test.RunTest("A", t)
}

func TestBIG(t *testing.T) {
	g := GenGraphTree(20)
	t.Log("Big", len(g))
	if hasCycle(len(g), g) {
		t.Fail()
	}
	u := rand.Intn(len(g))
	v := rand.Intn(len(g))
	g[u] = append(g[u], v)
	if !hasCycle(len(g), g) {
		t.Fail()
	}
}

func GenGraphTree(level int) [][]int {
	nextId := 0
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
	GenGraphTreeDfs(g, 0, level, &nextId)
	return g
}

func GenGraphTreeDfs(g [][]int, from int, level int, nextId *int) {
	if level == 1 {
		return
	}
	for i := 0; i < 2; i++ {
		*nextId++
		g[from] = append(g[from], *nextId)
		GenGraphTreeDfs(g, *nextId, level-1, nextId)
	}
}
